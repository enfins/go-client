package response

type Balance struct {
	Id         int32   `json:"id,omitempty"`
	ClientId   int32   `json:"client_id,omitempty"`
	Balance    float32 `json:"balance,string,omitempty"`
	Currency   string  `json:"currency,omitempty"`
	Active     bool    `json:"active,omitempty"`
	CreateDate int32   `json:"create_date,omitempty"`
	UpdateDate int32   `json:"update_date,omitempty"`
}
