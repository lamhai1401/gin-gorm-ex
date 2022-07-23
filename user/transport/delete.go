package usertrpt

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userbiz "github.com/lamhai1401/gin-gorm-ex/user/business"
	userstorage "github.com/lamhai1401/gin-gorm-ex/user/storage"
	"gorm.io/gorm"
)

func HandleDeleteAnUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := userstorage.NewMySQLStorage(db)
		biz := userbiz.NewDeleteUserBiz(storage)

		err = biz.DeleteUser(c.Request.Context(), map[string]interface{}{"id": id})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
