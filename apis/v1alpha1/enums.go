// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AuthenticationType string

const (
	AuthenticationType_DIRECTORY_IDENTITY AuthenticationType = "DIRECTORY_IDENTITY"
)

type CalculationExecutionState string

const (
	CalculationExecutionState_CREATING  CalculationExecutionState = "CREATING"
	CalculationExecutionState_CREATED   CalculationExecutionState = "CREATED"
	CalculationExecutionState_QUEUED    CalculationExecutionState = "QUEUED"
	CalculationExecutionState_RUNNING   CalculationExecutionState = "RUNNING"
	CalculationExecutionState_CANCELING CalculationExecutionState = "CANCELING"
	CalculationExecutionState_CANCELED  CalculationExecutionState = "CANCELED"
	CalculationExecutionState_COMPLETED CalculationExecutionState = "COMPLETED"
	CalculationExecutionState_FAILED    CalculationExecutionState = "FAILED"
)

type CapacityAllocationStatus string

const (
	CapacityAllocationStatus_PENDING   CapacityAllocationStatus = "PENDING"
	CapacityAllocationStatus_SUCCEEDED CapacityAllocationStatus = "SUCCEEDED"
	CapacityAllocationStatus_FAILED    CapacityAllocationStatus = "FAILED"
)

type CapacityReservationStatus string

const (
	CapacityReservationStatus_PENDING        CapacityReservationStatus = "PENDING"
	CapacityReservationStatus_ACTIVE         CapacityReservationStatus = "ACTIVE"
	CapacityReservationStatus_CANCELLING     CapacityReservationStatus = "CANCELLING"
	CapacityReservationStatus_CANCELLED      CapacityReservationStatus = "CANCELLED"
	CapacityReservationStatus_FAILED         CapacityReservationStatus = "FAILED"
	CapacityReservationStatus_UPDATE_PENDING CapacityReservationStatus = "UPDATE_PENDING"
)

type ColumnNullable string

const (
	ColumnNullable_NOT_NULL ColumnNullable = "NOT_NULL"
	ColumnNullable_NULLABLE ColumnNullable = "NULLABLE"
	ColumnNullable_UNKNOWN  ColumnNullable = "UNKNOWN"
)

type DataCatalogType string

const (
	DataCatalogType_LAMBDA DataCatalogType = "LAMBDA"
	DataCatalogType_GLUE   DataCatalogType = "GLUE"
	DataCatalogType_HIVE   DataCatalogType = "HIVE"
)

type EncryptionOption string

const (
	EncryptionOption_SSE_S3  EncryptionOption = "SSE_S3"
	EncryptionOption_SSE_KMS EncryptionOption = "SSE_KMS"
	EncryptionOption_CSE_KMS EncryptionOption = "CSE_KMS"
)

type ExecutorState string

const (
	ExecutorState_CREATING    ExecutorState = "CREATING"
	ExecutorState_CREATED     ExecutorState = "CREATED"
	ExecutorState_REGISTERED  ExecutorState = "REGISTERED"
	ExecutorState_TERMINATING ExecutorState = "TERMINATING"
	ExecutorState_TERMINATED  ExecutorState = "TERMINATED"
	ExecutorState_FAILED      ExecutorState = "FAILED"
)

type ExecutorType string

const (
	ExecutorType_COORDINATOR ExecutorType = "COORDINATOR"
	ExecutorType_GATEWAY     ExecutorType = "GATEWAY"
	ExecutorType_WORKER      ExecutorType = "WORKER"
)

type NotebookType string

const (
	NotebookType_IPYNB NotebookType = "IPYNB"
)

type QueryExecutionState string

const (
	QueryExecutionState_QUEUED    QueryExecutionState = "QUEUED"
	QueryExecutionState_RUNNING   QueryExecutionState = "RUNNING"
	QueryExecutionState_SUCCEEDED QueryExecutionState = "SUCCEEDED"
	QueryExecutionState_FAILED    QueryExecutionState = "FAILED"
	QueryExecutionState_CANCELLED QueryExecutionState = "CANCELLED"
)

type S3ACLOption string

const (
	S3ACLOption_BUCKET_OWNER_FULL_CONTROL S3ACLOption = "BUCKET_OWNER_FULL_CONTROL"
)

type SessionState string

const (
	SessionState_CREATING    SessionState = "CREATING"
	SessionState_CREATED     SessionState = "CREATED"
	SessionState_IDLE        SessionState = "IDLE"
	SessionState_BUSY        SessionState = "BUSY"
	SessionState_TERMINATING SessionState = "TERMINATING"
	SessionState_TERMINATED  SessionState = "TERMINATED"
	SessionState_DEGRADED    SessionState = "DEGRADED"
	SessionState_FAILED      SessionState = "FAILED"
)

type StatementType string

const (
	StatementType_DDL     StatementType = "DDL"
	StatementType_DML     StatementType = "DML"
	StatementType_UTILITY StatementType = "UTILITY"
)

type ThrottleReason string

const (
	ThrottleReason_CONCURRENT_QUERY_LIMIT_EXCEEDED ThrottleReason = "CONCURRENT_QUERY_LIMIT_EXCEEDED"
)

type WorkGroupState string

const (
	WorkGroupState_ENABLED  WorkGroupState = "ENABLED"
	WorkGroupState_DISABLED WorkGroupState = "DISABLED"
)
