package apimodel

type RequestCpProfile struct {
	CorrelatedId      string `json:"correlatedId" validate:"required"`
	PrimResourceValue string `json:"primResourceValue" validate:"required"`
	BusinessLine      string `json:"businessLine" validate:"required,eq=PAYTV|eq=MOBILE|eq=ONLINE|eq=ALL"`
}

type ResponseCpProfile struct {
	Uuid                string               `json:"uuid,omitempty"`
	ErrorCode           string               `json:"errorCode,omitempty"`
	Message             string               `json:"message,omitempty"`
	BackendResponseList *BackendResponseList `json:"backendResponseList,omitempty"`
	CpProfile           *CpProfile           `json:"cpProfile,omitempty"`
}

type NameInfo struct {
	Title     string `json:"title,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}
type CpProfile struct {
	Payroll           string    `json:"payroll,omitempty"`
	CompanyGroup      string    `json:"companyGroup,omitempty"`
	CompanyName       string    `json:"companyName,omitempty"`
	EmployeeID        string    `json:"employeeId,omitempty"`
	CertificateNumber string    `json:"certificateNumber,omitempty"`
	BirthDate         string    `json:"birthDate,omitempty"`
	Email             string    `json:"email,omitempty"`
	ThaiName          *NameInfo `json:"thaiName,omitempty"`
	EngName           *NameInfo `json:"engName,omitempty"`
}
