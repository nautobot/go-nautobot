/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-nautobot
import (
	"time"
)

// Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type Job struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	// Dotted name of the Python module providing this job
	ModuleName string `json:"module_name"`
	// Name of the Python class providing this job
	JobClassName string `json:"job_class_name"`
	// Human-readable grouping that this job belongs to
	Grouping string `json:"grouping"`
	// Human-readable name of this job
	Name string `json:"name"`
	// Markdown formatting and a limited subset of HTML are supported
	Description string `json:"description,omitempty"`
	// Whether the Python module and class providing this job are presently installed and loadable
	Installed bool `json:"installed"`
	// Whether this job can be executed by users
	Enabled bool `json:"enabled,omitempty"`
	// Whether this job is a job hook receiver
	IsJobHookReceiver bool `json:"is_job_hook_receiver"`
	// Whether this job is a job button receiver
	IsJobButtonReceiver bool `json:"is_job_button_receiver"`
	// Whether this job contains sensitive variables
	HasSensitiveVariables bool `json:"has_sensitive_variables,omitempty"`
	// Whether the job requires approval from another user before running
	ApprovalRequired bool `json:"approval_required,omitempty"`
	// Whether the job defaults to not being shown in the UI
	Hidden bool `json:"hidden,omitempty"`
	// Whether the job defaults to running with dryrun argument set to true
	DryrunDefault bool `json:"dryrun_default,omitempty"`
	// Set to true if the job does not make any changes to the environment
	ReadOnly bool `json:"read_only"`
	// Maximum runtime in seconds before the job will receive a <code>SoftTimeLimitExceeded</code> exception.<br>Set to 0 to use Nautobot system default
	SoftTimeLimit float64 `json:"soft_time_limit,omitempty"`
	// Maximum runtime in seconds before the job will be forcibly terminated.<br>Set to 0 to use Nautobot system default
	TimeLimit float64 `json:"time_limit,omitempty"`
	// Comma separated list of task queues that this job can run on. A blank list will use the default queue
	TaskQueues map[string]Object `json:"task_queues,omitempty"`
	// If supported, allows the job to bypass approval when running with dryrun argument set to true
	SupportsDryrun bool `json:"supports_dryrun"`
	// If set, the configured grouping will remain even if the underlying Job source code changes
	GroupingOverride bool `json:"grouping_override,omitempty"`
	// If set, the configured name will remain even if the underlying Job source code changes
	NameOverride bool `json:"name_override,omitempty"`
	// If set, the configured description will remain even if the underlying Job source code changes
	DescriptionOverride bool `json:"description_override,omitempty"`
	// If set, the configured value will remain even if the underlying Job source code changes
	ApprovalRequiredOverride bool `json:"approval_required_override,omitempty"`
	// If set, the configured value will remain even if the underlying Job source code changes
	DryrunDefaultOverride bool `json:"dryrun_default_override,omitempty"`
	// If set, the configured value will remain even if the underlying Job source code changes
	HiddenOverride bool `json:"hidden_override,omitempty"`
	// If set, the configured value will remain even if the underlying Job source code changes
	SoftTimeLimitOverride bool `json:"soft_time_limit_override,omitempty"`
	// If set, the configured value will remain even if the underlying Job source code changes
	TimeLimitOverride bool `json:"time_limit_override,omitempty"`
	// If set, the configured value will remain even if the underlying Job source code changes
	HasSensitiveVariablesOverride bool `json:"has_sensitive_variables_override,omitempty"`
	// If set, the configured value will remain even if the underlying Job source code changes
	TaskQueuesOverride bool `json:"task_queues_override,omitempty"`
	Created time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
}
