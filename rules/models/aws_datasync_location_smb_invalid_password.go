// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDatasyncLocationSmbInvalidPasswordRule checks the pattern is valid
type AwsDatasyncLocationSmbInvalidPasswordRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationSmbInvalidPasswordRule returns new rule with default attributes
func NewAwsDatasyncLocationSmbInvalidPasswordRule() *AwsDatasyncLocationSmbInvalidPasswordRule {
	return &AwsDatasyncLocationSmbInvalidPasswordRule{
		resourceType:  "aws_datasync_location_smb",
		attributeName: "password",
		max:           104,
		pattern:       regexp.MustCompile(`^.{0,104}$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationSmbInvalidPasswordRule) Name() string {
	return "aws_datasync_location_smb_invalid_password"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationSmbInvalidPasswordRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncLocationSmbInvalidPasswordRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationSmbInvalidPasswordRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationSmbInvalidPasswordRule) Check(runner tflint.Runner) error {
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
					"password must be 104 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`password does not match valid pattern ^.{0,104}$`,
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
