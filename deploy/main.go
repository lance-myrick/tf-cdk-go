package main

import (
	stack "cdk.tf/go/stack/serviceResources"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(nil)

	stack.EcsDemoAppStack(app, "deploy")

	app.Synth()
}
