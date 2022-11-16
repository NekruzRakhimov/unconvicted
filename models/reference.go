package models

type Reference struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Surname            string `json:"surname"`
	PhoneNumber        string `json:"phone_number"`
	Email              string `json:"email"`
	BirthDate          string `json:"birth_date"`
	ResidentialAddress string `json:"residential_address"`
	Inn                string `json:"inn"`
	PassportFront      string `json:"passport_front"`
	PassportBack       string `json:"passport_back"`
	PassportSelfie     string `json:"passport_selfie"`
	PaymentReceipt     string `json:"payment_receipt"`
	ReferenceLanguage  string `json:"reference_language"`
	ReferenceTariff    string `json:"reference_tariff"`
	Status             string `json:"status"`
	CreatedAt          string `json:"created_at"`
}
