package enfins

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-client/enfins/model/response"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const SCHEMA = "https"
const HOSTNAME = "api.enfins.com"
const VERSION = "v1"

// Scheme use "http" or "https" without semicolon. Example "http" not "http://"
// Host is URL or IP address with or without port. Example: "api.example.com", "127.0.0.1", "127.0.0.1:8888"
type Configuration struct {
	Scheme        string
	Host          string
	Ident         string
	SecretKey     string
	Version       string
	DefaultHeader map[string]string
}

type StatsOpt struct {
	Currency      string
	Begin         int
	End           int
	OperationType string
	ShowTesting   bool
}

type HistoryOpt struct {
	Begin         int
	End           int
	OperationType string
	Limit         int
	Offset        int
	ShowTesting   bool
}

type RatesOpt struct {
	From          string
	To            string
	Amount        float32
	ReceiveAmount float32
}

type PayoutOpt struct {
	Currency    string
	ToCurrency  string
	Amount      float32
	Description string
	Account     string
}
type PayoutCardOpt struct {
	Currency    string
	ToCurrency  string
	Amount      float32
	Description string
	CardNumber  string
}

type CreateBillPostOpts struct {
	Currency    string
	Amount      float32
	Description string
	MOrder      string
	Optional    *CreateBillOptional
}

type CreateBillOptional struct {
	Payeer     string
	SuccessUrl string
	FailUrl    string
	StatusUrl  string
	MName      string
	ExpireTtl  string
	ConvertTo  string
	Testing    bool
}

// Use NewAPIClient for default config options
// also you can create new with your configuration and http.Client
type APIClient struct {
	cfg        *Configuration
	httpClient *http.Client
}

// Create new client to make simple requests
func NewAPIClient(ident string, secret string) *APIClient {
	cfg := &Configuration{
		SCHEMA,
		HOSTNAME,
		ident,
		secret,
		VERSION,
		nil,
	}
	c := &APIClient{
		cfg,
		&http.Client{
			Timeout: 15*time.Second,
		},
	}
	return c
}

// Use to set HTTP client
func (api APIClient) SetHttpClient(client *http.Client) {
	api.httpClient = client
}


func (api *APIClient) GetBalance() ([]response.Balance, *response.Error, error) {
	qb, err := NewQuery("balance", api.cfg)
	if err != nil {
		return nil, nil, err
	}
	res, err := api.httpClient.Get(qb.Url())
	if err != nil {
		return nil, nil, err
	}

	var r []response.Balance
	e, err := decode(&r, res)
	if err != nil {
		return nil, nil, err
	}
	if e != nil {
		return nil, e, nil
	}
	return r, nil, nil
}

func (api *APIClient) GetStats(opt *StatsOpt) (*response.Stats, *response.Error, error) {
	qb, err := NewQuery("stats", api.cfg)
	if err != nil {
		return nil, nil, err
	}
	qb.AddParam("currency", opt.Currency)
	if opt.Begin > 0 {
		qb.AddParam("begin", strconv.Itoa(opt.Begin))
	}
	if opt.End > 0 {
		qb.AddParam("end", strconv.Itoa(opt.End))
	}
	if opt.OperationType != "" {
		qb.AddParam("operation_type", opt.OperationType)
	}
	if opt.ShowTesting {
		qb.AddParam("show_testing", "true")
	}
	res, err := api.httpClient.Get(qb.Url())
	if err != nil {
		return nil, nil, err
	}

	r := &response.Stats{}
	e, err := decode(r, res)
	if err != nil {
		return nil, nil, err
	}
	if e != nil {
		return nil, e, nil
	}
	return r, nil, nil
}

// Deprecated
// Temporary unavailable
func (api *APIClient) GetRates(opt *RatesOpt) (*response.Rates, *response.Error, error) {
	qb, err := NewQuery("rates", api.cfg)
	if err != nil {
		return nil, nil, err
	}
	qb.AddParam("from", opt.From)
	qb.AddParam("to", opt.To)
	if opt.Amount <= 0 && opt.ReceiveAmount <= 0 {
		return nil, nil, errors.New("input must have Amount or ReceiveAmount greater then 0")
	}
	if opt.Amount > 0 {
		qb.AddParam("amount", fmt.Sprintf("%f", opt.Amount))
	}
	if opt.ReceiveAmount > 0 {
		qb.AddParam("receive_amount", fmt.Sprintf("%f", opt.ReceiveAmount))
	}
	res, err := api.httpClient.Get(qb.Url())
	if err != nil {
		return nil, nil, err
	}

	r := &response.Rates{}
	e, err := decode(r, res)
	if err != nil {
		return nil, nil, err
	}
	if e != nil {
		return nil, e, nil
	}
	return r, nil, nil
}

func (api *APIClient) CreateBill(opt *CreateBillPostOpts) (*response.CreatedBill, *response.Error, error) {
	qb, err := NewQuery("create_bill", api.cfg)
	if err != nil {
		return nil, nil, err
	}
	qb.AddParam("currency", opt.Currency)
	qb.AddParam("m_order", opt.MOrder)
	qb.AddParam("amount", fmt.Sprintf("%f", opt.Amount))
	qb.AddParam("description", opt.Description)
	if opt.Optional != nil {
		if opt.Optional.Payeer != "" {
			qb.AddParam("payeer", opt.Optional.Payeer)
		}
		if opt.Optional.SuccessUrl != "" {
			qb.AddParam("success_url", opt.Optional.SuccessUrl)
		}
		if opt.Optional.FailUrl != "" {
			qb.AddParam("fail_url", opt.Optional.FailUrl)
		}
		if opt.Optional.StatusUrl != "" {
			qb.AddParam("status_url", opt.Optional.StatusUrl)
		}
		if opt.Optional.MName != "" {
			qb.AddParam("m_name", opt.Optional.MName)
		}
		if opt.Optional.ExpireTtl != "" {
			qb.AddParam("expire_ttl", opt.Optional.ExpireTtl)
		}
		if opt.Optional.ConvertTo != "" {
			qb.AddParam("convert_to", opt.Optional.ConvertTo)
		}
		if opt.Optional.Testing {
			qb.AddParam("testing", "true")
		}
	}
	res, err := api.httpClient.PostForm(qb.Url(), qb.params)
	if err != nil {
		return nil, nil, err
	}

	r := &response.CreatedBill{}
	e, err := decode(r, res)
	if err == nil && e == nil {
		return r, nil, nil
	} else {
		return nil, e, err
	}
}

