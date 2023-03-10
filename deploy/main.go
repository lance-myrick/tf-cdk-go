package main

import (
	stack "cdk.tf/go/stack/cloudResources"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(nil)

	stack.EcsDemoAppStack(app, "EcsDemoAppStackDev")

	app.Synth()
}
