/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// ExportTemplateOwner - struct for ExportTemplateOwner
type ExportTemplateOwner struct {
	NestedGitRepository *NestedGitRepository
}

// NestedGitRepositoryAsExportTemplateOwner is a convenience function that returns NestedGitRepository wrapped in ExportTemplateOwner
func NestedGitRepositoryAsExportTemplateOwner(v *NestedGitRepository) ExportTemplateOwner {
	return ExportTemplateOwner{
		NestedGitRepository: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ExportTemplateOwner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NestedGitRepository
	err = newStrictDecoder(data).Decode(&dst.NestedGitRepository)
	if err == nil {
		jsonNestedGitRepository, _ := json.Marshal(dst.NestedGitRepository)
		if string(jsonNestedGitRepository) == "{}" { // empty struct
			dst.NestedGitRepository = nil
		} else {
			if err = validator.Validate(dst.NestedGitRepository); err != nil {
				dst.NestedGitRepository = nil
			} else {
				match++
			}
		}
	} else {
		dst.NestedGitRepository = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.NestedGitRepository = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ExportTemplateOwner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ExportTemplateOwner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ExportTemplateOwner) MarshalJSON() ([]byte, error) {
	if src.NestedGitRepository != nil {
		return json.Marshal(&src.NestedGitRepository)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ExportTemplateOwner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.NestedGitRepository != nil {
		return obj.NestedGitRepository
	}

	// all schemas are nil
	return nil
}

type NullableExportTemplateOwner struct {
	value *ExportTemplateOwner
	isSet bool
}

func (v NullableExportTemplateOwner) Get() *ExportTemplateOwner {
	return v.value
}

func (v *NullableExportTemplateOwner) Set(val *ExportTemplateOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableExportTemplateOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableExportTemplateOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportTemplateOwner(val *ExportTemplateOwner) *NullableExportTemplateOwner {
	return &NullableExportTemplateOwner{value: val, isSet: true}
}

func (v NullableExportTemplateOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportTemplateOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


