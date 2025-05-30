// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package videointelligence

import (
	"context"
	"log/slog"
	"math"
	"time"

	videointelligencepb "cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newStreamingVideoIntelligenceClientHook clientHook

// StreamingVideoIntelligenceCallOptions contains the retry settings for each method of StreamingVideoIntelligenceClient.
type StreamingVideoIntelligenceCallOptions struct {
	StreamingAnnotateVideo []gax.CallOption
}

func defaultStreamingVideoIntelligenceGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("videointelligence.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("videointelligence.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("videointelligence.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://videointelligence.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultStreamingVideoIntelligenceCallOptions() *StreamingVideoIntelligenceCallOptions {
	return &StreamingVideoIntelligenceCallOptions{
		StreamingAnnotateVideo: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalStreamingVideoIntelligenceClient is an interface that defines the methods available from Cloud Video Intelligence API.
type internalStreamingVideoIntelligenceClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	StreamingAnnotateVideo(context.Context, ...gax.CallOption) (videointelligencepb.StreamingVideoIntelligenceService_StreamingAnnotateVideoClient, error)
}

// StreamingVideoIntelligenceClient is a client for interacting with Cloud Video Intelligence API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service that implements streaming Video Intelligence API.
type StreamingVideoIntelligenceClient struct {
	// The internal transport-dependent client.
	internalClient internalStreamingVideoIntelligenceClient

	// The call options for this service.
	CallOptions *StreamingVideoIntelligenceCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *StreamingVideoIntelligenceClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *StreamingVideoIntelligenceClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *StreamingVideoIntelligenceClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// StreamingAnnotateVideo performs video annotation with bidirectional streaming: emitting results
// while sending video/audio bytes.
// This method is only available via the gRPC API (not REST).
func (c *StreamingVideoIntelligenceClient) StreamingAnnotateVideo(ctx context.Context, opts ...gax.CallOption) (videointelligencepb.StreamingVideoIntelligenceService_StreamingAnnotateVideoClient, error) {
	return c.internalClient.StreamingAnnotateVideo(ctx, opts...)
}

// streamingVideoIntelligenceGRPCClient is a client for interacting with Cloud Video Intelligence API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type streamingVideoIntelligenceGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing StreamingVideoIntelligenceClient
	CallOptions **StreamingVideoIntelligenceCallOptions

	// The gRPC API client.
	streamingVideoIntelligenceClient videointelligencepb.StreamingVideoIntelligenceServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewStreamingVideoIntelligenceClient creates a new streaming video intelligence service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service that implements streaming Video Intelligence API.
func NewStreamingVideoIntelligenceClient(ctx context.Context, opts ...option.ClientOption) (*StreamingVideoIntelligenceClient, error) {
	clientOpts := defaultStreamingVideoIntelligenceGRPCClientOptions()
	if newStreamingVideoIntelligenceClientHook != nil {
		hookOpts, err := newStreamingVideoIntelligenceClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := StreamingVideoIntelligenceClient{CallOptions: defaultStreamingVideoIntelligenceCallOptions()}

	c := &streamingVideoIntelligenceGRPCClient{
		connPool:                         connPool,
		streamingVideoIntelligenceClient: videointelligencepb.NewStreamingVideoIntelligenceServiceClient(connPool),
		CallOptions:                      &client.CallOptions,
		logger:                           internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *streamingVideoIntelligenceGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *streamingVideoIntelligenceGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *streamingVideoIntelligenceGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *streamingVideoIntelligenceGRPCClient) StreamingAnnotateVideo(ctx context.Context, opts ...gax.CallOption) (videointelligencepb.StreamingVideoIntelligenceService_StreamingAnnotateVideoClient, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	var resp videointelligencepb.StreamingVideoIntelligenceService_StreamingAnnotateVideoClient
	opts = append((*c.CallOptions).StreamingAnnotateVideo[0:len((*c.CallOptions).StreamingAnnotateVideo):len((*c.CallOptions).StreamingAnnotateVideo)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		c.logger.DebugContext(ctx, "api streaming client request", "serviceName", serviceName, "rpcName", "StreamingAnnotateVideo")
		resp, err = c.streamingVideoIntelligenceClient.StreamingAnnotateVideo(ctx, settings.GRPC...)
		c.logger.DebugContext(ctx, "api streaming client response", "serviceName", serviceName, "rpcName", "StreamingAnnotateVideo")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
