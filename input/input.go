/**
 * User: jackong
 * Date: 10/18/13
 * Time: 3:35 PM
 */
package input

import "net/http"

type Input struct {
	req *http.Request
}

func New(req *http.Request) Input {
    return Input{req}
}
func (this Input) Get(name string) string {
	return name
}
