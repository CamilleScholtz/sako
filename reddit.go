package main

import "github.com/jzelinskie/geddit"

// Submissions represents some of the values found in a CryptoCompare `/data/news`
// response.
type Submissions []struct {
	Time   float64
	Title  string
	URL    string
	Source string
}

func redditSubmissions(sub string) (Submissions, error) {
	data := Submissions{}

	submission, err := geddit.NewSession("sako").SubredditSubmissions(sub,
		"hot", geddit.ListingOptions{Limit: 6})
	if err != nil {
		return data, err
	}

	for i, d := range submission {
		// Filter out pinned submissions.
		// TODO: See https://github.com/jzelinskie/geddit/issues/39
		if i < 2 {
			continue
		}

		data = append(data, struct {
			Time   float64
			Title  string
			URL    string
			Source string
		}{
			d.DateCreated,
			d.Title,
			"https://reddit.com" + d.Permalink,
			d.Domain,
		})
	}

	return data, nil
}
