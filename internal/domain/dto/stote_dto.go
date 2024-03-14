package dto

type StoreDTO struct {
	Id      int        `json:"id"`
	Address AddressDTO `json:"address"`
}
