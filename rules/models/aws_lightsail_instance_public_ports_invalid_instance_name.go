// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLightsailInstancePublicPortsInvalidInstanceNameRule checks the pattern is valid
type AwsLightsailInstancePublicPortsInvalidInstanceNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLightsailInstancePublicPortsInvalidInstanceNameRule returns new rule with default attributes
func NewAwsLightsailInstancePublicPortsInvalidInstanceNameRule() *AwsLightsailInstancePublicPortsInvalidInstanceNameRule {
	return &AwsLightsailInstancePublicPortsInvalidInstanceNameRule{
		resourceType:  "aws_lightsail_instance_public_ports",
		attributeName: "instance_name",
		pattern:       regexp.MustCompile(`^\w[\w\-]*\w$`),
	}
}

// Name returns the rule name
func (r *AwsLightsailInstancePublicPortsInvalidInstanceNameRule) Name() string {
	return "aws_lightsail_instance_public_ports_invalid_instance_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLightsailInstancePublicPortsInvalidInstanceNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLightsailInstancePublicPortsInvalidInstanceNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLightsailInstancePublicPortsInvalidInstanceNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLightsailInstancePublicPortsInvalidInstanceNameRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\w[\w\-]*\w$`),
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
