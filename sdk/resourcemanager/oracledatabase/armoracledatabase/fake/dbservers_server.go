// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
	"net/http"
	"net/url"
	"regexp"
)

// DbServersServer is a fake server for instances of the armoracledatabase.DbServersClient type.
type DbServersServer struct {
	// Get is the fake for method DbServersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, cloudexadatainfrastructurename string, dbserverocid string, options *armoracledatabase.DbServersClientGetOptions) (resp azfake.Responder[armoracledatabase.DbServersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByCloudExadataInfrastructurePager is the fake for method DbServersClient.NewListByCloudExadataInfrastructurePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByCloudExadataInfrastructurePager func(resourceGroupName string, cloudexadatainfrastructurename string, options *armoracledatabase.DbServersClientListByCloudExadataInfrastructureOptions) (resp azfake.PagerResponder[armoracledatabase.DbServersClientListByCloudExadataInfrastructureResponse])
}

// NewDbServersServerTransport creates a new instance of DbServersServerTransport with the provided implementation.
// The returned DbServersServerTransport instance is connected to an instance of armoracledatabase.DbServersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDbServersServerTransport(srv *DbServersServer) *DbServersServerTransport {
	return &DbServersServerTransport{
		srv:                                      srv,
		newListByCloudExadataInfrastructurePager: newTracker[azfake.PagerResponder[armoracledatabase.DbServersClientListByCloudExadataInfrastructureResponse]](),
	}
}

// DbServersServerTransport connects instances of armoracledatabase.DbServersClient to instances of DbServersServer.
// Don't use this type directly, use NewDbServersServerTransport instead.
type DbServersServerTransport struct {
	srv                                      *DbServersServer
	newListByCloudExadataInfrastructurePager *tracker[azfake.PagerResponder[armoracledatabase.DbServersClientListByCloudExadataInfrastructureResponse]]
}

// Do implements the policy.Transporter interface for DbServersServerTransport.
func (d *DbServersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DbServersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if dbServersServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = dbServersServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DbServersClient.Get":
				res.resp, res.err = d.dispatchGet(req)
			case "DbServersClient.NewListByCloudExadataInfrastructurePager":
				res.resp, res.err = d.dispatchNewListByCloudExadataInfrastructurePager(req)
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

func (d *DbServersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/cloudExadataInfrastructures/(?P<cloudexadatainfrastructurename>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dbServers/(?P<dbserverocid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudexadatainfrastructurenameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudexadatainfrastructurename")])
	if err != nil {
		return nil, err
	}
	dbserverocidParam, err := url.PathUnescape(matches[regex.SubexpIndex("dbserverocid")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, cloudexadatainfrastructurenameParam, dbserverocidParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DbServer, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DbServersServerTransport) dispatchNewListByCloudExadataInfrastructurePager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByCloudExadataInfrastructurePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByCloudExadataInfrastructurePager not implemented")}
	}
	newListByCloudExadataInfrastructurePager := d.newListByCloudExadataInfrastructurePager.get(req)
	if newListByCloudExadataInfrastructurePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/cloudExadataInfrastructures/(?P<cloudexadatainfrastructurename>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dbServers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudexadatainfrastructurenameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudexadatainfrastructurename")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByCloudExadataInfrastructurePager(resourceGroupNameParam, cloudexadatainfrastructurenameParam, nil)
		newListByCloudExadataInfrastructurePager = &resp
		d.newListByCloudExadataInfrastructurePager.add(req, newListByCloudExadataInfrastructurePager)
		server.PagerResponderInjectNextLinks(newListByCloudExadataInfrastructurePager, req, func(page *armoracledatabase.DbServersClientListByCloudExadataInfrastructureResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByCloudExadataInfrastructurePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByCloudExadataInfrastructurePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByCloudExadataInfrastructurePager) {
		d.newListByCloudExadataInfrastructurePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to DbServersServerTransport
var dbServersServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
