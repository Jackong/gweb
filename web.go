/**
 * User: jackong
 * Date: 10/18/13
 * Time: 11:53 AM
 */
package gweb

import (
	"net/http"
	"fmt"
	"github.com/Jackong/gweb/config"
	"github.com/Jackong/gweb/router"
	"github.com/Jackong/gweb/input"
	"github.com/Jackong/gweb/err"
)


func Go() {
	for _, path := range config.Project["view"]["paths"].([]interface {}) {
		http.Handle(path.(string), http.FileServer(http.Dir(config.Project["view"]["root"].(string))))
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(config.Project["server"]["addr"].(string), nil)
}

func handler(writer http.ResponseWriter, req *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			err := e.(err.Input)
			fmt.Println(err)
			http.Error(writer, "400 Bad Request", http.StatusBadRequest)
			return
		}
	}()

	if !router.IsSupportMethod(req.Method) {
		http.Error(writer, "405 Method Not Allowed", http.StatusMethodNotAllowed)
	}

	policy := router.Router(req.Method, req.URL.Path)
	if policy == nil {
		http.NotFound(writer, req)
		return
	}

	input := input.New(req)
	if ok, output := policy.RunBefore(input); ok == false {
		fmt.Fprint(writer, output)
		return
	}
	output := policy.Handler(input)
	fmt.Fprint(writer, output)
}
