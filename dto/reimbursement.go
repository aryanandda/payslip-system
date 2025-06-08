package dto

type ReimbursementRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}
