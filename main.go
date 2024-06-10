// packaga similar to namespace use
package main

import (
	_ "github.com/aiuoe/go_learn/basics"
	_ "github.com/aiuoe/go_learn/data"
	_ "github.com/aiuoe/go_learn/flow_controls"
	"github.com/aiuoe/go_learn/func_advanced"
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

	// _errors.Handler()
	// _errors.Defer()
	// _errors.Panic()
	// _errors.Logger()
	// _errors.Contacts()

	// database.TaskManager()

	// gomodule.Exec()
	// poo.ExecStruct()
	// poo.ExecInterface()

	// routines.Api()

	// func_advanced.Variadic()
	// func_advanced.Recursive()
	func_advanced.Clousure()
}
