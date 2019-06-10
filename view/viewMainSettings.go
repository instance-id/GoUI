package view

import (
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
)

func CreateViewMainSettings() {

	var tmpDiscordToken = DiscordToken
	var tmpCommandPrefix = CommandPrefix

	// --- Main Settings Frame -------------------------------------------
	FrmMainSettings = ui.CreateFrame(FrameContent, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	FrmMainSettings.SetPack(ui.Vertical)

	// --- Main Settings Content -----------------------------------------
	settingsFrame := ui.CreateFrame(FrmMainSettings, 130, ui.AutoSize, ui.BorderThin, ui.AutoSize)
	settingsFrame.SetPaddings(2, 2)
	settingsFrame.SetTitle(TxtMainSettings)
	settingsFrame.SetPack(ui.Vertical)

	// --- Discord Token -------------------------------------------------
	tokenFrame := NewFramedInput(settingsFrame, TxtDiscordToken, nil)
	tokenFrame.SetPaddings(2, 2)
	ui.CreateEditField(tokenFrame, ui.AutoSize, tmpDiscordToken, ui.Fixed)
	ui.CreateLabel(tokenFrame, ui.AutoSize, ui.AutoSize, TxtDiscordTokenDesc, ui.Fixed)

	// --- Command Prefix ------------------------------------------------
	cmdPrefixFrame := NewFramedInput(settingsFrame, TxtCmdPrefix, nil)
	cmdPrefixFrame.SetPaddings(2, 2)
	ui.CreateEditField(cmdPrefixFrame, ui.AutoSize, tmpCommandPrefix, ui.Fixed)
	ui.CreateLabel(cmdPrefixFrame, ui.AutoSize, ui.AutoSize, TxtCmdPrefixDesc, ui.Fixed)

	// --- Require Email -------------------------------------------------
	requireEmail := NewFramedInput(settingsFrame, TxtRequireEmail, nil)
	requireEmail.SetPaddings(2, 2)
	ui.CreateLabel(requireEmail, ui.AutoSize, ui.AutoSize, TxtRequireEmailDesc, ui.Fixed)
	ui.CreateCheckBox(requireEmail, 10, " Check for Yes, unchecked for No ", ui.Fixed)

	// --- Select Log Level ----------------------------------------------
	var logParams = FramedInputParams{Orientation: ui.Vertical, Width: 25, Height: 0, Scale: ui.Fixed, Border: ui.BorderThin, PadX: 1, PadY: 1}
	logLevel := NewFramedInput(settingsFrame, TxtLogLevel, &logParams)
	var logLevelParams = FramedInputParams{Orientation: ui.Vertical, Width: 25, Height: 0, Scale: ui.Fixed, Border: ui.BorderThin, PadX: 0, PadY: 0}
	logLevelTxtFrame := NewFramedInput(logLevel, "", &logLevelParams)
	BtnLogLevel = ui.CreateButton_NoShadow(logLevelTxtFrame, 25, ui.AutoSize, TxtLogLevelBtn, ui.Fixed)
	BtnLogLevel.SetAlign(ui.AlignLeft)
	BtnLogLevel.OnClick(func(ev ui.Event) {
		BtnLogLevel.SetEnabled(false)
		SelectLogLevel(BtnLogLevel)
	})
	ui.CreateLabel(logLevel, ui.AutoSize, ui.AutoSize, TxtLogLevelDesc, ui.Fixed)

	// --- Save Settings ------------------------------------------------
	btnFrame := ui.CreateFrame(settingsFrame, 10, 1, ui.BorderNone, ui.Fixed)
	btnFrame.SetPaddings(2, 2)
	var params = FramedInputParams{Orientation: ui.Vertical, Width: 10, Height: 4, Scale: ui.Fixed}
	saveSettings := NewFramedInput(btnFrame, TxtSaveDesc, &params)
	BtnMainSettingsSave = ui.CreateButton_NoShadow(saveSettings, 25, ui.AutoSize, TxtSave, ui.Fixed)
	BtnMainSettings.SetSize(10, ui.AutoSize)
	BtnMainSettingsSave.SetAlign(ui.AlignLeft)
	BtnMainSettingsSave.OnClick(func(ev ui.Event) {
		DiscordToken = tmpDiscordToken
		CommandPrefix = tmpCommandPrefix
	})
	FrmMainSettings.SetVisible(false)
	BtnLogLevel.SetActive(false)
	BtnMainSettingsSave.SetActive(false)
}
