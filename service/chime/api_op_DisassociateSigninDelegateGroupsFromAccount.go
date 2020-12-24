// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates the specified sign-in delegate groups from the specified Amazon
// Chime account.
func (c *Client) DisassociateSigninDelegateGroupsFromAccount(ctx context.Context, params *DisassociateSigninDelegateGroupsFromAccountInput, optFns ...func(*Options)) (*DisassociateSigninDelegateGroupsFromAccountOutput, error) {
	if params == nil {
		params = &DisassociateSigninDelegateGroupsFromAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateSigninDelegateGroupsFromAccount", params, optFns, addOperationDisassociateSigninDelegateGroupsFromAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateSigninDelegateGroupsFromAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateSigninDelegateGroupsFromAccountInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The sign-in delegate group names.
	//
	// This member is required.
	GroupNames []string
}

type DisassociateSigninDelegateGroupsFromAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateSigninDelegateGroupsFromAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateSigninDelegateGroupsFromAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateSigninDelegateGroupsFromAccount{}, middleware.After)
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
	if err = addOpDisassociateSigninDelegateGroupsFromAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateSigninDelegateGroupsFromAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateSigninDelegateGroupsFromAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DisassociateSigninDelegateGroupsFromAccount",
	}
}