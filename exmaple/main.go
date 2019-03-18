package exmaple

import (
	"fmt"
	"github.com/enfins/go-client/enfins"
)

var client *enfins.APIClient

func main() {

	client = enfins.NewAPIClient("YOUR_IDENT", "YOUR_SECRET_ACCESS_KEY")

	balanceExample()
	historyExample()
}

func balanceExample() {
	balanceResponse, errResponse, err := client.GetBalance()
	if err != nil {
		//TODO do something with internal error
	}
	if errResponse != nil {
		//TODO do something with response error
	}
	//TODO something with Valid response
	fmt.Println(balanceResponse)
}

func historyExample() {
	balanceResponse, errResponse, err := client.GetHistory(enfins.HistoryOpt{
		OperationType: "withdraw",
	})
	if err != nil {
		//TODO do something with internal error
	}
	if errResponse != nil {
		//TODO do something with response error
	}
	//TODO something with Valid response
	fmt.Println(balanceResponse)
}