/**
 * User: jackong
 * Date: 10/18/13
 * Time: 3:13 PM
 */
package mapper

import (
	"github.com/Jackong/gweb/router"
)


var (
	routers map[string]*router.PolicyBase
)

func init() {
    routers = make(map[string]*router.PolicyBase)
}

func Route(pattern string, core router.Core) *router.PolicyBase {
	policy := router.New(core)
	routers[pattern] = policy
	return policy
}

func Get(pattern string) *router.PolicyBase {
	router, ok := routers[pattern]
	if ok {
		return router
	}
	return nil
}
