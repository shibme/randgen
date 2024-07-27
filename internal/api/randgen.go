package api

import (
	"net/http"
	"strings"
	"time"

	"dev.shib.me/randgen"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
)

func genRandFile(c *gin.Context) {
	sizeStr := c.Query("size")
	if sizeStr == "" {
		c.String(http.StatusBadRequest, "Missing size parameter")
		return
	}

	size, err := humanize.ParseBytes(sizeStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid size parameter")
		return
	}

	name := c.Query("name")
	if name == "" {
		name = c.Query("filename")
		if name == "" {
			name = "randfile_" + time.Now().Format("2006_01_02_15_04_05_000")
		}
	}

	secure := strings.ToLower(c.Query("secure")) == "true" || strings.ToLower(c.Query("secure_random")) == "true"

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+name)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate")
	c.Header("Pragma", "public")

	if err := randgen.WriteRand(c.Writer, int(size), secure); err != nil {
		c.String(http.StatusInternalServerError, "Unable to download file")
		return
	}
}
