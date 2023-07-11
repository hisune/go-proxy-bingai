package api

import (
	"adams549659584/go-proxy-bingai/api/helper"
	"adams549659584/go-proxy-bingai/common"
	"net/http"
	"fmt"
	"io/ioutil"
)

func Sydney(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckAuth(r) {
		helper.UnauthorizedResult(w)
		return
	}
	// Get the request body
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(body)
	common.NewSingleHostReverseProxy(common.BING_SYDNEY_URL).ServeHTTP(w, r)
}
