package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/juridigo/juridigo_api_interacao/models"
	"github.com/juridigo/juridigo_api_interacao/utils"
	"gopkg.in/mgo.v2/bson"

	"github.com/juridigo/juridigo_api_interacao/helpers"
)

/*
JobDisperser - Distribuidor de chamadas
*/
func JobDisperser(w http.ResponseWriter, r *http.Request) {
	if helpers.ReqRefuse(w, r, "POST", "PUT", "GET") != nil {
		return
	}

	if r.Method == "POST" {
		createJob(w, r)
	} else if r.Method == "PUT" {
		updateJob(w, r)
	} else if r.Method == "GET" {
		getJobByUser(w, r)
	}
}

/*
GetJob - Função de obtenção de serviço
*/
func GetJob(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.String(), "trabalho/")[1]
	if id == "" {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Identificador deve ser passado", "erro": "id"}`))
		return
	}

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Identificador possui formato inválido", "erro": "id"}`))
		return
	}

	itens, err := helpers.Db().FindOne("trabalhos", bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["NOT_FOUND"])
		w.Write([]byte(`{"msg": "Identificador não encontrado", "erro": "id"}`))
		return
	}
	listItens, _ := bson.MarshalJSON(itens)
	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write(listItens)

}

/*
AcceptJob - Atualiza trabalho e proposta
*/
func AcceptJob(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("trabalho")

	var userInfo models.UserInfo
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&userInfo)

	if id == "" {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Identificador não encontrado", "erro": "trabalho"}`))
		return
	}
	if userInfo.UserID == "" {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Identificador não encontrado", "erro": "userId"}`))
		return
	}

	err := helpers.Db().Update("trabalhos", bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{
		"status":           "1",
		"usuarioAtribuido": userInfo.UserID,
	}})

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["NOT_FOUND"])
		w.Write([]byte(`{"msg": "Identificador não encontrado", "erro": "id"}`))
		return
	}

	if helpers.Db().Insert("propostas", bson.M{
		"idTrabalho":         id,
		"usuarioRelacionado": userInfo.UserID,
		"status":             "0",
		"valor":              userInfo.Valor,
		"prazo":              userInfo.Prazo,
		"longitude":          userInfo.Longitude,
		"latitude":           userInfo.Latitude,
		"rotulo":             userInfo.Rotulo,
	}) != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(lintErro(err.Error())))
		return
	}
	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Proposta criada com sucesso"}`))
}

func getJobByUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("usuario")
	status := r.URL.Query().Get("status")

	var err error
	var itens []interface{}

	if status != "" {

		statusFilter := strings.Split(status, ",")

		var statusQuery []bson.M
		for _, status := range statusFilter {
			statusQuery = append(statusQuery, bson.M{"status": status})
		}

		if id == "" {
			itens, err = helpers.Db().Find("trabalhos", bson.M{
				"usuarioAtribuido": "",
				"$and":             statusQuery,
			}, -1)
		} else {
			itens, err = helpers.Db().Find("trabalhos", bson.M{
				"usuarioAtribuido": id,
				"$or":              statusQuery,
			}, -1)
		}

	} else {
		if id != "" {
			itens, err = helpers.Db().Find("trabalhos", bson.M{"usuarioAtribuido": id}, -1)
		} else {
			w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
			w.Write([]byte(`{"msg": "Identificador não encontrado", "erro": "id"}`))
			return
		}
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

func updateJob(w http.ResponseWriter, r *http.Request) {
	job := models.Trabalho{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&job)

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

	err = helpers.Db().Update("trabalhos", bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": job})
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["NOT_FOUND"])
		w.Write([]byte(`{"msg": "Identificador não encontrado", "erro": "id"}`))
		return
	}
	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Trabalho atualizado com sucesso"}`))
}

func createJob(w http.ResponseWriter, r *http.Request) {
	job := models.Trabalho{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&job)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(lintErro(err.Error())))
		return
	}

	job.Atribuido = false
	job.Situacao = "iniciado"
	job.DataAtualizado = strconv.Itoa(int(time.Now().Unix()))
	job.Status = "0"

	if helpers.Db().Insert("trabalhos", &job) != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(lintErro(err.Error())))
		return
	}
	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Trabalho criado com sucesso"}`))
}

func lintErro(erro string) string {
	fieldError := strings.Trim(strings.Split(strings.Split(erro, "struct field")[1], "of type")[0], " ")
	formatError := strings.Trim(strings.Split(strings.Split(erro, "cannot unmarshal")[1], "into Go struct")[0], " ")
	formatCorrect := strings.Trim(strings.Split(erro, "of type")[1], " ")
	textErro := fmt.Sprintf(`{"msg": "Erro ao parsear o campo %s", "enviado": "%s", "esperado": "%s"}`, fieldError, formatError, formatCorrect)
	return textErro
}
