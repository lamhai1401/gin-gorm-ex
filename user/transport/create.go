package usertrpt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	userbiz "github.com/lamhai1401/gin-gorm-ex/user/business"
	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
	userstorage "github.com/lamhai1401/gin-gorm-ex/user/storage"
	"gorm.io/gorm"
)

func HandleCreateUser(db *gorm.DB, validate *validator.Validate) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem usermodels.User

		if err := c.ShouldBindJSON(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// setup dependencies
		storage := userstorage.NewMySQLStorage(db, validate)
		biz := userbiz.NewCreateUserBiz(storage)

		err := biz.CreateNewUser(c.Request.Context(), &dataItem)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataItem})
	}
}
