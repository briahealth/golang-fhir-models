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

// ExampleScenario is documented here http://hl7.org/fhir/StructureDefinition/ExampleScenario
type ExampleScenario struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                   `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                   `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string                   `bson:"name,omitempty" json:"name,omitempty"`
	Status            PublicationStatus         `bson:"status" json:"status"`
	Experimental      *bool                     `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                   `bson:"date,omitempty" gorm:"type:timestamp" json:"date,omitempty"`
	Publisher         *string                   `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail           `bson:"contact,omitempty" json:"contact,omitempty"`
	UseContext        []UsageContext            `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept         `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Copyright         *string                   `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Purpose           *string                   `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Actor             []ExampleScenarioActor    `bson:"actor,omitempty" json:"actor,omitempty"`
	Instance          []ExampleScenarioInstance `bson:"instance,omitempty" json:"instance,omitempty"`
	Process           []ExampleScenarioProcess  `bson:"process,omitempty" json:"process,omitempty"`
	Workflow          []string                  `bson:"workflow,omitempty" json:"workflow,omitempty"`
}
type ExampleScenarioActor struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ActorId           string                   `bson:"actorId" json:"actorId"`
	Type              ExampleScenarioActorType `bson:"type" json:"type"`
	Name              *string                  `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                  `bson:"description,omitempty" json:"description,omitempty"`
}
type ExampleScenarioInstance struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ResourceId        string                                     `bson:"resourceId" json:"resourceId"`
	ResourceType      ResourceType                               `bson:"resourceType" json:"resourceType"`
	Name              *string                                    `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                                    `bson:"description,omitempty" json:"description,omitempty"`
	Version           []ExampleScenarioInstanceVersion           `bson:"version,omitempty" json:"version,omitempty"`
	ContainedInstance []ExampleScenarioInstanceContainedInstance `bson:"containedInstance,omitempty" json:"containedInstance,omitempty"`
}
type ExampleScenarioInstanceVersion struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	VersionId         string      `bson:"versionId" json:"versionId"`
	Description       string      `bson:"description" json:"description"`
}
type ExampleScenarioInstanceContainedInstance struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ResourceId        string      `bson:"resourceId" json:"resourceId"`
	VersionId         *string     `bson:"versionId,omitempty" json:"versionId,omitempty"`
}
type ExampleScenarioProcess struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Title             string                       `bson:"title" json:"title"`
	Description       *string                      `bson:"description,omitempty" json:"description,omitempty"`
	PreConditions     *string                      `bson:"preConditions,omitempty" json:"preConditions,omitempty"`
	PostConditions    *string                      `bson:"postConditions,omitempty" json:"postConditions,omitempty"`
	Step              []ExampleScenarioProcessStep `bson:"step,omitempty" json:"step,omitempty"`
}
type ExampleScenarioProcessStep struct {
	Id                *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Process           []ExampleScenarioProcess                `bson:"process,omitempty" json:"process,omitempty"`
	Pause             *bool                                   `bson:"pause,omitempty" json:"pause,omitempty"`
	Operation         *ExampleScenarioProcessStepOperation    `bson:"operation,omitempty" json:"operation,omitempty"`
	Alternative       []ExampleScenarioProcessStepAlternative `bson:"alternative,omitempty" json:"alternative,omitempty"`
}
type ExampleScenarioProcessStepOperation struct {
	Id                *string                                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Number            string                                    `bson:"number" json:"number"`
	Type              *string                                   `bson:"type,omitempty" json:"type,omitempty"`
	Name              *string                                   `bson:"name,omitempty" json:"name,omitempty"`
	Initiator         *string                                   `bson:"initiator,omitempty" json:"initiator,omitempty"`
	Receiver          *string                                   `bson:"receiver,omitempty" json:"receiver,omitempty"`
	Description       *string                                   `bson:"description,omitempty" json:"description,omitempty"`
	InitiatorActive   *bool                                     `bson:"initiatorActive,omitempty" json:"initiatorActive,omitempty"`
	ReceiverActive    *bool                                     `bson:"receiverActive,omitempty" json:"receiverActive,omitempty"`
	Request           *ExampleScenarioInstanceContainedInstance `bson:"request,omitempty" json:"request,omitempty"`
	Response          *ExampleScenarioInstanceContainedInstance `bson:"response,omitempty" json:"response,omitempty"`
}
type ExampleScenarioProcessStepAlternative struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Title             string                       `bson:"title" json:"title"`
	Description       *string                      `bson:"description,omitempty" json:"description,omitempty"`
	Step              []ExampleScenarioProcessStep `bson:"step,omitempty" json:"step,omitempty"`
}
type OtherExampleScenario ExampleScenario

// MarshalJSON marshals the given ExampleScenario as JSON into a byte slice
func (r ExampleScenario) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExampleScenario
		ResourceType string `json:"resourceType"`
	}{
		OtherExampleScenario: OtherExampleScenario(r),
		ResourceType:         "ExampleScenario",
	})
}

// UnmarshalExampleScenario unmarshals a ExampleScenario.
func UnmarshalExampleScenario(b []byte) (ExampleScenario, error) {
	var exampleScenario ExampleScenario
	if err := json.Unmarshal(b, &exampleScenario); err != nil {
		return exampleScenario, err
	}
	return exampleScenario, nil
}
