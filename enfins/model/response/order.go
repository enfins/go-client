package response

type Order struct {
	Amount    string `json:"amount"`
	BillId    int    `json:"bill_id"`
	CreatedAt int32  `json:"created_at"`
	Currency  string `json:"currency"`
	Fee       string `json:"fee"`
	OrderId   int32  `json:"order_id"`
	Status    string `json:"status"`
	Testing   bool   `json:"testing,boolean"`
	//recipient
	//sender
}
