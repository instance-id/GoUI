package view

import (
	"fmt"
	. "github.com/instance-id/GoUI/components"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
)

var pendingMain = struct {
	token     bool
	cmdPrefix bool
	logLevel  bool
}{
	token:     false,
	cmdPrefix: false,
	logLevel:  false,
}

func CreateViewMainSettings() /*(*ui.Frame, *ui.EditField)*/ {

	var tmpDiscordToken = Cntnrs.Dac.System.Token
	var tmpCommandPrefix = Cntnrs.Dac.System.CommandPrefix
	Log.CurrentLogLevel = Cntnrs.Dac.System.FileLogLevel

	// --- Main Settings Frame -------------------------------------------
	FrmMainSettings = ui.CreateFrame(FrameContent, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	FrmMainSettings.SetPack(ui.Vertical)
	FrmMainSettings.SetBackColor(236)

	// --- Main Settings Content -----------------------------------------
	settingsFrame := ui.CreateFrame(FrmMainSettings, 130, ui.AutoSize, ui.BorderThin, ui.AutoSize)
	settingsFrame.SetPaddings(2, 2)
	settingsFrame.SetTitle(TxtMainSettings)
	settingsFrame.SetPack(ui.Vertical)
	settingsFrame.SetBackColor(236)

	// --- Discord Token -------------------------------------------------
	tokenFrame := NewFramedInput(settingsFrame, TxtDiscordToken, nil)
	tokenFrame.SetPaddings(2, 2)
	tokenFrame.SetBackColor(236)
	tokenEdit := ui.CreateEditField(tokenFrame, ui.AutoSize, tmpDiscordToken, ui.Fixed)
	tokenEdit.OnChange(func(event ui.Event) {
		tmpDiscordToken = tokenEdit.Title()
		if tmpDiscordToken != Cntnrs.Dac.System.Token {
			pendingMain.token = true
			SavePendingMainSettings()
		} else {
			pendingMain.token = false
			SavePendingMainSettings()
		}
	})
	ui.CreateLabel(tokenFrame, ui.AutoSize, ui.AutoSize, TxtDiscordTokenDesc, ui.Fixed)

	// --- Command Prefix ------------------------------------------------
	cmdPrefixFrame := NewFramedInput(settingsFrame, TxtCmdPrefix, nil)
	cmdPrefixFrame.SetPaddings(2, 2)
	cmdPrefixFrame.SetBackColor(236)
	prefixResult := ui.CreateEditField(cmdPrefixFrame, ui.AutoSize, tmpCommandPrefix, ui.Fixed)
	prefixResult.OnChange(func(event ui.Event) {
		tmpCommandPrefix = prefixResult.Title()
		if tmpCommandPrefix != Cntnrs.Dac.System.CommandPrefix {
			pendingMain.cmdPrefix = true
			SavePendingMainSettings()
		} else {
			pendingMain.cmdPrefix = false
			SavePendingMainSettings()
		}
	})
	ui.CreateLabel(cmdPrefixFrame, ui.AutoSize, ui.AutoSize, TxtCmdPrefixDesc, ui.Fixed)

	//// --- Require Email -------------------------------------------------
	//requireEmail := NewFramedInput(settingsFrame, TxtRequireEmail, nil)
	//requireEmail.SetPaddings(2, 2)
	//requireEmail.SetBackColor(236)
	//ui.CreateLabel(requireEmail, ui.AutoSize, ui.AutoSize, TxtRequireEmailDesc, ui.Fixed)
	//ui.CreateCheckBox(requireEmail, 10, " Check for Yes, unchecked for No ", ui.Fixed)

	// --- Select Log Level ----------------------------------------------
	var logParams = FramedInputParams{Orientation: ui.Vertical, Width: 25, Height: 0, Scale: ui.Fixed, Border: ui.BorderThin, PadX: 1, PadY: 1}
	logLevel := NewFramedInput(settingsFrame, TxtLogLevel, &logParams)
	logLevel.SetBackColor(236)
	BtnLogLevel = ui.CreateButton(logLevel, 25, ui.AutoSize, fmt.Sprintf("%s %s", TxtLogLevelBtn, Log.LogLevel[Log.CurrentLogLevel]), ui.Fixed)
	BtnLogLevel.SetAlign(ui.AlignLeft)
	BtnLogLevel.SetShadowType(ui.ShadowHalf)
	BtnLogLevel.OnClick(func(ev ui.Event) {
		BtnLogLevel.SetEnabled(false)
		Log.CurrentLogLevel = SelectLogLevel(BtnLogLevel, Log.CurrentLogLevel)
	})
	ui.CreateLabel(logLevel, ui.AutoSize, ui.AutoSize, TxtLogLevelDesc, ui.Fixed)

	// --- Save Settings ------------------------------------------------
	btnFrame := ui.CreateFrame(settingsFrame, 10, 1, ui.BorderNone, ui.Fixed)
	btnFrame.SetPaddings(2, 2)
	btnFrame.SetBackColor(236)

	var params = FramedInputParams{Orientation: ui.Vertical, Width: 10, Height: 4, Scale: ui.Fixed}
	saveSettings := NewFramedInput(btnFrame, TxtSaveDesc, &params)
	saveSettings.SetBackColor(236)
	BtnMainSettingsSave = ui.CreateButton(saveSettings, 25, ui.AutoSize, TxtSaveBtn, ui.Fixed)
	BtnMainSettingsSave.SetAlign(ui.AlignLeft)
	BtnMainSettingsSave.SetShadowType(ui.ShadowHalf)
	BtnMainSettingsSave.OnClick(func(ev ui.Event) {
		Cntnrs.Dac.System.Token = tmpDiscordToken
		Cntnrs.Dac.System.CommandPrefix = tmpCommandPrefix
		Cntnrs.Dac.System.FileLogLevel = Log.CurrentLogLevel
		_, err := Cntnrs.Wtr.SetConfig()
		if err != nil {
			ui.CreateAlertDialog(ErrCouldNotSaveCfg, fmt.Sprintf("Error Could not save config: %s", err), TxtCloseBtn)
		}
		BtnMainSettingsSave.SetTitle(TxtSaveBtn)
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	})

}

func SavePendingMainSettings() {
	if pendingMain.token || pendingMain.cmdPrefix || pendingMain.logLevel == true {
		BtnMainSettingsSave.SetTitle(TxtSavePendingBtn)
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	} else {
		BtnMainSettingsSave.SetTitle(TxtSaveBtn)
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	}
}
