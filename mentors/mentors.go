package mentors

type Mentor struct {
	Name string
}

func List() []Mentor {
	return []Mentor{
		{"Lucia"},
		{"Julio"},
		{"Marcus"},
	}
}
