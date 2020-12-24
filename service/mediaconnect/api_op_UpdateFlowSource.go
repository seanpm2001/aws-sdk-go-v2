// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the source of a flow.
func (c *Client) UpdateFlowSource(ctx context.Context, params *UpdateFlowSourceInput, optFns ...func(*Options)) (*UpdateFlowSourceOutput, error) {
	if params == nil {
		params = &UpdateFlowSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFlowSource", params, optFns, addOperationUpdateFlowSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFlowSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update the source of a flow.
type UpdateFlowSourceInput struct {

	// The flow that is associated with the source that you want to update.
	//
	// This member is required.
	FlowArn *string

	// The ARN of the source that you want to update.
	//
	// This member is required.
	SourceArn *string

	// The type of encryption used on the content ingested from this source.
	Decryption *types.UpdateEncryption

	// A description for the source. This value is not used or seen outside of the
	// current AWS Elemental MediaConnect account.
	Description *string

	// The ARN of the entitlement that allows you to subscribe to this flow. The
	// entitlement is set by the flow originator, and the ARN is generated as part of
	// the originator's flow.
	EntitlementArn *string

	// The port that the flow will be listening on for incoming content.
	IngestPort int32

	// The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
	MaxBitrate int32

	// The maximum latency in milliseconds. This parameter applies only to RIST-based
	// and Zixi-based streams.
	MaxLatency int32

	// The protocol that is used by the source.
	Protocol types.Protocol

	// The stream ID that you want to use for this transport. This parameter applies
	// only to Zixi-based streams.
	StreamId *string

	// The name of the VPC Interface to configure this Source with.
	VpcInterfaceName *string

	// The range of IP addresses that should be allowed to contribute content to your
	// source. These IP addresses should be in the form of a Classless Inter-Domain
	// Routing (CIDR) block; for example, 10.0.0.0/16.
	WhitelistCidr *string
}

type UpdateFlowSourceOutput struct {

	// The ARN of the flow that you want to update.
	FlowArn *string

	// The settings for the source of the flow.
	Source *types.Source

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateFlowSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFlowSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFlowSource{}, middleware.After)
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
	if err = addOpUpdateFlowSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFlowSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFlowSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "UpdateFlowSource",
	}
}