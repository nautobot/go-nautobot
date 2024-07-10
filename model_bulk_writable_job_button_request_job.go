/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-nautobot

// Job this button will run
type BulkWritableJobButtonRequestJob struct {
	Id *OneOfBulkWritableJobButtonRequestJobId `json:"id,omitempty"`
	ObjectType string `json:"object_type,omitempty"`
	Url string `json:"url,omitempty"`
}
