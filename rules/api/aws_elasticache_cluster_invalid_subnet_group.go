// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsElastiCacheClusterInvalidSubnetGroupRule checks whether attribute value actually exists
type AwsElastiCacheClusterInvalidSubnetGroupRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsElastiCacheClusterInvalidSubnetGroupRule returns new rule with default attributes
func NewAwsElastiCacheClusterInvalidSubnetGroupRule() *AwsElastiCacheClusterInvalidSubnetGroupRule {
	return &AwsElastiCacheClusterInvalidSubnetGroupRule{
		resourceType:  "aws_elasticache_cluster",
		attributeName: "subnet_group_name",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsElastiCacheClusterInvalidSubnetGroupRule) Name() string {
	return "aws_elasticache_cluster_invalid_subnet_group"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastiCacheClusterInvalidSubnetGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastiCacheClusterInvalidSubnetGroupRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastiCacheClusterInvalidSubnetGroupRule) Link() string {
	return ""
}

// Metadata returns the metadata about deep checking
func (r *AwsElastiCacheClusterInvalidSubnetGroupRule) Metadata() interface{} {
	return map[string]bool{"deep": true}
}

// Check checks whether the attributes are included in the list retrieved by DescribeCacheSubnetGroups
func (r *AwsElastiCacheClusterInvalidSubnetGroupRule) Check(rr tflint.Runner) error {
	runner := rr.(*aws.Runner)

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
			{Name: "provider"},
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

		if !r.dataPrepared {
			awsClient, err := runner.AwsClient(resource.Body.Attributes)
			if err != nil {
				return err
			}
			logger.Debug("invoking DescribeCacheSubnetGroups")
			r.data, err = awsClient.DescribeCacheSubnetGroups()
			if err != nil {
				err := fmt.Errorf("An error occurred while invoking DescribeCacheSubnetGroups; %w", err)
				logger.Error("%s", err)
				return err
			}
			r.dataPrepared = true
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.data[val] {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is invalid subnet group name.`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	}

	return nil
}
