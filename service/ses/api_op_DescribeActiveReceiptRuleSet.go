// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the metadata and receipt rules for the receipt rule set that is
// currently active. For information about setting up receipt rule sets, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rule-set.html).
// You can execute this operation no more than once per second.
func (c *Client) DescribeActiveReceiptRuleSet(ctx context.Context, params *DescribeActiveReceiptRuleSetInput, optFns ...func(*Options)) (*DescribeActiveReceiptRuleSetOutput, error) {
	if params == nil {
		params = &DescribeActiveReceiptRuleSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeActiveReceiptRuleSet", params, optFns, addOperationDescribeActiveReceiptRuleSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeActiveReceiptRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return the metadata and receipt rules for the receipt
// rule set that is currently active. You use receipt rule sets to receive email
// with Amazon SES. For more information, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type DescribeActiveReceiptRuleSetInput struct {
}

// Represents the metadata and receipt rules for the receipt rule set that is
// currently active.
type DescribeActiveReceiptRuleSetOutput struct {

	// The metadata for the currently active receipt rule set. The metadata consists of
	// the rule set name and a timestamp of when the rule set was created.
	Metadata *types.ReceiptRuleSetMetadata

	// The receipt rules that belong to the active rule set.
	Rules []types.ReceiptRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeActiveReceiptRuleSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeActiveReceiptRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeActiveReceiptRuleSet{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeActiveReceiptRuleSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeActiveReceiptRuleSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DescribeActiveReceiptRuleSet",
	}
}