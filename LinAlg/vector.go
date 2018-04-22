package LinAlg

import "fmt"

type Vector struct {
	data []float64
}

func MakeVector(data []float64) Vector {
	return Vector{data: data}
}

func MakeEmptyVector(size int) Vector {
	return Vector{data: make([]float64, size)}
}

func (v *Vector) Size() int {
	return len(v.data)
}

func (v *Vector) Set(index int, value float64) {
	v.data[index] = value
}

func (v *Vector) Get(index int) float64 {
	return v.data[index]
}

func (v1 *Vector) DotProduct(v2 *Vector) float64 {
	if v1.Size() != v2.Size() {
		panic(fmt.Sprintf("LinAlg.Vector.DotProduct: Vector sizes %d and %d must be the same", v1.Size(), v2.Size()))
	}
	var d float64 = 0
	for i := 0; i < v1.Size(); i++ {
		e1 := v1.Get(i)
		e2 := v2.Get(i)
		d += e1 * e2
	}
	return d
}

func (v *Vector) ScalarMultiplication(scalar float64) {
	for idx := range v.data {
		v.data[idx] *= scalar
	}
}

func (v *Vector) F(f func(float64) float64) {
	for idx := range v.data {
		value := v.data[idx]
		value = f(value)
		v.data[idx] = value
	}
}

func (v *Vector) Hadamard(other *Vector) Vector {
	if v.Size() != other.Size() {
		panic(fmt.Sprintf("LinAlg.Vector.Hadamard: Vectors must have same size, but is %d and %d", v.Size(), other.Size()))
	}
	result := MakeEmptyVector(v.Size())
	for idx := range v.data {
		e1 := v.data[idx]
		e2 := other.data[idx]
		result.Set(idx, e1*e2)
	}
	return result
}