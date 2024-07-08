package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	CandidateSigningCommit[Op_Value_T any] struct {
		Time          *time.Time                    `bson:"time"`
		CamdidateUUID *uuid.UUID                    `bson:"candidateUUID"`
		Operations    []*JsonPatchRawValueOperation `bson:"operations"` //[]*JsonPatchOperation
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
