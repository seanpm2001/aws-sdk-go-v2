// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/worklink/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of fleets for the current account and Region.
func (c *Client) ListFleets(ctx context.Context, params *ListFleetsInput, optFns ...func(*Options)) (*ListFleetsOutput, error) {
	if params == nil {
		params = &ListFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFleets", params, optFns, addOperationListFleetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFleetsInput struct {

	// The maximum number of results to be included in the next page.
	MaxResults *int32

	// The pagination token used to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string
}

type ListFleetsOutput struct {

	// The summary list of the fleets.
	FleetSummaryList []types.FleetSummary

	// The pagination token used to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFleetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFleets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFleets(options.Region), middleware.Before); err != nil {
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

// ListFleetsAPIClient is a client that implements the ListFleets operation.
type ListFleetsAPIClient interface {
	ListFleets(context.Context, *ListFleetsInput, ...func(*Options)) (*ListFleetsOutput, error)
}

var _ ListFleetsAPIClient = (*Client)(nil)

// ListFleetsPaginatorOptions is the paginator options for ListFleets
type ListFleetsPaginatorOptions struct {
	// The maximum number of results to be included in the next page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFleetsPaginator is a paginator for ListFleets
type ListFleetsPaginator struct {
	options   ListFleetsPaginatorOptions
	client    ListFleetsAPIClient
	params    *ListFleetsInput
	nextToken *string
	firstPage bool
}

// NewListFleetsPaginator returns a new ListFleetsPaginator
func NewListFleetsPaginator(client ListFleetsAPIClient, params *ListFleetsInput, optFns ...func(*ListFleetsPaginatorOptions)) *ListFleetsPaginator {
	options := ListFleetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListFleetsInput{}
	}

	return &ListFleetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFleetsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListFleets page.
func (p *ListFleetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFleetsOutput, error) {
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

	result, err := p.client.ListFleets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFleets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "ListFleets",
	}
}