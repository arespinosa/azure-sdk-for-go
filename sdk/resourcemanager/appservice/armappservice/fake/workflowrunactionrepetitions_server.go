// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
	"net/http"
	"net/url"
	"regexp"
)

// WorkflowRunActionRepetitionsServer is a fake server for instances of the armappservice.WorkflowRunActionRepetitionsClient type.
type WorkflowRunActionRepetitionsServer struct {
	// Get is the fake for method WorkflowRunActionRepetitionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, actionName string, repetitionName string, options *armappservice.WorkflowRunActionRepetitionsClientGetOptions) (resp azfake.Responder[armappservice.WorkflowRunActionRepetitionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkflowRunActionRepetitionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, name string, workflowName string, runName string, actionName string, options *armappservice.WorkflowRunActionRepetitionsClientListOptions) (resp azfake.PagerResponder[armappservice.WorkflowRunActionRepetitionsClientListResponse])

	// NewListExpressionTracesPager is the fake for method WorkflowRunActionRepetitionsClient.NewListExpressionTracesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListExpressionTracesPager func(resourceGroupName string, name string, workflowName string, runName string, actionName string, repetitionName string, options *armappservice.WorkflowRunActionRepetitionsClientListExpressionTracesOptions) (resp azfake.PagerResponder[armappservice.WorkflowRunActionRepetitionsClientListExpressionTracesResponse])
}

// NewWorkflowRunActionRepetitionsServerTransport creates a new instance of WorkflowRunActionRepetitionsServerTransport with the provided implementation.
// The returned WorkflowRunActionRepetitionsServerTransport instance is connected to an instance of armappservice.WorkflowRunActionRepetitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkflowRunActionRepetitionsServerTransport(srv *WorkflowRunActionRepetitionsServer) *WorkflowRunActionRepetitionsServerTransport {
	return &WorkflowRunActionRepetitionsServerTransport{
		srv:                          srv,
		newListPager:                 newTracker[azfake.PagerResponder[armappservice.WorkflowRunActionRepetitionsClientListResponse]](),
		newListExpressionTracesPager: newTracker[azfake.PagerResponder[armappservice.WorkflowRunActionRepetitionsClientListExpressionTracesResponse]](),
	}
}

// WorkflowRunActionRepetitionsServerTransport connects instances of armappservice.WorkflowRunActionRepetitionsClient to instances of WorkflowRunActionRepetitionsServer.
// Don't use this type directly, use NewWorkflowRunActionRepetitionsServerTransport instead.
type WorkflowRunActionRepetitionsServerTransport struct {
	srv                          *WorkflowRunActionRepetitionsServer
	newListPager                 *tracker[azfake.PagerResponder[armappservice.WorkflowRunActionRepetitionsClientListResponse]]
	newListExpressionTracesPager *tracker[azfake.PagerResponder[armappservice.WorkflowRunActionRepetitionsClientListExpressionTracesResponse]]
}

// Do implements the policy.Transporter interface for WorkflowRunActionRepetitionsServerTransport.
func (w *WorkflowRunActionRepetitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return w.dispatchToMethodFake(req, method)
}

func (w *WorkflowRunActionRepetitionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if workflowRunActionRepetitionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = workflowRunActionRepetitionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "WorkflowRunActionRepetitionsClient.Get":
				res.resp, res.err = w.dispatchGet(req)
			case "WorkflowRunActionRepetitionsClient.NewListPager":
				res.resp, res.err = w.dispatchNewListPager(req)
			case "WorkflowRunActionRepetitionsClient.NewListExpressionTracesPager":
				res.resp, res.err = w.dispatchNewListExpressionTracesPager(req)
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

func (w *WorkflowRunActionRepetitionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runs/(?P<runName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/actions/(?P<actionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/repetitions/(?P<repetitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
	if err != nil {
		return nil, err
	}
	runNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("runName")])
	if err != nil {
		return nil, err
	}
	actionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("actionName")])
	if err != nil {
		return nil, err
	}
	repetitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("repetitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, nameParam, workflowNameParam, runNameParam, actionNameParam, repetitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkflowRunActionRepetitionDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkflowRunActionRepetitionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runs/(?P<runName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/actions/(?P<actionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/repetitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
		if err != nil {
			return nil, err
		}
		runNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("runName")])
		if err != nil {
			return nil, err
		}
		actionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("actionName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, nameParam, workflowNameParam, runNameParam, actionNameParam, nil)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappservice.WorkflowRunActionRepetitionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		w.newListPager.remove(req)
	}
	return resp, nil
}

func (w *WorkflowRunActionRepetitionsServerTransport) dispatchNewListExpressionTracesPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListExpressionTracesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListExpressionTracesPager not implemented")}
	}
	newListExpressionTracesPager := w.newListExpressionTracesPager.get(req)
	if newListExpressionTracesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runs/(?P<runName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/actions/(?P<actionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/repetitions/(?P<repetitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listExpressionTraces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 7 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
		if err != nil {
			return nil, err
		}
		runNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("runName")])
		if err != nil {
			return nil, err
		}
		actionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("actionName")])
		if err != nil {
			return nil, err
		}
		repetitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("repetitionName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListExpressionTracesPager(resourceGroupNameParam, nameParam, workflowNameParam, runNameParam, actionNameParam, repetitionNameParam, nil)
		newListExpressionTracesPager = &resp
		w.newListExpressionTracesPager.add(req, newListExpressionTracesPager)
		server.PagerResponderInjectNextLinks(newListExpressionTracesPager, req, func(page *armappservice.WorkflowRunActionRepetitionsClientListExpressionTracesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListExpressionTracesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListExpressionTracesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListExpressionTracesPager) {
		w.newListExpressionTracesPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to WorkflowRunActionRepetitionsServerTransport
var workflowRunActionRepetitionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
