package gin

import (
	"net/http"
	"server/src/core/service"

	"github.com/gin-gonic/gin"
)

type ginServer struct {
	host    string
	service service.Account
}

func NewGinServer(host string, service service.Account) *ginServer {
	return &ginServer{host, service}
}

func (s *ginServer) Serve() {
	r := gin.Default()

	api := r.Group("/api")

	api.GET("/account", func(c *gin.Context) {
		account, _ := s.service.Get("abc")
		c.JSON(http.StatusOK, account)
	})

	api.GET("/accounts", func(c *gin.Context) {
		accounts := s.service.GetAll()
		c.JSON(http.StatusOK, accounts)
	})

	r.Run(s.host)
}
