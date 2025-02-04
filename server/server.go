package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID       int
	Nombre   string
	Cantidad int
	Code     string
}

var product_list []Product
var changes bool = false
var new_id int = 0

func get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Product_list" : product_list,
	})
}

func comprobate(c *gin.Context) {
	fmt.Println("Comprobando cambios: ", changes)
	c.JSON(http.StatusOK, gin.H{
		"Changes": changes,
	})
}

func send_changes(c *gin.Context) {
	fmt.Println("Enviando cambios")
}

func post(c *gin.Context) {
	var input struct {
		Nombre string `json:"name"`
		Code string `json:"code"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	new_id++

	new_product := &Product{ID: new_id, Nombre: input.Nombre, Cantidad: 0, Code: input.Code}
	product_list = append(product_list, *new_product)

	changes = true

	c.JSON(http.StatusOK, gin.H{
		"Message": "Producto creado",
		"Product": new_product,
	})	
}

func put(c *gin.Context) {
	id, error_param := c.Params.Get("id")
	if !error_param {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error al obtener parámetros",
		})
		return
	}

	id_number, error_strconv := strconv.Atoi(id)
	if error_strconv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Parámetro no válido",
		})
		return
	}

	var input struct {
		Nombre string `json:"name"`
		Amount int `json:"amount"`
		Code string `json:"code"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	var product_edited Product

	i := 0
	for i < len(product_list) {
		if product_list[i].ID == id_number {
			product_list[i].Nombre = input.Nombre
			product_list[i].Cantidad = input.Amount
			product_list[i].Code = input.Code
			product_edited = product_list[i]
			i = len(product_list)
		}
		i++
	}

	changes = true

	c.JSON(http.StatusOK, gin.H{
		"Message": "Producto editado",
		"Product": product_edited,
	})
}

func delete(c *gin.Context) {
	id, error_param := c.Params.Get("id")
	if !error_param {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error al obtener parámetros",
		})
		return
	}

	id_number, error_strconv := strconv.Atoi(id)
	if error_strconv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Parámetro no válido",
		})
		return
	}

	i := 0
	for i < len(product_list) {
		if product_list[i].ID == id_number {
			deleteElement(product_list, i)
			i = len(product_list)
		}
		i++
	}

	changes = true

	c.JSON(http.StatusOK, gin.H{
		"Message": "Producto eliminado",
	})

}

func deleteElement(slice []Product, indice int) []Product {
    return append(slice[:indice], slice[indice+1:]...)
}