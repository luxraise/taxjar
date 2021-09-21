package taxjar

type TaxRequest struct {
	FromCountry string `json:"from_country" form:"from_country"`
	FromZip     string `json:"from_zip" form:"from_zip"`
	FromState   string `json:"from_state" form:"from_state"`
	ToCountry   string `json:"to_country" form:"to_country"`
	ToZip       string `json:"to_zip" form:"to_zip"`
	ToState     string `json:"to_state" form:"to_state"`
	Amount      string `json:"amount" form:"amount"`
	Shipping    string `json:"shipping" form:"shipping"`

	LineItems
}

type LineItems struct {
	Quantity       int    `json:"quantity" form:"quantity"`
	UnitPrice      int64  `json:"unit_price" form:"unit_price"`
	ProductTaxCode string `json:"product_tax_code" form:"product_tax_code"`
}
