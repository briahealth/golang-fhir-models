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

// ExplanationOfBenefit is documented here http://hl7.org/fhir/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefit struct {
	Id                    *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                                `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                             `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                ExplanationOfBenefitStatus             `bson:"status" json:"status"`
	Type                  CodeableConcept                        `bson:"type" json:"type"`
	SubType               *CodeableConcept                       `bson:"subType,omitempty" json:"subType,omitempty"`
	Use                   Use                                    `bson:"use" json:"use"`
	Patient               Reference                              `bson:"patient" json:"patient"`
	BillablePeriod        *Period                                `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	Created               string                                 `bson:"created" gorm:"type:dateTime" json:"created"`
	Enterer               *Reference                             `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Insurer               Reference                              `bson:"insurer" json:"insurer"`
	Provider              Reference                              `bson:"provider" json:"provider"`
	Priority              *CodeableConcept                       `bson:"priority,omitempty" json:"priority,omitempty"`
	FundsReserveRequested *CodeableConcept                       `bson:"fundsReserveRequested,omitempty" json:"fundsReserveRequested,omitempty"`
	FundsReserve          *CodeableConcept                       `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	Related               []ExplanationOfBenefitRelated          `bson:"related,omitempty" json:"related,omitempty"`
	Prescription          *Reference                             `bson:"prescription,omitempty" json:"prescription,omitempty"`
	OriginalPrescription  *Reference                             `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	Payee                 *ExplanationOfBenefitPayee             `bson:"payee,omitempty" json:"payee,omitempty"`
	Referral              *Reference                             `bson:"referral,omitempty" json:"referral,omitempty"`
	Facility              *Reference                             `bson:"facility,omitempty" json:"facility,omitempty"`
	Claim                 *Reference                             `bson:"claim,omitempty" json:"claim,omitempty"`
	ClaimResponse         *Reference                             `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
	Outcome               RemittanceOutcome                      `bson:"outcome" json:"outcome"`
	Disposition           *string                                `bson:"disposition,omitempty" json:"disposition,omitempty"`
	PreAuthRef            []string                               `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	PreAuthRefPeriod      []Period                               `bson:"preAuthRefPeriod,omitempty" json:"preAuthRefPeriod,omitempty"`
	CareTeam              []ExplanationOfBenefitCareTeam         `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	SupportingInfo        []ExplanationOfBenefitSupportingInfo   `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Diagnosis             []ExplanationOfBenefitDiagnosis        `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Procedure             []ExplanationOfBenefitProcedure        `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Precedence            *int                                   `bson:"precedence,omitempty" json:"precedence,omitempty"`
	Insurance             []ExplanationOfBenefitInsurance        `bson:"insurance" json:"insurance"`
	Accident              *ExplanationOfBenefitAccident          `bson:"accident,omitempty" json:"accident,omitempty"`
	Item                  []ExplanationOfBenefitItem             `bson:"item,omitempty" json:"item,omitempty"`
	AddItem               []ExplanationOfBenefitAddItem          `bson:"addItem,omitempty" json:"addItem,omitempty"`
	Adjudication          []ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Total                 []ExplanationOfBenefitTotal            `bson:"total,omitempty" json:"total,omitempty"`
	Payment               *ExplanationOfBenefitPayment           `bson:"payment,omitempty" json:"payment,omitempty"`
	FormCode              *CodeableConcept                       `bson:"formCode,omitempty" json:"formCode,omitempty"`
	Form                  *Attachment                            `bson:"form,omitempty" json:"form,omitempty"`
	ProcessNote           []ExplanationOfBenefitProcessNote      `bson:"processNote,omitempty" json:"processNote,omitempty"`
	BenefitPeriod         *Period                                `bson:"benefitPeriod,omitempty" json:"benefitPeriod,omitempty"`
	BenefitBalance        []ExplanationOfBenefitBenefitBalance   `bson:"benefitBalance,omitempty" json:"benefitBalance,omitempty"`
}
type ExplanationOfBenefitRelated struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Claim             *Reference       `bson:"claim,omitempty" json:"claim,omitempty"`
	Relationship      *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Reference         *Identifier      `bson:"reference,omitempty" json:"reference,omitempty"`
}
type ExplanationOfBenefitPayee struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Party             *Reference       `bson:"party,omitempty" json:"party,omitempty"`
}
type ExplanationOfBenefitCareTeam struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int              `bson:"sequence" json:"sequence"`
	Provider          Reference        `bson:"provider" json:"provider"`
	Responsible       *bool            `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Qualification     *CodeableConcept `bson:"qualification,omitempty" json:"qualification,omitempty"`
}
type ExplanationOfBenefitSupportingInfo struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int              `bson:"sequence" json:"sequence"`
	Category          CodeableConcept  `bson:"category" json:"category"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Reason            *Coding          `bson:"reason,omitempty" json:"reason,omitempty"`
}
type ExplanationOfBenefitDiagnosis struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int               `bson:"sequence" json:"sequence"`
	Type              []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	OnAdmission       *CodeableConcept  `bson:"onAdmission,omitempty" json:"onAdmission,omitempty"`
	PackageCode       *CodeableConcept  `bson:"packageCode,omitempty" json:"packageCode,omitempty"`
}
type ExplanationOfBenefitProcedure struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int               `bson:"sequence" json:"sequence"`
	Type              []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Date              *string           `bson:"date,omitempty" gorm:"type:dateTime" json:"date,omitempty"`
	Udi               []Reference       `bson:"udi,omitempty" json:"udi,omitempty"`
}
type ExplanationOfBenefitInsurance struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Focal             bool        `bson:"focal" json:"focal"`
	Coverage          Reference   `bson:"coverage" json:"coverage"`
	PreAuthRef        []string    `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
}
type ExplanationOfBenefitAccident struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Date              *string          `bson:"date,omitempty" gorm:"type:date" json:"date,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
}
type ExplanationOfBenefitItem struct {
	Id                  *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence            int                                    `bson:"sequence" json:"sequence"`
	CareTeamSequence    []int                                  `bson:"careTeamSequence,omitempty" json:"careTeamSequence,omitempty"`
	DiagnosisSequence   []int                                  `bson:"diagnosisSequence,omitempty" json:"diagnosisSequence,omitempty"`
	ProcedureSequence   []int                                  `bson:"procedureSequence,omitempty" json:"procedureSequence,omitempty"`
	InformationSequence []int                                  `bson:"informationSequence,omitempty" json:"informationSequence,omitempty"`
	Revenue             *CodeableConcept                       `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category            *CodeableConcept                       `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService    CodeableConcept                        `bson:"productOrService" json:"productOrService"`
	Modifier            []CodeableConcept                      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode         []CodeableConcept                      `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity            *Quantity                              `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice           *Money                                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor              *string                                `bson:"factor,omitempty" gorm:"type:decimal" json:"factor,omitempty"`
	Net                 *Money                                 `bson:"net,omitempty" json:"net,omitempty"`
	Udi                 []Reference                            `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite            *CodeableConcept                       `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite             []CodeableConcept                      `bson:"subSite,omitempty" json:"subSite,omitempty"`
	Encounter           []Reference                            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	NoteNumber          []int                                  `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication        []ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail              []ExplanationOfBenefitItemDetail       `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ExplanationOfBenefitItemAdjudication struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `bson:"category" json:"category"`
	Reason            *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount            *Money           `bson:"amount,omitempty" json:"amount,omitempty"`
	Value             *string          `bson:"value,omitempty" gorm:"type:decimal" json:"value,omitempty"`
}
type ExplanationOfBenefitItemDetail struct {
	Id                *string                                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int                                       `bson:"sequence" json:"sequence"`
	Revenue           *CodeableConcept                          `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                          `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService  CodeableConcept                           `bson:"productOrService" json:"productOrService"`
	Modifier          []CodeableConcept                         `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept                         `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity                                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                                    `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string                                   `bson:"factor,omitempty" gorm:"type:decimal" json:"factor,omitempty"`
	Net               *Money                                    `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []Reference                               `bson:"udi,omitempty" json:"udi,omitempty"`
	NoteNumber        []int                                     `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication    `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail         []ExplanationOfBenefitItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}
