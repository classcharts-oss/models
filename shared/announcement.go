package shared

type Announcement struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`

	SchoolName  string `json:"school_name"`
	TeacherName string `json:"teacher_name"`
	SchoolLogo  string `json:"school_logo"`

	Sticky YesNoBool `json:"sticky"`
	State  string    `json:"state"` // viewed, new

	Timestamp string `json:"timestamp"`

	Attachments []AnnouncementAttachment `json:"attachments"`

	CommentVisibility string `json:"comment_visibility"`

	AllowComments    YesNoBool `json:"allow_comments"`
	AllowReactions   YesNoBool `json:"allow_reactions"`
	AllowConsent     YesNoBool `json:"allow_consent"`
	PriorityPinned   YesNoBool `json:"priority_pinned"`
	RequiresConsent  YesNoBool `json:"requires_consent"`
	CanChangeConsent bool      `json:"can_change_consent"`

	Consent       Consent        `json:"consent"`
	PupilConsents []PupilConsent `json:"pupil_consents"`
}

type AnnouncementAttachment struct {
	Id       int    `json:"id"`
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

type Consent struct {
	Id           int       `json:"id"`
	ConsentGiven YesNoBool `json:"consent_given"`
	Comment      string    `json:"comment"`
	ParentName   string    `json:"parent_name"`
}

type PupilConsent struct {
	Pupil            Object  `json:"pupil"`
	CanChangeConsent bool    `json:"can_change_consent"`
	Consent          Consent `json:"consent"`
}
