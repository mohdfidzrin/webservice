package main

import (
	"net/http"

	"github.com/mohdfidzrin/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
