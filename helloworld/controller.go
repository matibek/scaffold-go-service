package helloworld

import (
	"net/http"

	"github.com/matibek/scaffold-go-service/core"
)

// ReplyHello will respond hello message
func ReplyHello(c *core.Context) {
	name := c.DefaultQuery("name", "World")
	result, err := replyHello(name)
	if err != nil {
		c.Error(err)
		return
	}
	c.String(http.StatusOK, result)
}
