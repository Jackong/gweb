/**
 * User: jackong
 * Date: 10/18/13
 * Time: 3:13 PM
 */
package router


var (
	routers map[string]map[string]*PolicyBase
	methods = []string{GET, PUT, POST, DELETE}
)

const (
	GET = "GET"
	POST = "POST"
	PUT = "PUT"
	DELETE = "DELETE"
)

func init() {
    routers = make(map[string]map[string]*PolicyBase)
}

func Route(method, pattern string, core Core) *PolicyBase {
	if routers[method] == nil {
		routers[method] = make(map[string]*PolicyBase)
	}
	policy := &PolicyBase {Core: core, before: []Before{}}
	routers[method][pattern] = policy
	return policy
}

func Get(pattern string, core Core) *PolicyBase {
    return Route(GET, pattern, core)
}

func Post(pattern string, core Core) *PolicyBase {
    return Route(POST, pattern, core)
}


func Put(pattern string, core Core) *PolicyBase {
	return Route(PUT, pattern, core)
}


func Delete(pattern string, core Core) *PolicyBase {
	return Route(DELETE, pattern, core)
}

func IsSupportMethod(method string) bool {
	return routers[method] != nil
}

func Router(method, pattern string) *PolicyBase {
	policy, ok := routers[method][pattern]
	if ok {
		return policy
	}
	//todo match by pattern
	return nil
}
