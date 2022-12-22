package main

import (
	"api-trunfo/controllers"
	"api-trunfo/database"
	"api-trunfo/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SeteupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriarVeiculoMock() {
	v := models.Veiculos{Nome: "Veiculo mock", VelocidadeMaxima: "100", Potencia: "100", Aceleracao: "100", Cilindradas: "100", Peso: "100", Super: "super"}
	database.DB.Create(&v)
	ID = int(v.ID)
}

func DeletaVeiculoMock() {
	var v models.Veiculos
	database.DB.Delete(&v, ID)
}

func TestRotaHome(t *testing.T) {
	r := SeteupDasRotasDeTeste()
	r.GET("/home", controllers.Home)

	req, _ := http.NewRequest("GET", "/home", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	mockRes := `{"message":"API rodando com sucesso"}`
	resBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, mockRes, string(resBody))

	fmt.Println(mockRes)
}

func TestListaTodosVeiculos(t *testing.T) {
	database.GetDatabase()
	CriarVeiculoMock()
	defer DeletaVeiculoMock()

	r := SeteupDasRotasDeTeste()
	r.GET("/trunfo/veiculo", controllers.GetAllVeiculos)

	req, _ := http.NewRequest("GET", "/trunfo/veiculo", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestListaUmVeiculo(t *testing.T) {
	database.GetDatabase()
	CriarVeiculoMock()
	defer DeletaVeiculoMock()

	r := SeteupDasRotasDeTeste()
	r.GET("/trunfo/veiculo/:id", controllers.GetIDVeiculos)

	pathDaBusca := "/trunfo/veiculo/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var veiculoMock models.Veiculos
	json.Unmarshal(res.Body.Bytes(), &veiculoMock)

	assert.Equal(t, "Veiculo mock", veiculoMock.Nome)
	assert.Equal(t, "100", veiculoMock.VelocidadeMaxima)
	assert.Equal(t, "100", veiculoMock.Potencia)
	assert.Equal(t, "100", veiculoMock.Aceleracao)
	assert.Equal(t, "100", veiculoMock.Cilindradas)
	assert.Equal(t, "100", veiculoMock.Peso)
	assert.Equal(t, "super", veiculoMock.Super)
}

func TestDeletaVeiculo(t *testing.T) {
	database.GetDatabase()
	CriarVeiculoMock()
	defer DeletaVeiculoMock()

	r := SeteupDasRotasDeTeste()
	r.DELETE("/trunfo/veiculo/:id", controllers.DeletaVeiculo)

	pathBusca := "/trunfo/veiculo/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathBusca, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

}

func TestEditaVeiculo(t *testing.T) {
	database.GetDatabase()
	CriarVeiculoMock()
	defer DeletaVeiculoMock()

	r := SeteupDasRotasDeTeste()
	r.PATCH("/trunfo/veiculo/:id", controllers.EditaVeiculo)

	v := models.Veiculos{Nome: "Veiculo teste edit", VelocidadeMaxima: "200", Potencia: "200", Aceleracao: "200", Cilindradas: "200", Peso: "200", Super: "Super test"}
	valorJson, _ := json.Marshal(v)
	pathParEditar := "/trunfo/veiculo/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParEditar, bytes.NewBuffer(valorJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var vMockAtualizado models.Veiculos
	json.Unmarshal(res.Body.Bytes(), &vMockAtualizado)

	fmt.Println(vMockAtualizado.Nome)

	assert.Equal(t, "Veiculo teste edit", vMockAtualizado.Nome)
	// assert.Equal(t, "200", vMockAtualizado.VelocidadeMaxima)
	// assert.Equal(t, "200", vMockAtualizado.Potencia)
	// assert.Equal(t, "200", vMockAtualizado.Aceleracao)
	// assert.Equal(t, "200", vMockAtualizado.Cilindradas)
	// assert.Equal(t, "200", vMockAtualizado.Peso)
	// assert.Equal(t, "Super test", vMockAtualizado.Super)
}
