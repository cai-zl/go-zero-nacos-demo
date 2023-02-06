package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-nacos-demo/user-api/internal/logic"
	"go-zero-nacos-demo/user-api/internal/svc"
	"go-zero-nacos-demo/user-api/internal/types"
)

func AddUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.User
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddUserLogic(r.Context(), svcCtx)
		resp, err := l.AddUser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
