package seabattle_resultslogapi

import "github.com/gocraft/web"

type HttpContext struct {
	AuthKey string
	BotID   int64
	UserID  int64
}

func (s *HttpContext) AuthMiddleWare(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	req.ParseForm()
	authkey := req.FormValue("auth")
	s.AuthKey = authkey
	next(rw, req)
}
