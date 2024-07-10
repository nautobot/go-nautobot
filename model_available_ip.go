/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-nautobot

// Representation of an IP address which does not exist in the database.
type AvailableIp struct {
	IpVersion int32 `json:"ip_version"`
	Address string `json:"address"`
}
