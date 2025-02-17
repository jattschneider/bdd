package bdd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Criteria struct {
	t            *testing.T
	testScenario string
}

func ToTest(t *testing.T, testScenario string) *Criteria {
	return &Criteria{t: t, testScenario: fmt.Sprint(testScenario, " Scenario")}
}

// Given the initial context at the beginning of the scenario, in one or more clauses.
func (c *Criteria) Given(scenarioInitialContext string, f func(t *testing.T)) *Criteria {
	c.t.Log(c.testScenario, ":", "Given", scenarioInitialContext)
	f(c.t)
	return c
}

// When the event that triggers the scenario.
func (c *Criteria) When(scenarioEventTrigger string, f func(t *testing.T)) *Criteria {
	c.t.Log(c.testScenario, ":", "When", scenarioEventTrigger)
	f(c.t)
	return c
}

// Then the expected outcome, in one or more clauses.
func (c *Criteria) Then(scenarioExpectedOutcome string, scenarioPass bool) *Criteria {
	c.t.Log(c.testScenario, ":", "Then", scenarioExpectedOutcome)
	if !scenarioPass {
		c.t.Fail()
	}
	return c
}

// Then the expected outcome, in one or more clauses.
func (c *Criteria) ThenShouldEqual(expected, actual any) *Criteria {
	c.t.Log(c.testScenario, ":", "Then actual(", actual, ") should equal expected(", expected, ")")
	if !assert.Equal(c.t, expected, actual) {
		c.t.Fail()
	}
	return c
}

// Then the expected outcome, in one or more clauses.
func (c *Criteria) ThenShouldFail(actualErr error) *Criteria {
	if !assert.Error(c.t, actualErr) {
		c.t.Fail()
	}
	return c
}

// Then the expected outcome, in one or more clauses.
func (c *Criteria) ThenShouldNotFail(actualErr error) *Criteria {
	if !assert.NoError(c.t, actualErr) {
		c.t.Fail()
	}
	return c
}

func (c *Criteria) Passed() {
	if !c.t.Failed() {
		c.t.Log(c.testScenario, ":", "PASS!")
		return
	}
	c.t.Log(c.testScenario, ":", "FAIL!")
}
