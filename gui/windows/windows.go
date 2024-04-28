package gui

import (
	"bufio"
	"fmt"
	"gobot/state"
	"image/color"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	WIDTH  = 640
	HEIGHT = 640
)

func MainRender(app *fyne.App, st *state.AppState, window fyne.Window, prev fyne.Window) {

	score := canvas.NewText("Score", color.White)
	score.Alignment = fyne.TextAlignCenter
	score.TextSize = 20
	score.TextStyle.Bold = false

	points := widget.NewLabel("0")

	informWindow := (*app).NewWindow("About")
	informWindow.SetCloseIntercept(func() {
		informWindow.Hide()
	})
	informWindow.Hide()
	informWindow.Resize(fyne.NewSize(640, 480))

	rect := canvas.NewRectangle(state.NothingColor)
	rect.SetMinSize(fyne.NewSize(65, 65))

	nC := getCeilRectangle(state.NothingColor)
	roC := getCeilRectangle(state.RobotColor)
	pC := getCeilRectangle(state.PitColor)
	rC := getCeilRectangle(state.ResourceColor)
	eC := getCeilRectangle(state.EndColor)

	captionRobot := getCeilCaption("Robot")
	captionPit := getCeilCaption("Pit")
	captionNothing := getCeilCaption("Nothing")
	captionResource := getCeilCaption("Resource")
	captionEnd := getCeilCaption("End the map")

	about := container.New(
		layout.NewGridLayoutWithColumns(2),
		nC,
		captionNothing,
		roC,
		captionRobot,
		pC,
		captionPit,
		rC,
		captionResource,
		eC,
		captionEnd,
	)

	wrapAbout := container.New(layout.NewCenterLayout(), about)

	rulesLabel := widget.NewRichTextFromMarkdown(string(resourceAboutpart.Content()))

	informContainer := container.New(
		layout.NewVBoxLayout(),
		rulesLabel,
		wrapAbout,
	)

	informWindow.SetContent(informContainer)

	informationBtn := widget.NewButton(
		"?",
		func() {
			informWindow.Show()
		},
	)

	backBtn := widget.NewButton(
		"Back",
		func() {
			window.Hide()
			informWindow.Hide()
			prev.Show()
		},
	)

	header := container.New(
		layout.NewGridLayout(4),
		backBtn,
		score,
		points,
		informationBtn,
	)

	inputCommands := widget.NewMultiLineEntry()

	partEntry := canvas.NewRectangle(color.Transparent)
	partEntry.SetMinSize(fyne.NewSize(320, 500))

	input := container.New(
		layout.NewStackLayout(),
		partEntry,
		inputCommands,
	)

	field := container.New(
		layout.NewGridLayoutWithColumns(st.Columns),
	)

	data := ">> "
	lastPos := len(data)
	inputCommands.SetText(data)
	window.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		if ke.Name == fyne.KeyReturn {
			data = inputCommands.Text
			for index := lastPos; index < len(data); index++ {
				fmt.Printf("%c", data[index])

			}
			fmt.Println()

			data = data + "\n" + ">> "
			lastPos = len(data)
			inputCommands.SetText(data)
			//fieldRender(field, state)
		}
	})

	fieldRender(field, st)
	wrapField := container.New(layout.NewCenterLayout(), field)

	body := container.NewHSplit(
		input,
		wrapField,
	)

	partWindow := canvas.NewRectangle(color.Transparent)
	partWindow.SetMinSize(fyne.NewSize(100, 25))

	main_container := container.New(
		layout.NewVBoxLayout(),
		header,
		partWindow,
		body,
	)

	window.SetContent(main_container)
}

func fieldRender(grid *fyne.Container, st *state.AppState) {
	grid.RemoveAll()
	for i := 0; i < len(st.CellularField); i++ {
		fill := getColor(int(st.CellularField[i]))

		box := canvas.NewRectangle(fill)
		box.SetMinSize(fyne.NewSize(75, 75))

		ceil := container.New(layout.NewStackLayout(), box)
		grid.Add(ceil)
	}
}

func getColor(value int) color.Color {
	res := state.WallColor

	switch value {
	case 1:
		res = state.ResourceColor
	case 2:
		res = state.PitColor
	case 3:
		res = state.NothingColor
	case 4:
		res = state.RobotColor
	case 5:
		res = state.EndColor
	}

	return res
}

