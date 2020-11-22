package main

import (
	"clx/cli"
	"clx/cmd"
	"clx/constructors"
	"clx/controller"
)

func main() {
	cmd.Execute()
	sc := constructor.NewScreenController()
	controller.InitializeScreenController(sc)
	controller.SetResizeFunction(sc.Application, sc.SubmissionStates, sc.MainView, sc.ApplicationState)

	cli.Clear()

	sc.Application.SetRoot(sc.MainView.Grid, true)

	if err := sc.Application.Run(); err != nil {
		panic(err)
	}

}