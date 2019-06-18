package main

import (
	"github.com/chzyer/readline"
	. "github.com/instance-id/GoUI/components"
	"github.com/instance-id/GoUI/rpcclient"

	//. "github.com/instance-id/GoUI/rpcclient"
	"os"

	. "github.com/instance-id/GoUI/text"
	"github.com/instance-id/GoUI/view"
	ui "github.com/instance-id/clui"
	term "github.com/nsf/termbox-go"
)

func InitData() {
	view.Log.LogLevel = []string{"INFO", "DEBUG", "WARNING", "ERROR"}
	view.Log.DefaultLogLevel = 0

}

func MainInitialSettings() {
	view.FrmMainSettings.SetVisible(true)
	view.FrmDiscordSettings.SetActive(false)
	view.FrmDatabaseSettings.SetActive(false)
	view.FrmPlugins.SetActive(false)

	view.BtnRunVerifier.SetActive(false)
	view.BtnMainSettings.SetActive(false)
	view.BtnDiscordSettings.SetActive(false)
	view.BtnDatabaseSettings.SetActive(false)
	view.BtnPlugins.SetActive(false)
	view.BtnLogs.SetActive(false)
	view.BtnQuit.SetActive(false)
}

func createView() {
	InitData()
	// --- Main Window ---------------------------------------------------
	view.WindowMain = ui.AddWindow(0, 0, 10, 7, TxtApplication)
	view.WindowMain.SetPack(ui.Horizontal)
	view.WindowMain.SetBackColor(236)

	// --- Main Menu Frame -----------------------------------------------
	view.CreateViewMenu()

	// --- Content Frame -------------------------------------------------
	view.CreateViewContent()

	// --- Settings Frames -----------------------------------------------
	/*tokenFrame, tokenEdit := */
	view.CreateViewVerifier()
	view.CreateViewMainSettings()
	view.CreateViewDiscordSettings()
	view.CreateViewDatabaseSettings()
	view.CreateViewPlugins()

	MainInitialSettings()

	//ui.ActivateControl(tokenFrame, tokenEdit)
	//tokenEdit.SetActive(true)
	//tokenEdit.SetEnabled(true)
	//tokenEdit.SetTabStop(true)
	view.CommandMainSettings()
	ui.MainLoop()
}

func mainLoop() {

	ui.InitLibrary()
	defer ui.DeinitLibrary()

	// --- Newly added ---------------------
	term.SetOutputMode(term.Output256)

	ui.SetThemePath("themes")

	// --- Changing theme won't do much ----
	// --- Many values are hard coded ------
	ui.SetCurrentTheme("verifier")

	createView()

	ui.MainLoop()
}

func main() {
	phrase, key := CmdInitialize()
	rpcclient.GetKey(phrase, key)

	mainLoop()
}

func termNative(c int) term.Attribute {
	return term.Attribute(c + 1)
}

type stderr struct{}

func init() {
	readline.Stdout = &stderr{}
}

func (s *stderr) Write(b []byte) (int, error) {
	if len(b) == 1 && b[0] == 7 {
		return 0, nil
	}
	return os.Stderr.Write(b)
}

func (s *stderr) Close() error {
	return os.Stderr.Close()
}
