# go-swagger-example

# pre-requisite:
- Install go

# run the program
execute "go run main.go"

# access swagger
http://localhost:8080/swagger/index.html


# add new api

STEP 1 - Add new code snippet in main.go,

```
// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /helloworld [get]
func Helloworld(g *gin.Context)  {
	g.JSON(http.StatusOK,"helloworld")
 }
 ```

STEP 2 - execute "swag init" in the root directory /go-swagger-example


