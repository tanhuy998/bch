package model

import (
	"time"

	"github.com/google/uuid"
)

type CandidateSigningInfo struct {
	//UUID       uuid.UUID       `json:"uuid" bson:"uuid" validate:"required"`
	Identitity CitizenIdentity `json:"civilIdentity" bson:"civilIdentity,omitempty" validate:"required"`
	Job        *string         `json:"job" bson:"job,omitempty"`
	Education  EducationDetail `json:"education" bson:"education,omitempty"`
	Politic    *PoliticDetail  `json:"politic" bson:"politic,omitempty"`
	Family     FamilyInfo      `json:"family" bson:"family,omitempty"`
	Version    *time.Time      `json:"version" bson:"version,omitempty"`
}

type Citizen struct {
}

type CivilIDCardKind int

const (
	CMND CivilIDCardKind = 1
	CCCD CivilIDCardKind = 2
)

type EducationDetail struct {
	PrimarySchool   *string    `json:"primarySchool" bson:"primarySchool,omitempty"`
	SecondarySchool *string    `json:"secondarySchool" bson:"secondarySchool,omitempty"`
	HighSchool      *string    `json:"highSchool" bson:"highSchool,omitempty"`
	HighestGrade    *byte      `json:"highestGrade" bson:"highestGrade,omitempty"`
	College         *string    `json:"college" bson:"college,omitempty"`
	GraduateAt      *time.Time `json:"graduateAt" bson:"graduateAt,omitempty"`
	Expertise       *string    `json:"expertise" bson:"expertise,omitempty"`
}

type CitizenIdentity struct {
	CardKind          *CivilIDCardKind `json:"kind" bson:"kind,omitempty" validate:"required"`
	IDNumber          *string          `json:"idNumber" bson:"idNumber,omitempty" validate:"required,len=12"`
	Name              *string          `json:"name" bson:"name,omitempty" validate:"required"`
	DateOfBirth       *time.Time       `json:"dateOfBirth" bson:"dateOfBirth,omitempty" validate:"required"`
	PlaceOfBirth      *string          `json:"birthPlace" bson:"birthPlace,omitempty" validate:"required"`
	Ethnicity         *Ethinicty       `json:"ethnicity" bson:"ethinicity,omitempty" validate:"required"`
	Religion          *Religion        `json:"religion" bson:"religion,omitempty" validate:"required"`
	PermanentResident *string          `json:"permanentResident" bson:"permanentResident,omitempty" validate:"required"`
	TemporaryResident *string          `json:"temporaryResident" bson:"temporaryResident,omitempty"`
	CurrentResident   *string          `json:"currentResident" bson:"currentResident,omitempty" validate:"required"`
}

type Family struct {
	Mother  *FamilyMember  `json:"mother" bson:"mother,omitempty"`
	Father  *FamilyMember  `json:"father" bson:"father,omitempty"`
	Another []FamilyMember `json:"another,omitempty" bson:"another"`
}

type FamilyMember struct {
	Identity    CitizenIdentity `json:"indentity" bson:"identity,omitempty"`
	Name        string          `json:"name" bson:"name" validate:"required"`
	DateOfBirth time.Time       `json:"dateOfBirth" bson:"dateOfBirth" validate:"required"`
	Job         *string         `json:"job" bson:"job,omitempty"`
	Education   EducationDetail `json:"education" bson:"education,omitempty"`
}

type CivilHistory struct {
	BeforeRevolution *string `json:"beforeRevolution" bson:"beforeRevolution,omitempty"`
	AfterRevoolution *string `json:"afterRevolution" bson:"afterRevolution,omitempty"`
}

type FamilyInfo struct {
	Members []FamilyMember `json:"members" bson:"members,omitempty"`
}

type PoliticDetail struct {
	History       CivilHistory `json:"history" bson:"history,omitempty"`
	UnionJoinDate *time.Time   `json:"union" bson:"union,omitempty"`
	PartyJoinDate *time.Time   `json:"party" bson:"party,omitempty"`
}

type BasicType struct {
	UUID uuid.UUID `json:"uuid" bson:"uuid,omitempty" validate:"required"`
	Name string    `json:"name" bson:"name,omitempty" validate:"required"`
}

type Ethinicty BasicType

type Religion BasicType
