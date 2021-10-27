package taxjar

type Breakdown struct {
	TaxableAmount                 int        `json:"taxable_amount"`
	TaxCollectable                float64    `json:"tax_collectable"`
	CombinedTaxRate               float64    `json:"combined_tax_rate"`
	StateTaxableAmount            int        `json:"state_taxable_amount"`
	StateTaxRate                  float64    `json:"state_tax_rate"`
	StateTaxCollectable           float64    `json:"state_tax_collectable"`
	CountyTaxableAmount           int        `json:"county_taxable_amount"`
	CountyTaxRate                 float64    `json:"county_tax_rate"`
	CountyTaxCollectable          float64    `json:"county_tax_collectable"`
	CityTaxableAmount             int        `json:"city_taxable_amount"`
	CityTaxRate                   int        `json:"city_tax_rate"`
	CityTaxCollectable            int        `json:"city_tax_collectable"`
	SpecialDistrictTaxableAmount  int        `json:"special_district_taxable_amount"`
	SpecialTaxRate                float64    `json:"special_tax_rate"`
	SpecialDistrictTaxCollectable float64    `json:"special_district_tax_collectable"`
	LineItems                     []LineItem `json:"line_items"`
}
