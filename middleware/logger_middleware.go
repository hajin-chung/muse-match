package middleware

import (
	"fmt"
	"musematch/utils"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	reqMsg := fmt.Sprintf("REQ %15s %s %s ", c.IP(), c.Method(), c.OriginalURL())
	for key, value := range c.GetReqHeaders() {
		reqMsg += fmt.Sprintf("%s:%s;", key, value)
	}
	reqMsg += fmt.Sprintf("%s", string(c.BodyRaw()))
	utils.Log(reqMsg)

	res := c.Next()

	resMsg := fmt.Sprintf("RES %15s %s %s ", c.IP(), c.Method(), c.OriginalURL())
	for key, value := range c.GetRespHeaders() {
		resMsg += fmt.Sprintf("%s:%s;", key, value)
	}
	utils.Log(resMsg)
	return res
}
