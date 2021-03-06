package clients

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

import (
	"github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/sdk/go/session"
)

// SSLProfileClient is a client for avi SSLProfile resource
type SSLProfileClient struct {
	aviSession *session.AviSession
}

// NewSSLProfileClient creates a new client for SSLProfile resource
func NewSSLProfileClient(aviSession *session.AviSession) *SSLProfileClient {
	return &SSLProfileClient{aviSession: aviSession}
}

func (client *SSLProfileClient) getAPIPath(uuid string) string {
	path := "api/sslprofile"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of SSLProfile objects
func (client *SSLProfileClient) GetAll() ([]*models.SSLProfile, error) {
	var plist []*models.SSLProfile
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist)
	return plist, err
}

// Get an existing SSLProfile by uuid
func (client *SSLProfileClient) Get(uuid string) (*models.SSLProfile, error) {
	var obj *models.SSLProfile
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj)
	return obj, err
}

// GetByName - Get an existing SSLProfile by name
func (client *SSLProfileClient) GetByName(name string) (*models.SSLProfile, error) {
	var obj *models.SSLProfile
	err := client.aviSession.GetObjectByName("sslprofile", name, &obj)
	return obj, err
}

// Create a new SSLProfile object
func (client *SSLProfileClient) Create(obj *models.SSLProfile) (*models.SSLProfile, error) {
	var robj *models.SSLProfile
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj)
	return robj, err
}

// Update an existing SSLProfile object
func (client *SSLProfileClient) Update(obj *models.SSLProfile) (*models.SSLProfile, error) {
	var robj *models.SSLProfile
	path := client.getAPIPath(obj.UUID)
	err := client.aviSession.Put(path, obj, &robj)
	return robj, err
}

// Delete an existing SSLProfile object with a given UUID
func (client *SSLProfileClient) Delete(uuid string) error {
	return client.aviSession.Delete(client.getAPIPath(uuid))
}

// DeleteByName - Delete an existing SSLProfile object with a given name
func (client *SSLProfileClient) DeleteByName(name string) error {
	res, err := client.GetByName(name)
	if err != nil {
		return err
	}
	return client.Delete(res.UUID)
}
