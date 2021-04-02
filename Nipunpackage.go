package Nipunpackage

import(
	"fmt"
	"github.com/levigross/grequests"

)
func Nipunfunc() {
  resp, err := grequests.Get("http://httpbin.org/get", nil)
  if err != nil {
    fmt.Errorf("Unable to make request:", err)
  }
  fmt.Println(resp.String())
}