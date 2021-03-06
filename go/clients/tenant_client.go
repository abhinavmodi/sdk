package clients

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

import (
	"github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/sdk/go/session"
)

// TenantClient is a client for avi Tenant resource
type TenantClient struct {
	aviSession *session.AviSession
}

// NewTenantClient creates a new client for Tenant resource
func NewTenantClient(aviSession *session.AviSession) *TenantClient {
	return &TenantClient{aviSession: aviSession}
}

func (client *TenantClient) getAPIPath(uuid string) string {
	path := "api/tenant"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of Tenant objects
func (client *TenantClient) GetAll() ([]*models.Tenant, error) {
	var plist []*models.Tenant
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist)
	return plist, err
}

// Get an existing Tenant by uuid
func (client *TenantClient) Get(uuid string) (*models.Tenant, error) {
	var obj *models.Tenant
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj)
	return obj, err
}

// GetByName - Get an existing Tenant by name
func (client *TenantClient) GetByName(name string) (*models.Tenant, error) {
	var obj *models.Tenant
	err := client.aviSession.GetObjectByName("tenant", name, &obj)
	return obj, err
}

// Create a new Tenant object
func (client *TenantClient) Create(obj *models.Tenant) (*models.Tenant, error) {
	var robj *models.Tenant
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj)
	return robj, err
}

// Update an existing Tenant object
func (client *TenantClient) Update(obj *models.Tenant) (*models.Tenant, error) {
	var robj *models.Tenant
	path := client.getAPIPath(obj.UUID)
	err := client.aviSession.Put(path, obj, &robj)
	return robj, err
}

// Delete an existing Tenant object with a given UUID
func (client *TenantClient) Delete(uuid string) error {
	return client.aviSession.Delete(client.getAPIPath(uuid))
}

// DeleteByName - Delete an existing Tenant object with a given name
func (client *TenantClient) DeleteByName(name string) error {
	res, err := client.GetByName(name)
	if err != nil {
		return err
	}
	return client.Delete(res.UUID)
}
