package capsule

type Capsule struct {
	content       string
	PublicContent string
}

func GetString() string {
	return "test123"
}

func (capsule *Capsule) GetContent() string {
	return capsule.content
}

func (capsule *Capsule) SetContent(content string) *Capsule {
	capsule.content = content
	return capsule
}

func (capsule *Capsule) GetPublicContent() string {
	return capsule.PublicContent
}
