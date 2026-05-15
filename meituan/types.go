package meituan

type SearchProductDto struct {
	Keyword    string   `json:"keyword"`
	FetchCount int      `json:"fetch_count"`
	Tags       []string `json:"tags"`
}

type ProductInfoDto struct {
	Brand     string `json:"brand"`
	Name      string `json:"name"`
	Standard  string `json:"standard"`
	PackQty   int32  `json:"pack_qty"`
	TotalCent int64  `json:"total_cent"`
	UnitCent  int64  `json:"unit_cent"`
	ShopName  string `json:"shop_name"`
}

type ProductSearchResultDto struct {
	Keyword      string           `json:"keyword"`
	ProductInfos []ProductInfoDto `json:"product_infos"`
}
