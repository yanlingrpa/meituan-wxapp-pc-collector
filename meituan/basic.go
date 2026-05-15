package meituan

import (
	"fmt"

	"github.com/yanlingrpa/wxapp-pc-toolkits/wxapputils"
	"yanlingrpa.com/yanling/protocol/script"
)

var subscribers map[string]script.Subscriber

func Initialize(rt script.ModuleRuntime) (bool, error) {
	if subscribers == nil {
		subscribers = make(map[string]script.Subscriber)
	}
	sb, err := rt.Subscribe("github.com/yanlingrpa/wxapp-pc-toolkits/wxapputils", "app_ready", onAppReady)
	if err != nil {
		return false, err
	}
	subscribers["app_ready"] = sb
	return true, nil
}

func Finalize(rt script.ModuleRuntime) (bool, error) {
	if subscribers == nil {
		return true, nil
	}
	for _, sb := range subscribers {
		rt.Unsubscribe(sb)
	}
	subscribers = nil
	return true, nil
}

func Prepare(rt script.ModuleRuntime) (bool, error) {
	guiId, err := getGuiId(rt)
	if err != nil {
		return false, err
	}
	location, err := getLocation(rt)
	if err != nil {
		return false, err
	}
	ready, err := wxapputils.CheckWxappReady(rt, guiId)
	if err != nil {
		return false, fmt.Errorf("failed to check wxapp readiness: %w", err)
	}

	if !ready {
		return false, fmt.Errorf("wxapp is not ready")
	}
	_, err = wxapputils.ChangeGPSLocation(rt, wxapputils.PreferedLocation{
		GuiId:   guiId,
		Keyword: location,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
