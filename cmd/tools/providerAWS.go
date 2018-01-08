package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

//#####################
// ****func ProviderAWS****
// Create seesion with AWS account
// Take tags filter as Args
// Listing all instances with filter tag
// LIST OF ALL INSTANCES RUNNING IN DEV ENV
// LIST OF ALL INSTANCES COMPUTES - LAST COMMIT 3 DAYS OLD  ...
//###############################

func ProviderAWS() []string {
	var s []string
	sess := session.Must(session.NewSession())
	nameFilter := os.Args[1]

	awsRegion := "ap-southeast-1"
	svc := ec2.New(sess, &aws.Config{Region: aws.String(awsRegion)})
	fmt.Printf("Listing instances with tag %v in: %v\n", nameFilter, awsRegion)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag-value"),
				Values: []*string{
					aws.String(strings.Join([]string{"*", nameFilter, "*"}, "")),
				},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("There was an error listing instances in", awsRegion, err.Error())
		log.Fatal(err.Error())
	}

	table := tablewriter.NewWriter(os.Stdout)
	tableDel := tablewriter.NewWriter(os.Stdout)
	tableDel.SetHeader([]string{"ENVIROMENT", "InstanceId", "PublicDnsName", "BRANCH", "GITURL", "SERVER_STATUS", "LAST"})
	tableDel.SetBorder(true)
	table.SetHeader([]string{"ENVIROMENT", "InstanceId", "PublicDnsName", "BRANCH", "GITURL", "SERVER_STATUS", "LAST"})
	table.SetBorder(true)

	for _, i := range resp.Reservations {
		var ntt string
		var scm string
		var envar string
		for _, t := range i.Instances[0].Tags {
			if *t.Key == "BRANCH" {
				ntt = *t.Value
			}
			if *t.Key == "GIT" {
				scm = *t.Value
			}
			if *t.Key == "ENV" {
				envar = *t.Value
			}
		}

		ntt = "refs/heads/" + ntt

		var branch = plumbing.ReferenceName(ntt)
		u, x := ScmGIT(scm, branch)
		table.Append([]string{envar, *i.Instances[0].InstanceId, *i.Instances[0].PrivateDnsName, ntt, scm, x, u.String()})
		if x == "Unused" {
			s = append(s, *i.Instances[0].InstanceId)
			tableDel.Append([]string{envar, *i.Instances[0].InstanceId, *i.Instances[0].PrivateDnsName, ntt, scm, x, u.String()})

		}

	}
	Info("LIST OF ALL INSTANCES RUNNING IN DEV ENV  ...")
	table.Render()
	Info("LIST OF ALL INSTANCES COMPUTES - LAST COMMIT 3 DAYS OLD  ...")
	tableDel.Render()
	return s

}
