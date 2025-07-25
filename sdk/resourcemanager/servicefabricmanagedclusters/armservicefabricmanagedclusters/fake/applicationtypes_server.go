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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
	"net/http"
	"net/url"
	"regexp"
)

// ApplicationTypesServer is a fake server for instances of the armservicefabricmanagedclusters.ApplicationTypesClient type.
type ApplicationTypesServer struct {
	// CreateOrUpdate is the fake for method ApplicationTypesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, parameters armservicefabricmanagedclusters.ApplicationTypeResource, options *armservicefabricmanagedclusters.ApplicationTypesClientCreateOrUpdateOptions) (resp azfake.Responder[armservicefabricmanagedclusters.ApplicationTypesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ApplicationTypesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *armservicefabricmanagedclusters.ApplicationTypesClientBeginDeleteOptions) (resp azfake.PollerResponder[armservicefabricmanagedclusters.ApplicationTypesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ApplicationTypesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *armservicefabricmanagedclusters.ApplicationTypesClientGetOptions) (resp azfake.Responder[armservicefabricmanagedclusters.ApplicationTypesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ApplicationTypesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, clusterName string, options *armservicefabricmanagedclusters.ApplicationTypesClientListOptions) (resp azfake.PagerResponder[armservicefabricmanagedclusters.ApplicationTypesClientListResponse])

	// Update is the fake for method ApplicationTypesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, parameters armservicefabricmanagedclusters.ApplicationTypeUpdateParameters, options *armservicefabricmanagedclusters.ApplicationTypesClientUpdateOptions) (resp azfake.Responder[armservicefabricmanagedclusters.ApplicationTypesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewApplicationTypesServerTransport creates a new instance of ApplicationTypesServerTransport with the provided implementation.
// The returned ApplicationTypesServerTransport instance is connected to an instance of armservicefabricmanagedclusters.ApplicationTypesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewApplicationTypesServerTransport(srv *ApplicationTypesServer) *ApplicationTypesServerTransport {
	return &ApplicationTypesServerTransport{
		srv:          srv,
		beginDelete:  newTracker[azfake.PollerResponder[armservicefabricmanagedclusters.ApplicationTypesClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armservicefabricmanagedclusters.ApplicationTypesClientListResponse]](),
	}
}

// ApplicationTypesServerTransport connects instances of armservicefabricmanagedclusters.ApplicationTypesClient to instances of ApplicationTypesServer.
// Don't use this type directly, use NewApplicationTypesServerTransport instead.
type ApplicationTypesServerTransport struct {
	srv          *ApplicationTypesServer
	beginDelete  *tracker[azfake.PollerResponder[armservicefabricmanagedclusters.ApplicationTypesClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armservicefabricmanagedclusters.ApplicationTypesClientListResponse]]
}

// Do implements the policy.Transporter interface for ApplicationTypesServerTransport.
func (a *ApplicationTypesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *ApplicationTypesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if applicationTypesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = applicationTypesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ApplicationTypesClient.CreateOrUpdate":
				res.resp, res.err = a.dispatchCreateOrUpdate(req)
			case "ApplicationTypesClient.BeginDelete":
				res.resp, res.err = a.dispatchBeginDelete(req)
			case "ApplicationTypesClient.Get":
				res.resp, res.err = a.dispatchGet(req)
			case "ApplicationTypesClient.NewListPager":
				res.resp, res.err = a.dispatchNewListPager(req)
			case "ApplicationTypesClient.Update":
				res.resp, res.err = a.dispatchUpdate(req)
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

func (a *ApplicationTypesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabric/managedClusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationTypes/(?P<applicationTypeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armservicefabricmanagedclusters.ApplicationTypeResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	applicationTypeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationTypeName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, clusterNameParam, applicationTypeNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationTypeResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationTypesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabric/managedClusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationTypes/(?P<applicationTypeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		applicationTypeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationTypeName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, clusterNameParam, applicationTypeNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *ApplicationTypesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabric/managedClusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationTypes/(?P<applicationTypeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	applicationTypeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationTypeName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, clusterNameParam, applicationTypeNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationTypeResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationTypesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabric/managedClusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationTypes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceGroupNameParam, clusterNameParam, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armservicefabricmanagedclusters.ApplicationTypesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}

func (a *ApplicationTypesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabric/managedClusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationTypes/(?P<applicationTypeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armservicefabricmanagedclusters.ApplicationTypeUpdateParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	applicationTypeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationTypeName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Update(req.Context(), resourceGroupNameParam, clusterNameParam, applicationTypeNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationTypeResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ApplicationTypesServerTransport
var applicationTypesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
