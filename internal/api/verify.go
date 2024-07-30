package api

import (
	"net/http"

	"dev.shib.me/randgen"
	"github.com/gin-gonic/gin"
)

func verifyData(c *gin.Context) {
	req := make(map[string][]byte)
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sum, err := randgen.VerifyData(req["data"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "verified", "checksum": sum})
}

func verifyFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	sum, err := randgen.Verify(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "verified", "checksum": sum})
}
