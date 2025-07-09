// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

type ContinuationTokenForPartitionKey struct {
	ResourceID   string        `json:"resourceId"`
	PartitionKey *PartitionKey `json:"partitionKey"`
	Continuation *azcore.ETag  `json:"continuation"`
}

// newContinuationTokenForPartitionKey creates a new ContinuationTokenForPartitionKey with the specified resource ID, partition key, and continuation token.
// This function is used to create a continuation token for the Cosmos DB change feed for a specific partition key.
// It is designed for internal use only and should not be used directly by clients.
func newContinuationTokenForPartitionKey(resourceID string, partitionKey *PartitionKey, continuation *azcore.ETag) ContinuationTokenForPartitionKey {
	return ContinuationTokenForPartitionKey{
		ResourceID:   resourceID,
		PartitionKey: partitionKey,
		Continuation: continuation,
	}
}
