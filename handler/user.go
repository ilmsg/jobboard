package handler

import (
	"jobboard/config"
	"jobboard/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type UserHandler struct {
	store *session.Store
}

func (uh *UserHandler) GetRegister(c *fiber.Ctx) error {
	return c.Render("users/register", fiber.Map{
		"hello": "world",
	}, config.ViewsLayout)
}

func (uh *UserHandler) PostRegister(c *fiber.Ctx) error {
	var register models.Register
	if err := c.BodyParser(&register); err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(register)
}

func (uh *UserHandler) GetLogin(c *fiber.Ctx) error {
	return c.Render("users/login", fiber.Map{
		"hello": "world",
	}, config.ViewsLayout)
}

func (uh *UserHandler) PostLogin(c *fiber.Ctx) error {
	sess, err := uh.store.Get(c)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	var login models.Login
	if err := c.BodyParser(&login); err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	sess.Set("user", login)
	if err := sess.Save(); err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(login)
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		store: session.New(),
	}
}
