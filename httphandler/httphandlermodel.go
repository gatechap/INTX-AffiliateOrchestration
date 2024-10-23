package httphandler

type HttpHeaderInfo struct {
	Authorization   string
	XUsername       string
	XLegacyUsername string
	XChannel        string
	ClientIP        string
	ClientHostName  string
	XGatewayType    string
}

type HttpRequestParamInfo struct {
	Uuid         string
	RestPath     string
	CorrelatedId string
	Username     string
}
