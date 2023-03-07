package tx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Td struct {
	Rcv []string
}

type ServiceEntries struct {
	Body        string `json:"body"`
	Flags       int    `json:"flags"`
	Instruction string `json:"instruction"`
	ServiceId   string `json:"servicw_id"`
}

type Transfers struct {
	Amount                 int    `json:"amount"`
	Comment                string `json:"comment"`
	Fee                    int    `json:"fee"`
	Height                 int    `json:"height"`
	IsIncome               bool   `json:"is_income"`
	IsMining               bool   `json:"is_mining"`
	IsMixing               bool   `json:"is_mixing"`
	IsService              bool   `json:"is_service"`
	PaymentId              string `json:"payment_id"`
	RemoteAddresses        []string
	ServiceEntries         ServiceEntries `json:"service_entriese"`
	ShowSender             bool           `json:"show_sender"`
	Td                     Td             `json:"td"`
	Timestamp              int            `json:"timestamp"`
	TimestampInternalIndex int            `json:"timestamp_internal_index"`
	TxBlobSize             int            `json:"tx_blob_size"`
	TxHash                 string         `json:"tx_hash"`
	TxType                 int            `json:"tx_type"`
	UnlockTime             int            `json:"unlock_time"`
}

type Pi struct {
	Balance              int `json:"balance"`
	CurrentHeight        int `json:"curent_height"`
	TransferEntriesCOunt int `json:"transfer_entries_count"`
	TransfersCount       int `json:"transfers_count"`
	UnlockedBalance      int `json:"unlocked_balance"`
}

type Result struct {
	LastItemIndex  int         `json:"last_item_index"`
	Pi             Pi          `json:"pi"`
	TotalTransfers int         `json:"total_transfers"`
	Transfers      []Transfers `json:"transfers"`
}

type Txs struct {
	Id      string `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result  Result `json:"result"`
}

func CheckTxs() (bool, string, string, string, string, string, string) {

	posturl := "http://localhost:11212/json_rpc"

	jsonData := fmt.Sprintln(`{
    "jsonrpc": "2.0",
    "id": 0,
    "method": "get_recent_txs_and_info",
    "params": {
      "offset": 0,
      "update_provision_info": true,
      "exclude_mining_txs": true,
      "count": 50,
      "order": "FROM_END_TO_BEGIN",
      "exclude_unconfirmed": true
    }
  }`)

	request, error := http.NewRequest("POST", posturl, bytes.NewBuffer([]byte(jsonData)))
	if error != nil {
		fmt.Println("error") // return meaningful statement
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	res, error := client.Do(request)
	if error != nil {
		fmt.Println("error") // return meaningful statement
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	data := Txs{}

	_ = json.Unmarshal([]byte(body), &data)

	//store price
	featuredPrice := 10000000000000 * 100 // 500 $ZANO
	enhancedPrice := 5000000000000 * 300  // 300 $ZANO
	basicPrice := 2000000000000 * 500     // 100 $ZANO

	var storeType string

	//check comment against regex
	for _, tx := range data.Result.Transfers {
		if tx.Comment != "" && strings.Contains(tx.Comment, ";") && tx.Amount >= 1000000000000 {

			result := strings.Split(tx.Comment, ";")

			//featured store
			if tx.Amount >= featuredPrice {
				storeType = "featured"
			}
			if tx.Amount >= enhancedPrice && tx.Amount < featuredPrice {
				storeType = "enhanced"
			}

			if tx.Amount <= basicPrice {
				storeType = "basic"
			}

			return true, result[0], result[1], result[2], result[3], storeType, tx.TxHash

		}
	}
	return false, "", "", "", "", "", ""
}

func TxWatch() (bool, string, string, string, string, string, string) {
	exists, image, title, description, secret, storeType, hash := CheckTxs()
	return exists, image, title, description, secret, storeType, hash
}
