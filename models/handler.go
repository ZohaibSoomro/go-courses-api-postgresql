package models

import (
	"github.com/gofiber/fiber/v2"
)

func (m *MyDb) Create(ct *fiber.Ctx) error {
	c := &Course{}
	ct.BodyParser(c)
	m.createCourse(c)
	return ct.Status(fiber.StatusOK).JSON(&fiber.Map{
		"msg":    "creation sucess",
		"course": c,
	})
}

func (m *MyDb) GetAll(ct *fiber.Ctx) error {

	c := m.getAllCourses()
	return ct.Status(fiber.StatusOK).JSON(&fiber.Map{
		"msg":    "success",
		"course": c,
	})
}

func (m *MyDb) GetById(ct *fiber.Ctx) error {
	id := ct.Params("id")

	c := m.getById(id)
	status := "id found"
	if c.Title == "" {
		status = "id not found"
	}
	return ct.Status(fiber.StatusOK).JSON(&fiber.Map{
		"msg":    status,
		"course": c,
	})
}

func (m *MyDb) Delete(ct *fiber.Ctx) error {
	id := ct.Params("id")
	c := m.getById(id)

	if c.Title == "" {
		return ct.Status(fiber.StatusOK).JSON(&fiber.Map{
			"msg":    "id not found",
			"course": c,
		})
	}

	m.deleteCourse(c)
	return ct.Status(fiber.StatusOK).JSON(&fiber.Map{
		"msg":       "deletion sucess",
		"course id": id,
	})
}

func (m *MyDb) Update(ct *fiber.Ctx) error {
	c := &Course{}
	ct.BodyParser(c)

	id := ct.Params("id")
	c2 := m.getById(id)

	if c2.Title == "" {
		return ct.Status(fiber.StatusOK).JSON(&fiber.Map{
			"msg":    "id not found",
			"course": c2,
		})
	}
	c.ID = c2.ID
	m.updateCourse(c)
	return ct.Status(fiber.StatusOK).JSON(&fiber.Map{
		"msg":    "update sucess",
		"course": c,
	})
}

func (m *MyDb) SetUpRoutes(app *fiber.App) {
	app.Get("/courses", m.GetAll)
	app.Get("/courses/:id", m.GetById)
	app.Post("/courses", m.Create)
	app.Put("/courses/:id", m.Update)
	app.Delete("/courses/:id", m.Delete)
}
