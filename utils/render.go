package utils

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx, component templ.Component) error {
	buffer := bytes.Buffer{}
	writer := io.Writer(&buffer)

	err := component.Render(context.Background(), writer)
	if err != nil {
		return err
	}

	return c.SendString(buffer.String())
}
