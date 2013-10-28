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
	before *Before
	forward BeforeFunc
}

func (this *Before) Before(before BeforeFunc) *Before {
	this.forward = before
	this.before = &Before{}
	return this.before
}

func (this *Before) IsForward(in input.Input) bool {
	if this.before == nil {
		return true
	}
	if !this.before.IsForward(in) {
		return false
	}
    if this.forward == nil {
		return true
	}
	return this.forward(in)
}
