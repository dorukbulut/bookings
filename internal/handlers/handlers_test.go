package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals-quarters", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors-suite", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "POST", []postData{
		{key: "start", value: "2022-01-01"},
		{key: "end", value: "2022-01-02"},
	}, http.StatusOK},
	{"search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2022-01-01"},
		{key: "end", value: "2022-01-02"},
	}, http.StatusOK},
	{"make-reservation post", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Doruk"},
		{key: "last_name", value: "Bulut"},
		{key: "email", value: "doruk@gmail.com"},
		{key: "phone", value: "555-555-5555"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			res, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			} else {
				if res.StatusCode != e.expectedStatusCode {
					t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, res.StatusCode)
				}
			}
		} else {
			values := url.Values{}

			for _, x := range e.params {
				values.Add(x.key, x.value)
			}

			res, err := ts.Client().PostForm(ts.URL+e.url, values)

			if err != nil {
				t.Log(err)
				t.Fatal(err)
			} else {
				if res.StatusCode != e.expectedStatusCode {
					t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, res.StatusCode)
				}

			}
		}
	}
}
