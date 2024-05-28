package manager

import requestclient "github.com/sarthakjdev/wapi.go/internal/request_client"

type MediaManager struct {
	requester requestclient.RequestClient
}

func NewMediaManager(requester requestclient.RequestClient) *MediaManager {
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
