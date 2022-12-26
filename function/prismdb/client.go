package prismdb

import (
	"github.com/knakk/sparql"
	"github.com/pkg/errors"
	"os"
	"strings"
	"time"
)

const (
	defaultEndpoint = "https://prismdb.takanakahiko.me/sparql"
)

// Client represents PrismDB API Client
type Client struct {
	repo *sparql.Repo
}

// NewClient create a Client instance
func NewClient() (*Client, error) {
	endpoint := defaultEndpoint
	if os.Getenv("SPARQL_ENDPOINT") != "" {
		endpoint = os.Getenv("SPARQL_ENDPOINT")
	}

	repo, err := sparql.NewRepo(endpoint,
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
  {
    SELECT
      sample(?prefecture) AS ?prefecture
      sample(?name) AS ?name
      sample(?address) AS ?address
      group_concat(distinct ?series2; separator=",") AS ?series
    WHERE {
      ?shop a prism:Shop;
        prism:series ?series;
        prism:group ?group;
        prism:prefecture ?prefecture;
        prism:name ?name;
        prism:address ?address.
      FILTER (?series IN("primagi"))

      # Add shop group (1. Normal PriMagi shop, 2. Real PriMagi Studio) at the end of a series (e.g. primagi_1, primagi_2)
      BIND(concat("primagi_", ?group) AS ?series2)
    }
    GROUP BY ?shop ?prefecture
    ORDER BY ?prefecture ?shop
  }
  UNION
  {
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
      FILTER (?series NOT IN("primagi"))
    }
    GROUP BY ?shop ?prefecture
    ORDER BY ?prefecture ?shop
  }
}
GROUP BY ?name ?prefecture
ORDER BY ?prefecture ?name
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
