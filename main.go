package main

import (
	"github.com/mutao-net/rakuten-api-client/rakuten"
)

func main() {
	params := rakuten.QueryParameters{
		ApplicationID: "XXXXX",
		Keyword:       "golang",
		GenreID:       "0",
		Sort:          "+reviewAverage",
	}
	rakuten.GetRakutenItems(params)
}
