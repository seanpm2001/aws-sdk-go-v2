// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the description of specific Amazon FSx for Lustre or Amazon File Cache
// data repository associations, if one or more AssociationIds values are provided
// in the request, or if filters are used in the request. Data repository
// associations are supported on Amazon File Cache resources and all FSx for Lustre
// 2.12 and 2,15 file systems, excluding scratch_1 deployment type. You can use
// filters to narrow the response to include just data repository associations for
// specific file systems (use the file-system-id filter with the ID of the file
// system) or caches (use the file-cache-id filter with the ID of the cache), or
// data repository associations for a specific repository type (use the
// data-repository-type filter with a value of S3 or NFS ). If you don't use
// filters, the response returns all data repository associations owned by your
// Amazon Web Services account in the Amazon Web Services Region of the endpoint
// that you're calling. When retrieving all data repository associations, you can
// paginate the response by using the optional MaxResults parameter to limit the
// number of data repository associations returned in a response. If more data
// repository associations remain, a NextToken value is returned in the response.
// In this case, send a later request with the NextToken request parameter set to
// the value of NextToken from the last response.
func (c *Client) DescribeDataRepositoryAssociations(ctx context.Context, params *DescribeDataRepositoryAssociationsInput, optFns ...func(*Options)) (*DescribeDataRepositoryAssociationsOutput, error) {
	if params == nil {
		params = &DescribeDataRepositoryAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDataRepositoryAssociations", params, optFns, c.addOperationDescribeDataRepositoryAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDataRepositoryAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDataRepositoryAssociationsInput struct {

	// IDs of the data repository associations whose descriptions you want to retrieve
	// (String).
	AssociationIds []string

	// A list of Filter elements.
	Filters []types.Filter

	// The maximum number of resources to return in the response. This value must be
	// an integer greater than zero.
	MaxResults *int32

	// (Optional) Opaque pagination token returned from a previous operation (String).
	// If present, this token indicates from what point you can continue processing the
	// request, where the previous NextToken value left off.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeDataRepositoryAssociationsOutput struct {

	// An array of one or more data repository association descriptions.
	Associations []types.DataRepositoryAssociation

	// (Optional) Opaque pagination token returned from a previous operation (String).
	// If present, this token indicates from what point you can continue processing the
	// request, where the previous NextToken value left off.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDataRepositoryAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDataRepositoryAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDataRepositoryAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addDescribeDataRepositoryAssociationsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataRepositoryAssociations(options.Region), middleware.Before); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeDataRepositoryAssociationsAPIClient is a client that implements the
// DescribeDataRepositoryAssociations operation.
type DescribeDataRepositoryAssociationsAPIClient interface {
	DescribeDataRepositoryAssociations(context.Context, *DescribeDataRepositoryAssociationsInput, ...func(*Options)) (*DescribeDataRepositoryAssociationsOutput, error)
}

var _ DescribeDataRepositoryAssociationsAPIClient = (*Client)(nil)

// DescribeDataRepositoryAssociationsPaginatorOptions is the paginator options for
// DescribeDataRepositoryAssociations
type DescribeDataRepositoryAssociationsPaginatorOptions struct {
	// The maximum number of resources to return in the response. This value must be
	// an integer greater than zero.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDataRepositoryAssociationsPaginator is a paginator for
// DescribeDataRepositoryAssociations
type DescribeDataRepositoryAssociationsPaginator struct {
	options   DescribeDataRepositoryAssociationsPaginatorOptions
	client    DescribeDataRepositoryAssociationsAPIClient
	params    *DescribeDataRepositoryAssociationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDataRepositoryAssociationsPaginator returns a new
// DescribeDataRepositoryAssociationsPaginator
func NewDescribeDataRepositoryAssociationsPaginator(client DescribeDataRepositoryAssociationsAPIClient, params *DescribeDataRepositoryAssociationsInput, optFns ...func(*DescribeDataRepositoryAssociationsPaginatorOptions)) *DescribeDataRepositoryAssociationsPaginator {
	if params == nil {
		params = &DescribeDataRepositoryAssociationsInput{}
	}

	options := DescribeDataRepositoryAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDataRepositoryAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDataRepositoryAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDataRepositoryAssociations page.
func (p *DescribeDataRepositoryAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDataRepositoryAssociationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeDataRepositoryAssociations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeDataRepositoryAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "DescribeDataRepositoryAssociations",
	}
}

type opDescribeDataRepositoryAssociationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeDataRepositoryAssociationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeDataRepositoryAssociationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "fsx"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "fsx"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("fsx")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addDescribeDataRepositoryAssociationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeDataRepositoryAssociationsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
