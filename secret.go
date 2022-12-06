package messagebird

type Secret struct {
	Key      string             `json:"key"`
	Channels map[string]Channel `json:"channels"`
}

type Channel struct {
	ID        string `json:"id"`
	Namespace string `json:"namespace,omitempty"`
}
