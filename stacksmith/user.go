package stacksmith

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// UserService provides methods for accessing Stacksmith User API endpoints.
type UserService struct {
	sling *sling.Sling
}

func newUserService(sling *sling.Sling) *UserService {
	return &UserService{
		sling: sling.Path("user/"),
	}
}

// EmailNotifications ...
type EmailNotifications struct {
	EmailNotificationsEnabled bool `json:"email_notifications_enabled"`
}

// SlackChannels ...
type SlackChannels struct {
	TotalEntries int       `json:"total_entries"`
	TotalPAges   int       `json:"total_pages"`
	Items        []Channel `json:"items"`
}

// Channel ...
type Channel struct {
	ID           string `json:"id"`
	SlackChannel string `json:"slack_channel"`
}

// UpdateNotifications Update your email notification settings
// https://stacksmith.bitnami.com/api/v1/#!/User/patch_user
func (s *UserService) UpdateNotifications(params *EmailNotifications) (*EmailNotifications, *http.Response, error) {
	status := new(EmailNotifications)
	apiError := new(APIError)
	resp, err := s.sling.New().Patch("").BodyJSON(params).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// ListSlackChannels List all slack channels you have added integrations to.
// https://stacksmith.bitnami.com/api/v1/#!/User/get_user_slack_channels
func (s *UserService) ListSlackChannels(params *PaginationParams) (*SlackChannels, *http.Response, error) {
	slackChannels := new(SlackChannels)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("slack_channels").QueryStruct(params).Receive(slackChannels, apiError)
	return slackChannels, resp, relevantError(err, *apiError)
}

// RemoveSlackChannel Remove a Slack channel integration.
// https://stacksmith.bitnami.com/api/v1/#!/User/delete_user_slack_channels_id
func (s *UserService) RemoveSlackChannel(slackChannelID string) (*StatusDeletion, *http.Response, error) {
	status := new(StatusDeletion)
	apiError := new(APIError)
	path := fmt.Sprintf("slack_channels/%s", slackChannelID)
	resp, err := s.sling.New().Delete(path).Receive(status, apiError)
	return status, resp, relevantError(err, *apiError)
}

// TestSlackIntegration Send a test notification to a Slack channel.
// https://stacksmith.bitnami.com/api/v1/#!/User/post_user_slack_channels_id_test
func (s *UserService) TestSlackIntegration(slackChannelID string) (*Channel, *http.Response, error) {
	channel := new(Channel)
	apiError := new(APIError)
	path := fmt.Sprintf("slack_channels/%s/test", slackChannelID)
	resp, err := s.sling.New().Post(path).Receive(channel, apiError)
	return channel, resp, relevantError(err, *apiError)
}
