// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDatasyncLocationEfsInvalidEfsFileSystemArnRule checks the pattern is valid
type AwsDatasyncLocationEfsInvalidEfsFileSystemArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationEfsInvalidEfsFileSystemArnRule returns new rule with default attributes
func NewAwsDatasyncLocationEfsInvalidEfsFileSystemArnRule() *AwsDatasyncLocationEfsInvalidEfsFileSystemArnRule {
	return &AwsDatasyncLocationEfsInvalidEfsFileSystemArnRule{
		resourceType:  "aws_datasync_location_efs",
		attributeName: "efs_file_system_arn",
		max:           128,
		pattern:       regexp.MustCompile(`^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):elasticfilesystem:[a-z\-0-9]*:[0-9]{12}:file-system/fs-.*$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationEfsInvalidEfsFileSystemArnRule) Name() string {
	return "aws_datasync_location_efs_invalid_efs_file_system_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationEfsInvalidEfsFileSystemArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncLocationEfsInvalidEfsFileSystemArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationEfsInvalidEfsFileSystemArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationEfsInvalidEfsFileSystemArnRule) Check(runner tflint.Runner) error {
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
					"efs_file_system_arn must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):elasticfilesystem:[a-z\-0-9]*:[0-9]{12}:file-system/fs-.*$`),
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
