package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID       int
	Nombre   string
	Cantidad int
	Code     string
}

var product_list []Product


func replicate_changes(c *gin.Context) {

	for {
		response, err := http.Get("http://localhost:3000/comprobate")

		if err != nil {
			fmt.Println("Error al hacer la solicitud HTTP:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		if response.StatusCode == http.StatusOK {
			type ResponseMap map[string]bool
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println("Error al leer el cuerpo de la respuesta:", err)
				response.Body.Close()
				time.Sleep(10 * time.Second)
				continue
			}
			
			type Response struct {
				Changes bool `json:"Changes"`
			}

			var result Response

			err = json.Unmarshal(body, &result)

			if err != nil {
				fmt.Println("Error al deserializar la respuesta:", err)
				response.Body.Close()
				time.Sleep(10 * time.Second)
				continue
			}

			fmt.Println("Cambios: ", result.Changes)

			if result.Changes {
				/*
				for {
					res, err := http.Get("http://localhost:3000/send_changes")


				}
					*/

			}
		}else {
			fmt.Printf("Respuesta inesperada del servidor: %d\n", response.StatusCode)
		}

		response.Body.Close()

		time.Sleep(10 * time.Second)
	}

}