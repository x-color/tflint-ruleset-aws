// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogProductInvalidSupportURLRule checks the pattern is valid
type AwsServicecatalogProductInvalidSupportURLRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsServicecatalogProductInvalidSupportURLRule returns new rule with default attributes
func NewAwsServicecatalogProductInvalidSupportURLRule() *AwsServicecatalogProductInvalidSupportURLRule {
	return &AwsServicecatalogProductInvalidSupportURLRule{
		resourceType:  "aws_servicecatalog_product",
		attributeName: "support_url",
		max:           2083,
	}
}

// Name returns the rule name
func (r *AwsServicecatalogProductInvalidSupportURLRule) Name() string {
	return "aws_servicecatalog_product_invalid_support_url"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogProductInvalidSupportURLRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogProductInvalidSupportURLRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogProductInvalidSupportURLRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogProductInvalidSupportURLRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"support_url must be 2083 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
