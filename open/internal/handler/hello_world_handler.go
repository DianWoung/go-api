package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"open/internal/logic"
	"open/internal/svc"
)

func HelloWorldHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHelloWorldLogic(r.Context(), svcCtx)
		resp, err := l.HelloWorld()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
