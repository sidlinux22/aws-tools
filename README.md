# AWS Tools


# Tool Details

Automated process of launching feature servers by all developers through a chatbot.

Automated process where following step occurs:

1. Check the development servers (for internal testing of features) running on the AWS cloud.
2. Each server has its specific code-base, corresponding to a feature branch in Github 
3. Check the last commit on that Github branch, and if it’s older than three days, terminate the server.


### Prerequisites
You should have aws account with  ec2 resource created:
* AWS access/secret keys
* All instances running in AWS should have proper Tags 

Basic required Tags ENV,GIT,BRANCH:

Example:

| Key   |      Value      |  
|----------|:-------------:|
| ENV |  DEV| 
| GIT |    https://github.com/sidlinux22/ci_environment.git  |
| BRANCH| feature/add-test | 

### Installing

```
git clone https://github.com/sidlinux22/aws-tools.git

bash-3.2$ make clean
rm -f bin/e2-unused-vmcleanup-*

bash-3.2$ make darwin
cd /Users/siddharthsharma/omsairam/aws-tools/cmd/tools/; \
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.VERSION=? -X main.COMMIT=d061b30fd7f1c15873c186ef7857ff485f5042e0 -X main.BRANCH=master" -o /Users/siddharthsharma/omsairam/aws-tools/bin/e2-unused-vmcleanup-darwin-amd64 . ; \
	cd - >/dev/null

bash-3.2$ ls -ld bin/e2-unused-vmcleanup-darwin-amd64
-rwxr-xr-x  1 siddharthsharma  staff  11390204 Jan  8 16:56 bin/e2-unused-vmcleanup-darwin-amd64

```

Cleanup:
```
bash-3.2$ make clean
rm -f bin/e2-unused-vmcleanup-*
```

### Usage :

1. Clone git project to workstation
2. Run build following step mention in installation

Script usage:
```
bin/e2-unused-vmcleanup-darwin-amd64 "DEV"

*This script take argument as to filter all ec2 instance running with "DEV" flag

```

## Demo

![gif](https://github.com/sidlinux22/aws-tools/blob/master/tmp/tty.gif)



### Description
 
*  This tools will list all in AWS region fliter by tags
As currently we are flitering all the instance with ENV key value set to DEV
*  Get the branch and github repo details provided as instance tags
*  Fetching all the objects of each repository one by one (everything in memory) and retrieves the commit history
*  Validate the last commit timestamp and marked all instance with 3 days old commit as "Unused"
*  List all the aws instance with filter tag (both Unused and Inused)
*  List all the aws instance marked as "Unused"
* Prompt user for confirmation before terminating "Unused" instance
* If  input "YES/y" terminate all the "Unused" instance and response back result.


### Tested

This tool is been tested with follow:
```
go version go1.8 darwin/amd64
Darwin Kernel Version 16.7.0
AWS EC2 resource 
```
### Troubleshooting


*  AWS_ACCESS_KEY/AWS_SECRET_KEY are exported as env variable


* Make sure you have passed tag filter as command ARG

Example: 
```
bin/e2-unused-vmcleanup-darwin-amd64 "DEV
```

<Error messsage>
"
panic: runtime error: index out of range

goroutine 1 [running]:
main.ProviderAWS(0x143bd80, 0x0, 0x0)
	/Users/siddharthsharma/aws-tools/cmd/tools/providerAWS.go:29 +0x13be
main.main()
	/Users/siddharthsharma/aws-tools/cmd/tools/main.go:17 +0x26
	panic: runtime error: index out of range
	"


* Error "error: reference not found"

> Instance branch tags value is incorrect or branch is delete.



* GIT  related Error
<Error messsage>
"error: authentication required "
> Make sure you have access to repo and valid repo - Validate the GIT tags key value for more details. 



### Reference


https://github.com/aws/aws-sdk-go
https://github.com/src-d/go-git 
https://github.com/olekukonko/tablewriter



