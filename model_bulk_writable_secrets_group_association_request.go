/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nautobot

// Serializer for `SecretsGroupAssociation` objects.
type BulkWritableSecretsGroupAssociationRequest struct {
	Id string `json:"id"`
	AccessType *AccessTypeEnum `json:"access_type"`
	SecretType *SecretTypeEnum `json:"secret_type"`
	SecretsGroup *BulkWritableCableRequestStatus `json:"secrets_group"`
	Secret *BulkWritableCableRequestStatus `json:"secret"`
}
