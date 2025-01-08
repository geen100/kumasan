// Code generated by goagen v1.6.2, DO NOT EDIT.
//
// API "BearSighting": bear TestHelpers
//
// Command:
// $ main

package test

import (
	"backend/webapi/app"
	"context"
	"fmt"
	goa "github.com/shogo82148/goa-v1"
	"github.com/shogo82148/goa-v1/goatest"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// AddBearBadRequest runs the method Add of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AddBearBadRequest(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.BearController, payload *app.AddBearPayload) http.ResponseWriter {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/bear/"),
	}
	req := httptest.NewRequest("POST", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BearTest"), rw, req, prms)
	addCtx, _err := app.NewAddBearContext(goaCtx, req, service)
	if _err != nil {
		_e, _ok := _err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil
	}
	addCtx.Payload = payload

	// Perform action
	_err = ctrl.Add(addCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}

	// Return results
	return rw
}

// AddBearCreated runs the method Add of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AddBearCreated(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.BearController, payload *app.AddBearPayload) (http.ResponseWriter, *app.Sighting) {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/bear/"),
	}
	req := httptest.NewRequest("POST", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BearTest"), rw, req, prms)
	addCtx, _err := app.NewAddBearContext(goaCtx, req, service)
	if _err != nil {
		_e, _ok := _err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil, nil
	}
	addCtx.Payload = payload

	// Perform action
	_err = ctrl.Add(addCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}
	var mt *app.Sighting
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(*app.Sighting)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.Sighting", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// AddBearInternalServerError runs the method Add of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AddBearInternalServerError(t testing.TB, ctx context.Context, service *goa.Service, ctrl app.BearController, payload *app.AddBearPayload) http.ResponseWriter {
	t.Helper()

	// Setup service
	var (
		logBuf strings.Builder

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil
	}

	// Setup request context
	if ctx == nil {
		ctx = context.Background()
	}
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/bear/"),
	}
	req := httptest.NewRequest("POST", u.String(), nil)
	req = req.WithContext(ctx)
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BearTest"), rw, req, prms)
	addCtx, _err := app.NewAddBearContext(goaCtx, req, service)
	if _err != nil {
		_e, _ok := _err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil
	}
	addCtx.Payload = payload

	// Perform action
	_err = ctrl.Add(addCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}

	// Return results
	return rw
}
