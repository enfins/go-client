package response

type Stats struct {
	Statistic []Stat `json:"statistic"`
}

type Stat struct {
	TotalIn    float32 `json:"total_in,string,omitempty"`
	TotalOut   float32 `json:"total_out,string,omitempty"`
	Currency   string  `json:"currency,omitempty"`
}
