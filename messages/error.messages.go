package messages

type Error struct {
	RunId   string
	Message string
}

func (e *Error) ErrorJSON() []byte {
	return []byte(`{"run_id": "` + e.RunId + `", "message": "` + e.Message + `"}`)
}
