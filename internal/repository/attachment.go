package repository

type (
	Attachment struct {
		Model
		Label string `json:"label"`
		URL   string `json:"url"`
		Type  string `json:"type"`
		DayID uint
	}
)
