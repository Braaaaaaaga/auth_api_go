package models

type Usuario struct {
	Nome        string `bson:"nome"`
	Usuario     string `bson:"usuario"`
	Senha       string `bson:"senha"`
	NivelAcesso string `bson:"nivel_acesso"`
}
