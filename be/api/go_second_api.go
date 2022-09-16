package api

import (
	"net/http"
	"project/be"
)

type GoSecondApiResp struct {
	Str string `json:"str"`
}

func GoSecondApi(req *Req, resp *Resp) {
	a := "ali"

	if (a == "b") {
		resp.Send(http.StatusBadRequest)
	}

	resp.SendData(be.RC_SECOND, &GoSecondApiResp{
		Str: a,
	})
}
