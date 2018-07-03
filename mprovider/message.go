package mprovider

type MessageProvider struct {
	msg string
}

func New() *MessageProvider {
	return &MessageProvider{"burn in hell!"}
}

func (m *MessageProvider) Message() string {
	return m.msg
}
