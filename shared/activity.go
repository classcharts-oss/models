package shared

type ActivityType string
type ActivityPolarity string

const (
	Event           = ActivityType("event")
	AttendanceEvent = ActivityType("attendance_event")
	DT              = ActivityType("detention")
	Notice          = ActivityType("notice")
)

const (
	Positive = ActivityPolarity("positive")
	Negative = ActivityPolarity("negative")
	Blank    = ActivityPolarity("blank")
)

type ActivityStyle struct {
	BorderColour *string `json:"border_color"`
	CustomClass  *string `json:"custom_class"`
}

type Activity struct {
	Id       int              `json:"id"`
	Type     ActivityType     `json:"type"`
	Polarity ActivityPolarity `json:"polarity"`
	Reason   string           `json:"reason"`
	Score    int              `json:"score"`

	Timestamp string `json:"timestamp"`

	Style ActivityStyle `json:"style"`

	PupilName   string  `json:"pupil_name"`
	LessonName  *string `json:"lesson_name"`
	TeacherName *string `json:"teacher_name"`
	RoomName    *string `json:"room_name"`

	Note       *string `json:"note"`
	ParentNote *string `json:"parent_note"`

	DetentionDate     *string `json:"detention_date"`
	DetentionTime     *string `json:"detention_time"`
	DetentionLocation *string `json:"detention_location"`
	DetentionType     *string `json:"detention_type"`
}
