package messages

type Client struct {
	Type            ClientType
	ChallengeAnswer int64
}

type Server struct {
	Type ServerType
	Data string
}
