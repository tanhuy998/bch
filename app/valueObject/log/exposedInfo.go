package log

type (
	RequestExposedInfo struct {
		RequestType    string `json:"request_type"`
		Protocol       string `json:"protocol"`
		IsSecure       bool   `json:"secure"`
		Verb           string `json:"verb"`
		Path           string `json:"path"`
		ResponseStatus int    `json:"response_status"`
		SourceIP       string `json:"source_ip"`
		UserAgent      string `json:"user_agent"`
	}
)
