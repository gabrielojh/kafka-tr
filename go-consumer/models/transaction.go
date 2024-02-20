package models

type TransactionList struct {
	Transactions []Transaction `json:"transactions" bson:",inline"`
}

type Transaction struct {
	Name     string `json:"name" bson:"name" example:"John Doe"`
	Credit   int    `json:"credit" bson:"credit" example:"100"`
	Category string `json:"category" bson:"category" example:"Food"`
}
