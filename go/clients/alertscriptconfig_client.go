package clients

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

import (
	"github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/sdk/go/session"
)

// AlertScriptConfigClient is a client for avi AlertScriptConfig resource
type AlertScriptConfigClient struct {
	aviSession *session.AviSession
}

// NewAlertScriptConfigClient creates a new client for AlertScriptConfig resource
func NewAlertScriptConfigClient(aviSession *session.AviSession) *AlertScriptConfigClient {
	return &AlertScriptConfigClient{aviSession: aviSession}
}

func (client *AlertScriptConfigClient) getAPIPath(uuid string) string {
	path := "api/alertscriptconfig"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of AlertScriptConfig objects
func (client *AlertScriptConfigClient) GetAll() ([]*models.AlertScriptConfig, error) {
	var plist []*models.AlertScriptConfig
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist)
	return plist, err
}

// Get an existing AlertScriptConfig by uuid
func (client *AlertScriptConfigClient) Get(uuid string) (*models.AlertScriptConfig, error) {
	var obj *models.AlertScriptConfig
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj)
	return obj, err
}

// GetByName - Get an existing AlertScriptConfig by name
func (client *AlertScriptConfigClient) GetByName(name string) (*models.AlertScriptConfig, error) {
	var obj *models.AlertScriptConfig
	err := client.aviSession.GetObjectByName("alertscriptconfig", name, &obj)
	return obj, err
}

// Create a new AlertScriptConfig object
func (client *AlertScriptConfigClient) Create(obj *models.AlertScriptConfig) (*models.AlertScriptConfig, error) {
	var robj *models.AlertScriptConfig
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj)
	return robj, err
}

// Update an existing AlertScriptConfig object
func (client *AlertScriptConfigClient) Update(obj *models.AlertScriptConfig) (*models.AlertScriptConfig, error) {
	var robj *models.AlertScriptConfig
	path := client.getAPIPath(obj.UUID)
	err := client.aviSession.Put(path, obj, &robj)
	return robj, err
}

// Delete an existing AlertScriptConfig object with a given UUID
func (client *AlertScriptConfigClient) Delete(uuid string) error {
	return client.aviSession.Delete(client.getAPIPath(uuid))
}

// DeleteByName - Delete an existing AlertScriptConfig object with a given name
func (client *AlertScriptConfigClient) DeleteByName(name string) error {
	res, err := client.GetByName(name)
	if err != nil {
		return err
	}
	return client.Delete(res.UUID)
}
