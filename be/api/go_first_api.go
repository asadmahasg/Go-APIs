package api

import (
	"net/http"
	"project/be"
)

type GoFirstApiResp struct {
	Value int64 `json:"value"`
}

func GoFirstApi(req *Req, resp *Resp) {
	val := 1
	for i := 1; i < 10; i++ {
		val = (val + val) * i
	}

	if val > 5000 {
		resp.Send(http.StatusBadRequest)
	}

	resp.SendData(be.RC_FIRST, &GoFirstApiResp{
		Value: int64(val),
	})
}
