// packaga similar to namespace use
package main

import (
	_ "github.com/aiuoe/go_learn/basics"
	_ "github.com/aiuoe/go_learn/data"
	_errors "github.com/aiuoe/go_learn/errors"
	_ "github.com/aiuoe/go_learn/flow_controls"
)

func main() {
	// basics.Vars()
	// basics.Consts()
	// basics.Types()
	// basics.Parse()
	// basics.Fmt()
	// basics.Math()

	// flowcontrols.IfElse()
	// flowcontrols.Switch()
	// flowcontrols.For()

	// fmt.Println("Sum:", flowcontrols.Sum(12, 12))
	// fmt.Println("Multiply", flowcontrols.Multiply(12, 12))

	// with 0 as b value throw error
	// divideResult, e := flowcontrols.Divide(12, 0)

	// if e != nil {
	// 	fmt.Println("Error:", e)
	// }

	// fmt.Println("Divide:", divideResult)

	// flowcontrols.Menu()

	// data.Array()
	// data.Slices()
	// data.Map()
	// data.Structs()
	// data.Pointers()
	// data.Menu()

	_errors.Handler()
}
