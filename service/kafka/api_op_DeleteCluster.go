// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the MSK cluster specified by the Amazon Resource Name (ARN) in the
// request.
func (c *Client) DeleteCluster(ctx context.Context, params *DeleteClusterInput, optFns ...func(*Options)) (*DeleteClusterOutput, error) {
	if params == nil {
		params = &DeleteClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCluster", params, optFns, addOperationDeleteClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteClusterInput struct {

	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	//
	// This member is required.
	ClusterArn *string

	// The current version of the MSK cluster.
	CurrentVersion *string
}

type DeleteClusterOutput struct {

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// The state of the cluster. The possible states are ACTIVE, CREATING, DELETING,
	// FAILED, HEALING, MAINTENANCE, REBOOTING_BROKER, and UPDATING.
	State types.ClusterState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteCluster{}, middleware.After)
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
	if err = addOpDeleteClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "DeleteCluster",
	}
}