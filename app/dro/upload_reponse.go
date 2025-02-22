package dro

type UploadResponseData struct {
	Error error  `json:"error,omitempty"`
	Msg   string `json:"msg"`
	Code  uint   `json:"code,omitempty"`
	Path  string `json:"path,o,omitempty"`
	Salt  string `json:"salt,omitempty"`
	Url   string `json:"url,omitempty"`
	Host  string `json:"host,omitempty"`
	Key   int    `json:"key,omitempty"`
}
