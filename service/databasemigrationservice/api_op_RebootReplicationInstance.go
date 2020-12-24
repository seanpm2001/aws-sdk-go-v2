// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Reboots a replication instance. Rebooting results in a momentary outage, until
// the replication instance becomes available again.
func (c *Client) RebootReplicationInstance(ctx context.Context, params *RebootReplicationInstanceInput, optFns ...func(*Options)) (*RebootReplicationInstanceOutput, error) {
	if params == nil {
		params = &RebootReplicationInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RebootReplicationInstance", params, optFns, addOperationRebootReplicationInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RebootReplicationInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RebootReplicationInstanceInput struct {

	// The Amazon Resource Name (ARN) of the replication instance.
	//
	// This member is required.
	ReplicationInstanceArn *string

	// If this parameter is true, the reboot is conducted through a Multi-AZ failover.
	// (If the instance isn't configured for Multi-AZ, then you can't specify true.)
	ForceFailover *bool
}

type RebootReplicationInstanceOutput struct {

	// The replication instance that is being rebooted.
	ReplicationInstance *types.ReplicationInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRebootReplicationInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRebootReplicationInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRebootReplicationInstance{}, middleware.After)
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
	if err = addOpRebootReplicationInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRebootReplicationInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRebootReplicationInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "RebootReplicationInstance",
	}
}