package steps

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
)

func TestSimple(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeSimpleScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../simple.feature"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeSimpleScenario(ctx *godog.ScenarioContext) {

	ctx.Step(`^the string "([^"]*)"$`, theStringValue)
	ctx.Step(`^I run echo "([^"]*)"$`, iRunEchoValue)
	ctx.Step(`^the output should contain "([^"]*)"$`, theOutputShouldContainValue)
	ctx.Step(`^clear all$`, clearAll)
	ctx.Step(`^fail step$`, failStep)
}
func theStringValue(ctx context.Context, value string) (context.Context, error) {
	return ctx, nil
}
func iRunEchoValue(ctx context.Context, value string) (context.Context, error) {
	if value == "hello" {
		return ctx, nil
	}
	return ctx, fmt.Errorf("some error")
}

func theOutputShouldContainValue(ctx context.Context, value string) (context.Context, error) {
	//return ctx, godog.ErrPending
	return ctx, nil
}
func clearAll(ctx context.Context) (context.Context, error) {
	//return ctx, godog.ErrPending
	return ctx, nil
}
func failStep(ctx context.Context) (context.Context, error) {
	return ctx, fmt.Errorf("step failed")
}
