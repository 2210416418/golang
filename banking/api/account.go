package api

type Address struct {
	Apt     int    `json:"apt"`
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode int    `json:"zipcode"`
}
type Account_create struct {
	AccountNumbr   int     `json:"accountNumber"`
	AccountName    string  `json:"accountName"`
	AccountBalance float64 `json:"accountBalance"`
	Address        `json:"address"`
}
