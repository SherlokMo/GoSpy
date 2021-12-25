package service

type Observer interface {
	Update(payload map[string]interface{})
}
