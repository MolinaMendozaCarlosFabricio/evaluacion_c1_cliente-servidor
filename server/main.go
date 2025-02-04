package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/get", get)
	r.GET("/comprobate", comprobate)
	r.POST("/create", post)
	r.PUT("/edit/:id", put)
	r.DELETE("/delete/:id", delete)

	srv := &http.Server{
		Addr:         ":3000",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  1 * time.Hour,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Error: Server Main hasn't begin")
	}
}