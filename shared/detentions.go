package shared

type Detention struct {
	Id       int         `json:"id"`
	Attended string      `json:"attended"`
	Date     string      `json:"date"`
	Length   int         `json:"length"`
	Location string      `json:"location"`
	Notes    interface{} `json:"notes"`
	Time     string      `json:"time"`
}
