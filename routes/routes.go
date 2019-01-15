package routes

import (
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
	helpers.APIDisperser("/trabalho",
		models.DefaultAPI{SubPath: "", Handler: controllers.JobDisperser, Auth: true},
		models.DefaultAPI{SubPath: "/", Handler: controllers.GetJob, Auth: true},
	)
}
