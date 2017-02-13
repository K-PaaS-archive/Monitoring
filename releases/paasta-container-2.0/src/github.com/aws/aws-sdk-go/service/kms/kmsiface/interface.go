// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package kmsiface provides an interface for the AWS Key Management Service.
package kmsiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kms"
)

// KMSAPI is the interface type for kms.KMS.
type KMSAPI interface {
	CreateAliasRequest(*kms.CreateAliasInput) (*request.Request, *kms.CreateAliasOutput)

	CreateAlias(*kms.CreateAliasInput) (*kms.CreateAliasOutput, error)

	CreateGrantRequest(*kms.CreateGrantInput) (*request.Request, *kms.CreateGrantOutput)

	CreateGrant(*kms.CreateGrantInput) (*kms.CreateGrantOutput, error)

	CreateKeyRequest(*kms.CreateKeyInput) (*request.Request, *kms.CreateKeyOutput)

	CreateKey(*kms.CreateKeyInput) (*kms.CreateKeyOutput, error)

	DecryptRequest(*kms.DecryptInput) (*request.Request, *kms.DecryptOutput)

	Decrypt(*kms.DecryptInput) (*kms.DecryptOutput, error)

	DeleteAliasRequest(*kms.DeleteAliasInput) (*request.Request, *kms.DeleteAliasOutput)

	DeleteAlias(*kms.DeleteAliasInput) (*kms.DeleteAliasOutput, error)

	DescribeKeyRequest(*kms.DescribeKeyInput) (*request.Request, *kms.DescribeKeyOutput)

	DescribeKey(*kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error)

	DisableKeyRequest(*kms.DisableKeyInput) (*request.Request, *kms.DisableKeyOutput)

	DisableKey(*kms.DisableKeyInput) (*kms.DisableKeyOutput, error)

	DisableKeyRotationRequest(*kms.DisableKeyRotationInput) (*request.Request, *kms.DisableKeyRotationOutput)

	DisableKeyRotation(*kms.DisableKeyRotationInput) (*kms.DisableKeyRotationOutput, error)

	EnableKeyRequest(*kms.EnableKeyInput) (*request.Request, *kms.EnableKeyOutput)

	EnableKey(*kms.EnableKeyInput) (*kms.EnableKeyOutput, error)

	EnableKeyRotationRequest(*kms.EnableKeyRotationInput) (*request.Request, *kms.EnableKeyRotationOutput)

	EnableKeyRotation(*kms.EnableKeyRotationInput) (*kms.EnableKeyRotationOutput, error)

	EncryptRequest(*kms.EncryptInput) (*request.Request, *kms.EncryptOutput)

	Encrypt(*kms.EncryptInput) (*kms.EncryptOutput, error)

	GenerateDataKeyRequest(*kms.GenerateDataKeyInput) (*request.Request, *kms.GenerateDataKeyOutput)

	GenerateDataKey(*kms.GenerateDataKeyInput) (*kms.GenerateDataKeyOutput, error)

	GenerateDataKeyWithoutPlaintextRequest(*kms.GenerateDataKeyWithoutPlaintextInput) (*request.Request, *kms.GenerateDataKeyWithoutPlaintextOutput)

	GenerateDataKeyWithoutPlaintext(*kms.GenerateDataKeyWithoutPlaintextInput) (*kms.GenerateDataKeyWithoutPlaintextOutput, error)

	GenerateRandomRequest(*kms.GenerateRandomInput) (*request.Request, *kms.GenerateRandomOutput)

	GenerateRandom(*kms.GenerateRandomInput) (*kms.GenerateRandomOutput, error)

	GetKeyPolicyRequest(*kms.GetKeyPolicyInput) (*request.Request, *kms.GetKeyPolicyOutput)

	GetKeyPolicy(*kms.GetKeyPolicyInput) (*kms.GetKeyPolicyOutput, error)

	GetKeyRotationStatusRequest(*kms.GetKeyRotationStatusInput) (*request.Request, *kms.GetKeyRotationStatusOutput)

	GetKeyRotationStatus(*kms.GetKeyRotationStatusInput) (*kms.GetKeyRotationStatusOutput, error)

	ListAliasesRequest(*kms.ListAliasesInput) (*request.Request, *kms.ListAliasesOutput)

	ListAliases(*kms.ListAliasesInput) (*kms.ListAliasesOutput, error)

	ListAliasesPages(*kms.ListAliasesInput, func(*kms.ListAliasesOutput, bool) bool) error

	ListGrantsRequest(*kms.ListGrantsInput) (*request.Request, *kms.ListGrantsOutput)

	ListGrants(*kms.ListGrantsInput) (*kms.ListGrantsOutput, error)

	ListGrantsPages(*kms.ListGrantsInput, func(*kms.ListGrantsOutput, bool) bool) error

	ListKeyPoliciesRequest(*kms.ListKeyPoliciesInput) (*request.Request, *kms.ListKeyPoliciesOutput)

	ListKeyPolicies(*kms.ListKeyPoliciesInput) (*kms.ListKeyPoliciesOutput, error)

	ListKeyPoliciesPages(*kms.ListKeyPoliciesInput, func(*kms.ListKeyPoliciesOutput, bool) bool) error

	ListKeysRequest(*kms.ListKeysInput) (*request.Request, *kms.ListKeysOutput)

	ListKeys(*kms.ListKeysInput) (*kms.ListKeysOutput, error)

	ListKeysPages(*kms.ListKeysInput, func(*kms.ListKeysOutput, bool) bool) error

	PutKeyPolicyRequest(*kms.PutKeyPolicyInput) (*request.Request, *kms.PutKeyPolicyOutput)

	PutKeyPolicy(*kms.PutKeyPolicyInput) (*kms.PutKeyPolicyOutput, error)

	ReEncryptRequest(*kms.ReEncryptInput) (*request.Request, *kms.ReEncryptOutput)

	ReEncrypt(*kms.ReEncryptInput) (*kms.ReEncryptOutput, error)

	RetireGrantRequest(*kms.RetireGrantInput) (*request.Request, *kms.RetireGrantOutput)

	RetireGrant(*kms.RetireGrantInput) (*kms.RetireGrantOutput, error)

	RevokeGrantRequest(*kms.RevokeGrantInput) (*request.Request, *kms.RevokeGrantOutput)

	RevokeGrant(*kms.RevokeGrantInput) (*kms.RevokeGrantOutput, error)

	UpdateAliasRequest(*kms.UpdateAliasInput) (*request.Request, *kms.UpdateAliasOutput)

	UpdateAlias(*kms.UpdateAliasInput) (*kms.UpdateAliasOutput, error)

	UpdateKeyDescriptionRequest(*kms.UpdateKeyDescriptionInput) (*request.Request, *kms.UpdateKeyDescriptionOutput)

	UpdateKeyDescription(*kms.UpdateKeyDescriptionInput) (*kms.UpdateKeyDescriptionOutput, error)
}
