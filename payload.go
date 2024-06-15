package flaresolverr

type Payload struct {
	Cmd         string  `json:"cmd"`
	URL         string  `json:"url"`
	Cookies     Cookies `json:"cookies,omitempty"`
	OnlyCookies bool    `json:"returnOnlyCookies,omitempty"`
	MaxTimeout  int     `json:"maxTimeout"`
	PostData    string  `json:"postData,omitempty"`
	Session     string  `json:"session,omitempty"`
}
