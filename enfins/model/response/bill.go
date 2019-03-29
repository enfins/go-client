package response

type Bill struct {
	Id          int32  `json:"id"`
	MOrder      string `json:"m_order"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Status      string `json:"status"`
	Description string `json:"description"`
	ExpireAt    int32  `json:"expire_at"`
	CreatedAt   int32  `json:"created_at"`
	IsTesting   bool   `json:"testing"`
}
