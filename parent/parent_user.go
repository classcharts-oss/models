package parent

import (
	"github.com/classcharts-oss/models/school"
	"github.com/classcharts-oss/models/student"
)

// ParentUser represents a parent.
type ParentUser struct {
	Id              int    `json:"id"`              // The ID of the parent. (e.g. 187633)
	Name            string `json:"name"`            // The name of the parent. (e.g. "Jane Doe")
	Email           string `json:"email"`           // The email of the parent. (e.g. "jane@example.com")
	Language        string `json:"language"`        // The language of the parent user. (e.g. "en")
	IsEmailVerified bool   `json:"isEmailVerified"` // Whether the email of the parent's email is verified.
}

type Pupil struct {
	student.StudentUser

	SchoolName string `json:"school_name"`
	SchoolLogo string `json:"school_logo"`

	Timezone string `json:"timezone"`

	DisplayCovidTests   bool `json:"display_covid_tests"`
	CanRecordCovidTests bool `json:"can_record_covid_tests"`

	DetentionYesCount      int `json:"detention_yes_count"`
	DetentionNoCount       int `json:"detention_no_count"`
	DetentionPendingCount  int `json:"detention_pending_count"`
	DetentionUpscaledCount int `json:"detention_upscaled_count"`

	HomeworkTodoCount         int `json:"homework_todo_count"`
	HomeworkLateCount         int `json:"homework_late_count"`
	HomeworkNotCompletedCount int `json:"homework_not_completed_count"`
	HomeworkExcusedCount      int `json:"homework_excused_count"`
	HomeworkCompletedCount    int `json:"homework_completed_count"`
	HomeworkSubmittedCount    int `json:"homework_submitted_count"`
}

func NewParentUser(id int, name, email, language string, isEmailVerified bool) ParentUser {
	return ParentUser{
		Id:              id,
		Name:            name,
		Email:           email,
		Language:        language,
		IsEmailVerified: isEmailVerified,
	}
}

func NewPupilsFromStudentsAndSchool(
	students []student.StudentUser, school school.School) []Pupil {
	pupils := make([]Pupil, len(students))

	for i, student := range students {
		pupils[i] = Pupil{
			StudentUser: student,

			SchoolName: school.Name,
			SchoolLogo: school.Logo,

			Timezone: "Europe/London",
		}
	}

	return pupils
}
