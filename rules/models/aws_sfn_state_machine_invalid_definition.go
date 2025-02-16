// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSfnStateMachineInvalidDefinitionRule checks the pattern is valid
type AwsSfnStateMachineInvalidDefinitionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSfnStateMachineInvalidDefinitionRule returns new rule with default attributes
func NewAwsSfnStateMachineInvalidDefinitionRule() *AwsSfnStateMachineInvalidDefinitionRule {
	return &AwsSfnStateMachineInvalidDefinitionRule{
		resourceType:  "aws_sfn_state_machine",
		attributeName: "definition",
		max:           1048576,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSfnStateMachineInvalidDefinitionRule) Name() string {
	return "aws_sfn_state_machine_invalid_definition"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSfnStateMachineInvalidDefinitionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSfnStateMachineInvalidDefinitionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSfnStateMachineInvalidDefinitionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSfnStateMachineInvalidDefinitionRule) Check(runner tflint.Runner) error {
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
					"definition must be 1048576 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"definition must be 1 characters or higher",
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
