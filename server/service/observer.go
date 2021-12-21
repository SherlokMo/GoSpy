package service

type Observer interface {
	update(payload map[string]interface{})
}
