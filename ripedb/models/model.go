package models

type Model interface {
	Class() string
	Key() string
}
