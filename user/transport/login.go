package usertrpt

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	userbiz "github.com/lamhai1401/gin-gorm-ex/user/business"
	userstorage "github.com/lamhai1401/gin-gorm-ex/user/storage"
	"github.com/lamhai1401/gin-gorm-ex/utils"
	"gorm.io/gorm"
)

func HandleLoginUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// parsing data from client
		var credentail *utils.Credentials
		if err := c.ShouldBindJSON(&credentail); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := userstorage.NewMySQLStorage(db)
		biz := userbiz.NewLoginUserBiz(storage)

		token, err := biz.Login(c.Request.Context(), credentail)
		if err != nil {
			if strings.Contains(err.Error(), "An Authenticate") {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
