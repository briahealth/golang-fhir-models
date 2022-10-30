// Copyright 2019 - 2021 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// RegulatedAuthorization is documented here http://hl7.org/fhir/StructureDefinition/RegulatedAuthorization
type RegulatedAuthorization struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject           []Reference                 `bson:"subject,omitempty" json:"subject,omitempty"`
	Type              *CodeableConcept            `bson:"type,omitempty" json:"type,omitempty"`
	Description       *string                     `bson:"description,omitempty" json:"description,omitempty"`
	Region            []CodeableConcept           `bson:"region,omitempty" json:"region,omitempty"`
	Status            *CodeableConcept            `bson:"status,omitempty" json:"status,omitempty"`
	StatusDate        *string                     `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	ValidityPeriod    *Period                     `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty"`
	Indication        *CodeableReference          `bson:"indication,omitempty" json:"indication,omitempty"`
	IntendedUse       *CodeableConcept            `bson:"intendedUse,omitempty" json:"intendedUse,omitempty"`
	Basis             []CodeableConcept           `bson:"basis,omitempty" json:"basis,omitempty"`
	Holder            *Reference                  `bson:"holder,omitempty" json:"holder,omitempty"`
	Regulator         *Reference                  `bson:"regulator,omitempty" json:"regulator,omitempty"`
	Case              *RegulatedAuthorizationCase `bson:"case,omitempty" json:"case,omitempty"`
}
type RegulatedAuthorizationCase struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type              *CodeableConcept             `bson:"type,omitempty" json:"type,omitempty"`
	Status            *CodeableConcept             `bson:"status,omitempty" json:"status,omitempty"`
	Application       []RegulatedAuthorizationCase `bson:"application,omitempty" json:"application,omitempty"`
}
type OtherRegulatedAuthorization RegulatedAuthorization

// MarshalJSON marshals the given RegulatedAuthorization as JSON into a byte slice
func (r RegulatedAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRegulatedAuthorization
		ResourceType string `json:"resourceType"`
	}{
		OtherRegulatedAuthorization: OtherRegulatedAuthorization(r),
		ResourceType:                "RegulatedAuthorization",
	})
}

// UnmarshalRegulatedAuthorization unmarshals a RegulatedAuthorization.
func UnmarshalRegulatedAuthorization(b []byte) (RegulatedAuthorization, error) {
	var regulatedAuthorization RegulatedAuthorization
	if err := json.Unmarshal(b, &regulatedAuthorization); err != nil {
		return regulatedAuthorization, err
	}
	return regulatedAuthorization, nil
}
