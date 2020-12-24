// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information to AWS CodePipeline about new revisions to a source.
func (c *Client) PutActionRevision(ctx context.Context, params *PutActionRevisionInput, optFns ...func(*Options)) (*PutActionRevisionOutput, error) {
	if params == nil {
		params = &PutActionRevisionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutActionRevision", params, optFns, addOperationPutActionRevisionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutActionRevisionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PutActionRevision action.
type PutActionRevisionInput struct {

	// The name of the action that processes the revision.
	//
	// This member is required.
	ActionName *string

	// Represents information about the version (or revision) of an action.
	//
	// This member is required.
	ActionRevision *types.ActionRevision

	// The name of the pipeline that starts processing the revision to the source.
	//
	// This member is required.
	PipelineName *string

	// The name of the stage that contains the action that acts on the revision.
	//
	// This member is required.
	StageName *string
}

// Represents the output of a PutActionRevision action.
type PutActionRevisionOutput struct {

	// Indicates whether the artifact revision was previously used in an execution of
	// the specified pipeline.
	NewRevision bool

	// The ID of the current workflow state of the pipeline.
	PipelineExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutActionRevisionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutActionRevision{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutActionRevision{}, middleware.After)
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
	if err = addOpPutActionRevisionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutActionRevision(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutActionRevision(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "PutActionRevision",
	}
}