package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Trabalho struct {
	ID                 bson.ObjectId      `bson:"_id,omitempty" json:"_id,omitempty"`
	Atribuido          bool               `bson:"atribuido,omitempty" json:"atribuido,omitempty"`
	Situacao           uint               `bson:"situacao,omitempty" json:"situacao,omitempty"`
	TipoTrabalho       uint               `bson:"tipoTrabalho,omitempty" json:"tipoTrabalho,omitempty"`
	CategoriaTrabalho  uint               `bson:"categoriaTrabalho,omitempty" json:"categoriaTrabalho,omitempty"`
	UsuarioResponsável UsuarioResponsável `bson:"usuarioResponsavel,omitempty" json:"usuarioResponsavel,omitempty"`
	UsuarioAtribuido   string             `bson:"usuarioAtribuido,omitempty" json:"usuarioAtribuido,omitempty"`
	Descricao          string             `bson:"descricao,omitempty" json:"descricao,omitempty"`
	Rotulo             string             `bson:"rotulo,omitempty" json:"rotulo,omitempty"`
	Prazo              string             `bson:"prazo,omitempty" json:"prazo,omitempty"`
	Valor              float64            `bson:"valor,omitempty" json:"valor,omitempty"`
	Taxa               float64            `bson:"taxa,omitempty" json:"taxa,omitempty"`
	Classificacao      uint               `bson:"classificacao,omitempty" json:"classificacao,omitempty"`
	Multiplicador      uint               `bson:"multiplicador,omitempty" json:"multiplicador,omitempty"`
	DataAtribuido      string             `bson:"dataAtribuido,omitempty" json:"dataAtribuido,omitempty"`
	DataAtualizado     string             `bson:"dataAtualizado,omitempty" json:"dataAtualizado,omitempty"`
	Localizacao        Localizacao        `bson:"localizacao,omitempty" json:"localizacao,omitempty"`
	Audiencia          Audiencia          `bson:"audiencia,omitempty" json:"audiencia,omitempty"`
}

type UsuarioResponsável struct {
	ID           string `bson:"id,omitempty" json:"id,omitempty"`
	Nome         string `bson:"nome,omitempty" json:"nome,omitempty"`
	Empresa      string `bson:"empresa,omitempty" json:"empresa,omitempty"`
	ImagemPerfil string `bson:"imagemPerfil,omitempty" json:"imagemPerfil,omitempty"`
}

type Localizacao struct {
	Pais        string  `bson:"pais,omitempty" json:"pais,omitempty"`
	Estado      string  `bson:"estado,omitempty" json:"estado,omitempty"`
	Cidade      string  `bson:"cidade,omitempty" json:"cidade,omitempty"`
	Regiao      string  `bson:"regiao,omitempty" json:"regiao,omitempty"`
	Rua         string  `bson:"rua,omitempty" json:"rua,omitempty"`
	Numero      string  `bson:"numero,omitempty" json:"numero,omitempty"`
	Complemento string  `bson:"complemento,omitempty" json:"complemento,omitempty"`
	Longitude   float64 `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude    float64 `bson:"latitude" json:"latitude"`
}

type Audiencia struct {
	Tipo                uint     `bson:"tipo,omitempty" json:"tipo,omitempty"`
	Numero              string   `bson:"numero,omitempty" json:"numero,omitempty"`
	Partes              []string `bson:"partes,omitempty" json:"partes,omitempty"`
	Vara                string   `bson:"vara,omitempty" json:"vara,omitempty"`
	NomeResponsavel     string   `bson:"nomeResponsavel,omitempty" json:"nomeResponsavel,omitempty"`
	TelefoneResponsavel string   `bson:"telefoneResponsavel,omitempty" json:"telefoneResponsavel,omitempty"`
	Forum               string   `bson:"forum,omitempty" json:"forum,omitempty"`
}
