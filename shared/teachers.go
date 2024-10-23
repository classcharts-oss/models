package shared

type Teacher struct {
	Title     string `json:"title"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// FAKE INFO:
//	Science - Miss T Aritus (10C/Sc1)
//	Maths - Mr M Smith (10C/Ma1)
//	Computer Science - Miss C Yite (10C/Cs1)
//	English - Miss S Prile (10C/En1)

//	General/Attendance/Other - Mr T Pott - (No class, no room)
//	Year Leaders Office - Detention Location
