// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the dedicated IP addresses that are associated with your AWS account.
func (c *Client) GetDedicatedIps(ctx context.Context, params *GetDedicatedIpsInput, optFns ...func(*Options)) (*GetDedicatedIpsOutput, error) {
	if params == nil {
		params = &GetDedicatedIpsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDedicatedIps", params, optFns, addOperationGetDedicatedIpsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDedicatedIpsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain more information about dedicated IP pools.
type GetDedicatedIpsInput struct {

	// A token returned from a previous call to GetDedicatedIps to indicate the
	// position of the dedicated IP pool in the list of IP pools.
	NextToken *string

	// The number of results to show in a single call to GetDedicatedIpsRequest. If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results.
	PageSize *int32

	// The name of the IP pool that the dedicated IP address is associated with.
	PoolName *string
}

// Information about the dedicated IP addresses that are associated with your AWS
// account.
type GetDedicatedIpsOutput struct {

	// A list of dedicated IP addresses that are associated with your AWS account.
	DedicatedIps []types.DedicatedIp

	// A token that indicates that there are additional dedicated IP addresses to list.
	// To view additional addresses, issue another request to GetDedicatedIps, passing
	// this token in the NextToken parameter.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDedicatedIpsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDedicatedIps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDedicatedIps{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDedicatedIps(options.Region), middleware.Before); err != nil {
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

// GetDedicatedIpsAPIClient is a client that implements the GetDedicatedIps
// operation.
type GetDedicatedIpsAPIClient interface {
	GetDedicatedIps(context.Context, *GetDedicatedIpsInput, ...func(*Options)) (*GetDedicatedIpsOutput, error)
}

var _ GetDedicatedIpsAPIClient = (*Client)(nil)

// GetDedicatedIpsPaginatorOptions is the paginator options for GetDedicatedIps
type GetDedicatedIpsPaginatorOptions struct {
	// The number of results to show in a single call to GetDedicatedIpsRequest. If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetDedicatedIpsPaginator is a paginator for GetDedicatedIps
type GetDedicatedIpsPaginator struct {
	options   GetDedicatedIpsPaginatorOptions
	client    GetDedicatedIpsAPIClient
	params    *GetDedicatedIpsInput
	nextToken *string
	firstPage bool
}

// NewGetDedicatedIpsPaginator returns a new GetDedicatedIpsPaginator
func NewGetDedicatedIpsPaginator(client GetDedicatedIpsAPIClient, params *GetDedicatedIpsInput, optFns ...func(*GetDedicatedIpsPaginatorOptions)) *GetDedicatedIpsPaginator {
	options := GetDedicatedIpsPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetDedicatedIpsInput{}
	}

	return &GetDedicatedIpsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetDedicatedIpsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetDedicatedIps page.
func (p *GetDedicatedIpsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetDedicatedIpsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.GetDedicatedIps(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetDedicatedIps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetDedicatedIps",
	}
}