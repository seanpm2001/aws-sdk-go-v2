// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Error related to a specific channel, specified by its ARN.
type BatchError struct {

	// Channel ARN.
	Arn *string

	// Error code.
	Code *string

	// Error message, determined by the application.
	Message *string
}

// Object specifying a channel.
type Channel struct {

	// Channel ARN.
	Arn *string

	// Channel ingest endpoint, part of the definition of an ingest server, used when
	// you set up streaming software.
	IngestEndpoint *string

	// Channel latency mode. Default: LOW.
	LatencyMode ChannelLatencyMode

	// Channel name.
	Name *string

	// Channel playback URL.
	PlaybackUrl *string

	// Array of 1-50 maps, each of the form string:string (key:value).
	Tags map[string]*string

	// Channel type, which determines the allowable resolution and bitrate. If you
	// exceed the allowable resolution or bitrate, the stream probably will disconnect
	// immediately. Valid values:
	//
	//     * STANDARD: Multiple qualities are generated
	// from the original input, to automatically give viewers the best experience for
	// their devices and network conditions. Vertical resolution can be up to 1080 and
	// bitrate can be up to 8.5 Mbps.
	//
	//     * BASIC: Amazon IVS delivers the original
	// input to viewers. The viewer’s video-quality choice is limited to the original
	// input. Vertical resolution can be up to 480 and bitrate can be up to 1.5
	// Mbps.
	//
	// Default: STANDARD.
	Type ChannelType
}

// Summary information about a channel.
type ChannelSummary struct {

	// Channel ARN.
	Arn *string

	// Channel latency mode. Default: LOW.
	LatencyMode ChannelLatencyMode

	// Channel name.
	Name *string

	// Array of 1-50 maps, each of the form string:string (key:value).
	Tags map[string]*string
}

// Specifies a live video stream that has been ingested and distributed.
type Stream struct {

	// Channel ARN for the stream.
	ChannelArn *string

	// The stream’s health.
	Health StreamHealth

	// URL of the video master manifest, required by the video player to play the HLS
	// stream.
	PlaybackUrl *string

	// ISO-8601 formatted timestamp of the stream’s start.
	StartTime *time.Time

	// The stream’s state.
	State StreamState

	// Number of current viewers of the stream.
	ViewerCount *int64
}

// Object specifying a stream key.
type StreamKey struct {

	// Stream-key ARN.
	Arn *string

	// Channel ARN for the stream.
	ChannelArn *string

	// Array of 1-50 maps, each of the form string:string (key:value)
	Tags map[string]*string

	// Stream-key value.
	Value *string
}

// Summary information about a stream key.
type StreamKeySummary struct {

	// Stream-key ARN.
	Arn *string

	// Channel ARN for the stream.
	ChannelArn *string

	// Array of 1-50 maps, each of the form string:string (key:value)
	Tags map[string]*string
}

// Summary information about a stream.
type StreamSummary struct {

	// Channel ARN for the stream.
	ChannelArn *string

	// The stream’s health.
	Health StreamHealth

	// ISO-8601 formatted timestamp of the stream’s start.
	StartTime *time.Time

	// The stream’s state.
	State StreamState

	// Number of current viewers of the stream.
	ViewerCount *int64
}