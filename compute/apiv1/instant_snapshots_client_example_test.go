// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package compute_test

import (
	"context"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/iterator"
)

func ExampleNewInstantSnapshotsRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewInstantSnapshotsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleInstantSnapshotsClient_AggregatedList() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewInstantSnapshotsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.AggregatedListInstantSnapshotsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AggregatedListInstantSnapshotsRequest.
	}
	it := c.AggregatedList(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*computepb.InstantSnapshotAggregatedList)
	}
}

func ExampleInstantSnapshotsClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewInstantSnapshotsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.DeleteInstantSnapshotRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteInstantSnapshotRequest.
	}
	op, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleInstantSnapshotsClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewInstantSnapshotsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.GetInstantSnapshotRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetInstantSnapshotRequest.
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstantSnapshotsClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewInstantSnapshotsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.GetIamPolicyInstantSnapshotRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetIamPolicyInstantSnapshotRequest.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstantSnapshotsClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewInstantSnapshotsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.InsertInstantSnapshotRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertInstantSnapshotRequest.
	}
	op, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleInstantSnapshotsClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewInstantSnapshotsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListInstantSnapshotsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListInstantSnapshotsRequest.
	}
	it := c.List(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*computepb.InstantSnapshotList)
	}
}

func ExampleInstantSnapshotsClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewInstantSnapshotsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.SetIamPolicyInstantSnapshotRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetIamPolicyInstantSnapshotRequest.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstantSnapshotsClient_SetLabels() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewInstantSnapshotsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.SetLabelsInstantSnapshotRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetLabelsInstantSnapshotRequest.
	}
	op, err := c.SetLabels(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleInstantSnapshotsClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := compute.NewInstantSnapshotsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.TestIamPermissionsInstantSnapshotRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#TestIamPermissionsInstantSnapshotRequest.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}