package types

// Message type express a message sent or received.
type Message string

// UserID type express a ID of opponent or me in chat.
type UserID string

// Chatroom type keeps the connections to user.
type Chatroom struct {
	In  chan Message
	Out chan Message
}

// IsTalked express whether the topic talked to user or not. If the value is false, the framework will try another topic.
type IsTalked bool
