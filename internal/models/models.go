package models

// structures for all project

type Employee struct {
	Name   string
	Email  string
	Skills []string
}

type Department struct {
	Name           string
	RequiredSkills []string
}
