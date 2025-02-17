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