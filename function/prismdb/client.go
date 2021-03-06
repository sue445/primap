package prismdb

import (
	"github.com/knakk/sparql"
	"github.com/pkg/errors"
	"strings"
	"time"
)

// Client represents PrismDB API Client
type Client struct {
	repo *sparql.Repo
}

// NewClient create a Client instance
func NewClient() (*Client, error) {
	repo, err := sparql.NewRepo("https://prismdb.takanakahiko.me/sparql",
		sparql.Timeout(time.Second*30),
	)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Client{repo: repo}, nil
}

// GetAllShops get all shops from PrismDB
func (c *Client) GetAllShops() ([]*Shop, error) {
	query := `
PREFIX prism: <https://prismdb.takanakahiko.me/prism-schema.ttl#>
SELECT
    sample(?prefecture) AS ?prefecture
    sample(?name) AS ?name
    sample(?address) AS ?address
    group_concat(distinct ?series; separator=",") AS ?series
WHERE {
    ?shop a prism:Shop;
        prism:series ?series;
        prism:prefecture ?prefecture;
        prism:name ?name;
        prism:address ?address.
}
GROUP BY ?shop ?prefecture
ORDER BY ?prefecture ?shop
`

	res, err := c.repo.Query(query)

	if err != nil {
		return []*Shop{}, errors.WithStack(err)
	}

	var shops []*Shop

	for _, row := range res.Solutions() {
		shop := &Shop{
			Name:       row["name"].String(),
			Prefecture: row["prefecture"].String(),
			Address:    row["address"].String(),
			Series:     strings.Split(row["series"].String(), ","),
		}
		shops = append(shops, shop)
	}

	return shops, nil
}
