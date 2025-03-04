// Application
//
// Application description
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /
//	Version: 0.0.1
//
//	Consumes:
//	- application/json
//	- application/xml
//
//	Produces:
//	- application/json
//	- application/xml
//
// swagger:meta
package route

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/viper"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/zubroide/go-api-boilerplate/controller"
	"github.com/zubroide/go-api-boilerplate/dic"
	_ "github.com/zubroide/go-api-boilerplate/route/description" // For Swagger
	"net/http"
	"time"
)

var db = make(map[string]string)

func Setup(builder *di.Builder) *gin.Engine {
	gin.SetMode(viper.GetString("GIN_MODE"))

	r := gin.New()
	r.Use(gin.Recovery())

	r.Use(sentrygin.New(sentrygin.Options{
		Repanic:         true,
		WaitForDelivery: true,
		Timeout:         time.Second * 5,
	}))

	// Display Swagger documentation
	r.StaticFile("doc/swagger.json", "doc/swagger.json")
	config := &ginSwagger.Config{
		URL: "/doc/swagger.json", //The url pointing to API definition
	}
	// use ginSwagger middleware to
	r.GET("/doc/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

	userController := dic.Container.Get(dic.UserController).(*controller.UserController)

	// swagger:route GET /ping common getPing
	//
	// Ping
	//
	// Get Ping and reply Pong
	//
	//     Responses:
	//       200:
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// swagger:route GET /users user GetUsers
	//
	// Users list
	//
	// Get users list data
	//
	//     Responses:
	//       200: UsersResponse
	r.GET("/users", userController.List)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}
