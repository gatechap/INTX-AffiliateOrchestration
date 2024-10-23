package service

import (
	"th.truecorp.it.dsm.intcom/affiliateorchestration/apimodel"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/errormsg"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/legacy"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/locallogging"
)

func handleCpProfileService(response apimodel.ResponseCpProfile, uuid, errorCode, errorMessage string) (result apimodel.ResponseCpProfile) {
	response.BackendResponseList.Size = int32(len(*response.BackendResponseList.BackendResponseInfoArray))
	result = apimodel.ResponseCpProfile{
		Uuid:                uuid,
		ErrorCode:           errorCode,
		Message:             errorMessage,
		BackendResponseList: response.BackendResponseList,
	}
	return
}

func CpProfileService(request apimodel.RequestCpProfile, uuid, tagsApp, user string, log locallogging.LocalLogging) (response apimodel.ResponseCpProfile, er error) {

	response.BackendResponseList = &apimodel.BackendResponseList{}
	response.BackendResponseList.BackendResponseInfoArray = &[]apimodel.BackendResponseInfoArray{}
	errorCode := errormsg.ERR_CD_SUCCESS
	errorMessage := errormsg.ERR_MSG_SUCCESS

	searchProductPre := []legacy.SearchInfoArray{}
	searchProductPre = append(searchProductPre, legacy.SearchInfoArray{Type: "PRIMRESOURCEVAL", Value: request.PrimResourceValue})
	searchProductPre = append(searchProductPre, legacy.SearchInfoArray{Type: "BUSINESSLINE", Value: request.BusinessLine})
	searchProductPre = append(searchProductPre, legacy.SearchInfoArray{Type: "SUBSTATUS", Value: "ACTIVEORSUSPEND"})
	searchProductPre = append(searchProductPre, legacy.SearchInfoArray{Type: "REALTHAIID", Value: "N"})

	requestProductPre := legacy.RequestGetProductPreferenceList{
		CorrelatedID: uuid,
		SearchList:   legacy.SearchList{SearchInfoArray: searchProductPre},
		PageSize:     "1",
		PageNumber:   "1",
	}

	resultProductPre, backend := legacy.GetProductPreferenceListClient(uuid, request.CorrelatedId, tagsApp, user, log, requestProductPre)
	*response.BackendResponseList.BackendResponseInfoArray = append(*response.BackendResponseList.BackendResponseInfoArray, *backend)

	if resultProductPre.GetProductPreferenceListResponse.Return.ErrorCode == "OSBbllngA10001" {
		response = handleCpProfileService(response, uuid, errormsg.ERR_CD_DNF, errormsg.ERR_MSG_DNF)
		return
	} else if resultProductPre.GetProductPreferenceListResponse.Return.ErrorCode != "OSBbllngA00001" {
		response = handleCpProfileService(response, uuid, errormsg.ERR_CD_BACKEND_ERROR, errormsg.GenBackendErrorMessage(backend.System, backend.APIName))
		return
	} else if resultProductPre.GetProductPreferenceListResponse.Return.ErrorCode == "OSBbllngA00001" && len(resultProductPre.GetProductPreferenceListResponse.Return.ProductPreferenceList.ProductPreferenceInfoArray) < 1 {
		response = handleCpProfileService(response, uuid, errormsg.ERR_CD_DNF, errormsg.ERR_MSG_DNF)
		return
	}

	cert := resultProductPre.GetProductPreferenceListResponse.Return.ProductPreferenceList.ProductPreferenceInfoArray[0].Customer.CertificateNumber

	resultCP, backend := legacy.GetCPProfileByThaiIDClient(uuid, request.CorrelatedId, tagsApp, cert, log)
	*response.BackendResponseList.BackendResponseInfoArray = append(*response.BackendResponseList.BackendResponseInfoArray, *backend)

	if resultCP.Code == 600 {
		response = handleCpProfileService(response, uuid, errormsg.ERR_CD_DNF, errormsg.ERR_MSG_DNF)
		return
	} else if resultCP.Code != 200 {
		response = handleCpProfileService(response, uuid, errormsg.ERR_CD_BACKEND_ERROR, errormsg.GenBackendErrorMessage(backend.System, backend.APIName))
		return
	}

	var cpProfile *apimodel.CpProfile
	var engName *apimodel.NameInfo
	var thaiName *apimodel.NameInfo

	engName = &apimodel.NameInfo{
		Title:     resultCP.Data.Salutatione,
		FirstName: resultCP.Data.EngFirstname,
		LastName:  resultCP.Data.EngLastname,
	}

	thaiName = &apimodel.NameInfo{
		Title:     resultCP.Data.Salutation,
		FirstName: resultCP.Data.ThaiFirstname,
		LastName:  resultCP.Data.ThaiLastname,
	}

	cpProfile = &apimodel.CpProfile{
		ThaiName:          thaiName,
		EngName:           engName,
		Payroll:           resultCP.Data.Payroll,
		CompanyGroup:      resultCP.Data.Companygroup,
		CompanyName:       resultCP.Data.Companyname,
		EmployeeID:        resultCP.Data.Empid,
		CertificateNumber: resultCP.Data.Identification,
		BirthDate:         resultCP.Data.Birthdate,
		Email:             resultCP.Data.Email,
	}

	response.BackendResponseList.Size = int32(len(*response.BackendResponseList.BackendResponseInfoArray))
	response = apimodel.ResponseCpProfile{
		Uuid:                uuid,
		ErrorCode:           errorCode,
		Message:             errorMessage,
		BackendResponseList: response.BackendResponseList,
		CpProfile:           cpProfile,
	}

	return
}
