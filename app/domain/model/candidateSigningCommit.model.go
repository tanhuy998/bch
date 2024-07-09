package model

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CandidateSigningCommit struct {
		ID            *primitive.ObjectID `json:"commit" bson:"_id,omitempty"`
		Time          *time.Time          `bson:"time"`
		CandidateUUID *uuid.UUID          `bson:"candidateUUID"`
		Operations    string              `bson:"operations"` //[]*JsonPatchOperation
	}
)

const (
	OP_ADD     OperationKind = "add"
	OP_REPLACE OperationKind = "replace"
	OP_REMOVE  OperationKind = "remove"
)

// func (this *CandidateSigningCommit) operationFilter(opKind OperationKind) (ret []*JsonPatchOperation) {

// 	for _, op := range this.Operations {

// 		if op.Op == opKind {

// 			ret = append(ret, op)
// 		}
// 	}

// 	return ret
// }

// func (this *CandidateSigningCommit) OperationAdd() []*JsonPatchOperation {

// 	return this.operationFilter(OP_ADD)
// }

// func (this *CandidateSigningCommit) OperationReplace() []*JsonPatchOperation {

// 	return this.operationFilter(OP_REPLACE)
// }

// func (this *CandidateSigningCommit) OperationRemove() []*JsonPatchOperation {

// 	return this.operationFilter(OP_REMOVE)
// }
