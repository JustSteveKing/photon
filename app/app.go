package app

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/juststeveking/photon/middleware"
)

type App struct {
	Router *httprouter.Router
}

func New() *App {
	return &App{
		Router: httprouter.New(),
	}
}

func (a *App) Get(path string, handle httprouter.Handle, middlewares ...middleware.Middleware) {
	wrappedHandler := handle
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrappedHandler = middlewares[i](wrappedHandler)
	}
	a.Router.GET(path, wrappedHandler)
}

func (a *App) Post(path string, handle httprouter.Handle, middlewares ...middleware.Middleware) {
	wrappedHandler := handle
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrappedHandler = middlewares[i](wrappedHandler)
	}
	a.Router.POST(path, wrappedHandler)
}

func (a *App) Put(path string, handle httprouter.Handle, middlewares ...middleware.Middleware) {
	wrappedHandler := handle
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrappedHandler = middlewares[i](wrappedHandler)
	}
	a.Router.PUT(path, wrappedHandler)
}

func (a *App) Patch(path string, handle httprouter.Handle, middlewares ...middleware.Middleware) {
	wrappedHandler := handle
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrappedHandler = middlewares[i](wrappedHandler)
	}
	a.Router.PATCH(path, wrappedHandler)
}

func (a *App) Delte(path string, handle httprouter.Handle, middlewares ...middleware.Middleware) {
	wrappedHandler := handle
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrappedHandler = middlewares[i](wrappedHandler)
	}
	a.Router.DELETE(path, wrappedHandler)
}

func (a *App) Options(path string, handle httprouter.Handle, middlewares ...middleware.Middleware) {
	wrappedHandler := handle
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrappedHandler = middlewares[i](wrappedHandler)
	}
	a.Router.OPTIONS(path, wrappedHandler)
}

func (a *App) Head(path string, handle httprouter.Handle, middlewares ...middleware.Middleware) {
	wrappedHandler := handle
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrappedHandler = middlewares[i](wrappedHandler)
	}
	a.Router.HEAD(path, wrappedHandler)
}

func (a *App) Run() {
	if err := http.ListenAndServe(":3000", a.Router); err != nil {
		log.Fatalf("Failed: %v\n\n", err)
	}
}
