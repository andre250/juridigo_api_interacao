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
		models.DefaultAPI{SubPath: "", Handler: controllers.GetUser, Auth: false},
	)
	helpers.APIDisperser("/trabalho",
		models.DefaultAPI{SubPath: "", Handler: controllers.JobDisperser, Auth: false},
		models.DefaultAPI{SubPath: "/", Handler: controllers.GetJob, Auth: false},
	)
	helpers.APIDisperser("/proposta",
		models.DefaultAPI{SubPath: "", Handler: controllers.ProposalDisperser, Auth: false},
	)
	helpers.APIDisperser("/fluxo",
		models.DefaultAPI{SubPath: "", Handler: controllers.FlowDisperser, Auth: false},
	)
}

func teste(w http.ResponseWriter, r *http.Request) {
	fmt.Println("oi")

}
