package models

// ShoppingItem represents a shopping item with a name and amount
type ShoppingItem struct {
	Name   string `json:"name" example:"Milk"`
	Amount int    `json:"amount" example:"2"`
}

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// ResponseMessage represents a standard response message
type ResponseMessage struct {
	Message string `json:"message"`
}
