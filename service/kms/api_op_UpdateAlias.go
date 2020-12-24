// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates an existing AWS KMS alias with a different customer master key (CMK).
// Each alias is associated with only one CMK at a time, although a CMK can have
// multiple aliases. The alias and the CMK must be in the same AWS account and
// region. The current and new CMK must be the same type (both symmetric or both
// asymmetric), and they must have the same key usage (ENCRYPT_DECRYPT or
// SIGN_VERIFY). This restriction prevents errors in code that uses aliases. If you
// must assign an alias to a different type of CMK, use DeleteAlias to delete the
// old alias and CreateAlias to create a new alias. You cannot use UpdateAlias to
// change an alias name. To change an alias name, use DeleteAlias to delete the old
// alias and CreateAlias to create a new alias. Because an alias is not a property
// of a CMK, you can create, update, and delete the aliases of a CMK without
// affecting the CMK. Also, aliases do not appear in the response from the
// DescribeKey operation. To get the aliases of all CMKs in the account, use the
// ListAliases operation. The CMK that you use for this operation must be in a
// compatible key state. For details, see How Key State Affects Use of a Customer
// Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide. Cross-account use: No. You cannot
// perform this operation on a CMK in a different AWS account. Required
// permissions
//
// * kms:UpdateAlias
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// on the alias (IAM policy).
//
// * kms:UpdateAlias
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// on the current CMK (key policy).
//
// * kms:UpdateAlias
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// on the new CMK (key policy).
//
// For details, see Controlling access to aliases
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html#alias-access)
// in the AWS Key Management Service Developer Guide. Related operations:
//
// *
// CreateAlias
//
// * DeleteAlias
//
// * ListAliases
func (c *Client) UpdateAlias(ctx context.Context, params *UpdateAliasInput, optFns ...func(*Options)) (*UpdateAliasOutput, error) {
	if params == nil {
		params = &UpdateAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAlias", params, optFns, addOperationUpdateAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAliasInput struct {

	// Identifies the alias that is changing its CMK. This value must begin with alias/
	// followed by the alias name, such as alias/ExampleAlias. You cannot use
	// UpdateAlias to change the alias name.
	//
	// This member is required.
	AliasName *string

	// Identifies the customer managed CMK
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk)
	// to associate with the alias. You don't have permission to associate an alias
	// with an AWS managed CMK
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
	// The CMK must be in the same AWS account and Region as the alias. Also, the new
	// target CMK must be the same type as the current target CMK (both symmetric or
	// both asymmetric) and they must have the same key usage. Specify the key ID or
	// the Amazon Resource Name (ARN) of the CMK. For example:
	//
	// * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys or DescribeKey. To verify
	// that the alias is mapped to the correct CMK, use ListAliases.
	//
	// This member is required.
	TargetKeyId *string
}

type UpdateAliasOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAlias{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAlias(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "UpdateAlias",
	}
}