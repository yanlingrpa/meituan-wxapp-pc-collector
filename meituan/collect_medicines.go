package meituan

import (
	"fmt"

	"github.com/yanlingrpa/wxapp-pc-toolkits/wxapputils"
	"yanlingrpa.com/yanling/protocol/script"
)

func CollectMedicine(rt script.ModuleRuntime, dto SearchProductDto) (*ProductSearchResultDto, error) {
	_, err := getGuiId(rt)
	if err != nil {
		return nil, err
	}

	pageInfo, err := wxapputils.GetPageInfo(rt)
	if err != nil {
		return nil, err
	}

	if !pageInfo.Searchable {
		return nil, fmt.Errorf("current page is not searchable")
	}

	results := []ProductInfoDto{}
	one := ProductInfoDto{
		Brand:     "阿斯利康",
		Name:      "盐酸达泊西汀片",
		Standard:  "10mg*3片/盒",
		PackQty:   1,
		TotalCent: 100,
		UnitCent:  100,
		ShopName:  "海星医药旗舰店",
	}
	results = append(results, one)
	two := ProductInfoDto{
		Brand:     "辉瑞",
		Name:      "莫沙拉唑片",
		Standard:  "7.5mg*3片/盒",
		PackQty:   1,
		TotalCent: 100,
		UnitCent:  100,
		ShopName:  "海星医药旗舰店",
	}
	results = append(results, two)
	three := ProductInfoDto{
		Brand:     "美赞臣",
		Name:      "奥美拉唑片",
		Standard:  "10mg*3片/盒",
		PackQty:   1,
		TotalCent: 100,
		UnitCent:  100,
		ShopName:  "海星医药旗舰店",
	}
	results = append(results, three)

	psr := ProductSearchResultDto{
		Keyword:      dto.Keyword,
		ProductInfos: results,
	}
	// 发送搜索到的结果
	rt.Publish("product_infos", &psr)

	return &psr, nil

}
