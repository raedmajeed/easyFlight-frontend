package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/frontend/config"
	"github.com/raedmajeed/frontend/pkg/api"
	"log"
)

func main() {
	params, err := config.Config()
	if err != nil {
		log.Println(err.Error())
	}

	r := gin.Default()
	r.LoadHTMLGlob("../../html_templates/*.html")
	api.NewFrontendRoutes(r, params)
	if err := r.Run(":" + params.PORT); err != nil {
		log.Println("unable to start gin server on port: ", params.PORT)
	}
}
