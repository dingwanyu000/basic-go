package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/testGo/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//IndexApi为一个Handler
	router.POST("/api/login", LoginApi)
	router.GET("/api/", IndexApi)
	router.POST("/api/addPerson", AddPersonApi)
	router.GET("/api/persons", GetPersonsApi)
	router.GET("/api/person/:id", GetPersonApi)
	router.POST("/api/updatePerson", ModPersonApi)
	router.POST("/api/deletePerson/", DelPersonApi)
	return router
}
