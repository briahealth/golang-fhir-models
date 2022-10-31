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

// ChargeItemDefinition is documented here http://hl7.org/fhir/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinition struct {
	Id                *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                             `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                          `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                              `bson:"url" json:"url"`
	Identifier        []Identifier                        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                             `bson:"version,omitempty" json:"version,omitempty"`
	Title             *string                             `bson:"title,omitempty" json:"title,omitempty"`
	DerivedFromUri    []string                            `bson:"derivedFromUri,omitempty" json:"derivedFromUri,omitempty"`
	PartOf            []string                            `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Replaces          []string                            `bson:"replaces,omitempty" json:"replaces,omitempty"`
	Status            PublicationStatus                   `bson:"status" json:"status"`
	Experimental      *bool                               `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                             `bson:"date,omitempty" gorm:"type:dateTime" json:"date,omitempty"`
	Publisher         *string                             `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                     `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                             `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext                      `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                   `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Copyright         *string                             `bson:"copyright,omitempty" json:"copyright,omitempty"`
	ApprovalDate      *string                             `bson:"approvalDate,omitempty" gorm:"type:date" json:"approvalDate,omitempty"`
	LastReviewDate    *string                             `bson:"lastReviewDate,omitempty" gorm:"type:date" json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                             `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Code              *CodeableConcept                    `bson:"code,omitempty" json:"code,omitempty"`
	Instance          []Reference                         `bson:"instance,omitempty" json:"instance,omitempty"`
	Applicability     []ChargeItemDefinitionApplicability `bson:"applicability,omitempty" json:"applicability,omitempty"`
	PropertyGroup     []ChargeItemDefinitionPropertyGroup `bson:"propertyGroup,omitempty" json:"propertyGroup,omitempty"`
}
type ChargeItemDefinitionApplicability struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Expression        *string     `bson:"expression,omitempty" json:"expression,omitempty"`
}
type ChargeItemDefinitionPropertyGroup struct {
	Id                *string                                           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Applicability     []ChargeItemDefinitionApplicability               `bson:"applicability,omitempty" json:"applicability,omitempty"`
	PriceComponent    []ChargeItemDefinitionPropertyGroupPriceComponent `bson:"priceComponent,omitempty" json:"priceComponent,omitempty"`
}
type ChargeItemDefinitionPropertyGroupPriceComponent struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              InvoicePriceComponentType `bson:"type" json:"type"`
	Code              *CodeableConcept          `bson:"code,omitempty" json:"code,omitempty"`
	Factor            *string                   `bson:"factor,omitempty" gorm:"type:decimal" json:"factor,omitempty"`
	Amount            *Money                    `bson:"amount,omitempty" json:"amount,omitempty"`
}
type OtherChargeItemDefinition ChargeItemDefinition

// MarshalJSON marshals the given ChargeItemDefinition as JSON into a byte slice
func (r ChargeItemDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherChargeItemDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherChargeItemDefinition: OtherChargeItemDefinition(r),
		ResourceType:              "ChargeItemDefinition",
	})
}

// UnmarshalChargeItemDefinition unmarshals a ChargeItemDefinition.
func UnmarshalChargeItemDefinition(b []byte) (ChargeItemDefinition, error) {
	var chargeItemDefinition ChargeItemDefinition
	if err := json.Unmarshal(b, &chargeItemDefinition); err != nil {
		return chargeItemDefinition, err
	}
	return chargeItemDefinition, nil
}
