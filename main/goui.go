package main

import (
	"github.com/chzyer/readline"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
	"github.com/instance-id/GoUI/view"
	ui "github.com/instance-id/clui"
	term "github.com/nsf/termbox-go"
	"os"
)

var (
	text, background term.Attribute
	lightness        int
	helpVisible      bool
)

func InitData() {
	Log.LogLevel = []string{"INFO", "DEBUG", "WARNING", "ERROR"}
	Log.DefaultLogLevel = "INFO"
	Log.CurrentLogLevel = Log.DefaultLogLevel
	DiscordToken = "123123123SDFSDFSDFSDF1234123123"
	CommandPrefix = "!cmd "

	//if Asset == nil {
	//	Asset = &AssetContainer{AD: []AssetDetails{
	//		{
	//			"",
	//			"",
	//			"",
	//			"",
	//			"",
	//			"",
	//			"",
	//		},
	//	}}
	//}

	// InitTmpAssetData()
}

func MainInitialSettings() {
	FrmMainSettings.SetVisible(true)
	FrmDiscordSettings.SetActive(false)
	FrmDatabaseSettings.SetActive(false)
	FrmPlugins.SetActive(false)

	BtnRunVerifier.SetActive(false)
	BtnMainSettings.SetActive(false)
	BtnDiscordSettings.SetActive(false)
	BtnDatabaseSettings.SetActive(false)
	BtnPlugins.SetActive(false)
	//BtnTheme.SetActive(false)
	BtnLogs.SetActive(false)
	BtnQuit.SetActive(false)
	InitData()
}

func createView() {

	// --- Main Window ---------------------------------------------------
	WindowMain = ui.AddWindow(0, 0, 10, 7, TxtApplication)
	WindowMain.SetPack(ui.Horizontal)
	WindowMain.SetBackColor(236)

	// --- Main Menu Frame -----------------------------------------------
	view.CreateViewMenu()

	// --- Content Frame -------------------------------------------------
	view.CreateViewContent()

	// --- Settings Frames -----------------------------------------------
	tokenEdit := view.CreateViewMainSettings()
	view.CreateViewDiscordSettings()
	view.CreateViewDatabaseSettings()
	view.CreateViewPlugins()

	MainInitialSettings()

	tokenEdit.SetActive(true)
	tokenEdit.SetEnabled(true)
	tokenEdit.SetTabStop(true)

	ui.MainLoop()
}

func mainLoop() {
	// Bool - Use 256 color or default?
	ui.InitLibrary()
	defer ui.DeinitLibrary()

	// --- Newly added ---------------------
	term.SetOutputMode(term.Output256)

	ui.SetThemePath("themes")
	ui.SetCurrentTheme("verifier")

	createView()

	ui.MainLoop()
}

func main() {
	mainLoop()
}

func TermNative(c int) term.Attribute {
	return term.Attribute(c + 1)
}

type stderr struct{}

func (s *stderr) Write(b []byte) (int, error) {
	if len(b) == 1 && b[0] == 7 {
		return 0, nil
	}
	return os.Stderr.Write(b)
}

func (s *stderr) Close() error {
	return os.Stderr.Close()
}

func init() {
	readline.Stdout = &stderr{}
}
