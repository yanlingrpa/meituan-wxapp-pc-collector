package main

// ProductSearchResultDto (method)
type Github_com__yanlingrpa__meituan_wxapp_pc_collector__ProductSearchResultDto struct {
	Keyword		string			`json:"keyword"`
	ProductInfos	[]ProductInfoDto	`json:"product_infos"`
}

// SearchProductDto (method)
type Github_com__yanlingrpa__meituan_wxapp_pc_collector__SearchProductDto struct {
	Keyword		string		`json:"keyword"`
	FetchCount	int		`json:"fetch_count"`
	Tags		[]string	`json:"tags"`
}

