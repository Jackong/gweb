/**
 * User: jackong
 * Date: 10/18/13
 * Time: 3:03 PM
 */
package router

import (
	"github.com/Jackong/gweb/input"
)
type Handle func(input.Input) interface {}

type Router interface {
	Handle(input.Input) interface{}
}

type proxy struct {
	Router
	proxy Handle
}

func (this proxy) Handle(in input.Input) interface {} {
	return this.proxy(in)
}

type BeforeFunc func(input.Input) bool
type Before struct {
	Forward BeforeFunc
}
func (this *Before) Before(before BeforeFunc) {
	this.Forward = before
}
