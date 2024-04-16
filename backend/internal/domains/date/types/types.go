package types

const (
	TypeText = iota
	TypeSticker
)

type Date struct {
	User1    int       `json:"user1" db:"user1"`
	User2    int       `json:"user2" db:"user2"`
	Messages []Message `json:"messages" db:"messages"`
}

type Message struct {
	Author  bool   `json:"author"`
	Type    int    `json:"type"`
	Content string `json:"content"`
	Time    string `json:"time"`
}
