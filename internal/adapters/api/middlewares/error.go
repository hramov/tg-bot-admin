package middlewares

import (
	appError "github.com/hramov/tg-bot-admin/internal/error"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type AppHandle func(w http.ResponseWriter, r *http.Request, params httprouter.Params) appError.IAppError

func Error(h AppHandle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		err := h(w, r, params)

		//TODO error handling
		if err != nil {
			log.Println(err)
		}
	}
}