func getCeilCaption(caption string) *fyne.Container {

	box := canvas.NewRectangle(color.Transparent)
	box.SetMinSize(fyne.NewSize(75, 75))

	text := canvas.NewText(caption, color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 16
	text.TextStyle.Bold = true

	return container.New(
		layout.NewStackLayout(),
		box,
		text,
	)
}

func getCeilRectangle(c color.Color) *fyne.Container {
	box := canvas.NewRectangle(c)
	box.SetMinSize(fyne.NewSize(65, 65))

	return container.New(layout.NewStackLayout(), box)
}

var resourceAboutpart = &fyne.StaticResource{
	StaticName: "About.md",
	StaticContent: []byte(
		"## HOW TO USE TRANSLATOR.\n\n You need collect all resources or go to the end labirinth\n\nOn field must be 1 robot.\n\n***\n\n### You can use this commands:\n\n#### Move\n\n**Up**: Moves the robot up to 1 ceil\n\n**Down**: Moves the robot down to 1 ceil\n\n**Right**: Moves the robot right to 1 ceil\n\n**Left**: Moves the robot left to 1 ceil\n\n***\n\n#### Peek\n\n**PeekL**: Peek one ceil left\n\n**PeekR**: Peek one ceil right\n\n**PeekD**: Peek one ceil down\n\n**PeekU**: Peek one ceil up\n\n***\n\n#### Action\n\nFill - fill up the ceil\n\n***\n\n### Description of the colors:\n"),
}

func StartRender(app *fyne.App, st *state.AppState, window fyne.Window, next fyne.Window) {
	startContainer := container.New(layout.NewVBoxLayout())

	rowsEntry := widget.NewEntry()
	rowsEntry.PlaceHolder = "..."
	rowsLabel := widget.NewLabel("Rows")

	columnsEntry := widget.NewEntry()
	columnsEntry.PlaceHolder = "..."
	columnsLabel := widget.NewLabel("Columns")

	sizeContainer := container.New(
		layout.NewGridLayoutWithColumns(2),
		rowsLabel,
		columnsLabel,
		rowsEntry,
		columnsEntry,
	)

	filepath := widget.NewLabel("Empty")

	openFileBtn := widget.NewButton(
		"Choose file",
		func() {
			oldw := window.Canvas().Size().Width
			oldh := window.Canvas().Size().Height
			window.Resize(fyne.NewSize(640, 640))
			dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
				if uc != nil {
					filepath.SetText(uc.URI().Path())
				}
				window.Resize(fyne.NewSize(oldw, oldh))
			},
				window,
			)

		},
	)

	inputFieldGrid := container.New(
		layout.NewGridLayoutWithColumns(st.Columns),
	)

	inputFieldGrid.Hide()

	checkSizeBtn := widget.NewButton(
		"Check size",
		func() {
			rows, _ := strconv.ParseInt(rowsEntry.Text, 10, 64)
			columns, _ := strconv.ParseInt(columnsEntry.Text, 10, 64)
			st.InputField = make([]*widget.Select, 0)
			if (rows < state.MIN_SIZE) || (rows > state.MAX_SIZE) || (columns < state.MIN_SIZE) || (columns > state.MAX_SIZE) {
				rowsEntry.SetText("")
				columnsEntry.SetText("")
				dialog.ShowError(
					fmt.Errorf("ERROR: %s", "Incorrect input size"),
					window,
				)
			} else {
				st.Columns = int(columns)
				st.Rows = int(rows)
				inputFieldGrid.RemoveAll()
				inputFieldGrid.Add(inputGridRender(st))
				inputFieldGrid.Show()
			}
		},
	)

	inputField := container.New(layout.NewVBoxLayout(), sizeContainer, checkSizeBtn, inputFieldGrid)

	openFile := container.New(layout.NewVBoxLayout(), filepath, openFileBtn)

	inputTypeRadio := widget.NewRadioGroup([]string{"File", "Input"}, func(s string) {
		if s == "File" {
			openFile.Show()
			inputField.Hide()
			window.Resize(fyne.NewSize(320, 160))
		} else {
			openFile.Hide()
			inputField.Show()
		}
	})
	inputTypeRadio.Selected = "File"
	inputField.Hide()

	confirmBtn := widget.NewButton("Confirm", func() {
		if inputTypeRadio.Selected == "Input" {
			rows, _ := strconv.ParseInt(rowsEntry.Text, 10, 64)
			columns, _ := strconv.ParseInt(columnsEntry.Text, 10, 64)
			if (rows < state.MIN_SIZE) ||
				(rows > state.MAX_SIZE) ||
				(columns < state.MIN_SIZE) ||
				(columns > state.MAX_SIZE) {
				dialog.ShowError(
					fmt.Errorf("ERROR: %s", "Incorrect input size"),
					window,
				)
				rowsEntry.SetText("")
				columnsEntry.SetText("")
				st.InputField = make([]*widget.Select, 0)
			} else if !checkField(st) {
				dialog.ShowError(
					fmt.Errorf("ERROR: %s", "Incorrect field"),
					window,
				)
			} else {
				st.Columns = int(columns)
				st.Rows = int(rows)
				window.Hide()
				MainRender(app, st, next, window)
				next.Show()
			}
		} else {
			if _, err := os.Stat(filepath.Text); err != nil {
				dialog.ShowError(
					fmt.Errorf("ERROR: %w", err),
					window,
				)
			} else if !readFieldFromFile(st, filepath.Text) {
				dialog.ShowError(
					fmt.Errorf("ERROR: %s", "Bad format for preset"),
					window,
				)
			} else {
				window.Hide()
				MainRender(app, st, next, window)
				next.Show()
			}
		}

	})

	icon, err := fyne.LoadResourceFromPath("../../assets/icon.png")
	if err != nil {
		panic(err)
	}

	exampleRect := canvas.NewRectangle(color.Transparent)
	exampleRect.SetMinSize(fyne.NewSize(75, 75))
	inputIcon := container.New(layout.NewStackLayout(), exampleRect, widget.NewIcon(icon))

	inputHeader := container.New(
		layout.NewBorderLayout(nil, nil, inputTypeRadio, inputIcon),
		inputTypeRadio,
		inputIcon,
	)

	startContainer.Add(inputHeader)
	startContainer.Add(inputField)
	startContainer.Add(openFile)
	startContainer.Add(confirmBtn)

	window.SetContent(startContainer)
}

