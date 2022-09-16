package api

import (
	"flag"
	"net/http"
	"project/be"
	"strings"
)

type Req struct {
	be.Req
}

type Resp struct {
	be.Resp
}

func Main(init bool) {
	if init {
		flag.Parse()
		c = Init()
	}
	//listenUrl := util.MustOsGetEnv("localhost://1578")
	listenUrl := "localhost:1301"
	srvr := http.NewServeMux()
	srvr.HandleFunc("/", MainHandler)
	c.L.Infof("listening on %s", listenUrl)
	c.L.Fatal(http.ListenAndServe(listenUrl, srvr))
}

func MainHandler(w http.ResponseWriter, httpreq *http.Request) {
	req := &Req{Req: be.Req{Request: httpreq}}
	resp := &Resp{Resp: be.Resp{ResponseWriter: w}}

	// Standard headers
	header := w.Header()
	header.Set("Content-Type", "application/json; charset=utf-8")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-XSS-Protection", "1; mode=block")
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("Referrer-Policy", "strict-origin-when-cross-origin")
	header.Set("Vary", "Origin")

	url := req.URL.Path
	if url != "/" {
		url = strings.TrimRight(req.URL.Path, "/")
	}

	// Routing & authorization
	if req.Method == "GET" {
		if url == "/api/u/v1/first" {
			GoFirstApi(req, resp)
			return
		} else if url == "/api/u/v1/second" {
			GoSecondApi(req, resp)
		}
	} else if req.Method == "POST" {
		if url == "api/u/v1/login" {
			c.L.Infof("Login hit")
		}
	}
}
