package models

// structures for all project

type Employee struct {
	Name   string
	Skills map[string]interface{}
}

type Department struct {
	Name           string
	RequiredSkills map[string]interface{}
}
