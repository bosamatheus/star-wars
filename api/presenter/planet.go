package presenter

type Planet struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Films   int    `json:"films"`
	Terrain string `json:"terrain"`
}
