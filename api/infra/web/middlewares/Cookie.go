package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Função/Handler que cria um cookie com o token JWT passado pelo último handler
func CreateCookie() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Locals("token")
		cookie := new(fiber.Cookie)

		cookie.Name = "jwt_token"
		cookie.Value = token.(string)
		cookie.Expires = time.Now().Add(time.Hour * 24)

		c.Cookie(cookie)

		return c.Status(200).JSON(fiber.Map{
			"status": "success",
		})
	}
}
