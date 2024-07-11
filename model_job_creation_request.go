/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nautobot
import (
	"time"
)

// Nested serializer specifically for use with `JobInputSerializer.schedule`.  We don't use `WritableNestedSerializer` here because this is not used to look up an existing `ScheduledJob`, but instead used to specify parameters for creating one.
type JobCreationRequest struct {
	Name string `json:"name,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	Interval *JobExecutionTypeIntervalChoices `json:"interval"`
	// Cronjob syntax string for custom scheduling
	Crontab string `json:"crontab,omitempty"`
}
