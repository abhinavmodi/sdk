package clients

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

import (
	"github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/sdk/go/session"
)

// AlertEmailConfigClient is a client for avi AlertEmailConfig resource
type AlertEmailConfigClient struct {
	aviSession *session.AviSession
}

// NewAlertEmailConfigClient creates a new client for AlertEmailConfig resource
func NewAlertEmailConfigClient(aviSession *session.AviSession) *AlertEmailConfigClient {
	return &AlertEmailConfigClient{aviSession: aviSession}
}

func (client *AlertEmailConfigClient) getAPIPath(uuid string) string {
	path := "api/alertemailconfig"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of AlertEmailConfig objects
func (client *AlertEmailConfigClient) GetAll() ([]*models.AlertEmailConfig, error) {
	var plist []*models.AlertEmailConfig
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist)
	return plist, err
}

// Get an existing AlertEmailConfig by uuid
func (client *AlertEmailConfigClient) Get(uuid string) (*models.AlertEmailConfig, error) {
	var obj *models.AlertEmailConfig
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj)
	return obj, err
}

// GetByName - Get an existing AlertEmailConfig by name
func (client *AlertEmailConfigClient) GetByName(name string) (*models.AlertEmailConfig, error) {
	var obj *models.AlertEmailConfig
	err := client.aviSession.GetObjectByName("alertemailconfig", name, &obj)
	return obj, err
}

// Create a new AlertEmailConfig object
func (client *AlertEmailConfigClient) Create(obj *models.AlertEmailConfig) (*models.AlertEmailConfig, error) {
	var robj *models.AlertEmailConfig
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj)
	return robj, err
}

// Update an existing AlertEmailConfig object
func (client *AlertEmailConfigClient) Update(obj *models.AlertEmailConfig) (*models.AlertEmailConfig, error) {
	var robj *models.AlertEmailConfig
	path := client.getAPIPath(obj.UUID)
	err := client.aviSession.Put(path, obj, &robj)
	return robj, err
}

// Delete an existing AlertEmailConfig object with a given UUID
func (client *AlertEmailConfigClient) Delete(uuid string) error {
	return client.aviSession.Delete(client.getAPIPath(uuid))
}

// DeleteByName - Delete an existing AlertEmailConfig object with a given name
func (client *AlertEmailConfigClient) DeleteByName(name string) error {
	res, err := client.GetByName(name)
	if err != nil {
		return err
	}
	return client.Delete(res.UUID)
}
