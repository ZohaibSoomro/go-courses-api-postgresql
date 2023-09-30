package models

import (
	"log"
	"strconv"

	"gorm.io/gorm"
)

type MyDb struct {
	DB *gorm.DB
}

func (m *MyDb) createCourse(c *Course) {
	d := m.DB.Create(c)
	if d.Error != nil {
		log.Println(d.Error)
		return
	}
	log.Println("Course inserted with id", c.ID)
}

func (m *MyDb) getAllCourses() []Course {
	c := []Course{}
	d := m.DB.Find(&c)
	if d.Error != nil {
		log.Println(d.Error)
		return nil
	}
	return c
}

func (m *MyDb) getById(id string) *Course {
	c := &Course{}
	Id, _ := strconv.Atoi(id)
	d := m.DB.Where("ID=?", Id).Find(c)
	if d.Error != nil {
		log.Println(d.Error)
		return nil
	}
	if c.Title != "" {
		log.Println("Course found with id", c.ID)
		return c
	}
	return nil
}

func (m *MyDb) deleteCourse(c *Course) {

	d := m.DB.Delete(c)
	if d.Error != nil {
		log.Println(d.Error)
		return
	}
	log.Println("Course deleted with id", c.ID)
}

func (m *MyDb) updateCourse(c *Course) {
	d := m.DB.Updates(c)
	if d.Error != nil {
		log.Println(d.Error)
		return

	}
	log.Println("Course updated with id", c.ID)
}
