package yscript

type Github_com_yanlingrpa_meituan_wxapp_pc_collector_SearchProductDto struct {
	Keyword    string   `json:"keyword"`
	FetchCount int      `json:"fetch_count"`
	Tags       []string `json:"tags"`
}

type Github_com_yanlingrpa_meituan_wxapp_pc_collector_ProductInfoDto struct {
	Brand      string `json:"brand"`
	Name       string `json:"name"`
	Standard   string `json:"standard"`
	Count      int32  `json:"count"`
	TotalPrice string `json:"total_price"`
	UnitPrice  string `json:"unit_price"`
	ShopName   string `json:"shop_name"`
}
