package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// RegC425 : Resumo de itens do movimento diário (código 02 e 2D)
type RegC425 struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg      string        `bson:"reg" json:"reg"`
	CodItem  string        `bson:"coditem" json:"coditem"`
	Qtd      string        `bson:"qtd" json:"qtd"`
	Unid     string        `bson:"unid" json:"unid"`
	VlItem   string        `bson:"vlitem" json:"vlitem"`
	VlPis    string        `bson:"vlpis" json:"vlpis"`
	VlCofins string        `bson:"vlcofins" json:"vlcofins"`
	DtIni    string        `bson:"dtini" json:"dtini"`
	DtFin    string        `bson:"dtfin" json:"dtfin"`
	Cnpj     string        `bson:"cnpj" json:"cnpj"`
}

// Populate: O métdodo é responsável por preencher os dados pelo sped
func (r *RegC425) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.CodItem = l[2]
	r.Qtd = l[3]
	r.Unid = l[4]
	r.VlItem = l[5]
	r.VlPis = l[6]
	r.VlCofins = l[7]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.Cnpj = reg0000.Cnpj
}