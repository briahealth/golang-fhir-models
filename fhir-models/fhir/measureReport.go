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

// MeasureReport is documented here http://hl7.org/fhir/StructureDefinition/MeasureReport
type MeasureReport struct {
	Id                  *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              MeasureReportStatus  `bson:"status" json:"status"`
	Type                MeasureReportType    `bson:"type" json:"type"`
	Measure             string               `bson:"measure" json:"measure"`
	Subject             *Reference           `bson:"subject,omitempty" json:"subject,omitempty"`
	Date                *string              `bson:"date,omitempty" gorm:"type:timestamp" json:"date,omitempty"`
	Reporter            *Reference           `bson:"reporter,omitempty" json:"reporter,omitempty"`
	Period              Period               `bson:"period" json:"period"`
	ImprovementNotation *CodeableConcept     `bson:"improvementNotation,omitempty" json:"improvementNotation,omitempty"`
	Group               []MeasureReportGroup `bson:"group,omitempty" json:"group,omitempty"`
	EvaluatedResource   []Reference          `bson:"evaluatedResource,omitempty" json:"evaluatedResource,omitempty"`
}
type MeasureReportGroup struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *CodeableConcept               `bson:"code,omitempty" json:"code,omitempty"`
	Population        []MeasureReportGroupPopulation `bson:"population,omitempty" json:"population,omitempty"`
	MeasureScore      *Quantity                      `bson:"measureScore,omitempty" json:"measureScore,omitempty"`
	Stratifier        []MeasureReportGroupStratifier `bson:"stratifier,omitempty" json:"stratifier,omitempty"`
}
type MeasureReportGroupPopulation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Count             *int             `bson:"count,omitempty" json:"count,omitempty"`
	SubjectResults    *Reference       `bson:"subjectResults,omitempty" json:"subjectResults,omitempty"`
}
type MeasureReportGroupStratifier struct {
	Id                *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              []CodeableConcept                     `bson:"code,omitempty" json:"code,omitempty"`
	Stratum           []MeasureReportGroupStratifierStratum `bson:"stratum,omitempty" json:"stratum,omitempty"`
}
type MeasureReportGroupStratifierStratum struct {
	Id                *string                                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Value             *CodeableConcept                                `bson:"value,omitempty" json:"value,omitempty"`
	Component         []MeasureReportGroupStratifierStratumComponent  `bson:"component,omitempty" json:"component,omitempty"`
	Population        []MeasureReportGroupStratifierStratumPopulation `bson:"population,omitempty" json:"population,omitempty"`
	MeasureScore      *Quantity                                       `bson:"measureScore,omitempty" json:"measureScore,omitempty"`
}
type MeasureReportGroupStratifierStratumComponent struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
	Value             CodeableConcept `bson:"value" json:"value"`
}
type MeasureReportGroupStratifierStratumPopulation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Count             *int             `bson:"count,omitempty" json:"count,omitempty"`
	SubjectResults    *Reference       `bson:"subjectResults,omitempty" json:"subjectResults,omitempty"`
}
type OtherMeasureReport MeasureReport

// MarshalJSON marshals the given MeasureReport as JSON into a byte slice
func (r MeasureReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMeasureReport
		ResourceType string `json:"resourceType"`
	}{
		OtherMeasureReport: OtherMeasureReport(r),
		ResourceType:       "MeasureReport",
	})
}

// UnmarshalMeasureReport unmarshals a MeasureReport.
func UnmarshalMeasureReport(b []byte) (MeasureReport, error) {
	var measureReport MeasureReport
	if err := json.Unmarshal(b, &measureReport); err != nil {
		return measureReport, err
	}
	return measureReport, nil
}
