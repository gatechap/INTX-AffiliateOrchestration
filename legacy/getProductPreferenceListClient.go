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

func GetProductPreferenceListClient(uuid, correlateId, tagsApp, user string, log locallogging.LocalLogging, requestClient RequestGetProductPreferenceList) (result ResponseGetProductPreferenceList, backend *apimodel.BackendResponseInfoArray) {

	var svcProductPre = config.GetService("getProductPreferenceList")
	appConfig, _ := config.LoadConfig()

	// Write log Legacy
	logLegacy := locallogging.LocalLoggingLegacy{}
	logLegacy.SetLegacyInputLoggerStart(correlateId, uuid, appConfig.Application.Profile, tagsApp, time.Now())

	var statusCodeStr string = "500"
	var resLegacy string
	reqLegacy, _ := json.Marshal(requestClient)

	client := &http.Client{}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, _ := http.NewRequest(http.MethodPost, svcProductPre.Endpoint, bytes.NewBuffer(reqLegacy))

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(svcProductPre.User, svcProductPre.Password)
	req.Header.Set("X-Channel", user)

	var resp *http.Response
	var errPrimary error

	resp, errPrimary = client.Do(req)

	if errPrimary != nil {

		fmt.Println(errPrimary.Error())

		backend = setBackend(svcProductPre, "", errPrimary.Error())

		resLegacy = errPrimary.Error()

		defer writeLogLegacy(logLegacy, string(reqLegacy), resLegacy, statusCodeStr, "", svcProductPre.Endpoint)

		if errPrimary != nil {
			defer panicLegacyHandle(log, errPrimary, appConfig.Application.Name, svcProductPre.Name)
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

		backend = setBackend(svcProductPre, result.GetProductPreferenceListResponse.Return.ErrorCode, errPrimary.Error())

	} else if resp.StatusCode == 200 {

		backend = setBackend(svcProductPre, result.GetProductPreferenceListResponse.Return.ErrorCode, result.GetProductPreferenceListResponse.Return.Message)

	}

	statusCodeStr = strconv.Itoa(resp.StatusCode)

	defer writeLogLegacy(logLegacy, string(reqLegacy), resLegacy, statusCodeStr, resp.Status, svcProductPre.Endpoint)

	if errPrimary != nil {
		defer panicLegacyHandle(log, errPrimary, appConfig.Application.Name, svcProductPre.Name)
		fmt.Println(errPrimary.Error())
		panic("panicLegacy")
	}

	defer resp.Body.Close()

	return
}

type RequestGetProductPreferenceList struct {
	CorrelatedID string     `json:"correlatedId"`
	SearchList   SearchList `json:"searchList"`
	PageSize     string     `json:"pageSize"`
	PageNumber   string     `json:"pageNumber"`
}

type ResponseGetProductPreferenceList struct {
	GetProductPreferenceListResponse GetProductPreferenceListResponse `json:"getProductPreferenceListResponse"`
}

type RfMigration struct {
	AccountID      string `json:"accountId"`
	ActivationDate string `json:"activationDate"`
	ServiceLevel   string `json:"serviceLevel"`
}
type CodeDescInfo struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type PortIndicator struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type RecipientZone struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type DonorOperator struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type RecipientOperator struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type porting struct {
	DonorZone         CodeDescInfo `json:"donorZone"`
	PortIndicator     CodeDescInfo `json:"portIndicator"`
	RecipientZone     CodeDescInfo `json:"recipientZone"`
	DonorOperator     CodeDescInfo `json:"donorOperator"`
	RecipientOperator CodeDescInfo `json:"recipientOperator"`
}
type Address struct {
	ZipCode      string `json:"zipCode"`
	RoomNo       string `json:"roomNo"`
	City         string `json:"city"`
	AddressType  string `json:"addressType"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	AddressLine3 string `json:"addressLine3"`
	AddressLine4 string `json:"addressLine4"`
	Street       string `json:"street"`
	District     string `json:"district"`
	HouseNo      string `json:"houseNo"`
	Floor        string `json:"floor"`
	Moo          string `json:"moo"`
	Soi          string `json:"soi"`
	SubDistrict  string `json:"subDistrict"`
	Building     string `json:"building"`
	SinceDate    string `json:"sinceDate"`
	Country      string `json:"country"`
}

type CreditStatus struct {
	ReasonDescription string `json:"reasonDescription"`
	LastActivityDate  string `json:"lastActivityDate"`
	ReasonCode        string `json:"reasonCode"`
	Status            string `json:"status"`
}
type RmvMigration struct {
	AccountID      string `json:"accountId"`
	ActivationDate string `json:"activationDate"`
	ServiceLevel   string `json:"serviceLevel"`
}
type TmvMigration struct {
	AccountID      string `json:"accountId"`
	ActivationDate string `json:"activationDate"`
	ServiceLevel   string `json:"serviceLevel"`
}
type CollectionStatus struct {
	ReasonDescription string `json:"reasonDescription"`
	LastActivityDate  string `json:"lastActivityDate"`
	ReasonCode        string `json:"reasonCode"`
	Status            string `json:"status"`
}
type Contact struct {
	PreferredContactNo string `json:"preferredContactNo"`
	HomePhone          string `json:"homePhone"`
	Language           string `json:"language"`
	Email              string `json:"email"`
	Fax                string `json:"fax"`
	OfficePhone        string `json:"officePhone"`
	OfficePhoneExt     string `json:"officePhoneExt"`
	PrivatePhone       string `json:"privatePhone"`
}
type CreditLimitAtSubStatus struct {
	ReasonDescription string `json:"reasonDescription"`
	LastActivityDate  string `json:"lastActivityDate"`
	ReasonCode        string `json:"reasonCode"`
	Status            string `json:"status"`
}

type Dealer struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type Name struct {
	BranchCode       string `json:"branchCode"`
	BranchName       string `json:"branchName"`
	MiddleName       string `json:"middleName"`
	OrganizationName string `json:"organizationName"`
	StoreID          string `json:"storeId"`
	NameType         string `json:"nameType"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Title            string `json:"title"`
}
type Proof struct {
	DealerApp     string `json:"dealerApp"`
	ProofDate     string `json:"proofDate"`
	ProofDocument string `json:"proofDocument"`
}

type Status struct {
	StatusDate              string `json:"statusDate"`
	StatusLastActivity      string `json:"statusLastActivity"`
	StatusDescription       string `json:"statusDescription"`
	StatusReasonDescription string `json:"statusReasonDescription"`
	StatusReasonCode        string `json:"statusReasonCode"`
	LastActivityPathID      string `json:"lastActivityPathId"`
	StatusCode              string `json:"statusCode"`
}
type Subscriber struct {
	BarringByRequestActivityDate      string                 `json:"barringByRequestActivityDate"`
	BarringByRequestIndicator         string                 `json:"barringByRequestIndicator"`
	BarringByRequestReasonCode        string                 `json:"barringByRequestReasonCode"`
	BarringByRequestReasonDescription string                 `json:"barringByRequestReasonDescription"`
	CalculatedPaymentCategory         string                 `json:"calculatedPaymentCategory"`
	CommissionVariant                 string                 `json:"commissionVariant"`
	ConvergenceCode                   string                 `json:"convergenceCode"`
	ConvergenceRunNo                  string                 `json:"convergenceRunNo"`
	ExpirationDate                    string                 `json:"expirationDate"`
	ImsiAliasName                     string                 `json:"imsiAliasName"`
	InitialDealerCode                 string                 `json:"initialDealerCode"`
	InstallationType                  string                 `json:"installationType"`
	MultiSIMIndicator                 string                 `json:"multiSIMIndicator"`
	MultiSIMLevel                     string                 `json:"multiSIMLevel"`
	NetworkCode                       string                 `json:"networkCode"`
	NextSubscriberChangedDate         string                 `json:"nextSubscriberChangedDate"`
	NextSubscriberID                  string                 `json:"nextSubscriberId"`
	OriginalActivationDate            string                 `json:"originalActivationDate"`
	PreviousSubscriberChangedDate     string                 `json:"previousSubscriberChangedDate"`
	PreviousSubscriberID              string                 `json:"previousSubscriberId"`
	PrimResourceParamProductType      string                 `json:"primResourceParamProductType"`
	RelatedSubscriberID               string                 `json:"relatedSubscriberId"`
	RfMigration                       RfMigration            `json:"rfMigration"`
	SubscriberPassword                string                 `json:"subscriberPassword"`
	SubscriberTypeDescription         string                 `json:"subscriberTypeDescription"`
	TrueLifeID                        string                 `json:"trueLifeId"`
	Porting                           porting                `json:"porting"`
	Address                           Address                `json:"address"`
	ChNodeID                          string                 `json:"chNodeId"`
	PrimResourceType                  string                 `json:"primResourceType"`
	SubscriberType                    string                 `json:"subscriberType"`
	SplitPeriod                       string                 `json:"splitPeriod"`
	ExternalID                        string                 `json:"externalId"`
	SubscriberID                      string                 `json:"subscriberId"`
	CreditStatus                      CreditStatus           `json:"creditStatus"`
	RmvMigration                      RmvMigration           `json:"rmvMigration"`
	TmvMigration                      TmvMigration           `json:"tmvMigration"`
	PrimResourceValue                 string                 `json:"primResourceValue"`
	CollectionStatus                  CollectionStatus       `json:"collectionStatus"`
	Contact                           Contact                `json:"contact"`
	CreditLimitAtSubStatus            CreditLimitAtSubStatus `json:"creditLimitAtSubStatus"`
	CustomerID                        string                 `json:"customerId"`
	Dealer                            CodeDescInfo           `json:"dealer"`
	Name                              Name                   `json:"name"`
	Proof                             Proof                  `json:"proof"`
	SmsIndicator                      string                 `json:"smsIndicator"`
	SmsLanguage                       string                 `json:"smsLanguage"`
	EffectiveDate                     string                 `json:"effectiveDate"`
	StartDate                         string                 `json:"startDate"`
	CreateDate                        string                 `json:"createDate"`
	Status                            Status                 `json:"status"`
}
type Ou struct {
	OuID        string `json:"ouId"`
	AgreementID string `json:"agreementId"`
}
type RecurringFrequency struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type PaymentMethod struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type PayChannel struct {
	RecurringFrequency RecurringFrequency `json:"recurringFrequency"`
	PaymentCategory    string             `json:"paymentCategory"`
	PaymentSubMethod   string             `json:"paymentSubMethod"`
	PaymentMethod      PaymentMethod      `json:"paymentMethod"`
}
type ConvergenceInfoArray struct {
	ConvergenceCode string `json:"convergenceCode"`
	AssetGroupID    string `json:"assetGroupId"`
}
type ConvergenceList struct {
	Size                 string                 `json:"size"`
	ConvergenceInfoArray []ConvergenceInfoArray `json:"convergenceInfoArray"`
}
type AccountSubType struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Ben struct {
	ConsolidateIndicator string       `json:"consolidateIndicator"`
	BillLanguage         string       `json:"billLanguage"`
	Ben                  string       `json:"ben"`
	BcBanID              string       `json:"bcBanId"`
	BillMedia            string       `json:"billMedia"`
	Status               CodeDescInfo `json:"status"`
}
type Company struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type Account struct {
	ConvergenceCode   string         `json:"convergenceCode"`
	AccountID         string         `json:"accountId"`
	Classify          string         `json:"classify"`
	StatusDescription string         `json:"statusDescription"`
	AccountSubType    AccountSubType `json:"accountSubType"`
	Contact           Contact        `json:"contact"`
	Name              Name           `json:"name"`
	Ben               Ben            `json:"ben"`
	Company           Company        `json:"company"`
	ArBalance         string         `json:"arBalance"`
	StatusCode        string         `json:"statusCode"`
	TaxID             string         `json:"taxId"`
	BranchNo          string         `json:"branchNo"`
}
type CustomerType struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type BillCycle struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type BillCycleInfo struct {
	ChangeCycleRequestStatus string    `json:"changeCycleRequestStatus"`
	BillCycle                BillCycle `json:"billCycle"`
	ChangeCycleIndicator     string    `json:"changeCycleIndicator"`
	ChangeCycleReqeustDate   string    `json:"changeCycleReqeustDate"`
	NewBillCycle             string    `json:"newBillCycle"`
}
type Customer struct {
	CustomerType      CustomerType  `json:"customerType"`
	CertificateNumber string        `json:"certificateNumber"`
	BillCycleInfo     BillCycleInfo `json:"billCycleInfo"`
	CustomerID        string        `json:"customerId"`
	BirthDate         string        `json:"birthDate"`
	CertificateType   string        `json:"certificateType"`
}
type ProductPreferenceInfoArray struct {
	Subscriber      Subscriber      `json:"subscriber"`
	System          string          `json:"system"`
	Ou              Ou              `json:"ou"`
	PayChannel      PayChannel      `json:"payChannel"`
	ConvergenceList ConvergenceList `json:"convergenceList"`
	Account         Account         `json:"account"`
	Customer        Customer        `json:"customer"`
	BusinessLine    string          `json:"businessLine"`
}
type ProductPreferenceList struct {
	ProductPreferenceInfoArray []ProductPreferenceInfoArray `json:"productPreferenceInfoArray"`
	Size                       string                       `json:"size"`
}
type ReturnGetProductPreferenceList struct {
	UUID                  string                `json:"uuid"`
	ErrorCode             string                `json:"errorCode"`
	Message               string                `json:"message"`
	CalculatedPageSize    string                `json:"calculatedPageSize"`
	TotalSize             string                `json:"totalSize"`
	ProductPreferenceList ProductPreferenceList `json:"productPreferenceList"`
}
type GetProductPreferenceListResponse struct {
	Return ReturnGetProductPreferenceList `json:"return"`
}
