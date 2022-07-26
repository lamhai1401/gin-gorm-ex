package usertrpt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userbiz "github.com/lamhai1401/gin-gorm-ex/user/business"
	userstorage "github.com/lamhai1401/gin-gorm-ex/user/storage"
	"gorm.io/gorm"
)

func HandleFindAnUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")

		storage := userstorage.NewMySQLStorage(db)
		biz := userbiz.NewFindUserBiz(storage)

		data, err := biz.FindUser(c.Request.Context(), map[string]interface{}{"email": email})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
