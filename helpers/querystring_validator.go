package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/juridigo/juridigo_api_interacao/models"
	"github.com/juridigo/juridigo_api_interacao/utils"
)

/*
ValidateQueryString - Validator de processos passados por URL
*/
func ValidateQueryString(w http.ResponseWriter, reqParams url.Values) error {
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
