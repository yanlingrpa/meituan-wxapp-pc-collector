package yscript

import (
	"fmt"

	"yanlingrpa.com/yanling/protocol/script"
)

func CollectMedicine(rt script.ModuleRuntime, dto Github_com_yanlingrpa_meituan_wxapp_pc_collector_SearchProductDto) ([]string, error) {
	_rt_resp, err := rt.InvokeWorker("github.com/yanlingrpa/wxapp-pc-toolkits@latest", "CheckWxappReady", dto.Keyword)
	var ready bool
	if err == nil {
		ready = _rt_resp.(bool)
	}
	if err != nil {
		return nil, err
	}
	if !ready {
		return nil, fmt.Errorf("wxapp is not ready")
	}
	return nil, nil

}
