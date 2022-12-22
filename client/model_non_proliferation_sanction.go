/*
 * Watchman API
 *
 * Moov Watchman offers download, parse, and search functions over numerous U.S. trade sanction lists for complying with regional laws. Also included is a web UI and async webhook notification service to initiate processes on remote systems.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// NonProliferationSanction struct for NonProliferationSanction
type NonProliferationSanction struct {
	EntityID              string   `json:"entityID,omitempty"`
	Programs              []string `json:"programs,omitempty"`
	Name                  string   `json:"name,omitempty"`
	FederalRegisterNotice string   `json:"federalRegisterNotice,omitempty"`
	StartDate             string   `json:"startDate,omitempty"`
	Remarks               []string `json:"remarks,omitempty"`
	SourceListURL         string   `json:"sourceListURL,omitempty"`
	AlternateNames        []string `json:"alternateNames,omitempty"`
	SourceInfoURL         string   `json:"sourceInfoURL,omitempty"`
	// Match percentage of search query
	Match float32 `json:"match,omitempty"`
}
