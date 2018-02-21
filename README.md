# Ascension

## Initial Setup
Install go to your system. [Download Go](https://golang.org/doc/install)
Once installed, go should have setup a directory structure for your $GOPATH.
Check golangs page to see how a workspace is setup. [Go Workspace](https://golang.org/doc/code.html#Workspaces)

Then once your workspace is setup up clone this repository to ~HOME/go/src/github.com/

## How to use Hass cmdline test tool
cd to main/hack
The cmd line tool can do Get and Post requests.
Pass just the -url flag for a get request.
E.g. go run cmd.go -url http://10.5.0.99:8123/api/

For post requests add -post flag and enter a payload into data.json under the hack directory:  
go run cmd.go -url http://10.5.0.99:8123/api/ -post