package LinAlg

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Matrix struct {
	Rows int
	Cols int
	data []float64
}

//
// Implement interface 'GobEncoder'
//
func (m *Matrix) GobEncode() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(m.Rows)
	if err != nil {
		return nil, err
	}
	err = encoder.Encode(m.Cols)
	if err != nil {
		return nil, err
	}
	err = encoder.Encode(m.data)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

//
// Implement interface 'GobDecoder'
//
func (m *Matrix) GobDecode(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(&m.Rows)
	if err != nil {
		return err
	}
	err = decoder.Decode(&m.Cols)
	if err != nil {
		return err
	}
	return decoder.Decode(&m.data)
}

func MakeMatrix(rows int, cols int, data []float64) *Matrix {
	if size := rows * cols; size != len(data) {
		panic(fmt.Sprintf("LinAlg.Matrix.MakeMatrix: Matrix data has size %d, but %d expected", len(data), size))
	}
	return &Matrix{Rows: rows, Cols: cols, data: data}
}

func MakeEmptyMatrix(rows int, cols int) *Matrix {
	size := rows * cols
	return &Matrix{Rows: rows, Cols: cols, data: make([]float64, size)}
}

func (m *Matrix) index(row int, col int) int {
	return row*m.Cols + col
}

func (m Matrix) Set(row int, col int, value float64) {
	idx := m.index(row, col)
	m.data[idx] = value
}

func (m *Matrix) Get(row int, col int) float64 {
	idx := m.index(row, col)
	return m.data[idx]
}

func (m *Matrix) Transpose() *Matrix {
	t := MakeEmptyMatrix(m.Cols, m.Rows)
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			value := m.Get(row, col)
			t.Set(col, row, value)
		}
	}
	return t
}

func (m *Matrix) Ax(v *Vector) *Vector {
	if m.Cols != v.Size() {
		panic(fmt.Sprintf("LinAlg.Matrix.Ax: Matrix number of columns %d must equal vector size %d", m.Cols, v.Size()))
	}
	result := MakeEmptyVector(m.Rows)
	for row := 0; row < m.Rows; row++ {
		var value float64
		for col := 0; col < m.Cols; col++ {
			value += m.Get(row, col) * v.Get(col)
		}
		result.Set(row, value)
	}
	return result
}

func (m *Matrix) Am(other *Matrix) *Matrix {
	if m.Cols != other.Rows {
		panic(fmt.Sprint("LinAlg.Matrix.Am: Matrices not compatible"))
	}
	result := MakeEmptyMatrix(m.Rows, other.Cols)
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			var value float64
			for k := 0; k < m.Cols; k++ {
				value += m.Get(row, k) * other.Get(k, col)
			}
			result.Set(row, col, value)
		}
	}
	return result
}

func (m *Matrix) Scalar(scalar float64) *Matrix {
	for idx := range m.data {
		m.data[idx] *= scalar
	}
	return m
}

func (m *Matrix) Add(other *Matrix) *Matrix {
	if m.Rows != other.Rows {
		panic(fmt.Sprintf("LinAlg.Matrix.Add: Matrix number of rows %d and %d must equal", m.Rows, other.Rows))
	}
	if m.Cols != other.Cols {
		panic(fmt.Sprintf("LinAlg.Matrix.Add: Matrix number of columns %d and %d must equal", m.Cols, other.Cols))
	}
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			e1 := m.Get(row, col)
			e2 := other.Get(row, col)
			value := e1 + e2
			m.Set(row, col, value)
		}
	}
	return m
}

func (m *Matrix) Sub(other *Matrix) *Matrix {
	if m.Rows != other.Rows {
		panic(fmt.Sprintf("LinAlg.Matrix.Sub: Matrix number of rows %d and %d must equal", m.Rows, other.Rows))
	}
	if m.Cols != other.Cols {
		panic(fmt.Sprintf("LinAlg.Matrix.Sub: Matrix number of columns %d and %d must equal", m.Cols, other.Cols))
	}
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			e1 := m.Get(row, col)
			e2 := other.Get(row, col)
			value := e1 - e2
			m.Set(row, col, value)
		}
	}
	return m
}
