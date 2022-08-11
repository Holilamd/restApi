package entity

type Customer struct {
	ID           string `json:"id"`
	CustomerName string `json:"customer_name"`
	Cardno       string `json:"cardno"`
	Gender       string `json:"gender"`
}
