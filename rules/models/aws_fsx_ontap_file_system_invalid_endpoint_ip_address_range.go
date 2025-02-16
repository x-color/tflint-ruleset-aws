// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFsxOntapFileSystemInvalidEndpointIPAddressRangeRule checks the pattern is valid
type AwsFsxOntapFileSystemInvalidEndpointIPAddressRangeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFsxOntapFileSystemInvalidEndpointIPAddressRangeRule returns new rule with default attributes
func NewAwsFsxOntapFileSystemInvalidEndpointIPAddressRangeRule() *AwsFsxOntapFileSystemInvalidEndpointIPAddressRangeRule {
	return &AwsFsxOntapFileSystemInvalidEndpointIPAddressRangeRule{
		resourceType:  "aws_fsx_ontap_file_system",
		attributeName: "endpoint_ip_address_range",
		max:           17,
		min:           9,
		pattern:       regexp.MustCompile(`^[^\x{0000}\x{0085}\x{2028}\x{2029}\r\n]{9,17}$`),
	}
}

// Name returns the rule name
func (r *AwsFsxOntapFileSystemInvalidEndpointIPAddressRangeRule) Name() string {
	return "aws_fsx_ontap_file_system_invalid_endpoint_ip_address_range"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxOntapFileSystemInvalidEndpointIPAddressRangeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxOntapFileSystemInvalidEndpointIPAddressRangeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxOntapFileSystemInvalidEndpointIPAddressRangeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxOntapFileSystemInvalidEndpointIPAddressRangeRule) Check(runner tflint.Runner) error {
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
					"endpoint_ip_address_range must be 17 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"endpoint_ip_address_range must be 9 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[^\x{0000}\x{0085}\x{2028}\x{2029}\r\n]{9,17}$`),
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