// Returns true if successfully accepted
func (api *APIClient) Payout(opt *PayoutOpt) (*response.Payout, *response.Error, error) {
	qb, err := NewQuery("payout", api.cfg)
	if err != nil {
		return nil, nil, err
	}

	qb.AddParam("currency", opt.Currency)
	qb.AddParam("to_curr", opt.ToCurrency)
	qb.AddParam("amount", fmt.Sprintf("%f", opt.Amount))
	qb.AddParam("description", opt.Description)
	qb.AddParam("account", opt.Account)
	res, err := api.httpClient.PostForm(qb.Url(), qb.params)
	if err != nil {
		return nil, nil, err
	}

	r := &response.Payout{}
	e, err := decode(r, res)
	if err == nil && e == nil {
		return r, nil, nil
	} else {
		return nil, e, err
	}
}

// Returns true if successfully accepted
func (api *APIClient) PayoutCard(opt *PayoutCardOpt) (*response.Payout, *response.Error, error) {
	qb, err := NewQuery("payout_card", api.cfg)
	if err != nil {
		return nil, nil, err
	}
	qb.AddParam("currency", opt.Currency)
	qb.AddParam("to_curr", opt.ToCurrency)
	qb.AddParam("amount", fmt.Sprintf("%2f", opt.Amount))
	qb.AddParam("description", opt.Description)
	qb.AddParam("card_number", opt.CardNumber)
	res, err := api.httpClient.PostForm(qb.Url(), qb.params)
	if err != nil {
		return nil, nil, err
	}

	r := &response.Payout{}
	e, err := decode(r, res)
	if err == nil && e == nil {
		return r, nil, nil
	} else {
		return nil, e, err
	}
}

func (api *APIClient) GetHistory(opt *HistoryOpt) (*response.History, *response.Error, error) {
	qb, err := NewQuery("history", api.cfg)
	if err != nil {
		return nil, nil, err
	}
	if opt.Begin > 0 {
		qb.AddParam("begin", strconv.Itoa(opt.Begin))
	}
	if opt.End > 0 {
		qb.AddParam("end", strconv.Itoa(opt.End))
	}
	if opt.OperationType != "" {
		qb.AddParam("operation_type", opt.OperationType)
	}
	if opt.Limit > 0 {
		qb.AddParam("limit", strconv.Itoa(opt.Limit))
	}
	if opt.Offset > 0 {
		qb.AddParam("offset", strconv.Itoa(opt.Offset))
	}
	if opt.ShowTesting {
		qb.AddParam("show_testing", "true")
	}
	res, err := api.httpClient.Get(qb.Url())
	if err != nil {
		return nil, nil, err
	}

	r := &response.History{}
	e, err := decode(r, res)
	if err != nil {
		return nil, nil, err
	}
	if e != nil {
		return nil, e, nil
	}
	return r, nil, nil
}

func (api *APIClient) FindBill(mOrderId *string, billId *int) (*response.Bill, *response.Error, error) {
	qb, err := NewQuery("find/bill", api.cfg)
	if err != nil {
		return nil, nil, err
	}
	if mOrderId != nil {
		id := *mOrderId
		qb.AddParam("m_order", id)
	} else if billId != nil {
		bid := *billId
		qb.AddParam("bill_id", strconv.Itoa(bid))
	} else {
		return nil, nil, errors.New("one of 'ext_id' or 'bill_id' parameters must be not nil")
	}
	res, err := api.httpClient.Get(qb.Url())
	if err != nil {
		return nil, nil, err
	}

	r := &response.Bill{}
	e, err := decode(r, res)
	if err != nil {
		return nil, nil, err
	}
	if e != nil {
		return nil, e, nil
	}
	return r, nil, nil
}

func (api *APIClient) FindOrder(orderId int) (*response.Order, *response.Error, error) {
	qb, err := NewQuery("find/order", api.cfg)
	if err != nil {
		return nil, nil, err
	}
	qb.AddParam("order_id", strconv.Itoa(orderId))
	res, err := api.httpClient.Get(qb.Url())
	if err != nil {
		return nil, nil, err
	}

	r := &response.Order{}
	e, err := decode(r, res)
	if err != nil {
		return nil, nil, err
	}
	if e != nil {
		return nil, e, nil
	}
	return r, nil, nil
}

func decode(data interface{}, r *http.Response) (fail *response.Error, err error) {
	b, err := ioutil.ReadAll(r.Body)
	res := &response.Result{
		Data: data,
	}
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, res); err != nil {
		return nil, err
	}

	if !res.Result {
		return &res.Error, nil
	}

	if r.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("HTTP Error:\nHTTP CODE: %d\nREQUEST:%s\nBody: %s", r.StatusCode, r.Request.URL.RequestURI(), string(b)))
	}
	return nil, nil
}
