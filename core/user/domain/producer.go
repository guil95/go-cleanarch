package user

type Producer interface {
	Produce(message string, topic string)
}
