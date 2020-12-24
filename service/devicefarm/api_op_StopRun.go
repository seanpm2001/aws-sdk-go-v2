// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates a stop request for the current test run. AWS Device Farm immediately
// stops the run on devices where tests have not started. You are not billed for
// these devices. On devices where tests have started executing, setup suite and
// teardown suite tests run to completion on those devices. You are billed for
// setup, teardown, and any tests that were in progress or already completed.
func (c *Client) StopRun(ctx context.Context, params *StopRunInput, optFns ...func(*Options)) (*StopRunOutput, error) {
	if params == nil {
		params = &StopRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopRun", params, optFns, addOperationStopRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to stop a specific run.
type StopRunInput struct {

	// Represents the Amazon Resource Name (ARN) of the Device Farm run to stop.
	//
	// This member is required.
	Arn *string
}

// Represents the results of your stop run attempt.
type StopRunOutput struct {

	// The run that was stopped.
	Run *types.Run

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopRun{}, middleware.After)
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
	if err = addOpStopRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "StopRun",
	}
}