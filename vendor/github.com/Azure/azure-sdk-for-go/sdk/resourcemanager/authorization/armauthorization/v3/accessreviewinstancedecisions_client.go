//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AccessReviewInstanceDecisionsClient contains the methods for the AccessReviewInstanceDecisions group.
// Don't use this type directly, use NewAccessReviewInstanceDecisionsClient() instead.
type AccessReviewInstanceDecisionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccessReviewInstanceDecisionsClient creates a new instance of AccessReviewInstanceDecisionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccessReviewInstanceDecisionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessReviewInstanceDecisionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccessReviewInstanceDecisionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get access review instance decisions
//
// Generated from API version 2021-12-01-preview
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - AccessReviewInstanceDecisionsClientListOptions contains the optional parameters for the AccessReviewInstanceDecisionsClient.NewListPager
//     method.
func (client *AccessReviewInstanceDecisionsClient) NewListPager(scheduleDefinitionID string, id string, options *AccessReviewInstanceDecisionsClientListOptions) *runtime.Pager[AccessReviewInstanceDecisionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessReviewInstanceDecisionsClientListResponse]{
		More: func(page AccessReviewInstanceDecisionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessReviewInstanceDecisionsClientListResponse) (AccessReviewInstanceDecisionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AccessReviewInstanceDecisionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, scheduleDefinitionID, id, options)
			}, nil)
			if err != nil {
				return AccessReviewInstanceDecisionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AccessReviewInstanceDecisionsClient) listCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceDecisionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/decisions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessReviewInstanceDecisionsClient) listHandleResponse(resp *http.Response) (AccessReviewInstanceDecisionsClientListResponse, error) {
	result := AccessReviewInstanceDecisionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewDecisionListResult); err != nil {
		return AccessReviewInstanceDecisionsClientListResponse{}, err
	}
	return result, nil
}