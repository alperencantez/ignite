package middleware

import (
	"github.com/alperencantez/ignite/pkg"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func BodyValidator[T any]() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var errors []*pkg.IgniteError
		body := new(T)

		parserErr := c.BodyParser(&body)
		if parserErr != nil {
			resp := pkg.ApiResponse{
				Success: false,
				Message: "Could not parse request body",
				Errors:  parserErr,
			}
			return c.Status(fiber.StatusBadRequest).JSON(resp)
		}

		err := Validator.Struct(body)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				var el pkg.IgniteError
				el.Field = err.Field()
				el.Tag = err.Tag()

				errors = append(errors, &el)
			}

			resp := pkg.ApiResponse{
				Success: false,
				Message: "Invalid registration values",
				Errors:  errors,
			}
			return c.Status(fiber.StatusBadRequest).JSON(resp)
		}

		return c.Next()
	}
}
