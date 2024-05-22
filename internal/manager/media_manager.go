package manager

type MediaManager struct {
	requester requestClient
}

func NewMediaManager(requester requestClient) *MediaManager {
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
