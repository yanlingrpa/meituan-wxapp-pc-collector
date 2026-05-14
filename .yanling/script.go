package main

import (
	"fmt"
	"github.com/yanlingrpa/wxapp-pc-toolkits/wxapputils"
	"yanlingrpa.com/yanling/protocol/script"
)

// File: meituan\collect_medicines.go
type SearchProductDto struct {
	Keyword		string		`json:"keyword"`
	FetchCount	int		`json:"fetch_count"`
	Tags		[]string	`json:"tags"`
}

type ProductInfoDto struct {
	Brand		string	`json:"brand"`
	Name		string	`json:"name"`
	Standard	string	`json:"standard"`
	PackQty		int32	`json:"pack_qty"`
	TotalCent	int64	`json:"total_cent"`
	UnitCent	int64	`json:"unit_cent"`
	ShopName	string	`json:"shop_name"`
}

type ProductSearchResultDto struct {
	Keyword		string			`json:"keyword"`
	ProductInfos	[]ProductInfoDto	`json:"product_infos"`
}

func CollectMedicine(rt script.ModuleRuntime, dto SearchProductDto) (*ProductSearchResultDto, error) {
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
	results := []ProductInfoDto{}
	one := ProductInfoDto{
		Brand:		"阿斯利康",
		Name:		"盐酸达泊西汀片",
		Standard:	"10mg*3片/盒",
		PackQty:	1,
		TotalCent:	100,
		UnitCent:	100,
		ShopName:	"海星医药旗舰店",
	}
	results = append(results, one)
	two := ProductInfoDto{
		Brand:		"辉瑞",
		Name:		"莫沙拉唑片",
		Standard:	"7.5mg*3片/盒",
		PackQty:	1,
		TotalCent:	100,
		UnitCent:	100,
		ShopName:	"海星医药旗舰店",
	}
	results = append(results, two)
	three := ProductInfoDto{
		Brand:		"美赞臣",
		Name:		"奥美拉唑片",
		Standard:	"10mg*3片/盒",
		PackQty:	1,
		TotalCent:	100,
		UnitCent:	100,
		ShopName:	"海星医药旗舰店",
	}
	results = append(results, three)

	psr := ProductSearchResultDto{
		Keyword:	dto.Keyword,
		ProductInfos:	results,
	}

	rt.Publish("product_infos", &psr)

	return &psr, nil

}

