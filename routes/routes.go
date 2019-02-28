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
		models.DefaultAPI{SubPath: "/", Handler: controllers.GetUserInfo, Auth: true},
	)
	helpers.APIDisperser("/trabalho",
		models.DefaultAPI{SubPath: "", Handler: controllers.JobDisperser, Auth: true},
		models.DefaultAPI{SubPath: "/", Handler: controllers.GetJob, Auth: true},
		models.DefaultAPI{SubPath: "/aceite", Handler: controllers.AcceptJob, Auth: true},
	)
	helpers.APIDisperser("/proposta",
		models.DefaultAPI{SubPath: "", Handler: controllers.ProposalDisperser, Auth: true},
		models.DefaultAPI{SubPath: "/recusa", Handler: controllers.RefuseProposal, Auth: true},
		models.DefaultAPI{SubPath: "/atualiza", Handler: controllers.UpdateProposalByStatus, Auth: true},
	)
	helpers.APIDisperser("/fluxo",
		models.DefaultAPI{SubPath: "", Handler: controllers.FlowDisperser, Auth: true},
	)
}

func teste(w http.ResponseWriter, r *http.Request) {
	fmt.Println("oi")

}
