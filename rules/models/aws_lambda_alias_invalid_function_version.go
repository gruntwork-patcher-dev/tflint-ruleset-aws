// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaAliasInvalidFunctionVersionRule checks the pattern is valid
type AwsLambdaAliasInvalidFunctionVersionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsLambdaAliasInvalidFunctionVersionRule returns new rule with default attributes
func NewAwsLambdaAliasInvalidFunctionVersionRule() *AwsLambdaAliasInvalidFunctionVersionRule {
	return &AwsLambdaAliasInvalidFunctionVersionRule{
		resourceType:  "aws_lambda_alias",
		attributeName: "function_version",
		max:           1024,
		min:           1,
		pattern:       regexp.MustCompile(`^(\$LATEST|[0-9]+)$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaAliasInvalidFunctionVersionRule) Name() string {
	return "aws_lambda_alias_invalid_function_version"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaAliasInvalidFunctionVersionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaAliasInvalidFunctionVersionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaAliasInvalidFunctionVersionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaAliasInvalidFunctionVersionRule) Check(runner tflint.Runner) error {
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
					"function_version must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"function_version must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(\$LATEST|[0-9]+)$`),
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
