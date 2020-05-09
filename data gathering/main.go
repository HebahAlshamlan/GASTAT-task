package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Page struct {
	Data struct {
		Count   float64
		Limit   float64
		Product []map[string]interface{}
	}
}

func (p Page) TotalPages() int {
	if p.Data.Limit == 0 {
		return 2
	}
	return int(math.Ceil(p.Data.Count / p.Data.Limit))
}

func main() {
	products, err := fetch()
	check(err)

	err = combineAndSave(products)
	check(err)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func combineAndSave(products []map[string]interface{}) error {
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}
	err = save(data)
	return err //wether it's nil or not that's the caller's problem
}

func fetch() ([]map[string]interface{}, error) {
	client := http.Client{Timeout: 30 * time.Second}
	endpoint, err := url.Parse("https://website-api.ecom-dev.tamimimarkets.com/api/product")
	if err != nil {
		return nil, err
	}

	categories := []string{"weekly-offers-1", "exclusive-deals", "fruits--vegetables", "bakery", "meat", "poultry", "seafood", "deli", "dairy", "floral", "food--beverages", "breakfast", "snacks--confectionery", "oil--ghee", "rice--pasta", "water--beverages", "baking", "frozen", "healthy-living", "health--beauty", "disposables", "kitchen-shop", "baby-1", "tamimi-product"}

	products := []map[string]interface{}{}

	for _, cat := range categories {

		q := endpoint.Query()
		q.Set("category", cat)
		endpoint.RawQuery = q.Encode()
		resp, err := client.Get(endpoint.String())
		if err != nil {
			return nil, err
		}

		page := Page{}
		json.NewDecoder(resp.Body).Decode(&page)
		products = append(products, page.Data.Product...)

		for i := 2; i < page.TotalPages(); i++ {
			fmt.Printf("Cat: %s,Page: %d\n", cat, i)
			q.Set("page", strconv.Itoa(i))
			endpoint.RawQuery = q.Encode()
			resp, err := client.Get(endpoint.String())
			if err != nil {
				return nil, err
			}

			page := Page{}
			json.NewDecoder(resp.Body).Decode(&page)
			products = append(products, page.Data.Product...)

		}
	}

	return products, nil
}

func save(data []byte) error {
	err := ioutil.WriteFile("tamimi.json", data, 0644)
	return err
}
