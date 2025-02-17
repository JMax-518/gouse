package geometry

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func SampleMathParallelogram() {
	println("Area of parallelogram (integer): ", math.AreaParallelogram(10, 20))
	println("Area of parallelogram (float): ", fmt.Sprintf("%f", math.AreaParallelogramF(10.0, 20.0)))
}
