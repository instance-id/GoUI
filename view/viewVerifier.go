package view

import (
	"fmt"
	"os"
	"runtime"

	"github.com/hpcloud/tail"

	"github.com/instance-id/GoUI/rpcclient"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
)

var LogRunning = false
var LogChan chan bool

type ServerStatus struct {
	VerifierRunning bool
	RPCRunning      bool
}

func CreateViewVerifier(status *ServerStatus) /*(*ui.Frame, *ui.EditField)*/ {
	LogViewer = new(LogDialog)
	var sts bool

	// --- Verifier Controls Frame ---------------------------------------
	FrmVerifier = ui.CreateFrame(FrameContent, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	FrmVerifier.SetPack(ui.Vertical)
	FrmVerifier.SetBackColor(236)

	// --- Verifier Controls Content -------------------------------------
	settingsFrame := ui.CreateFrame(FrmVerifier, 130, ui.AutoSize, ui.BorderThin, ui.AutoSize)
	settingsFrame.SetPaddings(2, 1)
	settingsFrame.SetTitle(TxtVerifier)
	settingsFrame.SetPack(ui.Vertical)
	settingsFrame.SetBackColor(236)

	// --- Verifier Controls Button Frame --------------------------------
	controlFrame := ui.CreateFrame(settingsFrame, 130, ui.AutoSize, ui.BorderThin, ui.Fixed)
	controlFrame.SetBackColor(236)
	controlFrame.SetPaddings(0, 0)

	// --- Start ---------------------------------------------------------
	startFrame := ui.CreateFrame(controlFrame, 40, 6, ui.BorderNone, ui.Fixed)
	startFrame.SetPaddings(1, 1)
	startFrame.SetBackColor(236)
	BtnVerifierStart = ui.CreateButton(startFrame, 40, ui.AutoSize, fmt.Sprintf("%s", TxtVerifierStartBtn), ui.Fixed)
	BtnVerifierStart.SetAlign(ui.AlignLeft)
	BtnVerifierStart.SetShadowType(ui.ShadowHalf)
	if status.VerifierRunning {
		BtnVerifierStart.SetEnabled(false)
	}
	BtnVerifierStart.OnClick(func(ev ui.Event) {
		if !LogRunning {
			go LoadVerifierLogs()
			runtime.Gosched()
		}
		go func(sts bool) bool { sts = rpcclient.StartServer(); return sts }(sts)
		runtime.Gosched()
		if sts {
			status.VerifierRunning = sts
			BtnVerifierStart.SetEnabled(false)
		}

	})

	// --- Restart -------------------------------------------------------
	restartFrame := ui.CreateFrame(controlFrame, 40, 6, ui.BorderNone, ui.Fixed)
	restartFrame.SetPaddings(1, 1)
	restartFrame.SetBackColor(236)
	BtnVerifierRestart = ui.CreateButton(restartFrame, 40, ui.AutoSize, fmt.Sprintf("%s", TxtVerifierRestartBtn), ui.Fixed)
	BtnVerifierRestart.SetAlign(ui.AlignLeft)
	BtnVerifierRestart.SetShadowType(ui.ShadowHalf)
	BtnVerifierRestart.OnClick(func(ev ui.Event) {
		go rpcclient.RestartServer()
		runtime.Gosched()
	})

	// --- Stop ----------------------------------------------------------
	stopFrame := ui.CreateFrame(controlFrame, 40, 6, ui.BorderNone, ui.Fixed)
	stopFrame.SetPaddings(1, 1)
	stopFrame.SetBackColor(236)
	BtnVerifierStop = ui.CreateButton(stopFrame, 40, ui.AutoSize, fmt.Sprintf("%s", TxtVerifierStopBtn), ui.Fixed)
	BtnVerifierStop.SetAlign(ui.AlignLeft)
	BtnVerifierStop.SetShadowType(ui.ShadowHalf)
	BtnVerifierStop.OnClick(func(ev ui.Event) {
		go rpcclient.StopServer()
		runtime.Gosched()
		if LogRunning {
			LogRunning = false
		}
	})

	// --- Autoscroll checkbox ------------------------------------------------
	logFrameParams := FramedInputParams{Orientation: ui.Vertical, Width: 10, Height: 0, Scale: ui.Fixed, Border: ui.BorderThin, PadX: 1, PadY: 1}
	LogViewer.Frame = NewFramedInput(settingsFrame, TxtLogLevel, &logFrameParams)
	LogViewer.Frame.SetBackColor(236)

	LogViewer.Log = ui.CreateTextView(LogViewer.Frame, 130, 25, 1)
	ui.ActivateControl(LogViewer.Frame, LogViewer.Log)
	LogViewer.Log.LoadFileMD(LogLocation)
	LogViewer.Log.SetAutoScroll(true)

	buttonFrame := ui.CreateFrame(LogViewer.Frame, 130, 6, ui.BorderNone, ui.AutoSize)

	autoScroll := ui.CreateCheckBox(buttonFrame, ui.AutoSize, TxtAutoScrollChk, ui.AutoSize)
	autoScroll.SetState(1)
	autoScroll.OnChange(func(i int) {
		LogViewer.Log.SetAutoScroll(func() bool {
			var result bool
			switch i {
			case 0:
				result = false
			case 1:
				result = true
			}
			return result
		}())
	})

	// --- Clear Logs ----------------------------------------------------
	clearLogFrame := ui.CreateFrame(buttonFrame, 30, 6, ui.BorderNone, ui.AutoSize)
	clearLogFrame.SetPaddings(1, 1)
	clearLogFrame.SetBackColor(236)
	BtnVerifierStart = ui.CreateButton(clearLogFrame, 40, ui.AutoSize, fmt.Sprintf("%s", TxtClearBtn), ui.AutoSize)
	BtnVerifierStart.SetAlign(ui.AlignLeft)
	BtnVerifierStart.SetShadowType(ui.ShadowHalf)
	BtnVerifierStart.OnClick(func(ev ui.Event) {
		_ = os.Remove("./logs/verifier.log")
		_, _ = os.Create("./logs/verifier.log")
	})

	LogViewer.Log.SetBackColor(238)
	LogViewer.Log.SetTextColor(250)

}

func LoadVerifierLogs() error {
	//LogChan = make(chan bool)
	defer Tails.Done()

	var _ error
	Tails, _ = tail.TailFile("./logs/verifier.log", tail.Config{Follow: true, ReOpen: true, Logger: tail.DiscardingLogger})
	LogRunning = true

	for LogRunning {
		if !LogRunning {
			_ = Tails.Stop()
			fmt.Printf("Stopping log viewer")

			break
		}
		for line := range Tails.Lines {
			LogViewer.Log.AddText([]string{line.Text})
			if []string{line.Text} != nil {
				ui.PutEvent(ui.Event{Type: ui.EventRedraw})
			}

		}
	}
	LogRunning = false
	return nil

	//for {
	//	select {
	//	case line, ok := <-Tails.Lines:
	//		if !ok {
	//			_ = Tails.Wait()
	//			break
	//		}
	//		if []string{line.Text} != nil {
	//			LogViewer.Log.AddText([]string{line.Text})
	//			ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	//		}
	//	case <-LogChan:
	//		LogRunning = false
	//		return Tails.Stop()
	//	}
	//}
}
