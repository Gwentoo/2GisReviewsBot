package utils

import "net/http"

func ExpandShortURL(shortURL string) (string, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(shortURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	finalURL := resp.Header.Get("Location")
	if finalURL == "" {
		return shortURL, nil
	}

	return finalURL, nil
}
