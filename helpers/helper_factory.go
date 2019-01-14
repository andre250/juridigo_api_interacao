package helpers

import (
	"github.com/juridigo/juridigo_api_interacao/config"
	"github.com/juridigo/juridigo_api_interacao/models"
)

var configuration models.Config

/*
InitiConfig - Inicializador de configurações
*/
func InitConfig() {

	configuration = config.GetConfig()
}
