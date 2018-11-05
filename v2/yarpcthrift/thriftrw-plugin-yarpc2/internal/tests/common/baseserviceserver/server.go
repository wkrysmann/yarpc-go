// Code generated by thriftrw-plugin-yarpc
// @generated

package baseserviceserver

import (
	"context"
	"go.uber.org/thriftrw/wire"
	yarpc "go.uber.org/yarpc/v2"
	"go.uber.org/yarpc/v2/yarpcthrift"
	"go.uber.org/yarpc/v2/yarpcthrift/thriftrw-plugin-yarpc2/internal/tests/common"
)

// Interface is the server-side interface for the BaseService service.
type Interface interface {
	Healthy(
		ctx context.Context,
	) (bool, error)
}

// New prepares an implementation of the BaseService service for
// registration.
//
// 	handler := BaseServiceHandler{}
// 	dispatcher.Register(baseserviceserver.New(handler))
func New(impl Interface, opts ...yarpcthrift.RegisterOption) []yarpc.TransportProcedure {
	h := handler{impl}
	service := yarpcthrift.Service{
		Name: "BaseService",
		Methods: []yarpcthrift.Method{

			yarpcthrift.Method{
				Name:         "healthy",
				Handler:      yarpcthrift.Handler(h.Healthy),
				Signature:    "Healthy() (bool)",
				ThriftModule: common.ThriftModule,
			},
		},
	}

	procedures := make([]yarpc.TransportProcedure, 0, 1)
	procedures = append(procedures, yarpcthrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

func (h handler) Healthy(ctx context.Context, body wire.Value) (yarpcthrift.Response, error) {
	var args common.BaseService_Healthy_Args
	if err := args.FromWire(body); err != nil {
		return yarpcthrift.Response{}, err
	}

	success, err := h.impl.Healthy(ctx)

	hadError := err != nil
	result, err := common.BaseService_Healthy_Helper.WrapResponse(success, err)

	var response yarpcthrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}