package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/juridigo/juridigo_api_interacao/helpers"
	"github.com/juridigo/juridigo_api_interacao/models"
	"github.com/juridigo/juridigo_api_interacao/utils"
	"gopkg.in/mgo.v2/bson"
)

func ProposalDisperser(w http.ResponseWriter, r *http.Request) {
	if helpers.ReqRefuse(w, r, "POST", "PUT", "GET") != nil {
		return
	}

	if r.Method == "POST" {
		createProposal(w, r)
	} else if r.Method == "GET" {
		getProposalByUser(w, r)
	} else if r.Method == "PUT" {
		updateProposal(w, r)
	}
}

func updateProposal(w http.ResponseWriter, r *http.Request) {
	proposal := models.Proposta{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&proposal)

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

	err = helpers.Db().Update("propostas", bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": proposal})
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["NOT_FOUND"])
		w.Write([]byte(`{"msg": "Identificador não encontrado", "erro": "id"}`))
		return
	}
	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Proposta atualizado com sucesso"}`))
}

func getProposalByUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("usuario")

	if id == "" {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Identificador deve ser passado", "erro": "id"}`))
		return
	}

	itens, err := helpers.Db().Find("propostas", bson.M{"usuarioRelacionado": id}, -1)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["NOT_FOUND"])
		w.Write([]byte(`{"msg": "Identificador não encontrado", "erro": "id"}`))
		return
	}
	listItens, _ := bson.MarshalJSON(itens)
	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write(listItens)

}

func createProposal(w http.ResponseWriter, r *http.Request) {
	proposal := models.Proposta{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&proposal)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(lintErro(err.Error())))
		return
	}
	proposal.Atribuido = true
	if helpers.Db().Insert("propostas", &proposal) != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(lintErro(err.Error())))
		return
	}
	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Proposta criada com sucesso"}`))
}
