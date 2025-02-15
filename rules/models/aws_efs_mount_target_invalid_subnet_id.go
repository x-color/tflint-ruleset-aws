// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEfsMountTargetInvalidSubnetIDRule checks the pattern is valid
type AwsEfsMountTargetInvalidSubnetIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsEfsMountTargetInvalidSubnetIDRule returns new rule with default attributes
func NewAwsEfsMountTargetInvalidSubnetIDRule() *AwsEfsMountTargetInvalidSubnetIDRule {
	return &AwsEfsMountTargetInvalidSubnetIDRule{
		resourceType:  "aws_efs_mount_target",
		attributeName: "subnet_id",
		max:           47,
		min:           15,
		pattern:       regexp.MustCompile(`^subnet-[0-9a-f]{8,40}$`),
	}
}

// Name returns the rule name
func (r *AwsEfsMountTargetInvalidSubnetIDRule) Name() string {
	return "aws_efs_mount_target_invalid_subnet_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEfsMountTargetInvalidSubnetIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEfsMountTargetInvalidSubnetIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEfsMountTargetInvalidSubnetIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEfsMountTargetInvalidSubnetIDRule) Check(runner tflint.Runner) error {
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
					"subnet_id must be 47 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"subnet_id must be 15 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^subnet-[0-9a-f]{8,40}$`),
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
