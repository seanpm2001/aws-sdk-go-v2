// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the contents of the specified folder, including its documents and
// subfolders. By default, Amazon WorkDocs returns the first 100 active document
// and folder metadata items. If there are more results, the response includes a
// marker that you can use to request the next set of results. You can also request
// initialized documents.
func (c *Client) DescribeFolderContents(ctx context.Context, params *DescribeFolderContentsInput, optFns ...func(*Options)) (*DescribeFolderContentsOutput, error) {
	if params == nil {
		params = &DescribeFolderContentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFolderContents", params, optFns, addOperationDescribeFolderContentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFolderContentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFolderContentsInput struct {

	// The ID of the folder.
	//
	// This member is required.
	FolderId *string

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string

	// The contents to include. Specify "INITIALIZED" to include initialized documents.
	Include *string

	// The maximum number of items to return with this call.
	Limit *int32

	// The marker for the next set of results. This marker was received from a previous
	// call.
	Marker *string

	// The order for the contents of the folder.
	Order types.OrderType

	// The sorting criteria.
	Sort types.ResourceSortType

	// The type of items.
	Type types.FolderContentType
}

type DescribeFolderContentsOutput struct {

	// The documents in the specified folder.
	Documents []types.DocumentMetadata

	// The subfolders in the specified folder.
	Folders []types.FolderMetadata

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeFolderContentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFolderContents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFolderContents{}, middleware.After)
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
	if err = addOpDescribeFolderContentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFolderContents(options.Region), middleware.Before); err != nil {
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

// DescribeFolderContentsAPIClient is a client that implements the
// DescribeFolderContents operation.
type DescribeFolderContentsAPIClient interface {
	DescribeFolderContents(context.Context, *DescribeFolderContentsInput, ...func(*Options)) (*DescribeFolderContentsOutput, error)
}

var _ DescribeFolderContentsAPIClient = (*Client)(nil)

// DescribeFolderContentsPaginatorOptions is the paginator options for
// DescribeFolderContents
type DescribeFolderContentsPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFolderContentsPaginator is a paginator for DescribeFolderContents
type DescribeFolderContentsPaginator struct {
	options   DescribeFolderContentsPaginatorOptions
	client    DescribeFolderContentsAPIClient
	params    *DescribeFolderContentsInput
	nextToken *string
	firstPage bool
}

// NewDescribeFolderContentsPaginator returns a new DescribeFolderContentsPaginator
func NewDescribeFolderContentsPaginator(client DescribeFolderContentsAPIClient, params *DescribeFolderContentsInput, optFns ...func(*DescribeFolderContentsPaginatorOptions)) *DescribeFolderContentsPaginator {
	options := DescribeFolderContentsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeFolderContentsInput{}
	}

	return &DescribeFolderContentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFolderContentsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeFolderContents page.
func (p *DescribeFolderContentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFolderContentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeFolderContents(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeFolderContents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "DescribeFolderContents",
	}
}