package esi

import (
	"fmt"
	"net/url"
)

func (e *Client) GetCorporationsCorporationID(id uint, etag string) (Response, error) {

	version := 4
	path := fmt.Sprintf("/v%d/corporations/%d/", version, id)

	url := url.URL{
		Scheme: "https",
		Host:   e.Host,
		Path:   path,
	}

	headers := make(map[string]string)

	if etag != "" {
		headers["If-None-Match"] = etag
	}

	request := Request{
		Method:  "GET",
		Path:    url,
		Headers: headers,
		Body:    []byte(""),
	}

	return e.Request(request)

}
