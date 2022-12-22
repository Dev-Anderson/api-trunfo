package models

import "gorm.io/gorm"

type Veiculos struct {
	gorm.Model
	Nome             string `json:"nome"`
	VelocidadeMaxima string `json:"velocidade"`
	Potencia         string `json:"potencia"`
	Aceleracao       string `json:"aceleracao"`
	Cilindradas      string `json:"cilindradas"`
	Peso             string `json:"peso"`
	Super            string `json:"super"`
}
