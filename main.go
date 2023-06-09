package main
import (
    "github.com/gin-gonic/gin"
	"github.com/balram1312/go-gin-api/routes"
	
)

func main(){
	routes.Connect()
	router := gin.New()
	routes.EmployeeRouter(router)

	router.Run(":8080")
}