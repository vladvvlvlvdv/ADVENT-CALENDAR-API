package repository

type (
	Attachment struct {
		ID    uint   `json:"id"`
		Label string `json:"label"`
		URL   string `json:"url"`
		Type  string `json:"type"`
		DayID uint   `json:"-"`
	}
)