func checkField(st *state.AppState) bool {
	wasRobot := false
	wasEnd := false
	st.CellularField = make([]state.Cell, 0)
	for _, sel := range st.InputField {

		if (sel.Selected == "ROBOT" && wasRobot) || (sel.Selected == "END" && wasEnd) {
			return false
		}

		switch sel.Selected {
		case "NONE":
			st.CellularField = append(st.CellularField, state.Nothing)
		case "WALL":
			st.CellularField = append(st.CellularField, state.Wall)
		case "RESOURCE":
			st.CellularField = append(st.CellularField, state.Resource)
		case "PIT":
			st.CellularField = append(st.CellularField, state.Pit)
		case "ROBOT":
			st.CellularField = append(st.CellularField, state.Robot)
			wasRobot = true
		case "END":
			st.CellularField = append(st.CellularField, state.End)
			wasEnd = true
		default:
			return false
		}
	}

	return wasRobot && wasEnd
}

func inputGridRender(st *state.AppState) *fyne.Container {
	grid := container.New(layout.NewGridLayoutWithColumns(st.Columns))
	for i := 0; i < st.Columns*st.Rows; i++ {
		sel := widget.NewSelect([]string{"NONE", "WALL", "RESOURCE", "PIT", "ROBOT", "END"}, nil)
		st.InputField = append(st.InputField, sel)
		grid.Add(sel)
	}
	return grid
}

func readFieldFromFile(st *state.AppState, path string) bool {

	file, err := os.Open(path)
	if err != nil {
		return false
	}

	defer file.Close()
	st.CellularField = make([]state.Cell, 0)

	sc := bufio.NewScanner(file)
	buffer := make([]string, 0)
	for sc.Scan() {
		buffer = append(buffer, sc.Text())
	}

	tmp := strings.Fields(buffer[0])
	tmpColumns, err := strconv.ParseInt(tmp[0], 10, 0)
	if err != nil {
		return false
	}
	tmpRows, err := strconv.ParseInt(tmp[1], 10, 0)
	if err != nil {
		return false
	}

	if tmpColumns > state.MAX_SIZE || tmpColumns < state.MIN_SIZE || tmpRows > state.MAX_SIZE || tmpRows < state.MIN_SIZE {
		return false
	}

	st.Columns = int(tmpColumns)
	st.Rows = int(tmpRows)

	wasRobot := false
	wasEnd := false

	for rows := 1; rows < len(buffer); rows++ {
		columns := strings.Fields(buffer[rows])
		for col := 0; col < st.Columns; col++ {
			tmp, err := strconv.ParseInt(columns[col], 10, 64)
			if err != nil || tmp < 0 || tmp > state.COUNT {
				return false
			}

			if tmp == int64(state.Robot) {
				if !wasRobot {
					wasRobot = true
				} else {
					return false
				}
			}

			if tmp == int64(state.End) {
				if !wasEnd {
					wasEnd = true
				} else {
					return false
				}
			}

			st.CellularField = append(st.CellularField, state.Cell(tmp))
		}
	}

	return wasRobot && wasEnd
}
