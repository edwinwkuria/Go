package main

import (
	"net/http"

	"sample.com/webproject/web/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
