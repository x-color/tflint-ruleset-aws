// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderImageInvalidDistributionConfigurationArnRule checks the pattern is valid
type AwsImagebuilderImageInvalidDistributionConfigurationArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsImagebuilderImageInvalidDistributionConfigurationArnRule returns new rule with default attributes
func NewAwsImagebuilderImageInvalidDistributionConfigurationArnRule() *AwsImagebuilderImageInvalidDistributionConfigurationArnRule {
	return &AwsImagebuilderImageInvalidDistributionConfigurationArnRule{
		resourceType:  "aws_imagebuilder_image",
		attributeName: "distribution_configuration_arn",
		pattern:       regexp.MustCompile(`^arn:aws[^:]*:imagebuilder:[^:]+:(?:[0-9]{12}|aws):distribution-configuration/[a-z0-9-_]+$`),
	}
}

// Name returns the rule name
func (r *AwsImagebuilderImageInvalidDistributionConfigurationArnRule) Name() string {
	return "aws_imagebuilder_image_invalid_distribution_configuration_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderImageInvalidDistributionConfigurationArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderImageInvalidDistributionConfigurationArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderImageInvalidDistributionConfigurationArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderImageInvalidDistributionConfigurationArnRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:aws[^:]*:imagebuilder:[^:]+:(?:[0-9]{12}|aws):distribution-configuration/[a-z0-9-_]+$`),
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
