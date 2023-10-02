package handler

import (
	"jobboard/config"
	"jobboard/models"

	"github.com/gofiber/fiber/v2"
)

type JobBoardHandler struct {
}

func (uh *JobBoardHandler) List(c *fiber.Ctx) error {
	var jobboards []models.JobBoard
	return c.Render("jobboard/list", fiber.Map{
		"jobboards": jobboards,
	}, config.ViewsLayout)
}

func (uh *JobBoardHandler) GetById(c *fiber.Ctx) error {
	var jobboard models.JobBoard
	return c.Render("jobboard/detail", fiber.Map{
		"jobboard": jobboard,
	}, config.ViewsLayout)
}

func (uh *JobBoardHandler) EditById(c *fiber.Ctx) error {
	var jobboard models.JobBoard
	return c.Render("jobboard/edit", fiber.Map{
		"jobboard": jobboard,
	}, config.ViewsLayout)
}

func (uh *JobBoardHandler) DeleteById(c *fiber.Ctx) error {
	var jobboard models.JobBoard
	return c.Render("jobboard/detail", fiber.Map{
		"jobboard": jobboard,
	}, config.ViewsLayout)
}

func (uh *JobBoardHandler) Create(c *fiber.Ctx) error {
	var jobboard models.JobBoard
	if err := c.BodyParser(&jobboard); err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(jobboard)
}

func NewJobBoardHandlerHandler() *JobBoardHandler {
	return &JobBoardHandler{}
}
