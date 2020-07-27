package tpl

import (
	"html/template"
)

// the function map
var fm = template.FuncMap{
	"testfunc": testfunc,
	// "login":    auth.Login,
	// "signup":   auth.Signup,
}

func testfunc() string {
	return "foo"
}
