package middleware

import (
	"reflect"

	"github.com/featriadi/golang-libs/validator"
	"github.com/gofiber/fiber/v2"
)

func BodyValidator(bodyStruct any) fiber.Handler {
	return func(c *fiber.Ctx) error {
		typ := reflect.TypeOf(bodyStruct)

		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}

		requestBody := reflect.New(typ).Interface()

		if err := c.BodyParser(requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		}

		if err := validator.Validator(requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": "Validation error",
				"error":   err,
			})
		}

		c.Locals("requestBody", requestBody)

		return c.Next()
	}
}
