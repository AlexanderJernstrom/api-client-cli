package main

import (
	"github.com/AlexanderJernstrom/api-client-cli/lib/api"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()

	inputValue := ""

	form := tview.NewForm().
					AddInputField("URL", "", 40, nil, func(text string) {
						inputValue = text
					}).
					AddButton("Send", func() {
						api.GetRequest(inputValue)
					}).SetButtonBackgroundColor(tcell.ColorBlueViolet)
	
	form.SetBorder(true).SetTitle("Blaze API client")

	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}

}