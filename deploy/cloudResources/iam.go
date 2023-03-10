package cloudResources

import (
	"encoding/json"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/iaminstanceprofile"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/iampolicy"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/iamrole"
)

func IamEc2InstanceProfile(scope constructs.Construct, id *string) error {

	tmpJson0, err := json.Marshal(map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			map[string]interface{}{
				"Action":    []string{"sts:AssumeRole"},
				"Effect":    "Allow",
				"Sid":       "AllowAssumeRole",
				"Principal": map[string]interface{}{"Service": "ec2.amazonaws.com"},
			},
		},
	})
	if err != nil {
		return err
	}
	jsonAssumeRolePolicy := string(tmpJson0)

	tmpJson1, err := json.Marshal(map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			map[string]interface{}{
				"Action": []string{"autoscaling:Describe*",
					"cloudwatch:Describe*",
					"cloudwatch:List*",
					"cloudwatch:PutMetricAlarm*",
					"cloudwatch:DeleteAlarms",
					"cloudwatch:Get*"},
				"Effect":    "Allow",
				"Resources": []string{"*"},
			},
		},
	})
	if err != nil {
		return err
	}
	jsonInlinePolicy := string(tmpJson1)

	iampolicy.NewIamPolicy(scope, id, &iampolicy.IamPolicyConfig{
		Description: jsii.String("Allows ec2 assume role"),
		Path:        jsii.String("/"),
		Policy:      jsii.String(jsonAssumeRolePolicy),
	})

	iamInstanceRole := iamrole.NewIamRole(scope, id, &iamrole.IamRoleConfig{
		AssumeRolePolicy:  jsii.String(jsonAssumeRolePolicy),
		Path:              jsii.String("/"),
		ManagedPolicyArns: jsii.Strings("arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"),
		InlinePolicy:      jsonInlinePolicy,
	})

	iaminstanceprofile.NewIamInstanceProfile(scope, id, &iaminstanceprofile.IamInstanceProfileConfig{
		Role: iamInstanceRole.Name(),
	})

	return nil
}
