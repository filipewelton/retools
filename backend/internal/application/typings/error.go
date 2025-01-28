package typings

type Error struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Reason     string `json:"reason"`
}
