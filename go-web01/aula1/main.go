package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Code         string  `json:"code"`
	CreationDate string  `json:"creationDate"`
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messagem": "Olá, mundo!",
	})
}

func GetOne(c *gin.Context) {

	data, err := readProductsFile(c)

	var unmProducts []product

	fmt.Println("erro:", err)

	if err == nil {
		json.Unmarshal(data, &unmProducts)
		fmt.Println("Entrou aq")
	}

	// if len(unmProducts) > 0 {
	// 	fmt.Println("Opa! Tem dados!")
	// } else {
	// 	fmt.Println("F total")
	// }
	// c.JSON(http.StatusOK, unmProducts[0])

}

// Trazer todos os produtos
func GetAll(c *gin.Context) {
	data, err := readProductsFile(c)

	if err == nil {
		var products []product
		json.Unmarshal(data, &products)
		c.JSON(http.StatusOK, products)
	}

}

func readProductsFile(c *gin.Context) (data []byte, err error) {
	data, err = ioutil.ReadFile("./go-web01/aula1/products.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro":        err.Error(),
			"mensagem":    "Ocorreu um erro ao ler arquivo  dos produtos.",
			"status_code": http.StatusInternalServerError,
		})
	}

	return data, err
}
func main() {
	router := gin.Default()

	router.GET("/hello", helloWorld)
	//router.GET("/products", GetOne)
	router.GET("/products", GetAll)
	router.Run()
}
