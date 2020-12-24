// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the permissions for a Systems Manager document. If you created the
// document, you are the owner. If a document is shared, it can either be shared
// privately (by specifying a user's AWS account ID) or publicly (All).
func (c *Client) DescribeDocumentPermission(ctx context.Context, params *DescribeDocumentPermissionInput, optFns ...func(*Options)) (*DescribeDocumentPermissionOutput, error) {
	if params == nil {
		params = &DescribeDocumentPermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDocumentPermission", params, optFns, addOperationDescribeDocumentPermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDocumentPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDocumentPermissionInput struct {

	// The name of the document for which you are the owner.
	//
	// This member is required.
	Name *string

	// The permission type for the document. The permission type can be Share.
	//
	// This member is required.
	PermissionType types.DocumentPermissionType
}

type DescribeDocumentPermissionOutput struct {

	// The account IDs that have permission to use this document. The ID can be either
	// an AWS account or All.
	AccountIds []string

	// A list of AWS accounts where the current document is shared and the version
	// shared with each account.
	AccountSharingInfoList []types.AccountSharingInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDocumentPermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDocumentPermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDocumentPermission{}, middleware.After)
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
	if err = addOpDescribeDocumentPermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDocumentPermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDocumentPermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeDocumentPermission",
	}
}