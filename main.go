package main

import (
   "github.com/gin-gonic/gin"
   docs "github.com/go-swagger-example/docs"
   api "github.com/go-swagger-example/api"
   swaggerfiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)

func main()  {
   r := gin.Default()
   docs.SwaggerInfo.BasePath = "/api/v1"
   v1 := r.Group("/api/v1")
   {
    	v1.GET("/helloworld",api.Helloworld)
   }
   r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
   r.Run(":8080")

}