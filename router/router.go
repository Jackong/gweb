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
	if this.forward == nil {
		this.forward = before
		return this
	}
	this.before = &Before{}
	return this.before.Before(before)
}

func (this *Before) IsForward(in input.Input) bool {
	if this.before != nil {
		if !this.before.IsForward(in) {
			return false
		}
	}
    if this.forward == nil {
		return true
	}
	return this.forward(in)
}
