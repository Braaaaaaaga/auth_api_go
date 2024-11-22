package models

type AplicacaoExterna struct {
	Nome     string `bson:"nome"`
	ChaveApi string `bson:"chave_api"`
}
