package caculate
import(
	"math"
)

func Add(a, b float64) float64{
	return a+b
}
func Minus(a, b float64) float64{
	return a-b
}
func Times(a, b float64) float64{
	return a*b
}
func Divide(a,b float64) float64{
	return a/b
}
func Mod(a, b float64) float64{
	return math.Mod(a, b)
}
func Sqrt(a float64) float64{
	return math.Sqrt(a)
}
func Pow(a float64) float64{
	return a*a
}
