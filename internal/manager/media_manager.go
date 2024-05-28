package manager

type MediaManager struct {
	requester RequestClient
}

func NewMediaManager(requester RequestClient) *MediaManager {
	return &MediaManager{
		requester: requester,
	}
}

func (mm *MediaManager) UploadMedia() {
	// upload media
}

func (mm *MediaManager) GetMediaUrlById() {

}

func (mm *MediaManager) GetMediaIdByUrl() {

}
