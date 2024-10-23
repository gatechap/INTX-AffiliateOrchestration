package legacy

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"th.truecorp.it.dsm.intcom/affiliateorchestration/apimodel"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/config"
	"th.truecorp.it.dsm.intcom/affiliateorchestration/locallogging"
)

func GetCPProfileByThaiIDClient(uuid, correlateId, tagsApp, cert string, log locallogging.LocalLogging) (result ResponseGetCPProfileByThaiID, backend *apimodel.BackendResponseInfoArray) {

	var svcCpProfile = config.GetService("getCPProfileByThaiID")
	appConfig, _ := config.LoadConfig()

	var requestClient = RequestGetCPProfileByThaiID{
		Format: "json",
		Key:    svcCpProfile.ApiKey,
		Thaiid: cert,
	}

	// Write log Legacy
	logLegacy := locallogging.LocalLoggingLegacy{}
	logLegacy.SetLegacyInputLoggerStart(correlateId, uuid, appConfig.Application.Profile, tagsApp, time.Now())

	var statusCodeStr string = "500"
	var resLegacy string
	reqLegacy, _ := json.Marshal(requestClient)

	client := &http.Client{}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, _ := http.NewRequest(http.MethodPost, svcCpProfile.Endpoint, bytes.NewBuffer(reqLegacy))

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(svcCpProfile.User, svcCpProfile.Password)

	var resp *http.Response
	var errPrimary error

	resp, errPrimary = client.Do(req)

	if errPrimary != nil {

		fmt.Println(errPrimary.Error())

		backend = setBackend(svcCpProfile, "", errPrimary.Error())

		resLegacy = errPrimary.Error()

		defer writeLogLegacy(logLegacy, string(reqLegacy), resLegacy, statusCodeStr, "", svcCpProfile.Endpoint)

		if errPrimary != nil {
			defer panicLegacyHandle(log, errPrimary, appConfig.Application.Name, svcCpProfile.Name)
			fmt.Println(errPrimary.Error())
			panic("panicLegacy")
		}

	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// for result
	if er := json.Unmarshal(bodyBytes, &result); er != nil {
		fmt.Println(er.Error())
	}

	var responseAPI interface{}
	// for logs
	if er := json.Unmarshal(bodyBytes, &responseAPI); er != nil {
		fmt.Println(er.Error())
	}

	jsonResponse, _ := json.Marshal(responseAPI)
	resLegacy = string(jsonResponse)

	if resp.StatusCode != 200 {

		errPrimary = errors.New(resLegacy)

		backend = setBackend(svcCpProfile, strconv.Itoa(result.Code), errPrimary.Error())

	} else if resp.StatusCode == 200 {

		backend = setBackend(svcCpProfile, strconv.Itoa(result.Code), result.Description)

	}

	statusCodeStr = strconv.Itoa(resp.StatusCode)

	defer writeLogLegacy(logLegacy, string(reqLegacy), resLegacy, statusCodeStr, resp.Status, svcCpProfile.Endpoint)

	if errPrimary != nil {
		defer panicLegacyHandle(log, errPrimary, appConfig.Application.Name, svcCpProfile.Name)
		fmt.Println(errPrimary.Error())
		panic("panicLegacy")
	}

	defer resp.Body.Close()

	return
}

type RequestGetCPProfileByThaiID struct {
	Key    string `json:"key"`
	Format string `json:"format"`
	Thaiid string `json:"thaiid"`
}

type ResponseGetCPProfileByThaiID struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Data        Data   `json:"data"`
}
type Data struct {
	Payroll        string `json:"PAYROLL"`
	Companygroup   string `json:"COMPANYGROUP"`
	Companyname    string `json:"COMPANYNAME"`
	Empid          string `json:"EMPID"`
	Identification string `json:"IDENTIFICATION"`
	BDd            string `json:"B_DD"`
	BMm            string `json:"B_MM"`
	BYyyy          string `json:"B_YYYY"`
	Birthdate      string `json:"BIRTHDATE"`
	Salutation     string `json:"SALUTATION"`
	ThaiFirstname  string `json:"THAI_FIRSTNAME"`
	ThaiLastname   string `json:"THAI_LASTNAME"`
	ThaiFullname   string `json:"THAI_FULLNAME"`
	Salutatione    string `json:"SALUTATIONE"`
	EngFirstname   string `json:"ENG_FIRSTNAME"`
	EngLastname    string `json:"ENG_LASTNAME"`
	Sal            string `json:"SAL"`
	Email          string `json:"EMAIL"`
	Pro            string `json:"PRO"`
	Eligible       string `json:"ELIGIBLE"`
	Positioncode   string `json:"POSITIONCODE"`
	Position       string `json:"POSITION"`
}
