package librenms

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type DevicesList struct {
	Status  string   `json:"status"`
	Count   int      `json:"count"`
	Devices []Device `json:"devices"`
}

type Device struct {
	DeviceID                 int         `json:"device_id,omitempty"`
	Inserted                 string      `json:"inserted,omitempty"`
	Hostname                 string      `json:"hostname"`
	SysName                  string      `json:"sysName,omitempty"`
	IP                       string      `json:"ip,omitempty"`
	OverwriteIP              interface{} `json:"overwrite_ip,omitempty"`
	Community                string      `json:"community"`
	Authlevel                interface{} `json:"authlevel,omitempty"`
	Authname                 interface{} `json:"authname,omitempty"`
	Authpass                 interface{} `json:"authpass,omitempty"`
	Authalgo                 interface{} `json:"authalgo,omitempty"`
	Cryptopass               interface{} `json:"cryptopass,omitempty"`
	Cryptoalgo               interface{} `json:"cryptoalgo,omitempty"`
	Snmpver                  string      `json:"snmpver,omitempty"`
	Port                     int         `json:"port"`
	Transport                string      `json:"transport,omitempty"`
	Timeout                  interface{} `json:"timeout,omitempty"`
	Retries                  interface{} `json:"retries,omitempty"`
	SnmpDisable              int         `json:"snmp_disable,omitempty"`
	BgpLocalAs               interface{} `json:"bgpLocalAs,omitempty"`
	SysObjectID              string      `json:"sysObjectID,omitempty"`
	SysDescr                 string      `json:"sysDescr,omitempty"`
	SysContact               interface{} `json:"sysContact,omitempty"`
	Version                  string      `json:"version"`
	Hardware                 string      `json:"hardware,omitempty"`
	Features                 string      `json:"features,omitempty"`
	LocationID               int         `json:"location_id,omitempty"`
	Os                       string      `json:"os,omitempty"`
	Status                   bool        `json:"status,omitempty"`
	StatusReason             string      `json:"status_reason,omitempty"`
	Ignore                   int         `json:"ignore,omitempty"`
	Disabled                 int         `json:"disabled,omitempty"`
	Uptime                   int         `json:"uptime,omitempty"`
	AgentUptime              int         `json:"agent_uptime,omitempty"`
	LastPolled               string      `json:"last_polled,omitempty"`
	LastPollAttempted        interface{} `json:"last_poll_attempted,omitempty"`
	LastPolledTimetaken      float64     `json:"last_polled_timetaken,omitempty"`
	LastDiscoveredTimetaken  float64     `json:"last_discovered_timetaken,omitempty"`
	LastDiscovered           string      `json:"last_discovered,omitempty"`
	LastPing                 string      `json:"last_ping,omitempty"`
	LastPingTimetaken        float64     `json:"last_ping_timetaken,omitempty"`
	Purpose                  interface{} `json:"purpose,omitempty"`
	Type                     string      `json:"type,omitempty"`
	Serial                   interface{} `json:"serial,omitempty"`
	Icon                     string      `json:"icon,omitempty"`
	PollerGroup              int         `json:"poller_group,omitempty"`
	OverrideSysLocation      int         `json:"override_sysLocation,omitempty"`
	Notes                    interface{} `json:"notes,omitempty"`
	PortAssociationMode      int         `json:"port_association_mode,omitempty"`
	MaxDepth                 int         `json:"max_depth,omitempty"`
	DisableNotify            int         `json:"disable_notify,omitempty"`
	DependencyParentID       interface{} `json:"dependency_parent_id,omitempty"`
	DependencyParentHostname interface{} `json:"dependency_parent_hostname,omitempty"`
	Location                 string      `json:"location,omitempty"`
	Lat                      interface{} `json:"lat,omitempty"`
	Lng                      interface{} `json:"lng,omitempty"`
}

// GetDevice - Returns a specific device
func (c *Client) GetDevice(hostname string) (*Device, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v0/devices/%s", c.HostURL, hostname), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	devices := DevicesList{}
	err = json.Unmarshal(body, &devices)
	if err != nil {
		return nil, err
	}

	return &devices.Devices[0], nil
}

// GetDevices - Returns a list of devices
func (c *Client) GetDevices(hostname string) (*DevicesList, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v0/devices", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	devices := DevicesList{}
	err = json.Unmarshal(body, &devices)
	if err != nil {
		return nil, err
	}

	return &devices, nil
}

// CreateDevice - Create new device
func (c *Client) CreateDevice(device Device) (*Device, error) {
	rb, err := json.Marshal(device)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v0/devices", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

// UpdateDevice - Updates a device
func (c *Client) UpdateDevice(device Device) (*Device, error) {
	//This only handles changing the community string and snmp port at this point...
	rb := "{\"field\": [\"community\",\"port\"], \"data\": [\"" + device.Community + "\"," + strconv.Itoa(device.Port) + "]}"
	//This is how we could do multiple fields at once in the future
	//{"field": ["notes","purpose"], "data": ["This server should be kept online", "For serving web traffic"]}

	// fieldsToUpdate := []string{"Community", "Port"}

	// v := reflect.ValueOf(device)
	// typeOfDevice := v.Type()
	// var field []string
	// var data []string

	// for i := 0; i < v.NumField(); i++ {
	// 	if contains(fieldsToUpdate, typeOfDevice.Field(i).Name) {
	// 		field = append(field, typeOfDevice.Field(i).Name)
	// 		data = append(data, v.Field(i).String())
	// 	}
	// }

	// fmt.Println(field)
	// fmt.Println(data)

	// need to build out the body and choose which fields we'll update

	// rb, err := json.Marshal(device)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(rb)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/v0/devices/%s", c.HostURL, device.Hostname), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}
	return &device, nil
}

// DeleteDevice - Deletes a device
func (c *Client) DeleteDevice(hostname string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v0/devices/%s", c.HostURL, hostname), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	devices := DevicesList{}
	err = json.Unmarshal(body, &devices)
	if err != nil {
		return err
	}

	if devices.Status != "ok" {
		return errors.New(string(body))
	}

	return nil
}
