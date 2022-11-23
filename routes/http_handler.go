package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HttpHandler = func(http.ResponseWriter, *http.Request, httprouter.Params)
