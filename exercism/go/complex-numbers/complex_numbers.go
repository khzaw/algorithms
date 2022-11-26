package complexnumbers

import (
	"math"
)

// Define the Number type here.
type Number struct {
  a float64
  b float64
}

func (n Number) Real() float64 {
  return n.a
}

func (n Number) Imaginary() float64 {
  return n.b
}

func (n1 Number) Add(n2 Number) Number {
  a, b, c, d := n1.a, n1.b, n2.a, n2.b
  return Number{a + c, b + d}
}

func (n1 Number) Subtract(n2 Number) Number {
  a, b, c, d := n1.a, n1.b, n2.a, n2.b
  return Number{a - c, b - d}
}

func (n1 Number) Multiply(n2 Number) Number {
  a, b, c, d := n1.a, n1.b, n2.a, n2.b
  return Number{a * c - b * d, b * c + a * d}
}

func (n Number) Times(factor float64) Number {
  return Number{n.a * factor, n.b * factor}
}

func (n1 Number) Divide(n2 Number) Number {
  a, b, c, d := n1.a, n1.b, n2.a, n2.b
  return Number{
    (a * c + b * d) / (c * c + d * d),
    (b * c - a * d) / (c * c + d * d),
  }
}

func (n Number) Conjugate() Number {
  return Number{n.a, -n.b}
}

func (n Number) Abs() float64 {
  return math.Abs(math.Sqrt(n.a * n.a + n.b * n.b))
}

func (n Number) Exp() Number {
  return Number{
    a: math.Exp(n.a) * math.Cos(n.b),
    b: math.Exp(n.a) * math.Sin(n.b),
  }
}
