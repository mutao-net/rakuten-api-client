package main

import (
	"github.com/mutao-net/rakuten-api-client/rakuten"
)

func main() {
	params := rakuten.QueryParameters{
		ApplicationID: "XXXXX",
		Title:         "wine",
		GenreID:       "100317",
		Sort:          "+reviewAverage",
	}
	rakuten.GetRakutenItems(params)
}
