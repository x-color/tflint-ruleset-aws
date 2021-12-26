// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalTypeRule checks the pattern is valid
type AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalTypeRule returns new rule with default attributes
func NewAwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalTypeRule() *AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalTypeRule {
	return &AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalTypeRule{
		resourceType:  "aws_servicecatalog_principal_portfolio_association",
		attributeName: "principal_type",
		enum: []string{
			"IAM",
		},
	}
}

// Name returns the rule name
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalTypeRule) Name() string {
	return "aws_servicecatalog_principal_portfolio_association_invalid_principal_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogPrincipalPortfolioAssociationInvalidPrincipalTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as principal_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
