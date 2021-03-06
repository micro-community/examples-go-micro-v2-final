// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: micro/examples/blog/search/proto/search/search.proto

package go_micro_service_search

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for SearchService service

func NewSearchServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SearchService service

type SearchService interface {
	Index(ctx context.Context, in *IndexRequest, opts ...client.CallOption) (*IndexResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
}

type searchService struct {
	c    client.Client
	name string
}

func NewSearchService(name string, c client.Client) SearchService {
	return &searchService{
		c:    c,
		name: name,
	}
}

func (c *searchService) Index(ctx context.Context, in *IndexRequest, opts ...client.CallOption) (*IndexResponse, error) {
	req := c.c.NewRequest(c.name, "SearchService.Index", in)
	out := new(IndexResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.name, "SearchService.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SearchService service

type SearchServiceHandler interface {
	Index(context.Context, *IndexRequest, *IndexResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
}

func RegisterSearchServiceHandler(s server.Server, hdlr SearchServiceHandler, opts ...server.HandlerOption) error {
	type searchService interface {
		Index(ctx context.Context, in *IndexRequest, out *IndexResponse) error
		Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error
	}
	type SearchService struct {
		searchService
	}
	h := &searchServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SearchService{h}, opts...))
}

type searchServiceHandler struct {
	SearchServiceHandler
}

func (h *searchServiceHandler) Index(ctx context.Context, in *IndexRequest, out *IndexResponse) error {
	return h.SearchServiceHandler.Index(ctx, in, out)
}

func (h *searchServiceHandler) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.SearchServiceHandler.Search(ctx, in, out)
}
