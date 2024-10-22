package school

// Class represents a class in a school.
type Class struct {
	Id int `json:"id"` // The ID of the class. (e.g. 858122)

	Name    string `json:"short_name"`  // The short name of the class. (e.g. "10C/En1" for Year 10, Band C, English, Set 1)
	Subject string `json:"lesson_name"` // The subject taught in this class. (e.g. "English" or "Maths")
}

// NewClass creates a new class.
func NewClass(id int, name, subject string) Class {
	return Class{
		Id: id,

		Name:    name,
		Subject: subject,
	}
}
