package taxjar

type LineItem struct {
	// Primary fields
	ID             string `json:"id"`
	Quantity       int    `json:"quantity"`
	ProductTaxCode string `json:"product_tax_code"`
	UnitPrice      int    `json:"unit_price"`
	Discount       int    `json:"discount"`

	TaxableAmount                int     `json:"taxable_amount"`
	TaxCollectable               float64 `json:"tax_collectable"`
	CombinedTaxRate              float64 `json:"combined_tax_rate"`
	StateTaxableAmount           int     `json:"state_taxable_amount"`
	StateSalesTaxRate            float64 `json:"state_sales_tax_rate"`
	StateAmount                  float64 `json:"state_amount"`
	CountyTaxableAmount          int     `json:"county_taxable_amount"`
	CountyTaxRate                float64 `json:"county_tax_rate"`
	CountyAmount                 float64 `json:"county_amount"`
	CityTaxableAmount            int     `json:"city_taxable_amount"`
	CityTaxRate                  int     `json:"city_tax_rate"`
	CityAmount                   int     `json:"city_amount"`
	SpecialDistrictTaxableAmount int     `json:"special_district_taxable_amount"`
	SpecialTaxRate               float64 `json:"special_tax_rate"`
	SpecialDistrictAmount        float64 `json:"special_district_amount"`
}
