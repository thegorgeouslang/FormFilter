// Author: James Mallon <jamesmallondev@gmail.com>
// layout package -
package formfilter

import (
	"errors"
	"net/http"
	"regexp"
)

// Struct type FormHelper -
type FormValidator struct {
	Rules    map[string]string
	Messages map[string]string
}

// FilterEmail method -
func (this *FormValidator) Filter(r *http.Request) (e []error) {
	_ = r.ParseForm()
	fields := r.Form
	for frule, patt := range this.Rules {
		for field, val := range fields {
			if frule == field { // if field rule = field
				cond := regexp.MustCompile(patt)
				if !cond.MatchString(val[0]) {
					if len(this.Messages[field]) > 0 {
						e = append(e, errors.New(this.Messages[field]))
					} else {
						e = append(e, errors.New("The field "+field+" doesn't match the condition"))
					}
				}
			}
		}
	}
	return
}

// CheckErrors method -
func (this *FormValidator) ErrString(es []error, w http.ResponseWriter) (emsg string) {
	for _, e := range es {
		emsg = emsg + e.Error()
	}
	return
}
