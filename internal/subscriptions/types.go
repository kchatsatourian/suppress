package subscriptions

type Subscription struct {
	Channels []int64 `json:"channels"`
	Name     string  `json:"name"`
}
