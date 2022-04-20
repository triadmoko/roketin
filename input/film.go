package input

type Film struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Duration    string `json:"duration" binding:"required"`
	Artist      string `json:"artist" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
	Filename    string `json:"filename" binding:"required"`
}
