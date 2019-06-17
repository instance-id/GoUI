package elements

import (
	"github.com/hpcloud/tail"
	ui "github.com/instance-id/clui"
)

var (
	Log LogLevels
)

var LogViewer *LogDialog
var Tails *tail.Tail

type DataDiscord struct {
	GuildId  string
	BotUsers []string
	Roles    map[string]string
}

type EditableDialog struct {
	View       *ui.Window
	Frame      *ui.Frame
	result     int
	value      int
	edtResult  string
	edit       *ui.EditField
	editDialog *ui.SelectDialog
	confirm    *ui.ConfirmationDialog
	onClose    func()
}

type EditDialog struct {
	View       *ui.Window
	Frame      *ui.Frame
	result     int
	value      int
	edtResult  string
	list       *ui.ListBox
	edit       *ui.EditField
	editDialog *ui.SelectDialog
	confirm    *ui.ConfirmationDialog
	onClose    func()
}

type LogDialog struct {
	View       *ui.Window
	Frame      *ui.Frame
	result     int
	value      int
	edtResult  string
	Log        *ui.TextView
	edit       *ui.EditField
	editDialog *ui.SelectDialog
	confirm    *ui.ConfirmationDialog
	onClose    func()
}

type LogLevels struct {
	LogLevel        []string
	CurrentLogLevel int
	DefaultLogLevel int
}

type TableDialog struct {
	View       *ui.Window
	Frame      *ui.Frame
	result     int
	value      int
	edtResult  string
	table      *ui.TableView
	edit       *ui.EditField
	editDialog *ui.SelectDialog
	confirm    *ui.ConfirmationDialog
	onClose    func()
}

type TableEdit struct {
	Row    int
	Col    int
	NewVal string
	OldVal string
}

const (
	// DialogClosed - a user clicked close button on the dialog title
	DialogClosed = -1
	// DialogAlive - a user does not close the dialog yet, exit code is unavailable
	DialogAlive = 0
	// DialogButton1 - a user clicked the first button in the dialog (by default, it is 'Yes' or 'OK')
	DialogButton1 = 1
	// DialogButton2 - a user clicked the second button in the dialog
	DialogButton2 = 2
	// DialogButton3 - a user clicked the third button in the dialog
	DialogButton3 = 3
	// DialogButton4 - a user clicked the fourth button in the dialog
	DialogButton4 = 4
)
