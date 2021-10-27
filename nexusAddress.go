package taxjar

type NexusAddress struct {
	ID      string `json:"id"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
	State   string `json:"state"`
	City    string `json:"city"`
	Street  string `json:"street"`
}
