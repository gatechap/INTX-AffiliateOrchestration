package legacy

import (
	"fmt"
	"time"

	"th.truecorp.it.dsm.intcom/affiliateorchestration/apimodel"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/config"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/errormsg"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/locallogging"
)

const hiddenURL = "[hidden]"

func panicLegacyHandle(log locallogging.LocalLogging, err error, appName, legacyName string) {

	recover := recover()
	if recover == "panicLegacy" {
		fmt.Println("panicLegacy")
	}
	log.SetErrorInputLogger(errormsg.ERR_CD_INTERNAL_FAILURE, err.Error(), err, appName,
		"Legacy", legacyName, legacyName)
	log.WriteLogError()
}

func writeLogLegacy(logLegacy locallogging.LocalLoggingLegacy, reqLegacy, resLegacy, statusCodeStr, status, endPoint string) {
	logLegacy.SetLegacyInputLoggerEnd(reqLegacy, resLegacy, statusCodeStr, status, endPoint, time.Now())
	logLegacy.WriteLogLegacy()
}

// func panicLegacy(log locallogging.LocalLogging, errPrimary error, appName, serviceName string) {
// 	defer panicLegacyHandle(log, errPrimary, appName, serviceName)
// 	fmt.Println(errPrimary.Error())
// 	panic("panicLegacy")
// }

func setBackend(service config.Service, errorCode, errorMessage string) (backend *apimodel.BackendResponseInfoArray) {

	if errorCode == "" {
		errorCode = errormsg.ERR_CD_BACKEND_ERROR
	}

	if errorMessage == "" {
		errorMessage = errormsg.GenBackendErrorMessage(service.System, service.Name)
	}

	backend = &apimodel.BackendResponseInfoArray{
		APIName:   service.Name,
		ErrorCode: errorCode,
		Message:   errorMessage,
		System:    service.System,
		// URL:       service.Endpoint,
		URL: hiddenURL,
	}

	return
}

type SearchInfoArray struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type SearchList struct {
	SearchInfoArray []SearchInfoArray `json:"searchInfoArray"`
}

type CodeDescriptionInfo struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
