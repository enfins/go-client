package enfins

import (
	"testing"
	"fmt"
	"math/rand"
	"net/http"
)

var cfg *Configuration
var client *APIClient

func init() {
	cfg =  &Configuration{
		"http",
		"62.80.163.18:9000",
		"GEpcqDRSwB",
		"gF-GnQZMayC1PJ0rvAU5",
		"v1",
		nil,
	}
	var err error
	client = &APIClient{
		cfg,
		http.DefaultClient,
	}
	if err != nil {
		fmt.Errorf("error initing tests: %s", err)
	}

}

func TestQueryBuilder_AddParam(t *testing.T) {
	qb, _ := NewQuery("/", cfg)
	qb.AddParam("test", "yes")
	if qb.params.Get("test") != "yes" {
		t.Error("Arg not added properly")
	}
}

func TestAPIClient_GetBalance(t *testing.T) {
	_,e,err := client.GetBalance()
	if e != nil {
		t.Errorf("error response with Message '%s'", e.Message)
	}
	if err != nil {
		t.Errorf("error executing with message '%s'", err.Error())
	}
}

func TestAPIClient_PostCreateBill(t *testing.T) {
	_, e, err := client.PostCreateBill(CreateBillPostOpts{
		"UAH",
		100,
		"Test amount",
		fmt.Sprintf( "EXT_ORDER_RAND_%d", rand.Int()),
		nil,
	})
	if e != nil {
		t.Errorf("error response with Message %s", e.Message)
	}
	if err != nil {
		t.Errorf("error executing: %s", err.Error())
	}
}

func TestAPIClient_GetStats(t *testing.T) {
	s,e,err := client.GetStats(StatsOpt{"UAH", 0,0, "", true})
	if e != nil {
		t.Errorf("error response with Message '%s'", e.Message)
	}
	if err != nil {
		t.Errorf("error executing with message '%s'", err.Error())
	}
	if s == nil {
		t.FailNow()
	}
}

func TestAPIClient_GetRates_Success(t *testing.T) {
	t.Skip("Disabled method")
	s,e,err := client.GetRates(RatesOpt{
		"USD",
		"UAH",
		100,
		0,
	})
	if e != nil {
		t.Errorf("error response with Message '%s'", e.Message)
	}
	if err != nil {
		t.Errorf("error executing with message '%s'", err.Error())
	}
	if s == nil {
		t.FailNow()
	}
}

func TestAPIClient_GetRates_Error(t *testing.T) {
	t.Skip("Disabled method")
	s,e,err := client.GetRates(RatesOpt{
		"USD",
		"UAH",
		0,
		0,
	})
	if e != nil {
		t.Errorf("error response with Message '%s'", e.Message)
	}
	if err == nil {
		t.Error("Must occur an error")
		t.FailNow()
	}
	if s != nil {
		t.FailNow()
	}
}

func TestAPIClient_PostPayout_Error(t *testing.T) {
	s,e,err := client.PostPayout(PayoutOpt{
		"UAH",
		"UAH",
		10.00,
		"Testing",
		"00000001", // Not exist
	})
	if e != nil {
		if e.Code == 10606 {// user not found
			return
		}
		t.Errorf("error response with Message '%s'", e.Message)
	}
	if err != nil {
		t.Errorf("error executing with message '%s'", err.Error())
	}
	if s == false  {
		t.Errorf("payout not accepted")
		t.FailNow()
	}
}

func TestAPIClient_PostPayoutCard_Error(t *testing.T) {
	s,e,err := client.PostPayoutCard(PayoutCardOpt{
		"UAH",
		"UAH",
		0.00,
		"Testing",
		"4111111111111111", // Not exist
	})
	if e != nil {
		if e.Code == 10563 {// rejected by limits
			return
		}
		t.Errorf("error response with Message '%s'", e.Message)
	}
	if err != nil {
		t.Errorf("error executing with message '%s'", err.Error())
	}
	if s == false  {
		t.Errorf("payout not accepted")
		t.FailNow()
	}
}

func TestAPIClient_GetHistory(t *testing.T) {
	s,e,err := client.GetHistory(HistoryOpt{
		1496416572,
		1652919375,
		"withdraw",
		100,
		0,
		false,
	})
	if e != nil {
		t.Errorf("error response with Message '%s'", e.Message)
	}
	if err != nil {
		t.Errorf("error executing with message '%s'", err.Error())
	}
	if s == nil {
		t.FailNow()
	}
}