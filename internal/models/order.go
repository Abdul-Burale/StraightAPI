package models

import (
	//"fmt"
	//"net/http"
	//"encoding/json"
	//"log"
	"time"
)

type Item struct {
	Name 		string	`json:"name"`
	Quantity	int 	`json:"quantity"`
	Price		float64	`json:"price"`
}

type Order struct {
	//ID				primitive.ObjectID  `json:"id"`
	UserID   		string			    `json:"user_id"`
	OrderNumber		string				`json:"order_number"`
	Items			[]Item				`json:"items"`
	TotalAmount		float64				`json:"total_amount"`
	Status			string				`json:"status"` // Pending, Paid, etc.
	CreatedAt		time.Time			`json:"created_at"`
	UpdatedAt		time.Time			`json:"updated_at"`
}
