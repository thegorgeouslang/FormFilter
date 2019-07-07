// Author: James Mallon <jamesmallondev@gmail.com>
//  package formfilter -
package formfilter

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var formHelper *FormValidator

// setUp function - data and process initialization
func setUp() {
	fh := FormValidator{}
	fh.Rules = map[string]string{
		"email": `\w{2,64}@\w{2,64}\.\w{2,64}(\.\w+)?`,
	}

	fh.Messages = map[string]string{
		"email": "Invalid email format",
	}
	formHelper = &fh
}

// Test function TestFilter to evaluate the Index action
func TestFilter(t *testing.T) {
	setUp()
	mux := http.NewServeMux()

	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			errs := formHelper.Filter(r)
			if len(errs) > 0 {
				http.Error(w, errs[0].Error(), http.StatusForbidden)
			}
		})
	// create a new responsewriter obj
	w := httptest.NewRecorder()

	// create the request arguments
	user := "email=jamesmcklyntar@test.com&"

	r := httptest.NewRequest("POST", "/", strings.NewReader(user)) // send a request to the get  handler
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	mux.ServeHTTP(w, r)

	if w.Code != 200 { // desired, it means (in the current implementation) that the request was redirected
		t.Errorf("Response code is %v", w.Code)
	}

	// create the request arguments
	user = "email=j@t.com"

	r = httptest.NewRequest("POST", "/", strings.NewReader(user)) // send a request to the get  handler
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	mux.ServeHTTP(w, r)

	if w.Code != 403 { // desired, it means (in the current implementation) that the request was redirected
		t.Errorf("Response code is %v", w.Code)
	}
}
