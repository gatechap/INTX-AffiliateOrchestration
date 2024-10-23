package locallogging

import (
	"encoding/json"
	"fmt"
	"time"

	"th.truecorp.it.dsm.intcom/affiliateorchestration/config"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/httphandler"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/intutilities"
)

type LocalLogging struct {
	uuid             string
	username         string
	legacyUsername   string
	inputParam       string
	channel          string
	correlateId      string
	ip               string
	tagsApp          string
	tagsEnv          string
	gatewayType      string
	errorCode        string
	message          string
	system           string
	resultStatus     string
	mainInputKey     string
	mainInputValue   string
	outputString     string
	startTime        time.Time
	endTime          time.Time
	stackTrace       string
	errorApplication string
	errorModule      string
	errorFile        string
	errorFunction    string
}

func (logging *LocalLogging) SetRequestInputLogger(httpHeader *httphandler.HttpHeaderInfo, httpReq *httphandler.HttpRequestParamInfo, appConfig *config.Config, bReq []byte, tStartTime time.Time, mainInputKey string, mainInputValue string) {
	logging.uuid = httpReq.Uuid
	logging.username = httpReq.Username
	logging.correlateId = httpReq.CorrelatedId
	logging.tagsApp = httpReq.RestPath
	logging.inputParam = string(bReq)
	logging.channel = httpHeader.XChannel
	logging.ip = httpHeader.ClientIP
	logging.gatewayType = httpHeader.XGatewayType
	logging.legacyUsername = httpHeader.XLegacyUsername
	logging.startTime = tStartTime
	logging.tagsEnv = appConfig.Application.Profile
	logging.mainInputKey = mainInputKey
	logging.mainInputValue = mainInputValue
}

func (logging *LocalLogging) SetResponseInputLogger(errCode string, message string, resultStatus string, system string, strOutput string, endTime time.Time) {
	logging.errorCode = errCode
	logging.message = message
	logging.resultStatus = resultStatus
	logging.system = system
	logging.outputString = strOutput
	logging.endTime = endTime
}

func (logging *LocalLogging) writeLogResponseTag() {
	logRes := ResponseInputLogger{
		CorrelateId:    logging.correlateId,
		Uuid:           logging.uuid,
		DateCompletion: intutilities.GetCurrentISO8601(),
		LogType:        LOGTYPE_MESSAGE,
		Timestamp:      intutilities.GetCurrentISO8601(),

		// Response Input Logger
		Username:     logging.username,
		ErrorCode:    logging.errorCode,
		Message:      logging.message,
		ResultStatus: logging.resultStatus,
		OutputString: logging.outputString,
	}

	// Set tags
	logRes.Tags[0] = logging.tagsEnv
	logRes.Tags[1] = logging.tagsApp
	logRes.Tags[2] = TAG_RESPONSE

	jsRes, err := json.Marshal(logRes)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsRes))
}

func (logging *LocalLogging) writeLogElapsedTag() {
	logElapsed := ElapsedLogger{
		CorrelateId:    logging.correlateId,
		Uuid:           logging.uuid,
		DateCompletion: intutilities.GetCurrentISO8601(),
		LogType:        LOGTYPE_MESSAGE,
		Timestamp:      intutilities.GetCurrentISO8601(),

		// Response Input Logger
		Username:       logging.username,
		ErrorCode:      logging.errorCode,
		Message:        logging.message,
		System:         logging.system,
		MainInputKey:   logging.mainInputKey,
		MainInputValue: logging.mainInputValue,
	}

	// Set tags
	logElapsed.Tags[0] = logging.tagsEnv
	logElapsed.Tags[1] = logging.tagsApp
	logElapsed.Tags[2] = TAG_ELAPSE_TIME

	// calulate elapsed time to be ms
	logElapsed.ElapsedTime = calculateElapsedTime(&logging.startTime, &logging.endTime)

	jsElapsed, err := json.Marshal(logElapsed)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsElapsed))
}

func (logging *LocalLogging) BuildErrorInputLoggerBeforeController(profile string, currentApplication string, errCode string, errMsg string, err error, application string, module string, file string, function string) *ErrorInputLogger {
	logErr := ErrorInputLogger{

		// Base Input Logger
		TagsEnv: profile,
		TagsApp: currentApplication,

		// Response Input Logger
		ErrorCode:        errCode,
		Message:          errMsg,
		ErrorApplication: application,
		ErrorModule:      module,
		ErrorFile:        file,
		ErrorFunction:    function,

		// Convert error to string stacktrace
		StackTrace: Wrap(err).StackTrace,
	}
	return &logErr
}

func (logging *LocalLogging) SetErrorInputLoggerBeforeController(errLogger *ErrorInputLogger) {
	// Base Input Logger
	logging.tagsEnv = errLogger.TagsEnv
	logging.tagsApp = errLogger.TagsApp

	// Response Input Logger
	logging.errorCode = errLogger.ErrorCode
	logging.message = errLogger.Message
	logging.errorApplication = errLogger.ErrorApplication
	logging.errorModule = errLogger.ErrorModule
	logging.errorFile = errLogger.ErrorFile
	logging.errorFunction = errLogger.ErrorFunction
}

func (logging *LocalLogging) SetErrorInputLogger(errCode string, msg string, err error, application string, module string, file string, function string) {

	// Response Input Logger
	logging.errorCode = errCode
	logging.message = msg
	logging.errorApplication = application
	logging.errorModule = module
	logging.errorFile = file
	logging.errorFunction = function

	// Convert error to string stacktrace
	logging.stackTrace = Wrap(err).StackTrace
}

