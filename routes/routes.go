package routes

import (
	"fmt"
	"net/http"

	"github.com/juridigo/juridigo_api_interacao/controllers"
	"github.com/juridigo/juridigo_api_interacao/helpers"
	"github.com/juridigo/juridigo_api_interacao/models"
)

/*
Routes - Controlador de rotas do microsserviço
*/
func Routes() {
	helpers.APIDisperser("/usuario",
		models.DefaultAPI{SubPath: "", Handler: controllers.GetUser, Auth: true},
	)
}

func teste(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
}
