// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation provides status information on enabling or disabling Health to
// work with your organization. To call this operation, you must use the
// organization's management account.
func (c *Client) DescribeHealthServiceStatusForOrganization(ctx context.Context, params *DescribeHealthServiceStatusForOrganizationInput, optFns ...func(*Options)) (*DescribeHealthServiceStatusForOrganizationOutput, error) {
	if params == nil {
		params = &DescribeHealthServiceStatusForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHealthServiceStatusForOrganization", params, optFns, c.addOperationDescribeHealthServiceStatusForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHealthServiceStatusForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHealthServiceStatusForOrganizationInput struct {
	noSmithyDocumentSerde
}

type DescribeHealthServiceStatusForOrganizationOutput struct {

	// Information about the status of enabling or disabling the Health organizational
	// view feature in your organization. Valid values are ENABLED | DISABLED | PENDING
	// .
	HealthServiceAccessStatusForOrganization *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHealthServiceStatusForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeHealthServiceStatusForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeHealthServiceStatusForOrganization{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHealthServiceStatusForOrganization(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opDescribeHealthServiceStatusForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "DescribeHealthServiceStatusForOrganization",
	}
}
