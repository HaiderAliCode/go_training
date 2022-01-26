package common

import "fmt"

type Student struct {
	ID                  int
	FirstName, LastName string
}

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

type College struct {
	database map[int]Student
}

func (c *College) Add(payload Student, reply *Student) error {
	if _, ok := c.database[payload.ID]; ok {
		return fmt.Errorf("error occured while adding user in database")
	}

	c.database[payload.ID] = payload
	*reply = payload
	return nil
}

func (c *College) Get(payload int, reply *Student) error {
	result, ok := c.database[payload]

	if !ok {
		fmt.Println("cannot fetch user with this specific id")
	}

	*reply = result
	return nil
}

func NewCollege() *College {
	return &College{
		database: make(map[int]Student),
	}
}
