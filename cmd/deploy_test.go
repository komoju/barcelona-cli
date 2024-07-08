package cmd

import (
	"io/ioutil"
	"os"

	"github.com/jarcoal/httpmock"
	"github.com/urfave/cli"
)

func Example_deploy_quiet_false() {
	app := newTestApp(DeployCommand)
	app.Writer = os.Stdout
	testArgs := []string{"bcn", "deploy", "-e", "test", "--quiet=false"}
	pwd, _ := os.Getwd()
	endpoint := os.Getenv("BARCELONA_ENDPOINT")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resJson, _ := readJsonResponse(pwd + "/test/create_heritage.json")

	httpmock.RegisterResponder("PATCH", endpoint+"/v1/heritages/barcelona",
		httpmock.NewStringResponder(200, resJson))

	app.Run(testArgs)

	// Output:
	// Name:          nginx
	// Image Name:    nginx
	// Image Tag :    latest
	// Version:       1
	// Before Deploy: echo hello
	// Token:         560d9e10-70ce-4f47-82b8-37d47761116d
	// Scheduled Tasks:
	// rate(1 minute)       echo hello
	// Environment Variables
}

func Example_deploy_quiet_true() {
	app := newTestApp(DeployCommand)
	testArgs := []string{"bcn", "deploy", "-e", "test"}
	pwd, _ := os.Getwd()
	app.Writer = os.Stdout
	endpoint := os.Getenv("BARCELONA_ENDPOINT")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", endpoint+"/v1/auth/github/login",
		httpmock.NewStringResponder(200, "{}"))

	resJson, _ := readJsonResponse(pwd + "/test/create_heritage.json")

	httpmock.RegisterResponder("PATCH", endpoint+"/v1/heritages/barcelona",
		httpmock.NewStringResponder(200, resJson))

	app.Run(testArgs)

	// Output:
}

func newTestApp(command cli.Command) *cli.App {
	a := cli.NewApp()
	a.Name = "bcn"
	a.Writer = ioutil.Discard
	a.Commands = []cli.Command{command}

	pwd, _ := os.Getwd()
	HeritageConfigFilePath = pwd + "/test/test-barcelona.yml"

	return a
}

func readJsonResponse(path string) (string, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return "", err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	return string(byteValue), err
}
