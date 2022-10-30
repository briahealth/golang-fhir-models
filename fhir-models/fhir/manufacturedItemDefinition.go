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

// ManufacturedItemDefinition is documented here http://hl7.org/fhir/StructureDefinition/ManufacturedItemDefinition
type ManufacturedItemDefinition struct {
	Id                   *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                              `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                           `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               PublicationStatus                    `bson:"status" json:"status"`
	ManufacturedDoseForm CodeableConcept                      `bson:"manufacturedDoseForm" json:"manufacturedDoseForm"`
	UnitOfPresentation   *CodeableConcept                     `bson:"unitOfPresentation,omitempty" json:"unitOfPresentation,omitempty"`
	Manufacturer         []Reference                          `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Ingredient           []CodeableConcept                    `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	Property             []ManufacturedItemDefinitionProperty `bson:"property,omitempty" json:"property,omitempty"`
}
type ManufacturedItemDefinitionProperty struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
}
type OtherManufacturedItemDefinition ManufacturedItemDefinition

// MarshalJSON marshals the given ManufacturedItemDefinition as JSON into a byte slice
func (r ManufacturedItemDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherManufacturedItemDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherManufacturedItemDefinition: OtherManufacturedItemDefinition(r),
		ResourceType:                    "ManufacturedItemDefinition",
	})
}

// UnmarshalManufacturedItemDefinition unmarshals a ManufacturedItemDefinition.
func UnmarshalManufacturedItemDefinition(b []byte) (ManufacturedItemDefinition, error) {
	var manufacturedItemDefinition ManufacturedItemDefinition
	if err := json.Unmarshal(b, &manufacturedItemDefinition); err != nil {
		return manufacturedItemDefinition, err
	}
	return manufacturedItemDefinition, nil
}
