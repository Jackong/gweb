/**
 * User: jackong
 * Date: 10/18/13
 * Time: 12:50 PM
 */
package config

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"fmt"
)

//2 layer is enough, using interface{} is so troublesome
type conf map[string]map[string]interface {}

var (
	dir string = os.Getenv("GOPATH") + "/src/github.com/Jackong/gweb/config"
	Project conf
)

func init() {
	Project = make(conf)
	data, err := ioutil.ReadFile(dir + "/project.json")
	if err == nil {
		json.Unmarshal(data, &Project)
		return
	}
	fmt.Println(err)
	os.Exit(1)
}
