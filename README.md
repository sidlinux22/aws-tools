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



### Description


### Tested


### Troubleshooting





