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

// SubstanceDefinition is documented here http://hl7.org/fhir/StructureDefinition/SubstanceDefinition
type SubstanceDefinition struct {
	Id                *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                              `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                           `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                              `bson:"version,omitempty" json:"version,omitempty"`
	Status            *CodeableConcept                     `bson:"status,omitempty" json:"status,omitempty"`
	Classification    []CodeableConcept                    `bson:"classification,omitempty" json:"classification,omitempty"`
	Domain            *CodeableConcept                     `bson:"domain,omitempty" json:"domain,omitempty"`
	Grade             []CodeableConcept                    `bson:"grade,omitempty" json:"grade,omitempty"`
	Description       *string                              `bson:"description,omitempty" json:"description,omitempty"`
	InformationSource []Reference                          `bson:"informationSource,omitempty" json:"informationSource,omitempty"`
	Note              []Annotation                         `bson:"note,omitempty" json:"note,omitempty"`
	Manufacturer      []Reference                          `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Supplier          []Reference                          `bson:"supplier,omitempty" json:"supplier,omitempty"`
	Moiety            []SubstanceDefinitionMoiety          `bson:"moiety,omitempty" json:"moiety,omitempty"`
	Property          []SubstanceDefinitionProperty        `bson:"property,omitempty" json:"property,omitempty"`
	MolecularWeight   []SubstanceDefinitionMolecularWeight `bson:"molecularWeight,omitempty" json:"molecularWeight,omitempty"`
	Structure         *SubstanceDefinitionStructure        `bson:"structure,omitempty" json:"structure,omitempty"`
	Code              []SubstanceDefinitionCode            `bson:"code,omitempty" json:"code,omitempty"`
	Name              []SubstanceDefinitionName            `bson:"name,omitempty" json:"name,omitempty"`
	Relationship      []SubstanceDefinitionRelationship    `bson:"relationship,omitempty" json:"relationship,omitempty"`
	SourceMaterial    *SubstanceDefinitionSourceMaterial   `bson:"sourceMaterial,omitempty" json:"sourceMaterial,omitempty"`
}
type SubstanceDefinitionMoiety struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Identifier        *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name              *string          `bson:"name,omitempty" json:"name,omitempty"`
	Stereochemistry   *CodeableConcept `bson:"stereochemistry,omitempty" json:"stereochemistry,omitempty"`
	OpticalActivity   *CodeableConcept `bson:"opticalActivity,omitempty" json:"opticalActivity,omitempty"`
	MolecularFormula  *string          `bson:"molecularFormula,omitempty" json:"molecularFormula,omitempty"`
	MeasurementType   *CodeableConcept `bson:"measurementType,omitempty" json:"measurementType,omitempty"`
}
type SubstanceDefinitionProperty struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
}
type SubstanceDefinitionMolecularWeight struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Method            *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Amount            Quantity         `bson:"amount" json:"amount"`
}
type SubstanceDefinitionStructure struct {
	Id                       *string                                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Stereochemistry          *CodeableConcept                             `bson:"stereochemistry,omitempty" json:"stereochemistry,omitempty"`
	OpticalActivity          *CodeableConcept                             `bson:"opticalActivity,omitempty" json:"opticalActivity,omitempty"`
	MolecularFormula         *string                                      `bson:"molecularFormula,omitempty" json:"molecularFormula,omitempty"`
	MolecularFormulaByMoiety *string                                      `bson:"molecularFormulaByMoiety,omitempty" json:"molecularFormulaByMoiety,omitempty"`
	MolecularWeight          *SubstanceDefinitionMolecularWeight          `bson:"molecularWeight,omitempty" json:"molecularWeight,omitempty"`
	Technique                []CodeableConcept                            `bson:"technique,omitempty" json:"technique,omitempty"`
	SourceDocument           []Reference                                  `bson:"sourceDocument,omitempty" json:"sourceDocument,omitempty"`
	Representation           []SubstanceDefinitionStructureRepresentation `bson:"representation,omitempty" json:"representation,omitempty"`
}
type SubstanceDefinitionStructureRepresentation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Representation    *string          `bson:"representation,omitempty" json:"representation,omitempty"`
	Format            *CodeableConcept `bson:"format,omitempty" json:"format,omitempty"`
	Document          *Reference       `bson:"document,omitempty" json:"document,omitempty"`
}
type SubstanceDefinitionCode struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Status            *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	StatusDate        *string          `bson:"statusDate,omitempty" gorm:"type:timestamp" json:"statusDate,omitempty"`
	Note              []Annotation     `bson:"note,omitempty" json:"note,omitempty"`
	Source            []Reference      `bson:"source,omitempty" json:"source,omitempty"`
}
type SubstanceDefinitionName struct {
	Id                *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                            `bson:"name" json:"name"`
	Type              *CodeableConcept                  `bson:"type,omitempty" json:"type,omitempty"`
	Status            *CodeableConcept                  `bson:"status,omitempty" json:"status,omitempty"`
	Preferred         *bool                             `bson:"preferred,omitempty" json:"preferred,omitempty"`
	Language          []CodeableConcept                 `bson:"language,omitempty" json:"language,omitempty"`
	Domain            []CodeableConcept                 `bson:"domain,omitempty" json:"domain,omitempty"`
	Jurisdiction      []CodeableConcept                 `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Synonym           []SubstanceDefinitionName         `bson:"synonym,omitempty" json:"synonym,omitempty"`
	Translation       []SubstanceDefinitionName         `bson:"translation,omitempty" json:"translation,omitempty"`
	Official          []SubstanceDefinitionNameOfficial `bson:"official,omitempty" json:"official,omitempty"`
	Source            []Reference                       `bson:"source,omitempty" json:"source,omitempty"`
}
type SubstanceDefinitionNameOfficial struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Authority         *CodeableConcept `bson:"authority,omitempty" json:"authority,omitempty"`
	Status            *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	Date              *string          `bson:"date,omitempty" gorm:"type:timestamp" json:"date,omitempty"`
}
type SubstanceDefinitionRelationship struct {
	Id                   *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `bson:"type" json:"type"`
	IsDefining           *bool            `bson:"isDefining,omitempty" json:"isDefining,omitempty"`
	RatioHighLimitAmount *Ratio           `bson:"ratioHighLimitAmount,omitempty" json:"ratioHighLimitAmount,omitempty"`
	Comparator           *CodeableConcept `bson:"comparator,omitempty" json:"comparator,omitempty"`
	Source               []Reference      `bson:"source,omitempty" json:"source,omitempty"`
}
type SubstanceDefinitionSourceMaterial struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Genus             *CodeableConcept  `bson:"genus,omitempty" json:"genus,omitempty"`
	Species           *CodeableConcept  `bson:"species,omitempty" json:"species,omitempty"`
	Part              *CodeableConcept  `bson:"part,omitempty" json:"part,omitempty"`
	CountryOfOrigin   []CodeableConcept `bson:"countryOfOrigin,omitempty" json:"countryOfOrigin,omitempty"`
}
type OtherSubstanceDefinition SubstanceDefinition

// MarshalJSON marshals the given SubstanceDefinition as JSON into a byte slice
func (r SubstanceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceDefinition: OtherSubstanceDefinition(r),
		ResourceType:             "SubstanceDefinition",
	})
}

// UnmarshalSubstanceDefinition unmarshals a SubstanceDefinition.
func UnmarshalSubstanceDefinition(b []byte) (SubstanceDefinition, error) {
	var substanceDefinition SubstanceDefinition
	if err := json.Unmarshal(b, &substanceDefinition); err != nil {
		return substanceDefinition, err
	}
	return substanceDefinition, nil
}
