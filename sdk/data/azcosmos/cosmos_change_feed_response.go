// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ChangeFeedResponse contains the result of a change feed request.
type ChangeFeedResponse struct {
	// ResourceID is the unique identifier for the resource.
	ResourceID string `json:"_rid"`
	// Documents is a list of changed documents returned in the change feed.
	Documents []json.RawMessage `json:"Documents"`
	// Count is the number of documents returned in this page.
	Count int `json:"_count"`

	// Selected HTTP headers that we're retrieving from the response
	ETag              string
	ContinuationToken string
	LSN               string

	// Store the feed range if it was used in the request
	FeedRange *FeedRange

	PartitionKey *PartitionKey

	Response
}

// newChangeFeedResponse creates a new ChangeFeedResponse from an HTTP response.
func newChangeFeedResponse(resp *http.Response) (ChangeFeedResponse, error) {
	response := ChangeFeedResponse{
		Response: newResponse(resp),
		ETag:     resp.Header.Get("etag"),
		LSN:      resp.Header.Get("lsn"),
	}

	if resp.StatusCode == http.StatusNotModified {
		// Handle 304 Not Modified response (no changes since the specified ETag)
		response.Documents = []json.RawMessage{}
		response.Count = 0
		return response, nil
	}

	// For non-304 responses, unmarshal the response body
	defer resp.Body.Close()
	body, err := azruntime.Payload(resp)
	if err != nil {
		return response, err
	}
	// Parse the response into our response structure
	if err := json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return response, nil
}

// GetContinuation from ChangeFeedResponse
func (c ChangeFeedResponse) GetContinuation() string {
	return c.ETag
}

// GetContRanges extracts the continuation token range from the ChangeFeedResponse.
func (c ChangeFeedResponse) GetContRanges() (min string, max string, ok bool) {
	// If FeedRange was set in the request, use it
	if c.FeedRange != nil {
		fmt.Printf("FeedRange is set: %s, %s\n", c.FeedRange.MinInclusive, c.FeedRange.MaxExclusive)
		return c.FeedRange.MinInclusive, c.FeedRange.MaxExclusive, true
	}

	// Otherwise, try to extract from continuation token (fallback)
	if c.ContinuationToken == "" {
		return "", "", false
	}

	return "", "", false
}

// GetPartitionKey returns the partition key used in the request, if any.
func (c ChangeFeedResponse) GetPartitionKey() (partitionKey *PartitionKey, ok bool) {
	if c.PartitionKey != nil {
		return c.PartitionKey, true
	}
	return nil, false
}

// getCompositeContinuationToken creates a composite continuation token from the response.
// This token combines the feed range information with the ETag for use in subsequent requests.
func (c ChangeFeedResponse) getCompositeContinuationToken() (string, error) {
	// Extract the range from the continuation token
	min, max, ok := c.GetContRanges()
	if !ok {
		// No valid range in continuation token
		return "", nil
	}

	// Get the ETag
	etag := c.GetContinuation()
	if etag == "" {
		// No ETag available
		return "", nil
	}

	// Create the change feed range with continuation
	etagValue := azcore.ETag(etag)
	cfRange := newChangeFeedRange(min, max, &ChangeFeedRangeOptions{
		ContinuationToken: &etagValue,
	})

	// Create composite token
	compositeToken := newCompositeContinuationToken(c.ResourceID, []changeFeedRange{cfRange})

	// Marshal to JSON
	tokenBytes, err := json.Marshal(compositeToken)
	if err != nil {
		return "", err
	}

	return string(tokenBytes), nil
}

// getContinuationTokenForPartitionKey creates a continuation token for a specific partition key.
func (c ChangeFeedResponse) getContinuationTokenForPartitionKey() (string, error) {
	// Get the partition key from the response
	partitionKey, ok := c.GetPartitionKey()
	if !ok {
		// No partition key available in the response
		return "", nil
	}

	// Get the ETag
	etag := c.GetContinuation()
	if etag == "" {
		// No ETag available
		return "", nil
	}
	etagValue := azcore.ETag(etag)

	continuationToken := newContinuationTokenForPartitionKey(c.ResourceID, partitionKey, &etagValue)

	// Marshal to JSON
	tokenBytes, err := json.Marshal(continuationToken)
	if err != nil {
		return "", err
	}

	return string(tokenBytes), nil
}
