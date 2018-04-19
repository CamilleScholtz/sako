package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
)

// Funding represents the information from the funding required getmonero
// webpage.
type Funding []struct {
	Title         string
	URL           string
	Current       float64
	Total         float64
	Contributions int64
}

// getMoneroFunding request and returns information from the funding required
// getmonero webpage.
// TODO: Use API, see https://github.com/monero-project/monero-site/issues/689
func getMoneroFunding() (Funding, error) {
	data := Funding{}

	res, err := soup.Get("https://forum.getmonero.org/8/funding-required")
	if err != nil {
		return data, err
	}
	doc := soup.HTMLParse(res)

	title := doc.FindAll("a", "class", "thread-title")
	info := doc.FindAll("div", "class", "funding-info-box")

	r1 := regexp.MustCompile("XMR([0-9\\.,]+)")
	r2 := regexp.MustCompile("([0-9]+) contributions")

	for i := range title {
		ct := r1.FindAllStringSubmatch(info[(2*i)].Text(), 2)
		current, err := strconv.ParseFloat(strings.Replace(ct[0][1], ",", "",
			-1), 32)
		if err != nil {
			return data, err
		}
		total, err := strconv.ParseFloat(strings.Replace(ct[1][1], ",", "",
			-1), 32)
		if err != nil {
			return data, err
		}
		contributions, err := strconv.ParseInt(r2.FindStringSubmatch(info[(2 *
			i)].Text())[1], 10, 0)
		if err != nil {
			return data, err
		}

		data = append(data, struct {
			Title         string
			URL           string
			Current       float64
			Total         float64
			Contributions int64
		}{
			title[i].Text(),
			"https://forum.getmonero.org" + title[i].Attrs()["href"],
			current,
			total,
			contributions,
		})
	}

	return data, nil
}
