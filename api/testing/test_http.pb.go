// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.3

package testing

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type TestingHTTPServer interface {
	ListTest(context.Context, *ListTestRequest) (*ListTestReply, error)
}

func RegisterTestingHTTPServer(s *http.Server, srv TestingHTTPServer) {
	r := s.Route("/")
	r.GET("/test/", _Testing_ListTest0_HTTP_Handler(srv))
}

func _Testing_ListTest0_HTTP_Handler(srv TestingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTestRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.testing.Testing/ListTest")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTest(ctx, req.(*ListTestRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTestReply)
		return ctx.Result(200, reply)
	}
}

type TestingHTTPClient interface {
	ListTest(ctx context.Context, req *ListTestRequest, opts ...http.CallOption) (rsp *ListTestReply, err error)
}

type TestingHTTPClientImpl struct {
	cc *http.Client
}

func NewTestingHTTPClient(client *http.Client) TestingHTTPClient {
	return &TestingHTTPClientImpl{client}
}

func (c *TestingHTTPClientImpl) ListTest(ctx context.Context, in *ListTestRequest, opts ...http.CallOption) (*ListTestReply, error) {
	var out ListTestReply
	pattern := "/test/"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.testing.Testing/ListTest"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
