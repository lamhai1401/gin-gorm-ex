package usertrpt

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userbiz "github.com/lamhai1401/gin-gorm-ex/user/business"
	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
	userstorage "github.com/lamhai1401/gin-gorm-ex/user/storage"
	"gorm.io/gorm"
)

func HandleUpdateAnUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dataItem usermodels.User

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := userstorage.NewMySQLStorage(db)
		biz := userbiz.NewUpdateUserBiz(storage)

		if err := biz.UpdateUser(c.Request.Context(), map[string]interface{}{"id": id}, &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
