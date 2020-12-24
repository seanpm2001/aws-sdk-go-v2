// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Accepts a phone number and indicates whether the phone holder has opted out of
// receiving SMS messages from your account. You cannot send SMS messages to a
// number that is opted out. To resume sending messages, you can opt in the number
// by using the OptInPhoneNumber action.
func (c *Client) CheckIfPhoneNumberIsOptedOut(ctx context.Context, params *CheckIfPhoneNumberIsOptedOutInput, optFns ...func(*Options)) (*CheckIfPhoneNumberIsOptedOutOutput, error) {
	if params == nil {
		params = &CheckIfPhoneNumberIsOptedOutInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CheckIfPhoneNumberIsOptedOut", params, optFns, addOperationCheckIfPhoneNumberIsOptedOutMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CheckIfPhoneNumberIsOptedOutOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CheckIfPhoneNumberIsOptedOut action.
type CheckIfPhoneNumberIsOptedOutInput struct {

	// The phone number for which you want to check the opt out status.
	//
	// This member is required.
	PhoneNumber *string
}

// The response from the CheckIfPhoneNumberIsOptedOut action.
type CheckIfPhoneNumberIsOptedOutOutput struct {

	// Indicates whether the phone number is opted out:
	//
	// * true – The phone number is
	// opted out, meaning you cannot publish SMS messages to it.
	//
	// * false – The phone
	// number is opted in, meaning you can publish SMS messages to it.
	IsOptedOut bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCheckIfPhoneNumberIsOptedOutMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCheckIfPhoneNumberIsOptedOut{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCheckIfPhoneNumberIsOptedOut{}, middleware.After)
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
	if err = addOpCheckIfPhoneNumberIsOptedOutValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCheckIfPhoneNumberIsOptedOut(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCheckIfPhoneNumberIsOptedOut(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "CheckIfPhoneNumberIsOptedOut",
	}
}