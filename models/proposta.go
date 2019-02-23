package models

import "gopkg.in/mgo.v2/bson"

/*
Proposta - modelo estrutral de proposta
*/
type Proposta struct {
	ID                 bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	IDTrabalho         string        `bson:"idTrabalho,omitempty" json:"idTrabalho,omitempty"`
	UsuarioRelacionado string        `bson:"usuarioRelacionado" json:"usuarioRelacionado"`
	Atribuido          bool          `bson:"atribuido,omitempty" json:"atribuido,omitempty"`
	TipoTrabalho       string        `bson:"tipoTrabalho,omitempty" json:"tipoTrabalho,omitempty"`
	CategoriaTrabalho  string        `bson:"categoriaTrabalho,omitempty" json:"categoriaTrabalho,omitempty"`
	Valor              string        `bson:"valor,omitempty" json:"valor,omitempty"`
	Empresa            string        `bson:"empresa,omitempty" json:"empresa,omitempty"`
	Rotulo             string        `bson:"rotulo,omitempty" json:"rotulo,omitempty"`
	Prazo              string        `bson:"prazo,omitempty" json:"prazo,omitempty"`
	Localizacao        Localizacao   `bson:"localizacao,omitempty" json:"localizacao,omitempty"`
	Audiencia          Audiencia     `bson:"audiencia,omitempty" json:"audiencia,omitempty"`
	Status             string        `bson:"status,omitempty" json:"status,omitempty"`
}
