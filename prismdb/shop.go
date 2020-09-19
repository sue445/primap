package prismdb

// Shop represents a shop with pripara or prichan
type Shop struct {
	Name       string   `json:"name"`
	Prefecture string   `json:"prefecture"`
	Address    string   `json:"address"`
	Series     []string `json:"series"`
}
