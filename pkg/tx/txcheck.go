package tx

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Tx struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		LastItemIndex int `json:"last_item_index"`
		Pi            struct {
			Balance              int64 `json:"balance"`
			CurentHeight         int   `json:"curent_height"`
			TransferEntriesCount int   `json:"transfer_entries_count"`
			TransfersCount       int   `json:"transfers_count"`
			UnlockedBalance      int64 `json:"unlocked_balance"`
		} `json:"pi"`
		TotalTransfers int `json:"total_transfers"`
		Transfers      []struct {
			Amount          int64    `json:"amount"`
			Comment         string   `json:"comment"`
			Fee             int64    `json:"fee"`
			Height          int      `json:"height"`
			IsIncome        bool     `json:"is_income"`
			IsMining        bool     `json:"is_mining"`
			IsMixing        bool     `json:"is_mixing"`
			IsService       bool     `json:"is_service"`
			PaymentId       string   `json:"payment_id"`
			RemoteAddresses []string `json:"remote_addresses,omitempty"`
			ShowSender      bool     `json:"show_sender"`
			Td              struct {
				Rcv []int64 `json:"rcv,omitempty"`
				Spn []int64 `json:"spn,omitempty"`
			} `json:"td"`
			Timestamp             int    `json:"timestamp"`
			TransferInternalIndex int    `json:"transfer_internal_index"`
			TxBlobSize            int    `json:"tx_blob_size"`
			TxHash                string `json:"tx_hash"`
			TxType                int    `json:"tx_type"`
			UnlockTime            int    `json:"unlock_time"`
			ServiceEntries        []struct {
				Body        string `json:"body"`
				Flags       int    `json:"flags"`
				Instruction string `json:"instruction"`
				ServiceId   string `json:"service_id"`
			} `json:"service_entries,omitempty"`
		} `json:"transfers"`
	} `json:"result"`
}

func CheckTxs() {

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
	fmt.Println("response Body:", string(body))
}

func TxWatch() bool {
	CheckTxs()
	return true
}
