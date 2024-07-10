/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-nautobot

type BulkWritableCableRequestSource struct {
	Label string `json:"label,omitempty"`
	ObjectType string `json:"object_type,omitempty"`
	Objects []BulkWritableCableRequestSourceObjects `json:"objects,omitempty"`
}
