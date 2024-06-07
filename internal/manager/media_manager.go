package manager

import requestclient "github.com/sarthakjdev/wapi.go/internal/request_client"

// MediaManager is responsible for managing media related operations.
type MediaManager struct {
	requester requestclient.RequestClient
}

// NewMediaManager creates a new instance of MediaManager.
func NewMediaManager(requester requestclient.RequestClient) *MediaManager {
	return &MediaManager{
		requester: requester,
	}
}

// GetMediaUrlById retrieves the media URL by its ID.
func (mm *MediaManager) GetMediaUrlById() {
}

// GetMediaIdByUrl retrieves the media ID by its URL.
func (mm *MediaManager) GetMediaIdByUrl() {
}
