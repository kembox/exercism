package complexnumbers

import "math"

// Define the Number type here.

type Number struct {
	real float64
	img float64	
}

var img_unit = math.Sqrt(-1)

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.img
}

func (n1 Number) Add(n2 Number) Number {
	return Number{(n1.Real() + n2.Real()), n1.Imaginary() + n2.Imaginary()}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{(n1.Real() - n2.Real()), n1.Imaginary() - n2.Imaginary()}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{n1.Real() * n2.Real() - n1.Imaginary() * n2.Imaginary(),n1.Imaginary()*n2.Real() + n1.Real() * n2.Imaginary()}
}
func (n Number) Times(factor float64) Number {
	return Number{ factor * n.real,
		factor * n.img,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	divisor := math.Pow(n2.real,2) + math.Pow(n2.img,2)
	return Number{(n1.Real() * n2.Real() + n1.Imaginary() * n2.Imaginary())/divisor,
		(n1.Imaginary()*n2.Real() - n1.Real()*n2.Imaginary())/divisor,
	}
}

func (n Number) Conjugate() Number {
	return Number{n.Real(), -n.Imaginary()}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.Real()*n.Real() + n.Imaginary()*n.Imaginary())
}

func (n Number) Exp() Number {
	factor := math.Exp(n.real)
	return Number{
		factor * math.Cos(n.img),
		factor * math.Sin(n.img),
	}
}