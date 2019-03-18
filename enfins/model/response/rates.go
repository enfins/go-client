package response

type Rates struct {
	WithdrawAmount float32 `json:"withdraw_amount"`
	WithdrawCurr   string  `json:"withdraw_curr"`
	ReceiveAmount  float32 `json:"receive_amount"`
	ReceiveCurr    string  `json:"receive_curr"`
	ExchangeRate   float32 `json:"exchange_rate"`
}
