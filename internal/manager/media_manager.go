package manager

import (
	"strings"

	"github.com/sarthakjdev/wapi.go/internal/request_client"
)

// MediaManager is responsible for managing media related operations.
type MediaManager struct {
	requester request_client.RequestClient
}

// NewMediaManager creates a new instance of MediaManager.
func NewMediaManager(requester request_client.RequestClient) *MediaManager {
	return &MediaManager{
		requester: requester,
	}
}

// GetMediaUrlById retrieves the media URL by its ID.
func (mm *MediaManager) GetMediaUrlById(id string) {
	apiRequest := mm.requester.NewApiRequest(strings.Join([]string{"media", id}, "/"), "GET")
	apiRequest.Execute()

}

// GetMediaIdByUrl retrieves the media ID by its URL.
func (mm *MediaManager) GetMediaIdByUrl(id string) {
}
