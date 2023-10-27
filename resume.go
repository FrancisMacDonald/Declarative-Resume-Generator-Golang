package main

type Resume struct {
	Name           string           `yaml:"name"`
	Email          string           `yaml:"email"`
	Phone          string           `yaml:"phone"`
	Address        Address          `yaml:"address"`
	Summary        string           `yaml:"summary"`
	Experience     []Experience     `yaml:"experience"`
	Education      []Education      `yaml:"education"`
	Skills         []string         `yaml:"skills"`
	Certifications []Certifications `yaml:"certifications"`
}

type Address struct {
	Street string `yaml:"street"`
	City   string `yaml:"city"`
	State  string `yaml:"state"`
	Zip    int    `yaml:"zip"`
}

type Experience struct {
	Company    string   `yaml:"company"`
	Position   string   `yaml:"position"`
	StartDate  string   `yaml:"start_date"`
	EndDate    string   `yaml:"end_date"`
	Highlights []string `yaml:"highlights"`
}

type Education struct {
	Institution string `yaml:"institution"`
	Degree      string `yaml:"degree"`
	StartDate   string `yaml:"start_date"`
	EndDate     string `yaml:"end_date"`
}

type Certifications struct {
	Name string `yaml:"name"`
	Date string `yaml:"date"`
}
