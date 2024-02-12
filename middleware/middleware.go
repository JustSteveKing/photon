package middleware

import "github.com/julienschmidt/httprouter"

type Middleware func(httprouter.Handle) httprouter.Handle
