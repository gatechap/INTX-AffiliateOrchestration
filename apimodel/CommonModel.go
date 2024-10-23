package apimodel

type BackendResponseInfoArray struct {
	APIName   string `json:"apiName,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
	Message   string `json:"message,omitempty"`
	System    string `json:"system,omitempty"`
	URL       string `json:"url,omitempty"`
}

type BackendResponseList struct {
	Size                     int32                       `json:"size,omitempty"`
	BackendResponseInfoArray *[]BackendResponseInfoArray `json:"backendResponseInfoArray,omitempty"`
}
