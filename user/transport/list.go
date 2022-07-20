package usertrpt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userbiz "github.com/lamhai1401/gin-gorm-ex/user/business"
	usermodels "github.com/lamhai1401/gin-gorm-ex/user/model"
	userstorage "github.com/lamhai1401/gin-gorm-ex/user/storage"
	"gorm.io/gorm"
)

func HandleListUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging usermodels.DataPaging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		storage := userstorage.NewMySQLStorage(db)
		biz := userbiz.NewListUserBiz(storage)

		result, err := biz.ListUsers(c.Request.Context(), nil, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
	}
}