type ExplanationOfBenefitItemDetailSubDetail struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int                                    `bson:"sequence" json:"sequence"`
	Revenue           *CodeableConcept                       `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                       `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService  CodeableConcept                        `bson:"productOrService" json:"productOrService"`
	Modifier          []CodeableConcept                      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept                      `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity                              `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string                                `bson:"factor,omitempty" gorm:"type:decimal" json:"factor,omitempty"`
	Net               *Money                                 `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []Reference                            `bson:"udi,omitempty" json:"udi,omitempty"`
	NoteNumber        []int                                  `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}
type ExplanationOfBenefitAddItem struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ItemSequence      []int                                  `bson:"itemSequence,omitempty" json:"itemSequence,omitempty"`
	DetailSequence    []int                                  `bson:"detailSequence,omitempty" json:"detailSequence,omitempty"`
	SubDetailSequence []int                                  `bson:"subDetailSequence,omitempty" json:"subDetailSequence,omitempty"`
	Provider          []Reference                            `bson:"provider,omitempty" json:"provider,omitempty"`
	ProductOrService  CodeableConcept                        `bson:"productOrService" json:"productOrService"`
	Modifier          []CodeableConcept                      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept                      `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity                              `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string                                `bson:"factor,omitempty" gorm:"type:decimal" json:"factor,omitempty"`
	Net               *Money                                 `bson:"net,omitempty" json:"net,omitempty"`
	BodySite          *CodeableConcept                       `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite           []CodeableConcept                      `bson:"subSite,omitempty" json:"subSite,omitempty"`
	NoteNumber        []int                                  `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail            []ExplanationOfBenefitAddItemDetail    `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ExplanationOfBenefitAddItemDetail struct {
	Id                *string                                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                              `bson:"productOrService" json:"productOrService"`
	Modifier          []CodeableConcept                            `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Quantity          *Quantity                                    `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                                       `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string                                      `bson:"factor,omitempty" gorm:"type:decimal" json:"factor,omitempty"`
	Net               *Money                                       `bson:"net,omitempty" json:"net,omitempty"`
	NoteNumber        []int                                        `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication       `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail         []ExplanationOfBenefitAddItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}
