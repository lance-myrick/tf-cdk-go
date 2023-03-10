package cloudResources

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func EcsDemoAppStack(scope constructs.Construct, id *string) (cdktf.TerraformStack, error) {
	stack := cdktf.NewTerraformStack(scope, id)

	err := IamEc2InstanceProfile(scope, id)
	if err != nil {
		return stack, err
	}

	return stack, err
}
