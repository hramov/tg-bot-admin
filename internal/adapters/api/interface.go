package api

import "github.com/julienschmidt/httprouter"

type Handler interface {
	Init(router *httprouter.Router)
}
