package matrix

import (
	"slices"
	"sync"
)

type PositionValue struct {
	X int
	Y int
}

type MatrixElement struct {
	Value any
	Type  string
}

type MatrixPoint struct {
	Position PositionValue
	Element  MatrixElement
	Char     string
}

type Matrix struct {
	values map[string]MatrixPoint
}

var lock = &sync.Mutex{}

var singleInstance *Matrix

func getInstance() *Matrix {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &Matrix{values: make(map[string]MatrixPoint)}
		}
	}

	return singleInstance
}

func genMatrixKey(x int, y int) string {
	return string(x) + "-" + string(y)
}

func CreateMatrixElement(value any, _type string) *MatrixElement {
	return &MatrixElement{
		Value: value,
		Type:  _type,
	}
}

func CreateMatrixPoint(matrixElement MatrixElement, char string, x int, y int) *MatrixPoint {
	return &MatrixPoint{
		Element:  matrixElement,
		Char:     char,
		Position: PositionValue{X: x, Y: y},
	}
}

func RegisterMatrixPoint(matrixPoint MatrixPoint) {
	instance := getInstance()

	instance.values[genMatrixKey(matrixPoint.Position.X, matrixPoint.Position.Y)] = matrixPoint
}

func GetMatrixPoint(x int, y int) *MatrixPoint {
	instance := getInstance()

	value, exists := instance.values[genMatrixKey(x, y)]

	if exists {
		return &value
	}

	return nil
}

func ListMatrixPointsByType(_type string) []MatrixPoint {
	instance := getInstance()

	var values []MatrixPoint

	for _, value := range instance.values {
		if value.Element.Type == _type {
			values = append(values, value)
		}
	}

	return values
}

func (matrixElement MatrixPoint) FindNeighborPoints() []MatrixPoint {
	directions := [8][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	var neighbors []MatrixPoint

	for _, direction := range directions {
		neighbor := GetMatrixPoint(matrixElement.Position.X+direction[0], matrixElement.Position.Y+direction[1])

		if neighbor != nil {
			neighbors = append(neighbors, *neighbor)
		}
	}

	return neighbors
}

func (matrixPoint MatrixPoint) FindNeighborElements() []MatrixElement {

	uniqNeighbors := make([]MatrixElement, 0)

	for _, n := range matrixPoint.FindNeighborPoints() {

		if !slices.Contains(uniqNeighbors, n.Element) {
			uniqNeighbors = append(uniqNeighbors, n.Element)
		}

	}
	return uniqNeighbors
}
