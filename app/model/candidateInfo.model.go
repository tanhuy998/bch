package model

import (
	"time"

	"github.com/google/uuid"
)

type CandidateSigningInfo struct {
	IModel
	Model
	//UUID       uuid.UUID       `json:"uuid" bson:"uuid" validate:"required"`
	Identitity CitizenIdentity `json:"civilIdentity" bson:"civilIdentity" validate:"required"`
	Job        string          `json:"job" bson:"job"`
	Education  EducationDetail `json:"education" bson:"education"`
	Family     FamilyInfo      `json:"family" bson:"family"`
}

type Citizen struct {
}

type CivilIDCardKind int

const (
	CMND CivilIDCardKind = 1
	CCCD CivilIDCardKind = 2
)

type EducationDetail struct {
	PrimarySchool   string    `json:"primarySchool" bson:"primarySchool"`
	SecondarySchool string    `json:"secondarySchool" bson:"secondarySchool"`
	HighSchool      string    `json:"highSchool" bson:"highSchool"`
	HighestGrade    byte      `json:"highestGrade" bson:"highestGrade"`
	College         string    `json:"college" bson:"college"`
	GraduateAt      time.Time `json:"graduateAt" bson:"graduateAt"`
	Expertise       string    `json:"expertise" bson:"expertise"`
}

type CitizenIdentity struct {
	CardKind          CivilIDCardKind `json:"kind" bson:"kind" validate:"required"`
	IDNumber          string          `json:"idNumber" bson:"idNumber" validate:"required,len=12"`
	Name              string          `json:"name" bson:"name" validate:"required"`
	DateOfBirth       time.Time       `json:"dateOfBirth" bson:"dateOfBirth" validate:"required"`
	PlaceOfBirth      string          `json:"birthPlace" bson:"birthPlace" validate:"required"`
	Ethinicity        Ethinicty       `json:"ethinicity" bson:"ethinicity" validate:"required"`
	Religion          Religion        `json:"religion" bson:"religion" validate:"required"`
	PermanentResident string          `json:"permanentResident" bson:"permanentResident" validate:"required"`
	TemporaryResident string          `json:"temporaryResident" bson:"temporaryResident"`
	CurrentResident   string          `json:"currentResident" bson:"currentResident" validate:"required"`
	Politic           PoliticDetail   `json:"politic" bson:"politic"`
}

type FamilyMember struct {
	Identity  CitizenIdentity `json:"indentity" bson:"identity"`
	Job       string          `json:"job" bson:"job"`
	Education EducationDetail `json:"education" bson:"education"`
}

type CivilHistory struct {
	BeforeRevolution string `json:"beforeRevolution" bson:"beforeRevolution"`
	AfterRevoolution string `json:"afterRevolution" bson:"afterRevolution"`
}

type FamilyInfo struct {
	Members []FamilyMember `json:"members" bson:"members"`
}

type PoliticDetail struct {
	History       CivilHistory `json:"history" bson:"history"`
	UnionJoinDate time.Time    `json:"union" bson:"union"`
	PartyJoinDate time.Time    `json:"party" bson:"party"`
}

type BasicType struct {
	UUID uuid.UUID `json:"uuid" bson:"uuid" validate:"required"`
	Name string    `json:"name" bson:"name" validate:"required"`
}

type Ethinicty BasicType

type Religion BasicType
