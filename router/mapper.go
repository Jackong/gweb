/**
 * User: jackong
 * Date: 10/18/13
 * Time: 3:13 PM
 */
package router

var (
	routers map[string]map[string]Router
	befores map[string]map[string]*Before
	methods = []string{GET, PUT, POST, DELETE}
)

const (
	GET = "GET"
	POST = "POST"
	PUT = "PUT"
	DELETE = "DELETE"
)

func init() {
    routers = make(map[string]map[string]Router)
	befores = make(map[string]map[string] *Before)
}

func RouteFunc(method, pattern string, handler Handle) *Before {
	return Route(method, pattern, proxy{proxy: handler})
}

func Route(method, pattern string, router Router) *Before{
	if routers[method] == nil {
		routers[method] = make(map[string]Router)
		befores[method] = make(map[string]*Before)
	}
	routers[method][pattern] = router
	before := &Before{}
	befores[method][pattern] = before
	return before
}

func IsSupportMethod(method string) bool {
	return routers[method] != nil
}

func GetRouter(method, pattern string) (*Before, Router) {
	router, ok := routers[method][pattern]
	if ok {
		return befores[method][pattern], router
	}
	//todo match by pattern
	return nil, nil
}
