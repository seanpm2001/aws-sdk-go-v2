// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the metadata document for an existing SAML provider resource object.
// This operation requires Signature Version 4
// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
func (c *Client) UpdateSAMLProvider(ctx context.Context, params *UpdateSAMLProviderInput, optFns ...func(*Options)) (*UpdateSAMLProviderOutput, error) {
	if params == nil {
		params = &UpdateSAMLProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSAMLProvider", params, optFns, addOperationUpdateSAMLProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSAMLProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSAMLProviderInput struct {

	// An XML document generated by an identity provider (IdP) that supports SAML 2.0.
	// The document includes the issuer's name, expiration information, and keys that
	// can be used to validate the SAML authentication response (assertions) that are
	// received from the IdP. You must generate the metadata document using the
	// identity management software that is used as your organization's IdP.
	//
	// This member is required.
	SAMLMetadataDocument *string

	// The Amazon Resource Name (ARN) of the SAML provider to update. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	SAMLProviderArn *string
}

// Contains the response to a successful UpdateSAMLProvider request.
type UpdateSAMLProviderOutput struct {

	// The Amazon Resource Name (ARN) of the SAML provider that was updated.
	SAMLProviderArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateSAMLProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateSAMLProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateSAMLProvider{}, middleware.After)
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
	if err = addOpUpdateSAMLProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSAMLProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSAMLProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateSAMLProvider",
	}
}