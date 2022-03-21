package prishops

// Shop represents a shop
type Shop struct {
	Name       string   `json:"name"`
	Prefecture string   `json:"prefecture"`
	Address    string   `json:"address"`
	Series     []string `json:"series"`
}
