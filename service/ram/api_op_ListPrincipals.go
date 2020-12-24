// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the principals that you have shared resources with or that have shared
// resources with you.
func (c *Client) ListPrincipals(ctx context.Context, params *ListPrincipalsInput, optFns ...func(*Options)) (*ListPrincipalsOutput, error) {
	if params == nil {
		params = &ListPrincipalsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPrincipals", params, optFns, addOperationListPrincipalsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPrincipalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPrincipalsInput struct {

	// The type of owner.
	//
	// This member is required.
	ResourceOwner types.ResourceOwner

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The principals.
	Principals []string

	// The Amazon Resource Name (ARN) of the resource.
	ResourceArn *string

	// The Amazon Resource Names (ARN) of the resource shares.
	ResourceShareArns []string

	// The resource type. Valid values: codebuild:Project | codebuild:ReportGroup |
	// ec2:CapacityReservation | ec2:DedicatedHost | ec2:Subnet |
	// ec2:TrafficMirrorTarget | ec2:TransitGateway | imagebuilder:Component |
	// imagebuilder:Image | imagebuilder:ImageRecipe |
	// license-manager:LicenseConfiguration I resource-groups:Group | rds:Cluster |
	// route53resolver:ResolverRule
	ResourceType *string
}

type ListPrincipalsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The principals.
	Principals []types.Principal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPrincipalsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPrincipals{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPrincipals{}, middleware.After)
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
	if err = addOpListPrincipalsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPrincipals(options.Region), middleware.Before); err != nil {
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

// ListPrincipalsAPIClient is a client that implements the ListPrincipals
// operation.
type ListPrincipalsAPIClient interface {
	ListPrincipals(context.Context, *ListPrincipalsInput, ...func(*Options)) (*ListPrincipalsOutput, error)
}

var _ ListPrincipalsAPIClient = (*Client)(nil)

// ListPrincipalsPaginatorOptions is the paginator options for ListPrincipals
type ListPrincipalsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPrincipalsPaginator is a paginator for ListPrincipals
type ListPrincipalsPaginator struct {
	options   ListPrincipalsPaginatorOptions
	client    ListPrincipalsAPIClient
	params    *ListPrincipalsInput
	nextToken *string
	firstPage bool
}

// NewListPrincipalsPaginator returns a new ListPrincipalsPaginator
func NewListPrincipalsPaginator(client ListPrincipalsAPIClient, params *ListPrincipalsInput, optFns ...func(*ListPrincipalsPaginatorOptions)) *ListPrincipalsPaginator {
	options := ListPrincipalsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListPrincipalsInput{}
	}

	return &ListPrincipalsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPrincipalsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListPrincipals page.
func (p *ListPrincipalsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPrincipalsOutput, error) {
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

	result, err := p.client.ListPrincipals(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListPrincipals(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "ListPrincipals",
	}
}