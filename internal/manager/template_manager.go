package manager

type TemplateManager struct{}

type TemplateManagerConfig struct{}

func NewTemplateManager(config *TemplateManagerConfig) *TemplateManager {
	return &TemplateManager{}
}

func (tm *TemplateManager) FetchAll() {
	// ! TODO: call this API endpoint
	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account/message_templates/
}

func (tm *TemplateManager) Fetch(Id string) {
	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-hsm/
}

func (tm *TemplateManager) Create() {
	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account/message_templates/#:~:text=You%20can%20make%20a%20POST%20request%20to%20message_templates%20edge%20from%20the%20following%20paths%3A
}

func (tm *TemplateManager) Delete(id string) {
	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account/message_templates/#:~:text=on%20this%20endpoint.-,Deleting,-You%20can%20dissociate
}

func (tm *TemplateManager) Update() {
	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-hsm/#:~:text=2.0%20Access%20Token-,Updating,-You%20can%20update
}
