package dto

type AddressDTO struct {
	Address    string `json:"address"`
	Address2   string `json:"address2"`
	District   string `json:"district"`
	PostalCode string `json:"postal_code"`
	Phone      string `json:"phone"`
}
