package meituan

import (
	"fmt"

	"yanlingrpa.com/yanling/protocol/script"
)

func getGuiId(rt script.ModuleRuntime) (string, error) {
	guiId, ok := rt.StringVariable("wxapp-meituan")
	if !ok {
		return "", fmt.Errorf("wxapp-meituan is not ready")
	}
	return guiId, nil
}

func getLocation(rt script.ModuleRuntime) (string, error) {
	location, ok := rt.StringVariable("location")
	if !ok {
		return "", fmt.Errorf("location is not ready")
	}
	if location == "" {
		return "", fmt.Errorf("location is empty")
	}
	return location, nil
}
