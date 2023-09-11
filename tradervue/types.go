package tradervue

// Tradervue types as per:
// https://github.com/tradervue/api-docs

type Execution struct {
	Datetime   string `json:"datetime"`
	Symbol     string `json:"symbol"`
	Quantity   string `json:"quantity"`
	Price      string `json:"price"`
	Option     string `json:"option"`
	Commission string `json:"commission"`
	Transfee   string `json:"transfee"`
	Ecnfee     string `json:"ecnfee"`
}

type ImportRequest struct {
	AllowDuplicates   bool        `json:"allow_duplicates"`
	OverlayCommisions bool        `json:"overlay_commisions"`
	Tags              []string    `json:"tags"`
	AccountTag        string      `json:"account_tag"`
	Executions        []Execution `json:"executions"`
}

type ImportResponse struct {
	Status string `json:"status"`
}
