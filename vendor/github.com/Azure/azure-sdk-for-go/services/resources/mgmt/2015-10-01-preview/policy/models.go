package policy

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Type enumerates the values for type.
type Type string

const (
	// BuiltIn ...
	BuiltIn Type = "BuiltIn"
	// Custom ...
	Custom Type = "Custom"
	// NotSpecified ...
	NotSpecified Type = "NotSpecified"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{BuiltIn, Custom, NotSpecified}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Assignment the policy assignment.
type Assignment struct {
	autorest.Response `json:"-"`
	// AssignmentProperties - Properties for the policy assignment.
	*AssignmentProperties `json:"properties,omitempty"`
	// ID - The ID of the policy assignment.
	ID *string `json:"id,omitempty"`
	// Type - The type of the policy assignment.
	Type *string `json:"type,omitempty"`
	// Name - The name of the policy assignment.
	Name *string `json:"name,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// MarshalJSON is the custom marshaler for Assignment.
func (a Assignment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.AssignmentProperties != nil {
		objectMap["properties"] = a.AssignmentProperties
	}
	if a.ID != nil {
		objectMap["id"] = a.ID
	}
	if a.Type != nil {
		objectMap["type"] = a.Type
	}
	if a.Name != nil {
		objectMap["name"] = a.Name
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// UnmarshalJSON is the custom unmarshaler for Assignment struct.
func (a *Assignment) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var assignmentProperties AssignmentProperties
				err = json.Unmarshal(*v, &assignmentProperties)
				if err != nil {
					return err
				}
				a.AssignmentProperties = &assignmentProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				a.ID = &ID
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				a.Type = &typeVar
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				a.Name = &name
			}
		}
	}

	return nil
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// AssignmentListResult list of policy assignments.
type AssignmentListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of policy assignments.
	Value *[]Assignment `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// AssignmentListResultIterator provides access to a complete listing of Assignment values.
type AssignmentListResultIterator struct {
	i    int
	page AssignmentListResultPage
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AssignmentListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AssignmentListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Response returns the raw server response from the last page request.
func (iter AssignmentListResultIterator) Response() AssignmentListResult {
	return iter.page.Response()
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AssignmentListResultIterator) Value() Assignment {
	if !iter.page.NotDone() {
		return Assignment{}
	}
	return iter.page.Values()[iter.i]
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// IsEmpty returns true if the ListResult contains no values.
func (alr AssignmentListResult) IsEmpty() bool {
	return alr.Value == nil || len(*alr.Value) == 0
}

// assignmentListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (alr AssignmentListResult) assignmentListResultPreparer() (*http.Request, error) {
	if alr.NextLink == nil || len(to.String(alr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(alr.NextLink)))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// AssignmentListResultPage contains a page of Assignment values.
type AssignmentListResultPage struct {
	fn  func(AssignmentListResult) (AssignmentListResult, error)
	alr AssignmentListResult
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AssignmentListResultPage) Next() error {
	next, err := page.fn(page.alr)
	if err != nil {
		return err
	}
	page.alr = next
	return nil
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AssignmentListResultPage) NotDone() bool {
	return !page.alr.IsEmpty()
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Response returns the raw server response from the last page request.
func (page AssignmentListResultPage) Response() AssignmentListResult {
	return page.alr
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Values returns the slice of values for the current page or nil if there are no values.
func (page AssignmentListResultPage) Values() []Assignment {
	if page.alr.IsEmpty() {
		return nil
	}
	return *page.alr.Value
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// AssignmentProperties the policy assignment properties.
type AssignmentProperties struct {
	// DisplayName - The display name of the policy assignment.
	DisplayName *string `json:"displayName,omitempty"`
	// PolicyDefinitionID - The ID of the policy definition.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`
	// Scope - The scope for the policy assignment.
	Scope *string `json:"scope,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Definition the policy definition.
type Definition struct {
	autorest.Response `json:"-"`
	// DefinitionProperties - The policy definition properties.
	*DefinitionProperties `json:"properties,omitempty"`
	// ID - The ID of the policy definition.
	ID *string `json:"id,omitempty"`
	// Name - The name of the policy definition. If you do not specify a value for name, the value is inferred from the name value in the request URI.
	Name *string `json:"name,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// MarshalJSON is the custom marshaler for Definition.
func (d Definition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if d.DefinitionProperties != nil {
		objectMap["properties"] = d.DefinitionProperties
	}
	if d.ID != nil {
		objectMap["id"] = d.ID
	}
	if d.Name != nil {
		objectMap["name"] = d.Name
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// UnmarshalJSON is the custom unmarshaler for Definition struct.
func (d *Definition) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var definitionProperties DefinitionProperties
				err = json.Unmarshal(*v, &definitionProperties)
				if err != nil {
					return err
				}
				d.DefinitionProperties = &definitionProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				d.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				d.Name = &name
			}
		}
	}

	return nil
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// DefinitionListResult list of policy definitions.
type DefinitionListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of policy definitions.
	Value *[]Definition `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// DefinitionListResultIterator provides access to a complete listing of Definition values.
type DefinitionListResultIterator struct {
	i    int
	page DefinitionListResultPage
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DefinitionListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DefinitionListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Response returns the raw server response from the last page request.
func (iter DefinitionListResultIterator) Response() DefinitionListResult {
	return iter.page.Response()
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DefinitionListResultIterator) Value() Definition {
	if !iter.page.NotDone() {
		return Definition{}
	}
	return iter.page.Values()[iter.i]
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// IsEmpty returns true if the ListResult contains no values.
func (dlr DefinitionListResult) IsEmpty() bool {
	return dlr.Value == nil || len(*dlr.Value) == 0
}

// definitionListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dlr DefinitionListResult) definitionListResultPreparer() (*http.Request, error) {
	if dlr.NextLink == nil || len(to.String(dlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dlr.NextLink)))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// DefinitionListResultPage contains a page of Definition values.
type DefinitionListResultPage struct {
	fn  func(DefinitionListResult) (DefinitionListResult, error)
	dlr DefinitionListResult
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DefinitionListResultPage) Next() error {
	next, err := page.fn(page.dlr)
	if err != nil {
		return err
	}
	page.dlr = next
	return nil
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DefinitionListResultPage) NotDone() bool {
	return !page.dlr.IsEmpty()
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Response returns the raw server response from the last page request.
func (page DefinitionListResultPage) Response() DefinitionListResult {
	return page.dlr
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// Values returns the slice of values for the current page or nil if there are no values.
func (page DefinitionListResultPage) Values() []Definition {
	if page.dlr.IsEmpty() {
		return nil
	}
	return *page.dlr.Value
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy instead.
// DefinitionProperties the policy definition properties.
type DefinitionProperties struct {
	// PolicyType - The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom. Possible values include: 'NotSpecified', 'BuiltIn', 'Custom'
	PolicyType Type `json:"policyType,omitempty"`
	// DisplayName - The display name of the policy definition.
	DisplayName *string `json:"displayName,omitempty"`
	// Description - The policy definition description.
	Description *string `json:"description,omitempty"`
	// PolicyRule - The policy rule.
	PolicyRule interface{} `json:"policyRule,omitempty"`
}