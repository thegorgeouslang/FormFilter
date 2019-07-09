# TheGorgeousFormFilter

Rules and error messages for form input fields are created by adding values to the Rules and Message properties:

**project/controllers/helpers/FormHelper.go**
```Go
package helpers

import (
	. "github.com/thegorgeouslang/formfilter"
)

var FormHelper = FormValidator{}

func init() {
	FormHelper.Rules = map[string]string{
		"email":    `\w{2,64}@\w{2,64}\.\w{2,64}(\.\w+)?`,
		"password": `[A-Za-z\d@$!%*#?&]{8,}`,
	}

	FormHelper.Messages = map[string]string{
		"email":    "Invalid email format",
		"password": `The password must contain at least 8 characters,
                    1 uppercase character [A-Z],
                    1 lowercase character [a-z],
                    1 digit [0-9],
                    1 special character (!, $, #, etc)`,
	}
}

```
**project/controllers/AuthController.go**

```Go
package controllers

import (
	"net/http"
	. "project/controllers/helpers"
)


type authController struct{}

func AuthController() *authController {
	return &authController{}
}

func (this *authController) Signup(w http.ResponseWriter, r *http.Request) {
	// filtering form inputs
	if errs := FormHelper.Filter(r); len(errs) > 0 {
		http.Error(w, FormHelper.ErrString(errs), http.StatusForbidden)
		return
	}

}
```
**main.go**
```Go
package main

import (
	"net/http"
	"project/controllers"
)

func main() {
	// curl -d "email=t@t.com&password=123" -X POST http://localhost:3000/signup
	http.HandleFunc("/signup", controllers.AuthController().Signup)

	panic(http.ListenAndServe(":3000", nil))
}
```

```
$ curl -d "email=t@t.com&password=123" -X POST http://localhost:3000/signup     

Invalid email format
The password must contain at least 8 characters,
                    1 uppercase character [A-Z],
                    1 lowercase character [a-z],
                    1 digit [0-9],
                    1 special character (!, $, #, etc)

```
**by [James Mallon]**

[James Mallon]: <https://www.linkedin.com/in/thiago-mallon/>
