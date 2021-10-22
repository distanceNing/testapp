package statemachine

type Storage interface {
	Get(key string) error
	Set() error
}
