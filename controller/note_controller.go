package controller

import (
	"github.com/ggorockee/airbnb-backend/data/request"
	"github.com/ggorockee/airbnb-backend/data/response"
	"github.com/ggorockee/airbnb-backend/helper"
	"github.com/ggorockee/airbnb-backend/service"
	"github.com/gofiber/fiber/v2"
)

type NoteController struct {
	noteService service.NoteService
}

func NewNoteController(service service.NoteService) *NoteController {
	return &NoteController{noteService: service}
}

func (controller *NoteController) Create(c *fiber.Ctx) error {
	createNoteRequest := request.CreateNoteRequest{}
	err := c.BodyParser(&createNoteRequest)
	helper.ErrorPanic(err)

	controller.noteService.Create(createNoteRequest)

	webResponse := response.Response{
		Code:    fiber.StatusOK,
		Status:  "ok",
		Message: "Successfully created notes data!",
		Data:    nil,
	}

	return c.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *NoteController) Update(c *fiber.Ctx) error {
	updateNodeRequest := request.UpdateNoteRequest{}
	err := c.BodyParser(&updateNodeRequest)
	helper.ErrorPanic(err)

	noteId, _ := c.ParamsInt("noteId")
	updateNodeRequest.Id = noteId

	controller.noteService.Update(updateNodeRequest)

	webResponse := response.Response{
		Code:    fiber.StatusOK,
		Status:  "ok",
		Message: "Successfully updated notes data!",
		Data:    nil,
	}

	return c.Status(fiber.StatusOK).JSON(webResponse)

}
func (controller *NoteController) Delete(c *fiber.Ctx) error {
	noteId, _ := c.ParamsInt("noteId")
	controller.noteService.Delete(noteId)
	webResponse := response.Response{
		Code:    fiber.StatusOK,
		Status:  "ok",
		Message: "Successfully updated notes data!",
		Data:    nil,
	}
	return c.Status(fiber.StatusCreated).JSON(webResponse)
}
func (controller *NoteController) FindById(c *fiber.Ctx) error {
	noteId, _ := c.ParamsInt("noteId")

	noteResponse := controller.noteService.FindById(noteId)

	webResponse := response.Response{
		Code:    fiber.StatusOK,
		Status:  "ok",
		Message: "Successfully get notes data by id!",
		Data:    noteResponse,
	}
	return c.Status(fiber.StatusCreated).JSON(webResponse)

}
func (controller *NoteController) FindAll(c *fiber.Ctx) error {
	noteResponse := controller.noteService.FindAll()

	webResponse := response.Response{
		Code:    fiber.StatusOK,
		Status:  "ok",
		Message: "Successfully get notes data",
		Data:    noteResponse,
	}
	return c.Status(fiber.StatusCreated).JSON(webResponse)

}
