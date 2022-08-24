package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/realnfcs/ultividros-project/api/infra/services"
)

// Middleware do fiber responsável pela validação do token JWT
func JWTAuth() func(ctx *fiber.Ctx) error {
	err := godotenv.Load()
	if err != nil {
		return func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusNetworkAuthenticationRequired).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusNetworkAuthenticationRequired).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	})
}

// TODO: deve-se criar um novo fluxo com cookies, desde a criação dele com o login até
// a sua exclusão na hora do logout
// Middleware do fiber para validação do token JWT (por meio de fluxo de cookies)
func JWTAuthCookie() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// const BearerSchema = "Bearer "
		// jwt := ctx.Cookies("jwt")
		// token := jwt[len(BearerSchema):]
		token := ctx.Cookies("jwt_token")
		if token == "" {
			return ctx.Status(fiber.StatusNetworkAuthenticationRequired).JSON(fiber.Map{
				"error": "JWT token don't exist",
			})
		}

		if !services.NewJWTService().ValidateToken(token) {
			return ctx.Status(fiber.StatusNetworkAuthenticationRequired).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		godotenv.Load()

		claims := jwt.MapClaims{}
		t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		claim := t.Claims.(jwt.MapClaims)

		ctx.Locals("id", claim["sub"])

		return ctx.Next()
	}
}

// Middleware do fiber responsável por pegar os dados do token JWT
// e passar para o próximo handler
// Para pegar os dados, usa-se ctx.Locals("<campo do token>")
func JWTdata() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		ctx.Locals("id", claims["sub"])
		return ctx.Next()
	}
}
