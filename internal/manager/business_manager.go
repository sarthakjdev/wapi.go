package manager

type BusinessManager struct {
}

type BusinessManagerConfig struct{}

func NewBusinessManager(config *BusinessManagerConfig) *BusinessManager {
	return &BusinessManager{}
}

func (bm *BusinessManager) Fetch() {
	// ! TODO: need to call this API here
	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account
}
