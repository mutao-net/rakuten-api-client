package rakuten

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type QueryParameters struct {
	ApplicationID string `label:"applicationId"`
	Title         string `label:"title"`
	GenreID       string `label:"genreId"`
	Sort          string `label:"sort"`
}

type ResponesResults struct {
	Items []Item `json:"Items"`
}

type Item struct {
	Item struct {
		ItemName    string `json:"itemName"`
		ItemCaption string `json:"itemCaption"`
		ItemPrice   int    `json:"itemPrice"`
		ShopName    string `json:"shopName"`
		ItemURL     string `json:"itemUrl"`
	} `json:"Item"`
}

const (
	api = "https://app.rakuten.co.jp/services/api/IchibaItem/Search/20170706"
)

func GetRakutenItems(query QueryParameters) {
	params := setParams(query)
	result := getItems(api + params)
	fmt.Println(api + params)
	fmt.Printf("result: %+v\n", result)
}
func setParams(params QueryParameters) string {
	queries := []string{}

	rv := reflect.ValueOf(params)
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		key := field.Tag.Get("label")
		value, err := rv.FieldByName(field.Name).Interface().(string)
		if !err {
			panic(err)
		}
		if value == "" {
			continue
		}

		query := key + "=" + url.QueryEscape(value)
		queries = append(queries, query)
	}
	return "?" + strings.Join(queries, "&")
}

func getItems(url string) ResponesResults {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		log.Fatal("Get Http Error:", err)
	}
	var results ResponesResults
	json.NewDecoder(response.Body).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}
