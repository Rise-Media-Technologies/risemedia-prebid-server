package openrtb_ext

type ExtImpRiseMediaTech struct {
	BidFloor float64 `json:"bidfloor,omitempty"`
	TestMode int     `json:"testMode,omitempty"`
	SspId    string  `json:"sspId,omitempty"`
	SiteId   string  `json:"siteId,omitempty"`
}
