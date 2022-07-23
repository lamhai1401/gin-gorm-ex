package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lamhai1401/gin-gorm-ex/caching"
	"github.com/lamhai1401/gin-gorm-ex/db"
	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
	usertrpt "github.com/lamhai1401/gin-gorm-ex/user/transport"
	"github.com/lamhai1401/gin-gorm-ex/utils"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "example"
const DB_NAME = "quotes"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	// init cache
	caching.InitLocalCaching(100)

	// init utils
	utils.InitSnowflakeGenerator(1)

	// init db
	database := db.InitDB(dsn, dsn, dsn)
	defer db.CloseDB()

	// migrate table
	database.AutoMigrate(&usermodels.User{})

	// init validatort
	validate := validator.New()

	// setup gin
	router := gin.Default()
	router.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	v1 := router.Group("/v1")
	user := v1.Group("/users")
	{
		session := database.NewSession()
		user.POST("/login", usertrpt.HandleLoginUser(session))       // login user return token
		user.POST("", usertrpt.HandleCreateUser(session, validate))  // create user
		user.GET("", usertrpt.HandleListUser(session))               // list
		user.GET("/:email", usertrpt.HandleFindAnUser(session))      // get an item by ID
		user.PUT("/:email", usertrpt.HandleUpdateAnUser(session))    // edit an item by ID
		user.DELETE("/:email", usertrpt.HandleDeleteAnUser(session)) // delete an item by ID
	}
	return router
}
