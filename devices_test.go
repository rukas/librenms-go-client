package librenms

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDevice(t *testing.T) {
	apiKey := os.Getenv("LIBRENMS_API_KEY")
	c, err := NewClient(nil, &apiKey)
	assert.Nil(t, err, "expecting nil error for calling newClient")

	var d Device
	d.Hostname = "localhost"
	d.Community = "community"
	d.Version = "v2c"
	d.Port = 161

	res, err := c.CreateDevice(d)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	assert.Equal(t, d.Hostname, res.Hostname, "expecting correct hostname")
}

func TestUpdateDevice(t *testing.T) {
	apiKey := os.Getenv("LIBRENMS_API_KEY")
	c, err := NewClient(nil, &apiKey)
	assert.Nil(t, err, "expecting nil error for calling newClient")

	var d Device
	d.Hostname = "localhost"
	d.Community = "test"
	d.Port = 166

	res, err := c.UpdateDevice(d)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	assert.Equal(t, d.Hostname, res.Hostname, "expecting correct hostname")
}

func TestGetDevice(t *testing.T) {
	apiKey := os.Getenv("LIBRENMS_API_KEY")
	c, err := NewClient(nil, &apiKey)
	assert.Nil(t, err, "expecting nil error for calling newClient")

	hostname := "localhost"

	res, err := c.GetDevice(hostname)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	assert.Equal(t, hostname, res.Hostname, "expecting correct hostname")
}

func TestDeleteDevice(t *testing.T) {
	apiKey := os.Getenv("LIBRENMS_API_KEY")
	c, err := NewClient(nil, &apiKey)
	assert.Nil(t, err, "expecting nil error for calling newClient")

	hostname := "localhost"

	a := c.DeleteDevice(hostname)

	assert.Nil(t, a, "expecting nil error")
}
