package response

type Result struct {
	Result bool      `json:"result"`
	Data   interface{} `json:"data,omitempty"`
	Error  Error       `json:"error,omitempty"`
}
