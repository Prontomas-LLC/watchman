/*
 * Watchman API
 *
 * Moov Watchman offers download, parse, and search functions over numerous U.S. trade sanction lists for complying with regional laws. Also included is a web UI and async webhook notification service to initiate processes on remote systems.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// ForeignSanctionsEvader struct for ForeignSanctionsEvader
type ForeignSanctionsEvader struct {
	EntityID      string   `json:"entityID,omitempty"`
	EntityNumber  string   `json:"entityNumber,omitempty"`
	Type          string   `json:"type,omitempty"`
	Programs      []string `json:"programs,omitempty"`
	Name          string   `json:"name,omitempty"`
	Addresses     []string `json:"addresses,omitempty"`
	SourceListURL string   `json:"sourceListURL,omitempty"`
	Citizenships  string   `json:"citizenships,omitempty"`
	DatesOfBirth  string   `json:"datesOfBirth,omitempty"`
	SourceInfoURL string   `json:"sourceInfoURL,omitempty"`
	IDs           []string `json:"IDs,omitempty"`
}
