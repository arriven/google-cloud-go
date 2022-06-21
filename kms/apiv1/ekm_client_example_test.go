// Copyright 2022 Google LLC
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

package kms_test

import (
	"context"

	kms "cloud.google.com/go/kms/apiv1"
	"google.golang.org/api/iterator"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func ExampleNewEkmClient() {
	ctx := context.Background()
	c, err := kms.NewEkmClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleEkmClient_ListEkmConnections() {
	ctx := context.Background()
	c, err := kms.NewEkmClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &kmspb.ListEkmConnectionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/kms/v1#ListEkmConnectionsRequest.
	}
	it := c.ListEkmConnections(ctx, req)
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
	}
}

func ExampleEkmClient_GetEkmConnection() {
	ctx := context.Background()
	c, err := kms.NewEkmClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &kmspb.GetEkmConnectionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/kms/v1#GetEkmConnectionRequest.
	}
	resp, err := c.GetEkmConnection(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEkmClient_CreateEkmConnection() {
	ctx := context.Background()
	c, err := kms.NewEkmClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &kmspb.CreateEkmConnectionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/kms/v1#CreateEkmConnectionRequest.
	}
	resp, err := c.CreateEkmConnection(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEkmClient_UpdateEkmConnection() {
	ctx := context.Background()
	c, err := kms.NewEkmClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &kmspb.UpdateEkmConnectionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/kms/v1#UpdateEkmConnectionRequest.
	}
	resp, err := c.UpdateEkmConnection(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEkmClient_GetIamPolicy() {
	ctx := context.Background()
	c, err := kms.NewEkmClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#GetIamPolicyRequest.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEkmClient_SetIamPolicy() {
	ctx := context.Background()
	c, err := kms.NewEkmClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#SetIamPolicyRequest.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEkmClient_TestIamPermissions() {
	ctx := context.Background()
	c, err := kms.NewEkmClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#TestIamPermissionsRequest.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
