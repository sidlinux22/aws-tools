package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	. "gopkg.in/src-d/go-git.v4/_examples"
)

func main() {
	s := ProviderAWS()
	if len(s) == 0 {
		Info("No unused instance for termination ...")
		os.Exit(0)
	}
	Info("Following instance are marked for termination ...")
	printSlice(s)
	c := askForConfirmation("Do you want termination all these instance ...?")
	if c {
		Warning("Deleting unused instances....")
		actionTerminate(s)

	}
}

func printSlice(s []string) {
	fmt.Println(s)
}

//
//  Func to terminated unused instances
//
func actionTerminate(sdel []string) {

	for _, s := range sdel {

		sess := session.Must(session.NewSession())
		awsRegion := "ap-southeast-1"
		svc := ec2.New(sess, &aws.Config{Region: aws.String(awsRegion)})

		input := &ec2.TerminateInstancesInput{
			InstanceIds: []*string{
				aws.String(s),
			},
		}
		result, err := svc.TerminateInstances(input)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("Success", result.TerminatingInstances)
		}

	}
}

//
//  Func askForConfirmation for confirmation before proceeding with termination
//

func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}
