// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables Security Hub in your account only in the current Region. To disable
// Security Hub in all Regions, you must submit one request per Region where you
// have enabled Security Hub. When you disable Security Hub for an administrator
// account, it doesn't disable Security Hub for any associated member accounts.
// When you disable Security Hub, your existing findings and insights and any
// Security Hub configuration settings are deleted after 90 days and cannot be
// recovered. Any standards that were enabled are disabled, and your administrator
// and member account associations are removed. If you want to save your existing
// findings, you must export them before you disable Security Hub.
func (c *Client) DisableSecurityHub(ctx context.Context, params *DisableSecurityHubInput, optFns ...func(*Options)) (*DisableSecurityHubOutput, error) {
	if params == nil {
		params = &DisableSecurityHubInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableSecurityHub", params, optFns, addOperationDisableSecurityHubMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableSecurityHubOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableSecurityHubInput struct {
}

type DisableSecurityHubOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisableSecurityHubMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisableSecurityHub{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisableSecurityHub{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableSecurityHub(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableSecurityHub(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DisableSecurityHub",
	}
}
