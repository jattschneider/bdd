# BDD

### Sample
```
import (
	"testing"

	"github.com/jattschneider/bdd"
)

func TestCalculator(t *testing.T) {
	calculator := New()

	t.Run("Add two positive numbers", func(t *testing.T) {
		var (
			result,
			number1,
			number2,
			expected int
		)

		bdd.ToTest(t, "Calculator").
			Given("two positive numbers", func(t *testing.T) {
				number1 = 2
				number2 = 3
				expected = 5
			}).
			When("we add them", func(t *testing.T) {
				result = calculator.Add(number1, number2)
			}).
			ThenShouldEqual(expected, result).
			Passed()
	})
}
```

```
type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func New() *Calculator {
	return &Calculator{}
}
```

### Running Tests

```
~ go test -v --race
=== RUN   TestCalculator
=== RUN   TestCalculator/Add_two_positive_numbers
    bdd.go:21: Calculator Scenario : Given two positive numbers
    bdd.go:28: Calculator Scenario : When we add them
    bdd.go:44: Calculator Scenario : Then actual( 5 ) should equal expected( 5 )
    bdd.go:69: Calculator Scenario : PASS!
=== RUN   TestCalculator/Add_positive_and_negative_number
    bdd.go:21: Calculator Scenario : Given positive and negative number
    bdd.go:28: Calculator Scenario : When we add them
    bdd.go:44: Calculator Scenario : Then actual( 2 ) should equal expected( 2 )
    bdd.go:69: Calculator Scenario : PASS!
=== RUN   TestCalculator/Add_two_negative_numbers
    bdd.go:21: Calculator Scenario : Given two negative numbers
    bdd.go:28: Calculator Scenario : When we add them
    bdd.go:44: Calculator Scenario : Then actual( -5 ) should equal expected( -5 )
    bdd.go:69: Calculator Scenario : PASS!
--- PASS: TestCalculator (0.00s)
    --- PASS: TestCalculator/Add_two_positive_numbers (0.00s)
    --- PASS: TestCalculator/Add_positive_and_negative_number (0.00s)
    --- PASS: TestCalculator/Add_two_negative_numbers (0.00s)
PASS
ok      github.com/jattschneider/codetests/calculator   1.022s
```