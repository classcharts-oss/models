package shared

type Homework struct {
	Id int `json:"id"`

	Lesson  string `json:"lesson"`
	Subject string `json:"subject"`
	Teacher string `json:"teacher"`

	HomeworkType string `json:"homework_type"`

	Title       string `json:"title"`
	MetaTitle   string `json:"meta_title"`
	Description string `json:"description"`

	IssueDate string `json:"issue_date"`
	DueDate   string `json:"due_date"`

	CompletionTimeUnit  string `json:"completion_time_unit"`
	CompletionTimeValue string `json:"completion_time_value"`

	PublishTime string `json:"publish_time"`

	Status HomeworkStatus `json:"status"`

	// TODO: Make proper types
	ProvidedLinks       []string `json:"validated_links"` // These 2 are named differently to what the API names them. The reason it is different is because it's not really clear on what is a student's attachment and a link/attachment that came with the Homework itself.
	ProvidedAttachments []string `json:"validated_attachments"`
}

type HomeworkStatus struct {
	Id int

	// TODO: Find out type of these 2 variables
	State        *string
	Mark         *string
	MarkRelative int

	Ticked YesNoBool

	AllowAttachments      bool
	AllowMarkingCompleted bool

	FirstSeenDate *string
	LastSeenDate  *string // Note: This is unused in prod. afaik

	// TODO: Make proper types
	StudentAttachments []string // As previously aforementioned, this is named differently purely for the benefit of using these models.
}
