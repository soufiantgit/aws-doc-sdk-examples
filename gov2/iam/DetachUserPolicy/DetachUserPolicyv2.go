// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX - License - Identifier: Apache - 2.0
// snippet-start:[iam.go-v2.DetachUserPolicy]
package main

import (
    "context"
    "flag"
    "fmt"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/iam"
)

// IAMDetachRolePolicyAPI defines the interface for the DetachRolePolicy function.
// We use this interface to test the function using a mocked service.
type IAMDetachRolePolicyAPI interface {
    DetachRolePolicy(ctx context.Context,
        params *iam.DetachRolePolicyInput,
        optFns ...func(*iam.Options)) (*iam.DetachRolePolicyOutput, error)
}

// DetachDynamoFullPolicy detaches an Amazon DynamoDB full-access policy from an IAM role.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a DetachRolePolicyOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to DetachRolePolicy.
func DetachDynamoFullPolicy(c context.Context, api IAMDetachRolePolicyAPI, input *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error) {
    result, err := api.DetachRolePolicy(c, input)

    return result, err
}

func main() {
    roleName := flag.String("r", "", "The name of the IAM role")
    flag.Parse()

    if *roleName == "" {
        fmt.Println("You must supply a role name (-r ROLE)")
        return
    }

    cfg, err := config.LoadDefaultConfig()
    if err != nil {
        panic("configuration error, " + err.Error())
    }

    client := iam.NewFromConfig(cfg)

    policyArn := "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess"
    input := &iam.DetachRolePolicyInput{
        PolicyArn: &policyArn,
        RoleName:  roleName,
    }

    _, err = DetachDynamoFullPolicy(context.Background(), client, input)
    if err != nil {
        fmt.Println("Unable to detach DynamoDB full-access role policy to role")
        return
    }
    fmt.Println("Role detached successfully")
}

// snippet-end:[iam.go-v2.DetachUserPolicy]