type ExplanationOfBenefitAddItemDetailSubDetail struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                        `bson:"productOrService" json:"productOrService"`
	Modifier          []CodeableConcept                      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Quantity          *Quantity                              `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string                                `bson:"factor,omitempty" gorm:"type:decimal" json:"factor,omitempty"`
	Net               *Money                                 `bson:"net,omitempty" json:"net,omitempty"`
	NoteNumber        []int                                  `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}
type ExplanationOfBenefitTotal struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept `bson:"category" json:"category"`
	Amount            Money           `bson:"amount" json:"amount"`
}
type ExplanationOfBenefitPayment struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Adjustment        *Money           `bson:"adjustment,omitempty" json:"adjustment,omitempty"`
	AdjustmentReason  *CodeableConcept `bson:"adjustmentReason,omitempty" json:"adjustmentReason,omitempty"`
	Date              *string          `bson:"date,omitempty" gorm:"type:date" json:"date,omitempty"`
	Amount            *Money           `bson:"amount,omitempty" json:"amount,omitempty"`
	Identifier        *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
}
type ExplanationOfBenefitProcessNote struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Number            *int             `bson:"number,omitempty" json:"number,omitempty"`
	Type              *NoteType        `bson:"type,omitempty" json:"type,omitempty"`
	Text              *string          `bson:"text,omitempty" json:"text,omitempty"`
	Language          *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
}
type ExplanationOfBenefitBenefitBalance struct {
	Id                *string                                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept                               `bson:"category" json:"category"`
	Excluded          *bool                                         `bson:"excluded,omitempty" json:"excluded,omitempty"`
	Name              *string                                       `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                                       `bson:"description,omitempty" json:"description,omitempty"`
	Network           *CodeableConcept                              `bson:"network,omitempty" json:"network,omitempty"`
	Unit              *CodeableConcept                              `bson:"unit,omitempty" json:"unit,omitempty"`
	Term              *CodeableConcept                              `bson:"term,omitempty" json:"term,omitempty"`
	Financial         []ExplanationOfBenefitBenefitBalanceFinancial `bson:"financial,omitempty" json:"financial,omitempty"`
}
type ExplanationOfBenefitBenefitBalanceFinancial struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
}
type OtherExplanationOfBenefit ExplanationOfBenefit

// MarshalJSON marshals the given ExplanationOfBenefit as JSON into a byte slice
func (r ExplanationOfBenefit) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExplanationOfBenefit
		ResourceType string `json:"resourceType"`
	}{
		OtherExplanationOfBenefit: OtherExplanationOfBenefit(r),
		ResourceType:              "ExplanationOfBenefit",
	})
}

// UnmarshalExplanationOfBenefit unmarshals a ExplanationOfBenefit.
func UnmarshalExplanationOfBenefit(b []byte) (ExplanationOfBenefit, error) {
	var explanationOfBenefit ExplanationOfBenefit
	if err := json.Unmarshal(b, &explanationOfBenefit); err != nil {
		return explanationOfBenefit, err
	}
	return explanationOfBenefit, nil
}
