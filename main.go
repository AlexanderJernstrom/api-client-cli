package main

import (
	"github.com/AlexanderJernstrom/api-client-cli/lib/api"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()

	inputValue := ""
	var apiResponse api.Response
	statusCodeText := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	form := tview.NewForm().AddInputField("URL", "", 40, nil, func(text string) {
		inputValue = text
	}).AddButton("Send", func() {
		apiResponse = api.GetRequest(inputValue)
		statusCodeText.SetText(apiResponse.RawResponse)
	})

	/*flexContainer := tview.NewFlex().AddItem(form, 0, 1, true).AddItem(resultBox, 0, 1, true).SetBorder(true).
	SetTitle("Blaze API client")
	*/
	testFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true).
		AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorDarkBlue), 0, 2, false).
		AddItem(statusCodeText, 0, 1, false)

	if err := app.SetRoot(testFlex, true).Run(); err != nil {
		panic(err)
	}

}
