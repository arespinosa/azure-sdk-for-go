// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
	"net/http"
)

// CertificateRegistrationProviderServer is a fake server for instances of the armappservice.CertificateRegistrationProviderClient type.
type CertificateRegistrationProviderServer struct {
	// NewListOperationsPager is the fake for method CertificateRegistrationProviderClient.NewListOperationsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListOperationsPager func(options *armappservice.CertificateRegistrationProviderClientListOperationsOptions) (resp azfake.PagerResponder[armappservice.CertificateRegistrationProviderClientListOperationsResponse])
}

// NewCertificateRegistrationProviderServerTransport creates a new instance of CertificateRegistrationProviderServerTransport with the provided implementation.
// The returned CertificateRegistrationProviderServerTransport instance is connected to an instance of armappservice.CertificateRegistrationProviderClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCertificateRegistrationProviderServerTransport(srv *CertificateRegistrationProviderServer) *CertificateRegistrationProviderServerTransport {
	return &CertificateRegistrationProviderServerTransport{
		srv:                    srv,
		newListOperationsPager: newTracker[azfake.PagerResponder[armappservice.CertificateRegistrationProviderClientListOperationsResponse]](),
	}
}

// CertificateRegistrationProviderServerTransport connects instances of armappservice.CertificateRegistrationProviderClient to instances of CertificateRegistrationProviderServer.
// Don't use this type directly, use NewCertificateRegistrationProviderServerTransport instead.
type CertificateRegistrationProviderServerTransport struct {
	srv                    *CertificateRegistrationProviderServer
	newListOperationsPager *tracker[azfake.PagerResponder[armappservice.CertificateRegistrationProviderClientListOperationsResponse]]
}

// Do implements the policy.Transporter interface for CertificateRegistrationProviderServerTransport.
func (c *CertificateRegistrationProviderServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CertificateRegistrationProviderServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if certificateRegistrationProviderServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = certificateRegistrationProviderServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "CertificateRegistrationProviderClient.NewListOperationsPager":
				res.resp, res.err = c.dispatchNewListOperationsPager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (c *CertificateRegistrationProviderServerTransport) dispatchNewListOperationsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListOperationsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListOperationsPager not implemented")}
	}
	newListOperationsPager := c.newListOperationsPager.get(req)
	if newListOperationsPager == nil {
		resp := c.srv.NewListOperationsPager(nil)
		newListOperationsPager = &resp
		c.newListOperationsPager.add(req, newListOperationsPager)
		server.PagerResponderInjectNextLinks(newListOperationsPager, req, func(page *armappservice.CertificateRegistrationProviderClientListOperationsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListOperationsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListOperationsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListOperationsPager) {
		c.newListOperationsPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to CertificateRegistrationProviderServerTransport
var certificateRegistrationProviderServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
