package models

type Model interface {
	ParseJson() interface{}
}
