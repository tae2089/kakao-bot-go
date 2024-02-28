package kakao

type Link struct {
	WebURL       string `json:"web,omitempty"`
	MobileWebURL string `json:"mobile,omitempty"`
	PcURL        string `json:"pc,omitempty"`
}

func NewLink(webUrl string, mobileWebUrl string, pcUrl string) *Link {
	return &Link{
		WebURL:       webUrl,
		MobileWebURL: mobileWebUrl,
		PcURL:        pcUrl,
	}
}
