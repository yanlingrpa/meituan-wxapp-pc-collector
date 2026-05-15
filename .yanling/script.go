package main

import (
	"fmt"
	"github.com/yanlingrpa/wxapp-pc-toolkits/wxapputils"
	"yanlingrpa.com/yanling/protocol/script"
)

// File: meituan\basic.go
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
		GuiId:		guiId,
		Keyword:	location,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// File: meituan\collect_medicines.go
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

// File: meituan\topic_listener.go
func onAppReady(event script.Event) {
	data := event.Data.(wxapputils.AppReadyData)
	fmt.Printf("app is ready, payload: %+v\n", data)
}

// File: meituan\types.go
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

// File: meituan\utils.go
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

