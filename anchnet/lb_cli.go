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

func execCreateLoadBalancer(cmd *cobra.Command, args []string, client *anchnet.Client, out io.Writer) {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Load balancer name and public ips required")
		os.Exit(1)
	}

	lb_type := getFlagInt(cmd, "type")

	refs := strings.Split(args[1], ",")
	ips := make([]anchnet.CreateLoadBalancerIP, len(refs))
	for i, ip := range refs {
		ips[i].RefID = ip
	}

	request := anchnet.CreateLoadBalancerRequest{
		Product: anchnet.CreateLoadBalancerProduct{
			Loadbalancer: anchnet.CreateLoadBalancerLB{
				Name: args[0],
				Type: anchnet.LoadBalancerType(lb_type),
			},
			Eips: ips,
		},
	}
	var response anchnet.CreateLoadBalancerResponse
	sendResult(&response, out, "CreateLoadBalancer", response.Code, client.SendRequest(request, &response))
}

func execDeleteLoadBalancer(cmd *cobra.Command, args []string, client *anchnet.Client, out io.Writer) {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Load balancer ids and public ips required")
		os.Exit(1)
	}

	lbs := strings.Split(args[0], ",")
	ips := strings.Split(args[1], ",")

	request := anchnet.DeleteLoadBalancersRequest{
		LoadbalancerIDs: lbs,
		EipIDs:          ips,
	}
	var response anchnet.DeleteLoadBalancersResponse
	sendResult(&response, out, "DeleteLoadBalancer", response.Code, client.SendRequest(request, &response))
}

func execSearchLoadBalancer(cmd *cobra.Command, args []string, client *anchnet.Client, out io.Writer) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "LoadBalancer name required")
		os.Exit(1)
	}

	var lb_status []anchnet.LoadBalancerStatus
	for _, status := range strings.Split(getFlagString(cmd, "status"), ",") {
		lb_status = append(lb_status, anchnet.LoadBalancerStatus(status))
	}

	request := anchnet.DescribeLoadBalancersRequest{
		SearchWord: args[0],
		Status:     lb_status,
		Verbose:    1,
	}
	var response anchnet.DescribeLoadBalancersResponse
	sendResult(&response, out, "SearchLoadBalancer", response.Code, client.SendRequest(request, &response))
}
