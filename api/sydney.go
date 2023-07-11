package api

import (
	"adams549659584/go-proxy-bingai/api/helper"
	"adams549659584/go-proxy-bingai/common"
	"net/http"
	"fmt"
)

func Sydney(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckAuth(r) {
		helper.UnauthorizedResult(w)
		return
	}
	q := req.URL.Query()
	fmt.Println(q)
	common.NewSingleHostReverseProxy(common.BING_SYDNEY_URL).ServeHTTP(w, r)
}
