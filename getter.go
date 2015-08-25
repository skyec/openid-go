package openid

import (
	"log"
	"net/http"
	"net/url"
)

// Interface that simplifies testing.
type httpGetter interface {
	Get(uri string, headers map[string]string) (resp *http.Response, err error)
	Post(uri string, form url.Values) (resp *http.Response, err error)
}

type defaultGetter struct{}

var urlGetter = &defaultGetter{}

func (*defaultGetter) Get(uri string, headers map[string]string) (resp *http.Response, err error) {
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Println("Error building request object", err)
		return
	}
	for h, v := range headers {
		request.Header.Add(h, v)
	}
	client := &http.Client{}
	resp, err = client.Do(request)
	if err != nil {
		log.Printf("Error making request to: %s: %s", request.URL.String(), err)
	}
	return
}

func (*defaultGetter) Post(uri string, form url.Values) (resp *http.Response, err error) {
	return http.PostForm(uri, form)
}
