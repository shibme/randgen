package api

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	envar_RANDGEN_SERVER_PORT        = "RANDGEN_SERVER_PORT"
	defaultPort               uint16 = 8080
)

func Serve(port uint16) error {
	if port <= 0 {
		portFromEnvar, err := strconv.ParseUint(os.Getenv(envar_RANDGEN_SERVER_PORT), 10, 16)
		if err != nil || portFromEnvar <= 0 {
			port = defaultPort
		} else {
			port = uint16(portFromEnvar)
		}
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/file", genRandFile)
	router.POST("/file", verifyFile)
	router.GET("/data", genRandData)
	router.POST("/data", verifyData)

	fmt.Printf("Serving RandGen API on port %d\n", port)
	return router.Run(":" + fmt.Sprintf("%d", port))
}
