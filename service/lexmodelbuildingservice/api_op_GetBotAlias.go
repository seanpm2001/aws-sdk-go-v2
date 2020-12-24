// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about an Amazon Lex bot alias. For more information about
// aliases, see versioning-aliases. This operation requires permissions for the
// lex:GetBotAlias action.
func (c *Client) GetBotAlias(ctx context.Context, params *GetBotAliasInput, optFns ...func(*Options)) (*GetBotAliasOutput, error) {
	if params == nil {
		params = &GetBotAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBotAlias", params, optFns, addOperationGetBotAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBotAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBotAliasInput struct {

	// The name of the bot.
	//
	// This member is required.
	BotName *string

	// The name of the bot alias. The name is case sensitive.
	//
	// This member is required.
	Name *string
}

type GetBotAliasOutput struct {

	// The name of the bot that the alias points to.
	BotName *string

	// The version of the bot that the alias points to.
	BotVersion *string

	// Checksum of the bot alias.
	Checksum *string

	// The settings that determine how Amazon Lex uses conversation logs for the alias.
	ConversationLogs *types.ConversationLogsResponse

	// The date that the bot alias was created.
	CreatedDate *time.Time

	// A description of the bot alias.
	Description *string

	// The date that the bot alias was updated. When you create a resource, the
	// creation date and the last updated date are the same.
	LastUpdatedDate *time.Time

	// The name of the bot alias.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBotAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBotAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBotAlias{}, middleware.After)
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
	if err = addOpGetBotAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBotAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBotAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBotAlias",
	}
}