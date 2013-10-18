/**
 * User: jackong
 * Date: 10/18/13
 * Time: 3:03 PM
 */
package router

import (
	"github.com/Jackong/gweb/input"
)


type Core interface {
	Handler(input input.Input) interface {}
}

type Before func(input.Input) bool
type PolicyBase struct {
	before[]Before
	Core
}

func (this *PolicyBase) Before(before Before) {
    this.before = append(this.before, before)
}

func (this *PolicyBase) RunBefore(in input.Input) (bool, output interface {}) {
	for _, before := range this.before {
		if !before(in) {
			return false, "error"
		}
	}
	return true, output
}
