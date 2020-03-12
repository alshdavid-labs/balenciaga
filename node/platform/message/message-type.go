package message

// MessageType describes accepted messages
var MessageType = struct {
	Register          string
	Subscribe         string
	Subscribed        string
	Authorize         string
	MoveConnection    string
	Terminate         string
	Challenge         string
	ChallengeResponse string
}{
	Register:          "REGISTER",
	Subscribe:         "SUBSCRIBE",
	Subscribed:        "SUBSCRIBED",
	Authorize:         "AUTHORIZE",
	MoveConnection:    "MOVE_CONNECTION",
	Terminate:         "TERMINATE",
	Challenge:         "CHALLENGE",
	ChallengeResponse: "CHALLENGE_RESPONSE",
}
