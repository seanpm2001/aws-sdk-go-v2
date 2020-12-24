// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves all sampling rules.
func (c *Client) GetSamplingRules(ctx context.Context, params *GetSamplingRulesInput, optFns ...func(*Options)) (*GetSamplingRulesOutput, error) {
	if params == nil {
		params = &GetSamplingRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSamplingRules", params, optFns, addOperationGetSamplingRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSamplingRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSamplingRulesInput struct {

	// Pagination token.
	NextToken *string
}

type GetSamplingRulesOutput struct {

	// Pagination token.
	NextToken *string

	// Rule definitions and metadata.
	SamplingRuleRecords []types.SamplingRuleRecord

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSamplingRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSamplingRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSamplingRules{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSamplingRules(options.Region), middleware.Before); err != nil {
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

// GetSamplingRulesAPIClient is a client that implements the GetSamplingRules
// operation.
type GetSamplingRulesAPIClient interface {
	GetSamplingRules(context.Context, *GetSamplingRulesInput, ...func(*Options)) (*GetSamplingRulesOutput, error)
}

var _ GetSamplingRulesAPIClient = (*Client)(nil)

// GetSamplingRulesPaginatorOptions is the paginator options for GetSamplingRules
type GetSamplingRulesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetSamplingRulesPaginator is a paginator for GetSamplingRules
type GetSamplingRulesPaginator struct {
	options   GetSamplingRulesPaginatorOptions
	client    GetSamplingRulesAPIClient
	params    *GetSamplingRulesInput
	nextToken *string
	firstPage bool
}

// NewGetSamplingRulesPaginator returns a new GetSamplingRulesPaginator
func NewGetSamplingRulesPaginator(client GetSamplingRulesAPIClient, params *GetSamplingRulesInput, optFns ...func(*GetSamplingRulesPaginatorOptions)) *GetSamplingRulesPaginator {
	options := GetSamplingRulesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetSamplingRulesInput{}
	}

	return &GetSamplingRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetSamplingRulesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetSamplingRules page.
func (p *GetSamplingRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetSamplingRulesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.GetSamplingRules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetSamplingRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "GetSamplingRules",
	}
}