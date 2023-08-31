package getalbums

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums string

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
