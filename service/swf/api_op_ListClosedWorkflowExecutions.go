// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of closed workflow executions in the specified domain that meet
// the filtering criteria. The results may be split into multiple pages. To
// retrieve subsequent pages, make the call again using the nextPageToken returned
// by the initial call. This operation is eventually consistent. The results are
// best effort and may not exactly reflect recent updates and changes. Access
// Control You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// * Use a Resource element with the domain name to limit
// the action to only specified domains.
//
// * Use an Action element to allow or deny
// permission to call this action.
//
// * Constrain the following parameters by using a
// Condition element with the appropriate keys.
//
// * tagFilter.tag: String
// constraint. The key is swf:tagFilter.tag.
//
// * typeFilter.name: String constraint.
// The key is swf:typeFilter.name.
//
// * typeFilter.version: String constraint. The
// key is swf:typeFilter.version.
//
// If the caller doesn't have sufficient
// permissions to invoke the action, or the parameter values fall outside the
// specified constraints, the action fails. The associated event attribute's cause
// parameter is set to OPERATION_NOT_PERMITTED. For details and example IAM
// policies, see Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) ListClosedWorkflowExecutions(ctx context.Context, params *ListClosedWorkflowExecutionsInput, optFns ...func(*Options)) (*ListClosedWorkflowExecutionsOutput, error) {
	if params == nil {
		params = &ListClosedWorkflowExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListClosedWorkflowExecutions", params, optFns, addOperationListClosedWorkflowExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListClosedWorkflowExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListClosedWorkflowExecutionsInput struct {

	// The name of the domain that contains the workflow executions to list.
	//
	// This member is required.
	Domain *string

	// If specified, only workflow executions that match this close status are listed.
	// For example, if TERMINATED is specified, then only TERMINATED workflow
	// executions are listed. closeStatusFilter, executionFilter, typeFilter and
	// tagFilter are mutually exclusive. You can specify at most one of these in a
	// request.
	CloseStatusFilter *types.CloseStatusFilter

	// If specified, the workflow executions are included in the returned results based
	// on whether their close times are within the range specified by this filter.
	// Also, if this parameter is specified, the returned results are ordered by their
	// close times. startTimeFilter and closeTimeFilter are mutually exclusive. You
	// must specify one of these in a request but not both.
	CloseTimeFilter *types.ExecutionTimeFilter

	// If specified, only workflow executions matching the workflow ID specified in the
	// filter are returned. closeStatusFilter, executionFilter, typeFilter and
	// tagFilter are mutually exclusive. You can specify at most one of these in a
	// request.
	ExecutionFilter *types.WorkflowExecutionFilter

	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results.
	MaximumPageSize int32

	// If NextPageToken is returned there are more results available. The value of
	// NextPageToken is a unique pagination token for each page. Make the call again
	// using the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 60 seconds. Using an expired
	// pagination token will return a 400 error: "Specified token has exceeded its
	// maximum lifetime". The configured maximumPageSize determines how many results
	// can be returned in a single call.
	NextPageToken *string

	// When set to true, returns the results in reverse order. By default the results
	// are returned in descending order of the start or the close time of the
	// executions.
	ReverseOrder bool

	// If specified, the workflow executions are included in the returned results based
	// on whether their start times are within the range specified by this filter.
	// Also, if this parameter is specified, the returned results are ordered by their
	// start times. startTimeFilter and closeTimeFilter are mutually exclusive. You
	// must specify one of these in a request but not both.
	StartTimeFilter *types.ExecutionTimeFilter

	// If specified, only executions that have the matching tag are listed.
	// closeStatusFilter, executionFilter, typeFilter and tagFilter are mutually
	// exclusive. You can specify at most one of these in a request.
	TagFilter *types.TagFilter

	// If specified, only executions of the type specified in the filter are returned.
	// closeStatusFilter, executionFilter, typeFilter and tagFilter are mutually
	// exclusive. You can specify at most one of these in a request.
	TypeFilter *types.WorkflowTypeFilter
}

// Contains a paginated list of information about workflow executions.
type ListClosedWorkflowExecutionsOutput struct {

	// The list of workflow information structures.
	//
	// This member is required.
	ExecutionInfos []types.WorkflowExecutionInfo

	// If a NextPageToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in nextPageToken. Keep all other arguments unchanged. The
	// configured maximumPageSize determines how many results can be returned in a
	// single call.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListClosedWorkflowExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListClosedWorkflowExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListClosedWorkflowExecutions{}, middleware.After)
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
	if err = addOpListClosedWorkflowExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListClosedWorkflowExecutions(options.Region), middleware.Before); err != nil {
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

// ListClosedWorkflowExecutionsAPIClient is a client that implements the
// ListClosedWorkflowExecutions operation.
type ListClosedWorkflowExecutionsAPIClient interface {
	ListClosedWorkflowExecutions(context.Context, *ListClosedWorkflowExecutionsInput, ...func(*Options)) (*ListClosedWorkflowExecutionsOutput, error)
}

var _ ListClosedWorkflowExecutionsAPIClient = (*Client)(nil)

// ListClosedWorkflowExecutionsPaginatorOptions is the paginator options for
// ListClosedWorkflowExecutions
type ListClosedWorkflowExecutionsPaginatorOptions struct {
	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListClosedWorkflowExecutionsPaginator is a paginator for
// ListClosedWorkflowExecutions
type ListClosedWorkflowExecutionsPaginator struct {
	options   ListClosedWorkflowExecutionsPaginatorOptions
	client    ListClosedWorkflowExecutionsAPIClient
	params    *ListClosedWorkflowExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListClosedWorkflowExecutionsPaginator returns a new
// ListClosedWorkflowExecutionsPaginator
func NewListClosedWorkflowExecutionsPaginator(client ListClosedWorkflowExecutionsAPIClient, params *ListClosedWorkflowExecutionsInput, optFns ...func(*ListClosedWorkflowExecutionsPaginatorOptions)) *ListClosedWorkflowExecutionsPaginator {
	options := ListClosedWorkflowExecutionsPaginatorOptions{}
	if params.MaximumPageSize != 0 {
		options.Limit = params.MaximumPageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListClosedWorkflowExecutionsInput{}
	}

	return &ListClosedWorkflowExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListClosedWorkflowExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListClosedWorkflowExecutions page.
func (p *ListClosedWorkflowExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListClosedWorkflowExecutionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextPageToken = p.nextToken

	params.MaximumPageSize = p.options.Limit

	result, err := p.client.ListClosedWorkflowExecutions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListClosedWorkflowExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "ListClosedWorkflowExecutions",
	}
}