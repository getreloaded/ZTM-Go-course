package main

import (
	"encoding/json"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var client *http.Client

type Fact struct {
	Text string `json:"text,omitempty"`
}

func GetFact() (Fact, error) {
	var f Fact
	resp, err := client.Get("https://uselessfacts.jsph.pl/api/v2/facts/random?language=en")
	if err != nil {
		return Fact{}, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&f)
	if err != nil {
		return Fact{}, err
	}

	return f, nil
}

func main() {
	app := app.New()
	win := app.NewWindow("Facts")
	win.Resize(fyne.NewSize(400, 300))

	client = &http.Client{
		Timeout: 10 * time.Second,
	}

	// fact, err := GetFact()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(fact.Text)
	// }

	title := widget.NewLabel("Get your useless facts")
	title.Alignment = fyne.TextAlignCenter
	title.Wrapping = fyne.TextWrapWord
	title.TextStyle.Bold = true

	sep := widget.NewSeparator()

	factText := widget.NewLabel("")
	factText.Alignment = fyne.TextAlignLeading
	factText.Wrapping = fyne.TextWrapWord

	button := widget.NewButton("Get the next facts", func() {
		f, err := GetFact()
		if err != nil {
			dialog.ShowError(err, win)
		} else {
			factText.SetText(f.Text)
		}
	})
	hbox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	vbox := container.New(layout.NewVBoxLayout(), title, sep, hbox, sep, factText)

	win.SetContent(vbox)
	win.Show()
	app.Run()

}
