package response

type CreatedBill struct {
	Url    string `json:"url,omitempty"`
	BillId int    `json:"bill"`
}
