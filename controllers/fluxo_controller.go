package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/juridigo/juridigo_api_interacao/helpers"
	"github.com/juridigo/juridigo_api_interacao/models"
	"github.com/juridigo/juridigo_api_interacao/utils"
	"gopkg.in/mgo.v2/bson"
)

func FlowDisperser(w http.ResponseWriter, r *http.Request) {
	if helpers.ReqRefuse(w, r, "POST", "PUT", "GET") != nil {
		return
	}

	if r.Method == "POST" {
		createFlow(w, r)
	} else if r.Method == "GET" {
		getFlowByJob(w, r)
	} else if r.Method == "PUT" {
		updateFlow(w, r)
	}
}

func getFlowByJob(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("trabalho")

	if id == "" {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Identificador deve ser passado", "erro": "id"}`))
		return
	}

	itens, err := helpers.Db().Find("fluxo", bson.M{"idTrabalho": id}, -1)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["NOT_FOUND"])
		w.Write([]byte(`{"msg": "Identificador não encontrado", "erro": "id"}`))
		return
	}
	listItens, _ := bson.MarshalJSON(itens)
	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write(listItens)

}

func createFlow(w http.ResponseWriter, r *http.Request) {
	flow := models.Fluxo{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&flow)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(lintErro(err.Error())))
		return
	}
	flow.Situacao = "iniciado"
	if helpers.Db().Insert("fluxos", &flow) != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(lintErro(err.Error())))
		return
	}
	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Fluxo criado com sucesso"}`))
}

func updateFlow(w http.ResponseWriter, r *http.Request) {
	flow := models.Fluxo{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&flow)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(lintErro(err.Error())))
		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Identificador deve ser passado", "erro": "id"}`))
		return
	}

	err = helpers.Db().Update("fluxo", bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": flow})
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["NOT_FOUND"])
		w.Write([]byte(`{"msg": "Identificador não encontrado", "erro": "id"}`))
		return
	}
	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Fluxo atualizado com sucesso"}`))
}
