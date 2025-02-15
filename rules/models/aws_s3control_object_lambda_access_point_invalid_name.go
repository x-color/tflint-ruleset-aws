// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsS3controlObjectLambdaAccessPointInvalidNameRule checks the pattern is valid
type AwsS3controlObjectLambdaAccessPointInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsS3controlObjectLambdaAccessPointInvalidNameRule returns new rule with default attributes
func NewAwsS3controlObjectLambdaAccessPointInvalidNameRule() *AwsS3controlObjectLambdaAccessPointInvalidNameRule {
	return &AwsS3controlObjectLambdaAccessPointInvalidNameRule{
		resourceType:  "aws_s3control_object_lambda_access_point",
		attributeName: "name",
		max:           45,
		min:           3,
		pattern:       regexp.MustCompile(`^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`),
	}
}

// Name returns the rule name
func (r *AwsS3controlObjectLambdaAccessPointInvalidNameRule) Name() string {
	return "aws_s3control_object_lambda_access_point_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3controlObjectLambdaAccessPointInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3controlObjectLambdaAccessPointInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3controlObjectLambdaAccessPointInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3controlObjectLambdaAccessPointInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 45 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 3 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`),
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
