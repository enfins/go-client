package response

type Operation struct {
	Amount float32
	Comment string
	Currency string
	Description string
	ExternalId string
	Time int64
	OperationTime string
	OrderType string
	PaymentSystemName string
	Status string
	ErrorCode int
	Recipient Recipient
}

type Recipient struct {
	RecipientName string
	RecipientAccount string
}

type History struct {
	Operation []Operation
	Statistic []Stats
	Limit     int
	Offset    int
	Total     int
}