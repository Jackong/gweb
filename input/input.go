/**
 * User: jackong
 * Date: 10/18/13
 * Time: 3:35 PM
 */
package input

import (
	"net/http"
	"regexp"
	"github.com/Jackong/gweb/err"
)

type Input struct {
	req *http.Request
}

func New(req *http.Request) Input {
    return Input{req}
}

func (this Input) Get(name, pattern, defo string) string {
	value := this.req.FormValue(name)
	if value == "" {
		this.req.ParseForm()
		value = this.req.Form.Get(name)
	}
	if value == "" {
		if defo != "" {
			return defo
		}
		panic(err.Input("Invalid param: " + name))
	}
	if pattern == "" {
		return value
	}
	if match, _ := regexp.MatchString(pattern, value); !match {
		panic(err.Input("Invalid param: " + name))
	}
	return value
}

func (this Input) Default(name, defo string) string {
	return this.Get(name, "", defo)
}

func (this Input) Pattern(name, pattern string) string {
	return this.Get(name, pattern, "")
}

func (this Input) Required(name string) string {
	return this.Pattern(name, "")
}
