package messagebird

type Secret struct {
	Key      string            `json:"key"`
	Channels map[string]string `json:"channels"`
}
