package handler

import (
	"fmt"
	"jobboard/config"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type IndexHandlerHttp interface {
	GetIndexApi(c *fiber.Ctx) error
	GetIndexWeb(c *fiber.Ctx) error
}

type IndexHandler struct {
	store *session.Store
}

func (ih *IndexHandler) GetIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "world",
	}, config.ViewsLayout)
}

func (ih *IndexHandler) GetAboutUs(c *fiber.Ctx) error {
	return c.Render("aboutus", fiber.Map{
		"Title": "world",
	})
}

func (ih *IndexHandler) GetCounter(c *fiber.Ctx) error {
	sess, err := ih.store.Get(c)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	counter, err := strconv.Atoi(fmt.Sprintf("%v", sess.Get("counter")))
	if err != nil {
		// return c.JSON(fiber.Map{"error": err.Error()})
		counter = 0
	}

	return c.Render("counter", fiber.Map{
		"Title":   "world",
		"Counter": counter,
	}, config.ViewsLayout)
}

func (ih *IndexHandler) PostCounter(c *fiber.Ctx) error {
	sess, err := ih.store.Get(c)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	counter, err := strconv.Atoi(fmt.Sprintf("%v", sess.Get("counter")))
	if err != nil {
		counter = 0
	}
	counter += 1

	sess.Set("counter", counter)
	if err := sess.Save(); err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.Redirect("/counter")
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{
		store: session.New(),
	}
}
