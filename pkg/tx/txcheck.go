package tx

type Result struct {
	Ap             string `json:"ap"`
	At             string `json:"at"`
	B              string `json:"b"`
	Cat            string `json:"cat"`
	Cnt            string `json:"cnt"`
	Com            string `json:"com"`
	Do             string `json:"do"`
	Et             int    `json:"et"`
	Fee            int64  `json:"fee"`
	IndexInTx      int    `json:"index_in_tx"`
	Lci            string `json:"lci"`
	Lco            string `json:"lco"`
	Ot             int    `json:"ot"`
	P              string `json:"p"`
	Pt             string `json:"pt"`
	Security       string `json:"security"`
	T              string `json:"t"`
	Timestamp      int64  `json:"timestamp"`
	TxHash         string `json:"tx_hash"`
	TxHashOriginal string `json:"tx_hash_original"`
	Url            string `json:"url"`
}
type Offers struct {
	Id      string `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result  Result `json:"result"`
}

func TxWatch() string {
	return "Hello world"
}
