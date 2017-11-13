package main 

import (
	"fmt"
	"strconv"
)

type Fraction struct {
	numerator int
	denominator int
}

func (a *Fraction) Mul(b *Fraction) *Fraction {
	c := &Fraction {a.numerator * b.numerator, a.denominator * b.denominator}

	c.simplify()

	return c
}

func (a *Fraction) Div(b *Fraction) *Fraction {
	c := &Fraction {a.numerator * b.denominator, a.denominator * b.numerator}

	c.simplify()

	return c
}

func (a *Fraction) Add(b *Fraction) *Fraction {	
	a1 := &Fraction {a.numerator * b.denominator, a.denominator * b.denominator}
	b1 := &Fraction {b.numerator * a.denominator, b.denominator * a.denominator}

	c := &Fraction {a1.numerator + b1.numerator, a1.denominator}

	c.simplify()

	return c
}

func (f *Fraction) String() string {
	f.simplify ()

	return strconv.Itoa (f.numerator) + "/" + strconv.Itoa (f.denominator)
}

func (f *Fraction) Approx () float64 {
	return float64 (f.numerator/f.denominator)
}

func NewFraction (num, denom int) *Fraction {
	f := &Fraction {num, denom}
	
	f.simplify()

	return f
}

func (f *Fraction) simplify () {
	c := &Fraction {f.numerator/gcd (f.numerator, f.denominator), f.denominator/gcd (f.numerator, f.denominator)}

	f.numerator = c.numerator
	f.denominator = c.denominator
}

func gcd (a, b int) int {
	if (a % b == 0) {
		return b
	}

	return gcd (b, a % b)
}

func main() {
	f := &Fraction {6, 13}
	f1 := &Fraction {8, 14}

	f = f.Add (f1)

	fmt.Println (f.String())
}