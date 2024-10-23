package errormsg

import (
	"strconv"
	"strings"
)

func GenRequiredFieldMsg(pname string) string {
	return strings.Replace(ERR_MSG_REQUIRED_FIELD, "{0}", pname, 1)
}

func GenParamInvalidMsg(pname string, validvalues string) string {
	message := strings.Replace(ERR_MSG_PARAM_INVALID, "{0}", pname, 1)
	message = strings.Replace(message, "{1}", validvalues, 1)
	return message
}

func GenCanNotGreaterThan(pname string, validvalue int) string {
	message := strings.Replace(ERR_MSG_CAN_NOT_GREATER_THAN, "{0}", pname, 1)
	message = strings.Replace(message, "{1}", strconv.Itoa(validvalue), 1)
	return message
}

func GenInternalFailureMsg(message string) string {
	return strings.Replace(ERR_MSG_INTERNAL_FAILURE, "{0}", message, 1)
}

func GenBackendErrorMessage(system, apiname string) (message string) {
	message = strings.Replace(ERR_MSG_BACKEND_ERROR, "{0}", system, 1)
	message = strings.Replace(message, "{1}", apiname, 1)
	return
}
