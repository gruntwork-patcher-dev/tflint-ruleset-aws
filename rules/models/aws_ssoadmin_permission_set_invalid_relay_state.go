// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsoadminPermissionSetInvalidRelayStateRule checks the pattern is valid
type AwsSsoadminPermissionSetInvalidRelayStateRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsoadminPermissionSetInvalidRelayStateRule returns new rule with default attributes
func NewAwsSsoadminPermissionSetInvalidRelayStateRule() *AwsSsoadminPermissionSetInvalidRelayStateRule {
	return &AwsSsoadminPermissionSetInvalidRelayStateRule{
		resourceType:  "aws_ssoadmin_permission_set",
		attributeName: "relay_state",
		max:           240,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9&$@#\\\/%?=~\-_'"|!:,.;*+\[\]\ \(\)\{\}]+$`),
	}
}

// Name returns the rule name
func (r *AwsSsoadminPermissionSetInvalidRelayStateRule) Name() string {
	return "aws_ssoadmin_permission_set_invalid_relay_state"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsoadminPermissionSetInvalidRelayStateRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsoadminPermissionSetInvalidRelayStateRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsoadminPermissionSetInvalidRelayStateRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsoadminPermissionSetInvalidRelayStateRule) Check(runner tflint.Runner) error {
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
					"relay_state must be 240 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"relay_state must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9&$@#\\\/%?=~\-_'"|!:,.;*+\[\]\ \(\)\{\}]+$`),
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
