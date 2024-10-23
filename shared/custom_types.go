package shared

type YesNoBool string
type AttachmentScanStatus string

const (
	Yes YesNoBool = "yes"
	No  YesNoBool = "no"
)

const (
	Pass AttachmentScanStatus = "pass"
)

func (ynb YesNoBool) AsBool() bool {
	return ynb == Yes
}