func (logging *LocalLogging) WriteLogRequest() {
	logReq := RequestInputLogger{
		CorrelateId:    logging.correlateId,
		Uuid:           logging.uuid,
		DateCompletion: intutilities.GetCurrentISO8601(),
		LogType:        LOGTYPE_MESSAGE,
		Timestamp:      intutilities.GetCurrentISO8601(),
		Username:       logging.username,
		InputParamLog:  logging.inputParam,
		Ip:             logging.ip,
		Channel:        logging.channel,
		GatewayType:    logging.gatewayType,
		LegacyUsername: logging.legacyUsername,
	}

	// Set tags
	logReq.Tags[0] = logging.tagsEnv
	logReq.Tags[1] = logging.tagsApp
	logReq.Tags[2] = TAG_REQUEST

	jsReq, err := json.Marshal(logReq)

	if err != nil {
		// executes a function asynchronously
		go fmt.Println(err)
		return
	}

	// executes a function asynchronously
	go fmt.Println(string(jsReq))
}

func (logging *LocalLogging) WriteLogResponse() {

	// executes a function asynchronously
	go logging.writeLogResponseTag()

	// executes a function asynchronously
	go logging.writeLogElapsedTag()
}

func (logging *LocalLogging) WriteLogError() {
	logError := ErrorInputLogger{
		CorrelateId:    logging.correlateId,
		Uuid:           logging.uuid,
		DateCompletion: intutilities.GetCurrentISO8601(),
		LogType:        LOGTYPE_ERROR,
		Timestamp:      intutilities.GetCurrentISO8601(),

		// Response Input Logger
		ErrorCode:        logging.errorCode,
		Message:          logging.message,
		StackTrace:       logging.stackTrace,
		ErrorApplication: logging.errorApplication,
		ErrorModule:      logging.errorModule,
		ErrorFile:        logging.errorFile,
		ErrorFunction:    logging.errorFunction,
		ResultStatus:     "F",
	}

	// Set tags
	logError.Tags[0] = logging.tagsEnv
	logError.Tags[1] = logging.tagsApp
	logError.Tags[2] = TAG_ERROR

	jsErr, err := json.Marshal(logError)

	if err != nil {

		// executes a function asynchronously
		go fmt.Println(err)
		return
	}

	fmt.Println(string(jsErr))
}

func (logging *LocalLogging) WriteLogErrorBeforeController(errLogger *ErrorInputLogger) {

	errLogger.CorrelateId = logging.correlateId
	errLogger.Uuid = logging.uuid
	errLogger.DateCompletion = intutilities.GetCurrentISO8601()
	errLogger.Timestamp = intutilities.GetCurrentISO8601()
	errLogger.LogType = LOGTYPE_ERROR
	errLogger.ResultStatus = "F"

	// Set tags
	errLogger.Tags[0] = logging.tagsEnv
	errLogger.Tags[1] = logging.tagsApp
	errLogger.Tags[2] = TAG_ERROR

	// Reset tagapp & tagenv
	errLogger.TagsEnv = ""
	errLogger.TagsApp = ""

	jsErr, err := json.Marshal(errLogger)

	if err != nil {
		// executes a function asynchronously
		go fmt.Println(err)
		return
	}

	// executes a function asynchronously
	go fmt.Println(string(jsErr))
}

func (logging *LocalLoggingLegacy) SetLegacyInputLoggerStart(correlateId string, uuid string, tagsEnv string, tagsApp string, tStartTime time.Time) {
	logging.correlateId = correlateId
	logging.uuid = uuid
	logging.timestamp = intutilities.GetCurrentISO8601()
	logging.tags = []string{tagsEnv, tagsApp, TAG_LEGACY}
	logging.startRequest = intutilities.GetCurrentISO8601()
	logging.startTime = tStartTime

}
func (logging *LocalLoggingLegacy) SetLegacyInputLoggerEnd(request string, response string, httpResponseCode string, httpResponseMessage string, targetEp string, tEndTime time.Time) {
	logging.request = request
	logging.response = response
	logging.httpResponseCode = httpResponseCode
	logging.httpResponseMessage = httpResponseMessage
	logging.targetEp = targetEp
	logging.endRequest = intutilities.GetCurrentISO8601()
	logging.endTime = tEndTime
}

func (logging *LocalLoggingLegacy) WriteLogLegacy() {

	start := logging.startTime.UnixNano() / int64(time.Millisecond)
	end := logging.endTime.UnixNano() / int64(time.Millisecond)
	diff := end - start

	logRes := LegacyInputLogger{
		CorrelateId:    logging.correlateId,
		Uuid:           logging.uuid,
		DateCompletion: intutilities.GetCurrentISO8601(),
		LogType:        LOGTYPE_APPLICATION,
		Timestamp:      intutilities.GetCurrentISO8601(),

		// Legacy Input Logger
		Request:             logging.request,
		Response:            logging.response,
		HttpResponseCode:    logging.httpResponseCode,
		HttpResponseMessage: logging.httpResponseMessage,
		StartRequest:        logging.startRequest,
		EndRequest:          logging.endRequest,
		TargetEp:            logging.targetEp,
		ElapsedTime:         diff,
	}

	// Set tags
	logRes.Tags[0] = logging.tags[0]
	logRes.Tags[1] = logging.tags[1]
	logRes.Tags[2] = logging.tags[2]

	jsRes, err := json.Marshal(logRes)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsRes))

}

type LocalLoggingLegacy struct {
	uuid        string
	correlateId string
	timestamp   string
	tags        []string
	// dateCompletion      string
	// logType             int
	httpResponseCode    string
	httpResponseMessage string
	request             string
	response            string
	targetEp            string
	// elapsedTime         string
	startRequest string
	endRequest   string
	startTime    time.Time
	endTime      time.Time
}
