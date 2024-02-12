# Photon

Photon is a fast to setup and fast to run micro-framework style thingy for Go, inspired by SlimPHP

```go
func main() {
	app := app.New()

	app.Get("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	}, middleware.ExampleMiddlewareFactory("test"))

	app.Run()
}
```