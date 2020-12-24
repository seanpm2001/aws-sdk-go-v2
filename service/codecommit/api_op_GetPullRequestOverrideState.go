// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about whether approval rules have been set aside
// (overridden) for a pull request, and if so, the Amazon Resource Name (ARN) of
// the user or identity that overrode the rules and their requirements for the pull
// request.
func (c *Client) GetPullRequestOverrideState(ctx context.Context, params *GetPullRequestOverrideStateInput, optFns ...func(*Options)) (*GetPullRequestOverrideStateOutput, error) {
	if params == nil {
		params = &GetPullRequestOverrideStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPullRequestOverrideState", params, optFns, addOperationGetPullRequestOverrideStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPullRequestOverrideStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPullRequestOverrideStateInput struct {

	// The ID of the pull request for which you want to get information about whether
	// approval rules have been set aside (overridden).
	//
	// This member is required.
	PullRequestId *string

	// The system-generated ID of the revision for the pull request. To retrieve the
	// most recent revision ID, use GetPullRequest.
	//
	// This member is required.
	RevisionId *string
}

type GetPullRequestOverrideStateOutput struct {

	// A Boolean value that indicates whether a pull request has had its rules set
	// aside (TRUE) or whether all approval rules still apply (FALSE).
	Overridden bool

	// The Amazon Resource Name (ARN) of the user or identity that overrode the rules
	// and their requirements for the pull request.
	Overrider *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetPullRequestOverrideStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPullRequestOverrideState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPullRequestOverrideState{}, middleware.After)
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
	if err = addOpGetPullRequestOverrideStateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPullRequestOverrideState(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPullRequestOverrideState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetPullRequestOverrideState",
	}
}