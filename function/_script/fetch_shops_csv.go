package main

// Usage: go run fetch_shops_csv.go > /path/to/shops.csv

import (
	"encoding/csv"
	"github.com/sue445/primap/config"
	"github.com/sue445/primap/prismdb"
	"log"
	"os"
	"strings"
)

func main() {
	prismdbClient, err := prismdb.NewClient()
	if err != nil {
		log.Fatalln(err)
	}

	shops, err := prismdbClient.GetAllShops()
	if err != nil {
		log.Fatalln(err)
	}

	shops = config.AggregateShops(shops)

	header := []string{"name", "prefecture", "address", "series"}

	w := csv.NewWriter(os.Stdout)

	if err := w.Write(header); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	for _, shop := range shops {
		record := []string{
			shop.Name,
			shop.Prefecture,
			shop.Address,
			strings.Join(shop.Series, ","),
		}
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
