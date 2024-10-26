package school

import "github.com/classcharts-oss/models/shared"

// School represents a school, of course.
type School struct {
	Id   int    `json:"id"`   // The ID of the school. (e.g. 968093)
	Name string `json:"name"` // The name of the school. (e.g. "Springfield Secondary School")

	Logo string `json:"logo"` // The URL of the school's logo. (e.g. "https://placehold.co/360")

	Classes       []Class               `json:"classes"`       // The classes in the school.
	Announcements []shared.Announcement `json:"announcements"` // The announcements in the school.
}

func NewSchool(id int, name, logo string) School {
	return School{
		Id:   id,
		Name: name,

		Logo: logo,

		Classes:       []Class{},
		Announcements: []shared.Announcement{},
	}
}
