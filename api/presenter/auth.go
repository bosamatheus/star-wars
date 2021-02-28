package presenter

type Auth struct {
	Token      string `json:"token"`
	Authorized bool   `json:"authorized"`
}
