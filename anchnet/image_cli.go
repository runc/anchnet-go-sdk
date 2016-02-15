// Copyright 2015 anchnet-go authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	anchnet "github.com/caicloud/anchnet-go"
	"github.com/spf13/cobra"
)

func execCaptureInstance(cmd *cobra.Command, args []string, client *anchnet.Client, out io.Writer) {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Image name and instance id required")
		os.Exit(1)
	}

	request := anchnet.CaptureInstanceRequest{
		ImageName: args[0],
		Instance:  args[1],
	}
	var response anchnet.CaptureInstanceResponse
	sendResult(&response, out, "CaptureInstance", response.Code, client.SendRequest(request, &response))
}

func execGrantImageToUsers(cmd *cobra.Command, args []string, client *anchnet.Client, out io.Writer) {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Image id and user ids required")
		os.Exit(1)
	}

	request := anchnet.GrantImageToUsersRequest{
		ImageID: args[0],
		UserIDs: strings.Split(args[1], ","),
	}
	var response anchnet.GrantImageToUsersResponse
	sendResult(&response, out, "GrantImageToUsers", response.Code, client.SendRequest(request, &response))
}

func execRevokeImageFromUsers(cmd *cobra.Command, args []string, client *anchnet.Client, out io.Writer) {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Image id and user ids required")
		os.Exit(1)
	}

	request := anchnet.RevokeImageFromUsersRequest{
		ImageIDs: []string{args[0]},
		UserIDs:  strings.Split(args[1], ","),
	}
	var response anchnet.RevokeImageFromUsersResponse
	sendResult(&response, out, "RevokeImageToUsers", response.Code, client.SendRequest(request, &response))
}

func execDescribeImageUsers(cmd *cobra.Command, args []string, client *anchnet.Client, out io.Writer) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Image id required")
		os.Exit(1)
	}

	request := anchnet.DescribeImageUsersRequest{
		ImageIDs: []string{args[0]},
	}
	var response anchnet.DescribeImageUsersResponse
	sendResult(&response, out, "DescribeImageUsers", response.Code, client.SendRequest(request, &response))
}
