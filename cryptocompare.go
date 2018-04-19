package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

// cryptoCompareRequest requests and parses JSON from a specified URL into a
// specified interface.
func cryptoCompareRequest(url string, t interface{}) error {
	c := http.Client{Timeout: time.Second * 5}

	req, err := http.NewRequest(http.MethodGet,
		"https://min-api.cryptocompare.com/data/"+url+"&extraParams=sako", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request: Returned invalid statuscode %d",
			res.StatusCode)
	}

	return json.NewDecoder(res.Body).Decode(t)
}

// Graph represents some of the values found in a CryptoCompare
// `/data/histohour` response.
type Graph struct {
	Time  []uint64
	Value []float64
}

// cryptoCompareGraph request and returns graphing information from the
// CryptoCompare API.
func cryptoCompareGraph(crypto string) (Graph, error) {
	data := Graph{}

	var histohour = struct {
		Data []struct {
			Close float64 `json:"close"`
			Time  uint64  `json:"time"`
		} `json:"Data"`
	}{}
	if err := cryptoCompareRequest("histohour?fsym="+crypto+"&limit=48&tsym="+
		config.Currency, &histohour); err != nil {
		return data, err
	}

	for _, d := range histohour.Data {
		data.Time = append(data.Time, d.Time)
		data.Value = append(data.Value, d.Close)
	}

	return data, nil
}

// News represents some of the values found in a CryptoCompare `/data/news`
// response.
type News []struct {
	Time   uint64
	Title  string
	URL    string
	Source string
}

// cryptoCompareNews request and returns news information from the CryptoCompare
// API.
func cryptoCompareNews(category string, max int) (News, error) {
	data := News{}

	var news = []struct {
		PublishedOn uint64 `json:"published_on"`
		Title       string `json:"title"`
		URL         string `json:"url"`
		SourceInfo  struct {
			Name string `json:"name"`
		} `json:"source_info"`
	}{}
	if err := cryptoCompareRequest("news/?categories="+category,
		&news); err != nil {
		return data, err
	}

	for i, d := range news {
		if i == max {
			break
		}

		data = append(data, struct {
			Time   uint64
			Title  string
			URL    string
			Source string
		}{
			d.PublishedOn,
			d.Title,
			d.URL,
			d.SourceInfo.Name,
		})
	}

	return data, nil
}

// Price represents some of the values found in a CryptoCompare `/data/price`
// response. It also includes a currency symbol.
type Price struct {
	Symbol string
	Value  float64
}

// cryptoComparePrice request and returns price information from the
// CryptoCompare API.
func cryptoComparePrice(crypto string) (Price, error) {
	data := Price{}

	var price = struct {
		USD float64 `json:"USD"`
		EUR float64 `json:"EUR"`
	}{}
	if err := cryptoCompareRequest("price?fsym="+crypto+"&tsyms="+
		config.Currency, &price); err != nil {
		return data, err
	}

	switch config.Currency {
	case "EUR":
		data.Symbol = "€"
	case "USD":
		data.Symbol = "$"
	}
	data.Value = reflect.Indirect(reflect.ValueOf(price)).FieldByName(config.
		Currency).Float()

	return data, nil
}
