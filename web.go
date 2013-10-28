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
	"github.com/Jackong/gweb/log"
)


func Go() {
	log.Info("Starting server...")
	for _, path := range config.Project["view"]["paths"].([]interface {}) {
		http.Handle(path.(string), http.FileServer(http.Dir(config.Project["view"]["root"].(string))))
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(config.Project["server"]["addr"].(string), nil)
}

func handler(writer http.ResponseWriter, req *http.Request) {
	//todo log
	defer func() {
		if e := recover(); e != nil {
			err := e.(err.Input)
			log.Error(err)
			http.Error(writer, "400 Bad Request", http.StatusBadRequest)
			return
		}
	}()

	if !router.IsSupportMethod(req.Method) {
		log.Error("Not support method: ", req.Method)
		http.Error(writer, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	before, rt := router.GetRouter(req.Method, req.URL.Path)
	if rt == nil {
		log.Error("redirct to /html/", req.URL.Path)
		http.Redirect(writer, req, "/html" + req.URL.Path, http.StatusFound)
		return
	}

	input := input.New(req)
	if ok := before.IsForward(input); !ok {
		log.Error("Can not forward")
		http.Error(writer, "400 Bad Request", http.StatusBadRequest)
		return
	}
	output := rt.Handle(input)
	fmt.Fprint(writer, output)
}
