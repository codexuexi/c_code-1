package c_code

type FindWailianInfo struct {
	Domain       string            `json:"domain"`
	WaiLianCount int64             `json:"wai_lian_count"`
	Walian       []FindDomaWaiLian `json:"walian"`
	Yu           map[string]int    `json:"yu"`
}

type FindDomaWaiLian struct {
	Ci         string                `json:"ci"`
	DomainList []FindDomaWaiLianInfo `json:"domain_list"`
	Count      int                   `json:"-"`
}
type FindDomaWaiLianInfo struct {
	Url string `json:"url"`
	C   string `json:"c"`
	U   string `json:"u"`
	R   string `json:"r"`
}
