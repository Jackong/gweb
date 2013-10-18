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
	"github.com/Jackong/gweb/mapper"
	"github.com/Jackong/gweb/input"
)


func Go() {
	for _, path := range config.Project["view"]["paths"].([]interface {}) {
		http.Handle(path.(string), http.FileServer(http.Dir(config.Project["view"]["root"].(string))))
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(config.Project["server"]["addr"].(string), nil)
}

func handler(writer http.ResponseWriter, req *http.Request) {
	router := mapper.Get(req.URL.Path)
	if router == nil {
		http.NotFound(writer, req)
		return
	}
	input := input.New(req)
	if ok, output := router.RunBefore(input); ok == false {
		fmt.Fprint(writer, output)
		return
	}
	output := router.Handler(input)
	fmt.Fprint(writer, output)
}
