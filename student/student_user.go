package student

import (
	"github.com/classcharts-oss/models/shared"
	"strings"
)

type StudentUser struct {
	Id                           int     `json:"id"`
	Name                         string  `json:"name"`
	FirstName                    string  `json:"first_name"`
	LastName                     string  `json:"last_name"`
	AvatarURL                    string  `json:"avatar_url"`
	DisplayBehaviour             bool    `json:"display_behaviour"`
	DisplayParentBehaviour       bool    `json:"display_parent_behaviour"`
	DisplayHomework              bool    `json:"display_homework"`
	DisplayRewards               bool    `json:"display_rewards"`
	DisplayDetentions            bool    `json:"display_detentions"`
	DisplayReportCards           bool    `json:"display_report_cards"`
	DisplayClasses               bool    `json:"display_classes"`
	DisplayAnnouncements         bool    `json:"display_announcements"`
	DisplayAcademicReports       bool    `json:"display_academic_reports"`
	DisplayAttendance            bool    `json:"display_attendance"`
	DisplayAttendanceType        bool    `json:"display_attendance_type"`
	DisplayAttendancePercentage  bool    `json:"display_attendance_percentage"`
	DisplayActivity              bool    `json:"display_activity"`
	DisplayMentalHeath           bool    `json:"display_mental_health"`
	DisplayMentalHealthNoTracker bool    `json:"display_mental_health_no_tracker"`
	DisplayTimetable             bool    `json:"display_timetable"`
	IsDisabled                   bool    `json:"is_disabled"`
	DisplayTwoWayCommunications  bool    `json:"display_two_way_communications"`
	DisplayAbsences              bool    `json:"display_absences"`
	CanUploadAttachments         bool    `json:"can_upload_attachments"`
	DisplayEventBadges           bool    `json:"display_event_badges"`
	DisplayAvatars               bool    `json:"display_avatars"`
	DisplayConcernSubmission     bool    `json:"display_concern_submission"`
	DisplayCustomFields          bool    `json:"display_custom_fields"`
	PupilConcernsHelpText        string  `json:"pupil_concerns_help_text"`
	AllowPupilsAddTimetableNotes bool    `json:"allow_pupils_add_timetable_notes"`
	DetentionAliasPluralUC       string  `json:"detention_alias_plural_uc"`
	AnnouncementsCount           int     `json:"announcements_count"`
	MessagesCount                int     `json:"messages_count"`
	PusherChannelName            string  `json:"pusher_channel_name"`
	HasBirthday                  bool    `json:"has_birthday"`
	HasNewSurvey                 bool    `json:"has_new_survery"`
	SurveyId                     *string `json:"survey_id"`
}

type DBStudentUser struct {
	StudentUser

	SchoolId int `json:"school_id"`

	Code string  `json:"code"`
	DOB  *string `json:"dob"`

	Concerns   []string          `json:"concerns"`
	Activities []shared.Activity `json:"activities"`
}

func NewUser(id int, name string, avatarUrl string) StudentUser {
	nameParts := strings.Split(name, " ")

	return StudentUser{
		Id:                           id,
		Name:                         name,
		FirstName:                    nameParts[0],
		LastName:                     nameParts[len(nameParts)-1],
		AvatarURL:                    avatarUrl,
		DisplayBehaviour:             true,
		DisplayParentBehaviour:       false,
		DisplayHomework:              true,
		DisplayRewards:               true,
		DisplayDetentions:            true,
		DisplayReportCards:           true,
		DisplayClasses:               true,
		DisplayAnnouncements:         true,
		DisplayAcademicReports:       true,
		DisplayAttendance:            true,
		DisplayAttendanceType:        true,
		DisplayAttendancePercentage:  true,
		DisplayActivity:              true,
		DisplayMentalHeath:           true,
		DisplayMentalHealthNoTracker: true,
		DisplayTimetable:             true,
		IsDisabled:                   false,
		DisplayTwoWayCommunications:  true,
		DisplayAbsences:              true,
		CanUploadAttachments:         true,
		DisplayEventBadges:           true,
		DisplayAvatars:               true,
		DisplayConcernSubmission:     true,
		DisplayCustomFields:          true,
		PupilConcernsHelpText:        "Please report any concerns you have.",
		AllowPupilsAddTimetableNotes: true,
		DetentionAliasPluralUC:       "Detentions",
		AnnouncementsCount:           0,
		MessagesCount:                0,
		PusherChannelName:            "Pusher_none",
		HasBirthday:                  true,
		HasNewSurvey:                 false,
		SurveyId:                     nil,
	}
}

func (dbu *DBStudentUser) GetAllPositive() []shared.Activity {
	var positive []shared.Activity

	for _, a := range dbu.Activities {
		if a.Polarity == "positive" {
			positive = append(positive, a)
		}
	}

	return positive
}

func (dbu *DBStudentUser) GetAllNegative() []shared.Activity {
	var negative []shared.Activity

	for _, a := range dbu.Activities {
		if a.Polarity == "negative" {
			negative = append(negative, a)
		}
	}

	return negative
}

func (dbu *DBStudentUser) GetTotalScore() int {
	total := 0

	for _, a := range dbu.Activities {
		total += a.Score
	}

	return total
}
