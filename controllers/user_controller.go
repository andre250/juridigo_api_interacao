package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/juridigo/juridigo_api_interacao/helpers"
	"github.com/juridigo/juridigo_api_interacao/models"
	"github.com/juridigo/juridigo_api_interacao/utils"
	"gopkg.in/mgo.v2/bson"
)

/*
GetUser - Função responsável por achar usuario através do parametros
*/
func GetUser(w http.ResponseWriter, r *http.Request) {
	if helpers.ReqRefuse(w, r, "GET") != nil {
		return
	}
	if validateQueryString(w, r.URL.Query()) != nil {
		return
	}

	var rankingQuery []bson.M

	if r.URL.Query().Get("rank") != "" {
		rankingStars := strings.Split(r.URL.Query().Get("rank"), ",")

		for _, stars := range rankingStars {
			i, _ := strconv.Atoi(stars)
			rankingQuery = append(rankingQuery, bson.M{"ranking": i})
		}
	} else {
		for i := 1; i <= 5; i++ {
			rankingQuery = append(rankingQuery, bson.M{"ranking": i})
		}
	}

	dist, _ := strconv.ParseFloat(r.URL.Query().Get("dist"), 64)
	value := getDegree(dist)
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	long, _ := strconv.ParseFloat(r.URL.Query().Get("long"), 64)

	distance := make(map[string]float64)
	distance["maxLat"] = lat + value
	distance["minLat"] = lat - value
	distance["maxLong"] = long + value
	distance["minLong"] = long - value

	itens, _ := helpers.Db().Find("usuarios", bson.M{
		"cadastrais.latitude":  bson.M{"$gte": distance["minLat"], "$lte": distance["maxLat"]},
		"cadastrais.longitude": bson.M{"$gte": distance["minLong"], "$lte": distance["maxLong"]},
		"$or":                  rankingQuery,
	}, -1)

	var resultado []models.Usuario
	for _, iten := range itens {
		var usuario models.Usuario
		json, _ := bson.MarshalJSON(iten)
		bson.UnmarshalJSON(json, &usuario)
		resultado = append(resultado, usuario)
	}

	lita, _ := json.Marshal(resultado)

	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write(lita)
}

func getDegree(dist float64) float64 {
	return dist / 111
}

func validateQueryString(w http.ResponseWriter, reqParams url.Values) error {
	var erros []models.ErroItem

	distance, err := strconv.ParseFloat(reqParams.Get("dist"), 64)
	if err != nil {
		erros = append(erros, models.ErroItem{
			Msg:  "Parametro 'dist' inválido",
			Erro: "dist",
		})
	} else if distance < 0 {
		erros = append(erros, models.ErroItem{
			Msg:  "Parametro 'dist' deve ser maior que 0 ",
			Erro: "dist",
		})
	}
	longitude, err := strconv.ParseFloat(reqParams.Get("long"), 64)
	if err != nil {
		erros = append(erros, models.ErroItem{
			Msg:  "Parametro 'long' inválido",
			Erro: "long",
		})
	} else if longitude > 90 || longitude < -90 {
		erros = append(erros, models.ErroItem{
			Msg:  "Parametro 'long' deve ser maior que -90 e menor que 90",
			Erro: "long",
		})
	}

	latitude, err := strconv.ParseFloat(reqParams.Get("lat"), 64)
	if err != nil {
		erros = append(erros, models.ErroItem{
			Msg:  "Parametro 'lat' inválido",
			Erro: "lat",
		})
	} else if latitude > 90 || latitude < -90 {
		erros = append(erros, models.ErroItem{
			Msg:  "Parametro 'lat' deve ser maior que -90 e menor que 90",
			Erro: "lat",
		})
	}

	ranking := reqParams.Get("rank")
	if ranking != "" {
		rankingStars := strings.Split(ranking, ",")

		for _, stars := range rankingStars {
			i, err := strconv.Atoi(stars)
			if err != nil {
				erros = append(erros, models.ErroItem{
					Msg:  fmt.Sprintf("Valor '%s' inválido como ranking", stars),
					Erro: "ranking",
				})
			} else if i < 1 || i > 5 {
				erros = append(erros, models.ErroItem{
					Msg:  "Valor de ranking deve estar entre 1 e 5",
					Erro: "ranking",
				})
			}
		}

	}

	if len(erros) != 0 {
		listError := models.ErroList{Erros: erros}
		j, _ := json.Marshal(listError)
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write(j)
		return errors.New("erro")
	}
	return nil
}
