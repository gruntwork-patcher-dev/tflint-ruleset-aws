// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSignerSigningJobInvalidProfileNameRule checks the pattern is valid
type AwsSignerSigningJobInvalidProfileNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSignerSigningJobInvalidProfileNameRule returns new rule with default attributes
func NewAwsSignerSigningJobInvalidProfileNameRule() *AwsSignerSigningJobInvalidProfileNameRule {
	return &AwsSignerSigningJobInvalidProfileNameRule{
		resourceType:  "aws_signer_signing_job",
		attributeName: "profile_name",
		max:           64,
		min:           2,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_]{2,}`),
	}
}

// Name returns the rule name
func (r *AwsSignerSigningJobInvalidProfileNameRule) Name() string {
	return "aws_signer_signing_job_invalid_profile_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSignerSigningJobInvalidProfileNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSignerSigningJobInvalidProfileNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSignerSigningJobInvalidProfileNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSignerSigningJobInvalidProfileNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"profile_name must be 64 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"profile_name must be 2 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_]{2,}`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
