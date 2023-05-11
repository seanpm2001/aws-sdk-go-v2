// Code generated by smithy-go-codegen DO NOT EDIT.

package ivsrealtime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivsrealtime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all participants in a specified stage session.
func (c *Client) ListParticipants(ctx context.Context, params *ListParticipantsInput, optFns ...func(*Options)) (*ListParticipantsOutput, error) {
	if params == nil {
		params = &ListParticipantsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListParticipants", params, optFns, c.addOperationListParticipantsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListParticipantsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListParticipantsInput struct {

	// ID of the session within the stage.
	//
	// This member is required.
	SessionId *string

	// Stage ARN.
	//
	// This member is required.
	StageArn *string

	// Filters the response list to only show participants who published during the
	// stage session. Only one of filterByUserId , filterByPublished , or filterByState
	// can be provided per request.
	FilterByPublished bool

	// Filters the response list to only show participants in the specified state.
	// Only one of filterByUserId , filterByPublished , or filterByState can be
	// provided per request.
	FilterByState types.ParticipantState

	// Filters the response list to match the specified user ID. Only one of
	// filterByUserId , filterByPublished , or filterByState can be provided per
	// request. A userId is a customer-assigned name to help identify the token; this
	// can be used to link a participant to a user in the customer’s own systems.
	FilterByUserId *string

	// Maximum number of results to return. Default: 50.
	MaxResults int32

	// The first participant to retrieve. This is used for pagination; see the
	// nextToken response field.
	NextToken *string

	noSmithyDocumentSerde
}

type ListParticipantsOutput struct {

	// List of the matching participants (summary information only).
	//
	// This member is required.
	Participants []types.ParticipantSummary

	// If there are more rooms than maxResults , use nextToken in the request to get
	// the next set.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListParticipantsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListParticipants{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListParticipants{}, middleware.After)
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
	if err = addOpListParticipantsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListParticipants(options.Region), middleware.Before); err != nil {
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

// ListParticipantsAPIClient is a client that implements the ListParticipants
// operation.
type ListParticipantsAPIClient interface {
	ListParticipants(context.Context, *ListParticipantsInput, ...func(*Options)) (*ListParticipantsOutput, error)
}

var _ ListParticipantsAPIClient = (*Client)(nil)

// ListParticipantsPaginatorOptions is the paginator options for ListParticipants
type ListParticipantsPaginatorOptions struct {
	// Maximum number of results to return. Default: 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListParticipantsPaginator is a paginator for ListParticipants
type ListParticipantsPaginator struct {
	options   ListParticipantsPaginatorOptions
	client    ListParticipantsAPIClient
	params    *ListParticipantsInput
	nextToken *string
	firstPage bool
}

// NewListParticipantsPaginator returns a new ListParticipantsPaginator
func NewListParticipantsPaginator(client ListParticipantsAPIClient, params *ListParticipantsInput, optFns ...func(*ListParticipantsPaginatorOptions)) *ListParticipantsPaginator {
	if params == nil {
		params = &ListParticipantsInput{}
	}

	options := ListParticipantsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListParticipantsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListParticipantsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListParticipants page.
func (p *ListParticipantsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListParticipantsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListParticipants(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListParticipants(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "ListParticipants",
	}
}
