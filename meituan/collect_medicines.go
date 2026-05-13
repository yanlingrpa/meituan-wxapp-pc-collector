package meituan

import (
	"fmt"

	"github.com/yanlingrpa/wxapp-pc-toolkits/wxapputils"
	"yanlingrpa.com/yanling/protocol/script"
)

type SearchProductDto struct {
	Keyword    string   `json:"keyword"`
	FetchCount int      `json:"fetch_count"`
	Tags       []string `json:"tags"`
}

type ProductInfoDto struct {
	Brand      string `json:"brand"`
	Name       string `json:"name"`
	Standard   string `json:"standard"`
	Count      int32  `json:"count"`
	TotalPrice string `json:"total_price"`
	UnitPrice  string `json:"unit_price"`
	ShopName   string `json:"shop_name"`
}

func CollectMedicine(rt script.ModuleRuntime, dto SearchProductDto) ([]ProductInfoDto, error) {
	val, ok := rt.GetVariable("wxapp-meituan")
	if !ok {
		return nil, fmt.Errorf("wxapp-meituan is not ready")
	}
	guiId := val.(string)
	ready, err := wxapputils.CheckWxappReady(rt, guiId)
	if err != nil {
		return nil, err
	}
	if !ready {
		return nil, fmt.Errorf("wxapp is not ready")
	}

	return nil, nil

}
