package models

import (
	"gopkg.in/mgo.v2/bson"
)

/*
Fluxo - Modelo estrutural de fluxo
*/
type Fluxo struct {
	ID         bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	IDTrabalho string        `bson:"idTrabalho,omitempty" json:"idTrabalho,omitempty"`
	Situacao   string        `bson:"situacao,omitempty" json:"situacao,omitempty"`
	Etapas     []Etapa       `bson:"etapas,omitempty" json:"etapas,omitempty"`
	Status     string        `bson:"status,omitempty" json:"status,omitempty"`
}

/*
Etapa - Modelo construtor de model.fluxo
*/
type Etapa struct {
	EtapaID     string         `bson:"etapaId,omitempty" json:"etapaId,omitempty"`
	Prazo       string         `bson:"prazo,omitempty" json:"prazo,omitempty"`
	Status      string         `bson:"status,omitempty" json:"status,omitempty"`
	Nome        string         `bson:"nome,omitempty" json:"nome,omitempty"`
	Descricao   string         `bson:"descricao,omitempty" json:"descricao,omitempty"`
	Usuario     []UsuarioFluxo `bson:"usuario,omitempty" json:"usuario,omitempty"`
	Localizacao Localizacao    `bson:"localizacao,omitempty" json:"localizacao,omitempty"`
}

/*
UsuarioFluxo - Modelo Usuario model.fluxo
*/
type UsuarioFluxo struct {
	ID             string `bson:"id,omitempty" json:"id,omitempty"`
	Nome           string `bson:"nome,omitempty" json:"nome,omitempty"`
	ImagemPerfil   string `bson:"imagemPerfil,omitempty" json:"imagemPerfil,omitempty"`
	DataAtribuido  string `bson:"dataAtribuido,omitempty" json:"dataAtribuido,omitempty"`
	DataAtualizado string `bson:"dataAtualizado,omitempty" json:"dataAtualizado,omitempty"`
}
