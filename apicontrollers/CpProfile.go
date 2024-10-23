package apicontrollers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"th.truecorp.it.dsm.intcom/affiliateorchestration/apimodel"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/errormsg"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/httphandler"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/locallogging"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// affiliateorchestration go doc
// @Summary affiliateorchestration
// @Tags cpprofile-controller
// @Accept  json
// @Produce  json
// @param X-Channel header string false "X-Channel"
// @param X-GatewayType header string false "X-GatewayType"
// @param X-LegacyUsername header string false "X-LegacyUsername"
// @param X-Username header string true "X-Username"
// @Param filter body apimodel.RequestCpProfile true "requestInfo"
// @Success 200 {array} apimodel.ResponseCpProfile
// @Failure 400,404 {object} apimodel.ResponseCpProfile
// @Failure 500 {object} apimodel.ResponseCpProfile
// @Failure default {object} apimodel.ResponseCpProfile
// @Router /cpemployee/cpprofile/primresource [post]
func CpProfileByPrimResource(c *gin.Context) {
	const CURRENT_FUNCTION = "CpProfileByPrimResource"

	startTime := time.Now()

	var isError bool
	var errCode string
	var errMsg string
	var resultStatus string
	// var result string
	var isDNF bool
	var response apimodel.ResponseCpProfile
	var httpResponse int
	// var srv service.GetproductorchestratedService

	// initial error hanler info
	errHanlerInfo := errormsg.ErrorHandlerInfo{
		ErrorApplication: AppConfig.Application.Name,
		ErrorModule:      CURRENT_MODULE,
		ErrorFile:        CURRENT_FUNCTION,
		ErrorFunction:    CURRENT_FUNCTION,
	}

	// generate uuid
	struuid := uuid.New().String()

	// Validate request (bad request)
	var request apimodel.RequestCpProfile
	var bReq []byte
	var err error

	if err = c.ShouldBindJSON(&request); err != nil {
		// check bad request
		isError = true
		errCode = errormsg.ERR_CD_PARAM_INVALID
		errMsg = err.Error()
		errHanlerInfo.Error = err
	} else {
		bReq, err = json.Marshal(&request)
		if err != nil {
			isError = true
			errCode = errormsg.ERR_CD_INTERNAL_FAILURE
			errMsg = errormsg.GenInternalFailureMsg(err.Error())
			errHanlerInfo.Error = err
		}
	}

	if err = validator.New().Struct(&request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				isError = true
				errCode = errormsg.ERR_CD_REQUIRED_FIELD
				errMsg = errormsg.GenRequiredFieldMsg(err.StructField())
				errHanlerInfo.Error = err
				break
			} else {
				isError = true
				errCode = errormsg.ERR_CD_PARAM_INVALID
				errMsg = errormsg.GenParamInvalidMsg(err.StructField(), strings.ReplaceAll(err.Tag(), "eq=", ""))
				errHanlerInfo.Error = err
				break
			}
		}
	}

	// httpheader
	httpHeader, requestparam := httphandler.GetHttpHeaderInfo(c, &request.CorrelatedId, &struuid)
	username := requestparam.Username
	if username == "" {
		isError = true
		errCode = errormsg.ERR_CD_REQUIRED_FIELD
		errMsg = errormsg.GenRequiredFieldMsg("Username")
		errHanlerInfo.Error = err
	}

	// Assign mainInputKey and value
	mainInputKey := "primResourceValue"
	mainInputValue := request.PrimResourceValue

	// Write log request
	log := locallogging.LocalLogging{}
	log.SetRequestInputLogger(httpHeader, requestparam, AppConfig, bReq, startTime, mainInputKey, mainInputValue)
	log.WriteLogRequest()

	if !isError {
		// srv.RedisClient, err = redisservice.GetRedisClient(AppConfig)

		response, err = service.CpProfileService(request, struuid, "/cpprofile/primresource", requestparam.Username, log)

		if err != nil {
			isError = true
			errCode = errormsg.ERR_CD_INTERNAL_FAILURE
			errMsg = errormsg.GenInternalFailureMsg(err.Error())
			errHanlerInfo.Error = err
		}
	}

	if !isError {

		// result, isDNF, err = srv.GetByKey(request.Key)

		if err != nil {
			isError = true
			errCode = errormsg.ERR_CD_INTERNAL_FAILURE
			errMsg = errormsg.GenInternalFailureMsg(err.Error())
			errHanlerInfo.Error = err
		}
	}

	if isError {
		response = apimodel.ResponseCpProfile{
			Uuid:      struuid,
			ErrorCode: errCode,
			Message:   errMsg,
		}
		httpResponse = http.StatusOK
		resultStatus = "F"

		log.SetErrorInputLogger(errCode, errMsg, errHanlerInfo.Error, errHanlerInfo.ErrorApplication, errHanlerInfo.ErrorModule, errHanlerInfo.ErrorFile, errHanlerInfo.ErrorFunction)
		log.WriteLogError()
	} else {

		if isDNF {
			errCode = errormsg.ERR_CD_DNF
			errMsg = errormsg.ERR_MSG_DNF
			resultStatus = "NA"
		} else {
			errCode = errormsg.ERR_CD_SUCCESS
			errMsg = errormsg.ERR_MSG_SUCCESS
			resultStatus = "S"
		}

		// response = apimodel.ResponseCpProfileService{
		// 	Uuid:      struuid,
		// 	ErrorCode: errCode,
		// 	Message:   errMsg,
		// }

		httpResponse = http.StatusOK
	}

	log.SetResponseInputLogger(errCode, errMsg, resultStatus, "", "", time.Now())
	log.WriteLogResponse()
	c.JSON(httpResponse, response)
}
