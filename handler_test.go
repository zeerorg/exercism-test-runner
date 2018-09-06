package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Test struct {
	name string
	r *http.Request
	testFunc func(resp *http.Response) (err error)
}

func TestWelcome(t *testing.T) {
	ts := httptest.NewServer(getRouter())
	defer ts.Close()

	var tests = CreateTests(ts, t)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(test.r)
			defer resp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}
			if err := test.testFunc(resp); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func CreateTests(ts *httptest.Server, t *testing.T) []Test {
	return []Test{
		{
			name: "Welcome Page",
			r: func() *http.Request {
				req, err := http.NewRequest("GET", ts.URL+"/", nil)
				if err != nil {
					t.Fatal(err)
				}
				return req
			}(),
			testFunc: func(resp *http.Response) (err error) {
				if resp.StatusCode != 200 {
					return errors.New("Bad status code")
				}
				return nil
			},
		},
		{
			name: "Passing Python Test",
			r: func() *http.Request {
				req, err := http.NewRequest("GET", ts.URL+"/test/python/0e3134698b1242f3998a6509b9378a9e", nil)
				if err != nil {
					t.Fatal(err)
				}
				return req
			}(),
			testFunc: func(resp *http.Response) (err error) {
				if resp.StatusCode != 200 {
					return errors.New("Bad status Code")
				}
				var msg map[string]bool
				err = json.NewDecoder(resp.Body).Decode(&msg)
				if err != nil {
					return err
				} else if msg["result"] != true {
					return errors.New("Wrong Result expected true got false")
				}
				return nil
			},
		},
		{
			name: "Failing Python Test",
			r: func() *http.Request {
				req, err := http.NewRequest("GET", ts.URL+"/test/python/d7058833975546f9a5b1f20fcd48cb63", nil)
				if err != nil {
					t.Fatal(err)
				}
				return req
			}(),
			testFunc: func(resp *http.Response) (err error) {
				if resp.StatusCode != 200 {
					return errors.New("Bad status Code")
				}
				var msg map[string]bool
				err = json.NewDecoder(resp.Body).Decode(&msg)
				if err != nil {
					return err
				} else if msg["result"] != false {
					return errors.New("Wrong Result expected false got true")
				}
				return nil
			},
		},
		{
			name: "Wrong language test",
			r: func() *http.Request {
				req, err := http.NewRequest("GET", ts.URL+"/test/haskell/d7058833975546f9a5b1f20fcd48cb63", nil)
				if err != nil {
					t.Fatal(err)
				}
				return req
			}(),
			testFunc: func(resp *http.Response) (err error) {
				if resp.StatusCode != 405 {
					return errors.New("Bad status Code")
				}
				var msg map[string]StrArray
				err = json.NewDecoder(resp.Body).Decode(&msg)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}
