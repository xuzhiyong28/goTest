package napodate

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetEndpoint      endpoint.Endpoint
	StatusEndpoint   endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

func MakeGetEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(GetRequest) //// 我们只需要请求，不需要使用它的值
		d, err := srv.Get(ctx)
		if err != nil {
			return GetResponse{d, err.Error()}, nil
		}
		return GetResponse{d, ""}, nil
	}
}

func MakeStatusEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(StatusRequest)
		s, err := srv.Status(ctx)
		if err != nil {
			return StatusResponse{s}, err
		}
		return StatusResponse{s}, nil
	}
}

func MakeValidateEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ValidateRequest)
		b, err := srv.Validate(ctx, req.Date)
		if err != nil {
			return ValidateResponse{b, err.Error()}, nil
		}
		return ValidateResponse{b, ""}, nil
	}
}

// get 断点映射
func (e Endpoints) Get(ctx context.Context) (string, error) {
	req := GetRequest{}
	resp, err := e.GetEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(GetResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Date, nil
}

func (e Endpoints) Status(ctx context.Context) (string, error) {
	req := StatusRequest{}
	resp, err := e.StatusEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	statusResp := resp.(StatusResponse)
	return statusResp.Status, nil
}

func (e Endpoints) Validate(ctx context.Context, date string) (bool, error) {
	req := ValidateRequest{Date: date}
	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	validateResp := resp.(ValidateResponse)
	if validateResp.Err != "" {
		return false, errors.New(validateResp.Err)
	}
	return validateResp.Valid, nil
}
