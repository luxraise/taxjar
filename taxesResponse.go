package taxjar

type taxesResponse struct {
	Tax TaxesResponse `json:"tax"`
}

type TaxesResponse struct {
	OrderTotalAmount float64 `json:"order_total_amount"`
	Shipping         float64 `json:"shipping"`
	TaxableAmount    float64 `json:"taxable_amount"`

	AmountToCollect float64 `json:"amount_to_collect"`
	Rate            float64 `json:"rate"`

	HasNexus       bool `json:"has_nexus"`
	FreightTaxable bool `json:"freight_taxable"`

	TaxSource string `json:"tax_source"`

	Jurisdictions Jurisdictions `json:"jurisdictions"`
	Breakdown     Breakdown     `json:"breakdown"`
}
