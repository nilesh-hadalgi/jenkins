// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX - License - Identifier: Apache - 2.0
// snippet-start:[ec2.go-v2.StartInstances]
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/smithy-go"
)

// EC2StopInstancesAPI defines the interface for the StopInstances function.
// We use this interface to test the function using a mocked service.
type EC2StopInstancesAPI interface {
	StopInstances(ctx context.Context,
		params *ec2.StopInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.StopInstancesOutput, error)
}

// StopInstance stops an Amazon Elastic Compute Cloud (Amazon EC2) instance.
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region.
//	api is the interface that defines the method call.
//	input defines the input arguments to the service call.
//
// Output:
//
//	If success, a StopInstancesOutput object containing the result of the service call and nil.
//	Otherwise, nil and an error from the call to StopInstances.
func StopInstances(c context.Context, api EC2StopInstancesAPI, input *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	resp, err := api.StopInstances(c, input)

	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		fmt.Println("User has permission to stop instances.")
		input.DryRun = aws.Bool(false)
		return api.StopInstances(c, input)
	}

	return resp, err
}

func StopInstancesCmd(instanceIds []string) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ec2.NewFromConfig(cfg)

	input := &ec2.StopInstancesInput{
		InstanceIds: instanceIds,
		DryRun:      aws.Bool(true),
	}

	_, err = StopInstances(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error stopping the instance")
		fmt.Println(err)
		return
	}

	fmt.Println("Stopped instances with IDs " + strings.Join(instanceIds, ","))
}

// snippet-end:[ec2.go-v2.StopInstances]

// EC2StartInstancesAPI defines the interface for the StartInstances function.
// We use this interface to test the function using a mocked service.
type EC2StartInstancesAPI interface {
	StartInstances(ctx context.Context,
		params *ec2.StartInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.StartInstancesOutput, error)
}

// StartInstance starts an Amazon Elastic Compute Cloud (Amazon EC2) instance.
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region.
//	api is the interface that defines the method call.
//	input defines the input arguments to the service call.
//
// Output:
//
//	If success, a StartInstancesOutput object containing the result of the service call and nil.
//	Otherwise, nil and an error from the call to StartInstances.
func StartInstances(c context.Context, api EC2StartInstancesAPI, input *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	resp, err := api.StartInstances(c, input)

	var apiErr smithy.APIError
	if errors.As(err, &apiErr) && apiErr.ErrorCode() == "DryRunOperation" {
		fmt.Println("User has permission to start an instance.")
		input.DryRun = aws.Bool(false)
		return api.StartInstances(c, input)
	}

	return resp, err
}

func StartInstancesCmd(instanceIds []string) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ec2.NewFromConfig(cfg)

	input := &ec2.StartInstancesInput{
		InstanceIds: instanceIds,
		DryRun:      aws.Bool(true),
	}

	_, err = StartInstances(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error starting the instance")
		fmt.Println(err)
		return
	}

	fmt.Println("Started instances with IDs " + strings.Join(instanceIds, ","))
}

// snippet-end:[ec2.go-v2.StartInstances]

func main() {

	command := flag.String("c", "", "command  start or stop")
	instanceIds := flag.String("i", "", "The comma separated IDs of the instances to start/stop")

	flag.Parse()

	if *instanceIds == "" {
		fmt.Println("You must supply an instance ID (-i INSTANCE-ID or comma separated list of ids")
		return
	}

	if *command == "" {
		fmt.Println("You must supply an command  start or stop (-c start")
		return
	}

	instances := strings.Split(*instanceIds, ",")

	if *command == "stop" {
		StopInstancesCmd(instances)
	}

	if *command == "start" {
		StartInstancesCmd(instances)
	}
}
