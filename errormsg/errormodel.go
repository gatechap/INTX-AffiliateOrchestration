package errormsg

type ErrorHandlerInfo struct {
	Error            error
	ErrorApplication string
	ErrorModule      string
	ErrorFile        string
	ErrorFunction    string
}
