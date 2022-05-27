package models

import "gorm.io/gorm"

// Order record
type Order struct {
	gorm.Model
	Email        string  `json:"email" csv:"email"`
	PhoneNumber  string  `json:"phone_number" csv:"phone_number"`
	ParcelWeight float32 `json:"parcel_weight" csv:"parcel_weight"`
	Country      *string `json:"country,omitempty"`
}

// BasicError is returned to the API caller when an error occurs when processing request
// Code contains a short descriptive text indicating the kind of problem
// Message contains the full error message that will help the API caller understand the error
type BasicError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SaveResponse struct {
	Status string `json:"status"`
}
