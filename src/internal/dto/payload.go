package dto

import (
	"payment/pkg/tools"
)

// Payload ...
type Payload struct {
	State         string `json:"state" validate:"required"`
	Amount        string `json:"amount" validate:"required"`
	TransactionID string `json:"transactionId" validate:"required"`
	sourceType    string
	amountDecimal float64 // parse string amount to int value
}

const lost = "lost"

// AmountDecimal ...
func (p Payload) AmountDecimal() float64 { return p.amountDecimal }

// SetAmountInt ...
func (p *Payload) SetAmountDecimal(amountDecimal float64) {
	// We handle the situation when we lost a positive amount
	if p.State == lost {
		p.amountDecimal = tools.ReverseSign(amountDecimal)
	} else {
		p.amountDecimal = amountDecimal
	}

}

// SourceType ...
func (p Payload) SourceType() string { return p.sourceType }

// SetSourceType ...
func (p *Payload) SetSourceType(sourceType string) { p.sourceType = sourceType }
