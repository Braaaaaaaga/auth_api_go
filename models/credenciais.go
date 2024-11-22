package models

type Credenciais struct {
	Usuario  string `bson:"usuario"`
	Senha    string `bson:"senha"`
	ChaveApi string `bson:"chave_api,omitempty"`
}
