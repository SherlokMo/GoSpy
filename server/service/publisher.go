package service

type Pub interface {
	register(subscriber Observer)
	deregister(subscriber Observer)
	notifyAll()
}
