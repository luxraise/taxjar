# TaxJar
TaxJar is a Client SDK for the TaxJar API

## Usage
### New
```go
func ExampleNew() {
	var err error
	if testClient, err = New("[TaxJar API Key]"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TaxJar Client has been initialized! %v\n", testClient)
}

```

### Client.GetTaxes
```go
func ExampleClient_GetTaxes() {
	var err error
	if testClient, err = New("[Shippo API Key]"); err != nil {
		log.Fatal(err)
	}

	var req TaxesRequest
	req.FromCountry = "US"
	req.ToCountry = "US"
	req.ToState = "OR"
	req.ToZip = "97035"
	req.Amount = 100
	req.Shipping = 10

	var resp *TaxesResponse
	if resp, err = testClient.GetTaxes(req); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Taxes response! %+v\n", resp)
}
```
