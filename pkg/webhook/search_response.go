package webhook

type SearchResponse struct {
	Paging  PagingResponse `json:"pagination"`
	Results []Response     `json:"results"`
	Links   LinksResponse  `json:"_links"`
}

type PagingResponse struct {
	Page         int `json:"page"`
	PageSize     int `json:"page_size"`
	TotalResults int `json:"total_results"`
}

type LinksResponse struct {
	Previous Previous `json:"previous"`
	Self     Self     `json:"self"`
	Next     Next     `json:"next"`
}

type Previous struct {
	Href string `json:"href"`
}

type Self struct {
	Href string `json:"href"`
}
type Next struct {
	Href string `json:"href"`
}
