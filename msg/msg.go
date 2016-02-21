package msg

import (
	"user"
)

// #################### In ####################
type In struct {
	Sender user.Login
}

// #################### InChat ####################
type InChat struct {
	In
	Text string
}

// #################### OutChat ####################
type OutChat struct {
	SenderEmail string `json:"senderEmail"`
	Text        string `json:"text"`
}
