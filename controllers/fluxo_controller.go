package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

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
	status := r.URL.Query().Get("status")

	if id == "" {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Identificador deve ser passado", "erro": "id"}`))
		return
	}
	var err error
	var itens []interface{}

	if status != "" {
		statusFilter := strings.Split(status, ",")

		var statusQuery []bson.M
		for _, status := range statusFilter {
			statusQuery = append(statusQuery, bson.M{"status": status})
		}

		itens, err = helpers.Db().Find("propostas", bson.M{
			"idTrabalho": id,
			"$or":        statusQuery,
		}, -1)

	} else {
		itens, err = helpers.Db().Find("fluxos", bson.M{"idTrabalho": id}, -1)
	}
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
	flow.Status = "0"
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

	err = helpers.Db().Update("fluxos", bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": flow})
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["NOT_FOUND"])
		w.Write([]byte(`{"msg": "Identificador não encontrado", "erro": "id"}`))
		return
	}
	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Fluxo atualizado com sucesso"}`))
}
