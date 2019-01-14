package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/juridigo/juridigo_api_interacao/config"
	"github.com/juridigo/juridigo_api_interacao/helpers"
	"github.com/juridigo/juridigo_api_interacao/models"
	"github.com/juridigo/juridigo_api_interacao/routes"
)

var wg sync.WaitGroup
var configGlobal models.Config

func main() {
	// Inicialização das rotas
	routes.Routes()

	//Processo de definição das configurações
	wg.Add(1)
	config.SetConfig(&wg)
	wg.Wait()
	// Obtenção das configurações de ambiente
	configGlobal = config.GetConfig()
	// inicialização da conexão
	helpers.Connection()
	// Inicialização do servidor
	serverConfig := []string{":", configGlobal.App.Port}
	fmt.Printf("Juridigo [User] v%s ouvindo porta: %s\n", configGlobal.Version, configGlobal.App.Port)
	if http.ListenAndServe(strings.Join(serverConfig, ""), nil) != nil {
		fmt.Println("Porta já esta sendo utilizada")
		log.Fatal()
	}
}
