package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// BgpProfile bgp profile
// swagger:model BgpProfile
type BgpProfile struct {

	// Set the community attribute - 'internet', 'local-AS', 'no-advertise', 'no-export', <AS> <Val> One or more of these can be configured with 1 <= AS,Val <= 65535. Field introduced in 17.1.2.
	Community []string `json:"community,omitempty"`

	// Hold time for Peers. Allowed values are 3-7200.
	HoldTime int32 `json:"hold_time,omitempty"`

	// BGP peer type.
	// Required: true
	Ibgp bool `json:"ibgp"`

	// Keepalive interval for Peers. Allowed values are 0-3600.
	KeepaliveInterval int32 `json:"keepalive_interval,omitempty"`

	// Local Autonomous System ID. Allowed values are 1-4294967295.
	// Required: true
	LocalAs int32 `json:"local_as"`

	// BGP Peers.
	Peers []*BgpPeer `json:"peers,omitempty"`

	// Send community attribute to all peers. Field introduced in 17.1.2.
	SendCommunity bool `json:"send_community,omitempty"`
}
