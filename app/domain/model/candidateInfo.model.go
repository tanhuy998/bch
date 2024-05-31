package model

import (
	"time"

	"github.com/google/uuid"
)

type CandidateSigningInfo struct {
	//UUID       uuid.UUID       `json:"uuid" bson:"uuid" validate:"required"`
	CivilIndentity CitizenIdentity `json:"civilIdentity" bson:"civilIdentity,omitempty" validate:"required"`
	Politic        *PoliticDetail  `json:"politic" bson:"politic,omitempty"`
	Education      EducationDetail `json:"education" bson:"education,omitempty"`
	Job            *string         `json:"job" bson:"job,omitempty"`
	JobPlace       *string         `json:"jobPlace" bson:"jobPlace,omitempty"`
	Family         *FamilyInfo     `json:"family" bson:"family,omitempty"`
	Version        *time.Time      `json:"version" bson:"version,omitempty"`
}

type Citizen struct {
}

type CivilIDCardKind int

const (
	CMND CivilIDCardKind = 1
	CCCD CivilIDCardKind = 2
)

type EducationDetail struct {
	PrimarySchool   *string    `json:"primarySchool" bson:"primarySchool"`
	SecondarySchool *string    `json:"secondarySchool" bson:"secondarySchool"`
	HighSchool      *string    `json:"highSchool" bson:"highSchool"`
	HighestGrade    *byte      `json:"highestGrade" bson:"highestGrade"`
	College         *string    `json:"college" bson:"college,omitempty"`
	GraduateAt      *time.Time `json:"graduateAt" bson:"graduateAt,omitempty"`
	Expertise       *string    `json:"expertise" bson:"expertise,omitempty"`
}

type CitizenIdentity struct {
	//CardKind          *CivilIDCardKind `json:"kind" bson:"kind,omitempty"`
	Name              *string    `json:"name" bson:"name" validate:"required,alphaunicode"`
	DateOfBirth       *time.Time `json:"dateOfBirth" bson:"dateOfBirth" validate:"required"`
	Male              *bool      `json:"male" bson:"male,omitempty" validate:"required"`
	IDNumber          *string    `json:"idNumber" bson:"idNumber" validate:"required,number,len=12"`
	BirthPlace        *string    `json:"birthPlace" bson:"birthPlace" validate:"required"`
	PlaceOfOrigin     *string    `json:"placeOfOrigin" bson:"placeOfOrigin" validate:"required"`
	Ethnicity         *string    `json:"ethnicity" bson:"ethnicity" validate:"required"`
	Religion          *string    `json:"religion" bson:"religion" validate:"required"`
	Nationality       *string    `json:"nationality" bson:"nationality" validate:"required"`
	PermanentResident *string    `json:"permanentResident" bson:"permanentResident" validate:"required"`
	TemporaryResident *string    `json:"temporaryResident" bson:"temporaryResident"`
	CurrentResident   *string    `json:"currentResident" bson:"currentResident" validate:"required"`
}

type Family struct {
	Mother  *FamilyMember  `json:"mother" bson:"mother,omitempty"`
	Father  *FamilyMember  `json:"father" bson:"father,omitempty"`
	Another []FamilyMember `json:"another,omitempty" bson:"another"`
}

type FamilyMember struct {
	Name          string           `json:"name" bson:"name" validate:"required"`
	DateOfBirth   time.Time        `json:"dateOfBirth" bson:"dateOfBirth" validate:"required"`
	Dead          *bool            `json:"dead" bson:"dead,omitEmpty"`
	Job           *string          `json:"job" bson:"job,omitempty"`
	Politic       *PoliticDetail   `json:"politic" bson:"politic" validate:"required"`
	Education     *EducationDetail `json:"education" bson:"education,omitempty"`
	CivilIdentity *CitizenIdentity `json:"civilIdentity" bson:"civilIdentity,omitempty"`
}

type CivilHistory struct {
	BeforeReunification *string `json:"beforeReunification" bson:"beforeReunification,omitempty"`
	AfterReunification  *string `json:"afterReunification" bson:"beforeReunification,omitempty"`
}

type FamilyInfo struct {
	Father   FamilyMember   `json:"father" bson:"father,omitempty"`
	Mother   FamilyMember   `json:"mother" bson:"mother,omitempty"`
	Anothers []FamilyMember `json:"anothers" bson:"anothers,omitempty"`
}

type PoliticDetail struct {
	History       CivilHistory `json:"history" bson:"history,omitempty"`
	UnionJoinDate *time.Time   `json:"unionJoinDate" bson:"unionJoinDate,omitempty"`
	PartyJoinDate *time.Time   `json:"partyJoinDate" bson:"partyJoinDate,omitempty"`
}

type BasicType struct {
	UUID uuid.UUID `json:"uuid" bson:"uuid,omitempty" validate:"required"`
	Name string    `json:"name" bson:"name,omitempty" validate:"required"`
}

type Ethinicty BasicType

type Religion BasicType
