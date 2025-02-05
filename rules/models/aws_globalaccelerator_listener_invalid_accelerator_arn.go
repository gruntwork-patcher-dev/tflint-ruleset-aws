// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGlobalacceleratorListenerInvalidAcceleratorArnRule checks the pattern is valid
type AwsGlobalacceleratorListenerInvalidAcceleratorArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsGlobalacceleratorListenerInvalidAcceleratorArnRule returns new rule with default attributes
func NewAwsGlobalacceleratorListenerInvalidAcceleratorArnRule() *AwsGlobalacceleratorListenerInvalidAcceleratorArnRule {
	return &AwsGlobalacceleratorListenerInvalidAcceleratorArnRule{
		resourceType:  "aws_globalaccelerator_listener",
		attributeName: "accelerator_arn",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsGlobalacceleratorListenerInvalidAcceleratorArnRule) Name() string {
	return "aws_globalaccelerator_listener_invalid_accelerator_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlobalacceleratorListenerInvalidAcceleratorArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlobalacceleratorListenerInvalidAcceleratorArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlobalacceleratorListenerInvalidAcceleratorArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlobalacceleratorListenerInvalidAcceleratorArnRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"accelerator_arn must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
