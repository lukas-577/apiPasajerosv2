package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type pasajero struct {
	Rut     string  `json:"rut"`
	Nombre  string  `json:"nombre"`
	Origen  string  `json:"origen"`
	Destino string  `json:"destino"`
	Precio  float64 `json:"precio"`
}

var pasajeros = []pasajero{
	{Rut: "1234", Nombre: "Elias", Origen: "Santiago", Destino: "Iquique", Precio: 128.99},
	{Rut: "123", Nombre: "Pedro", Origen: "Santiago", Destino: "Holanda", Precio: 156.99},
	{Rut: "123456", Nombre: "Nico", Origen: "Santiago", Destino: "Puerto Natales", Precio: 100.99},
}

func getPasajeros(c *gin.Context) {
	//transforma la data en json
	c.IndentedJSON(http.StatusOK, pasajeros)
}

func postPasajeros(c *gin.Context) {
	var newPasajero pasajero
	//recibe una estructura y devuelve un error
	if error := c.BindJSON(&newPasajero); error != nil {
		return
	}

	pasajeros = append(pasajeros, newPasajero)

	//transforma la data en json
	c.IndentedJSON(http.StatusCreated, newPasajero)
}

func getRutPasajero(c *gin.Context) {
	rut := c.Param("rut")

	for _, a := range pasajeros {
		if a.Rut == rut {
			c.IndentedJSON(http.StatusOK, a)
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pasajero no encontrado"})
}

func putPasajero(c *gin.Context) {
	rut := c.Param("rut")

	var nuevoNombre struct {
		Nombre string `json:"nombre"`
	}

	if err := c.BindJSON(&nuevoNombre); err != nil {
		return
	}

	for i, a := range pasajeros {
		if a.Rut == rut {
			//para actualizar el nombre del pasajero
			pasajeros[i].Nombre = nuevoNombre.Nombre
			c.IndentedJSON(http.StatusOK, a)
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pasajero no actualizado"})
}

func deletePasajero(c *gin.Context) {
	rut := c.Param("rut")

	for i, a := range pasajeros {
		if a.Rut == rut {
			pasajeros = append(pasajeros[:i], pasajeros[i+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pasajero no eliminado"})
}

func main() {
	router := gin.Default()

	// Configurar el middleware CORS para permitir cualquier origen
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://react.lumonidy.studio"}
	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Pasajeros is ready!",
		})
	})
	router.POST("/pasajeros", postPasajeros)

	router.GET("/pasajeros", getPasajeros)

	router.GET("/pasajeros/:rut", getRutPasajero)

	router.PUT("/pasajeros/:rut", putPasajero)

	router.DELETE("pasajeros/:rut", deletePasajero)

	router.Run("0.0.0.0:8080")
}
