package controllers

import (
	"encoding/json"
	"net/http"
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
	if helpers.ValidateQueryString(w, r.URL.Query()) != nil {
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

	distance := map[string]float64{
		"maxLat":  lat + value,
		"minLat":  lat - value,
		"maxLong": long + value,
		"minLong": long - value,
	}

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

	w.WriteHeader(utils.HTTPStatusCode["OK"])
	if len(resultado) != 0 {
		lista, _ := json.Marshal(resultado)
		w.Write(lista)

	} else {
		w.Write([]byte("[]"))
	}
}

func getDegree(dist float64) float64 {
	return dist / 111
}
