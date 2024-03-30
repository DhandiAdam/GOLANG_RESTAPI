package productcontroller

import (
	"net/http"

	"github.com/DhandiAdam/GOLANG_RESTAPI/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var products []models.Product

	//Menggambil data semua produk
	models.DB.Find(&products)
	//Mengirim product ini ke json
	c.JSON(http.StatusOK, gin.H{"Products": products})

}

func Show(c *gin.Context) {

}

func Create(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
