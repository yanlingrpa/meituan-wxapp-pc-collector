package meituan

import (
	"fmt"

	"github.com/yanlingrpa/wxapp-pc-toolkits/wxapputils"
	"yanlingrpa.com/yanling/protocol/script"
)

func CollectMedicine(rt script.ModuleRuntime, keyword string) ([]string, error) {
	ready, err := wxapputils.CheckWxappReady(rt)
	if err != nil {
		return nil, err
	}
	if !ready {
		return nil, fmt.Errorf("wxapp is not ready")
	}

	return nil, nil

}
