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

// AdverseEvent is documented here http://hl7.org/fhir/StructureDefinition/AdverseEvent
type AdverseEvent struct {
	Id                    *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Actuality             AdverseEventActuality       `bson:"actuality" json:"actuality"`
	Category              []CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	Event                 *CodeableConcept            `bson:"event,omitempty" json:"event,omitempty"`
	Subject               Reference                   `bson:"subject" json:"subject"`
	Encounter             *Reference                  `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Date                  *string                     `bson:"date,omitempty" gorm:"type:timestamp" json:"date,omitempty"`
	Detected              *string                     `bson:"detected,omitempty" gorm:"type:timestamp" json:"detected,omitempty"`
	RecordedDate          *string                     `bson:"recordedDate,omitempty" gorm:"type:timestamp" json:"recordedDate,omitempty"`
	ResultingCondition    []Reference                 `bson:"resultingCondition,omitempty" json:"resultingCondition,omitempty"`
	Location              *Reference                  `bson:"location,omitempty" json:"location,omitempty"`
	Seriousness           *CodeableConcept            `bson:"seriousness,omitempty" json:"seriousness,omitempty"`
	Severity              *CodeableConcept            `bson:"severity,omitempty" json:"severity,omitempty"`
	Outcome               *CodeableConcept            `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Recorder              *Reference                  `bson:"recorder,omitempty" json:"recorder,omitempty"`
	Contributor           []Reference                 `bson:"contributor,omitempty" json:"contributor,omitempty"`
	SuspectEntity         []AdverseEventSuspectEntity `bson:"suspectEntity,omitempty" json:"suspectEntity,omitempty"`
	SubjectMedicalHistory []Reference                 `bson:"subjectMedicalHistory,omitempty" json:"subjectMedicalHistory,omitempty"`
	ReferenceDocument     []Reference                 `bson:"referenceDocument,omitempty" json:"referenceDocument,omitempty"`
	Study                 []Reference                 `bson:"study,omitempty" json:"study,omitempty"`
}
type AdverseEventSuspectEntity struct {
	Id                *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Instance          Reference                            `bson:"instance" json:"instance"`
	Causality         []AdverseEventSuspectEntityCausality `bson:"causality,omitempty" json:"causality,omitempty"`
}
type AdverseEventSuspectEntityCausality struct {
	Id                 *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Assessment         *CodeableConcept `bson:"assessment,omitempty" json:"assessment,omitempty"`
	ProductRelatedness *string          `bson:"productRelatedness,omitempty" json:"productRelatedness,omitempty"`
	Author             *Reference       `bson:"author,omitempty" json:"author,omitempty"`
	Method             *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
}
type OtherAdverseEvent AdverseEvent

// MarshalJSON marshals the given AdverseEvent as JSON into a byte slice
func (r AdverseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAdverseEvent
		ResourceType string `json:"resourceType"`
	}{
		OtherAdverseEvent: OtherAdverseEvent(r),
		ResourceType:      "AdverseEvent",
	})
}

// UnmarshalAdverseEvent unmarshals a AdverseEvent.
func UnmarshalAdverseEvent(b []byte) (AdverseEvent, error) {
	var adverseEvent AdverseEvent
	if err := json.Unmarshal(b, &adverseEvent); err != nil {
		return adverseEvent, err
	}
	return adverseEvent, nil
}
