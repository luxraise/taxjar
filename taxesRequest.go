package taxjar

type TaxesRequest struct {
	FromCountry string `json:"from_country"`
	FromZip     string `json:"from_zip"`
	FromState   string `json:"from_state"`
	FromCity    string `json:"from_city"`
	FromStreet  string `json:"from_street"`

	ToCountry string `json:"to_country"`
	ToZip     string `json:"to_zip"`
	ToState   string `json:"to_state"`
	ToCity    string `json:"to_city"`
	ToStreet  string `json:"to_street"`

	Amount   float64 `json:"amount"`
	Shipping float64 `json:"shipping"`

	NexusAddresses []NexusAddress `json:"nexus_addresses,omitempty"`
	LineItems      []LineItem     `json:"line_items,omitempty"`
}
