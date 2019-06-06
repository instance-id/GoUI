package view

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
)

func CreateViewAssetSettings() {

	var tmpGuidId =

	// --- Discord Settings Frame ------------------------------------------
	FrmDiscordSettings = ui.CreateFrame(WindowMain, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.AutoSize)
	FrmDiscordSettings.SetTitle("FrameTop")
	FrmDiscordSettings.SetPack(ui.Horizontal)

	// --- Discord Settings Content ----------------------------------------
	settingsFrame := ui.CreateFrame(FrmDiscordSettings, ui.AutoSize, ui.AutoSize, ui.BorderThin, ui.AutoSize)
	settingsFrame.SetPaddings(2, 2)
	settingsFrame.SetTitle(TxtMainSettings)
	settingsFrame.SetPack(ui.Vertical)

	// --- Discord Token -------------------------------------------------
	tokenFrame := NewFramedInput(settingsFrame, TxtDiscordToken, nil)
	ui.CreateEditField(tokenFrame, 70, tmpDiscordToken, ui.Fixed)
	ui.CreateLabel(tokenFrame, ui.AutoSize, ui.AutoSize, TxtDiscordTokenDesc, ui.Fixed)

	// --- Command Prefix ------------------------------------------------
	cmdPrefixFrame := NewFramedInput(settingsFrame, TxtCmdPrefix, nil)
	ui.CreateEditField(cmdPrefixFrame, 10, tmpCommandPrefix, ui.Fixed)
	ui.CreateLabel(cmdPrefixFrame, ui.AutoSize, ui.AutoSize, TxtCmdPrefixDesc, ui.Fixed)

	// --- Require Email -------------------------------------------------
	requireEmail := NewFramedInput(settingsFrame, TxtRequireEmail, nil)
	ui.CreateLabel(requireEmail, ui.AutoSize, ui.AutoSize, TxtRequireEmailDesc, ui.Fixed)
	ui.CreateCheckBox(requireEmail, 10, " Check for Yes, unchecked for No ", ui.Fixed)

	// --- Select Log Level ----------------------------------------------
	logLevel := NewFramedInput(settingsFrame, TxtLogLevel, nil)
	BtnLogLevel := ui.CreateButton(logLevel, ui.AutoSize, ui.AutoSize, TxtLogLevel, 1)
	BtnLogLevel.OnClick(func(ev ui.Event) {
		BtnLogLevel.SetEnabled(false)
		SelectLogLevel(BtnLogLevel)
	})
	ui.CreateLabel(logLevel, ui.AutoSize, ui.AutoSize, TxtLogLevelDesc, ui.Fixed)

	// --- Save Settings ------------------------------------------------
	var params = FramedInputParams{Orientation: ui.Vertical, Width: 4, Height: 4}
	saveSettings := NewFramedInput(settingsFrame, TxtSaveDesc, &params)
	BtnSave = ui.CreateButton(saveSettings, 4, ui.AutoSize, TxtSave, ui.Fixed)
	BtnSave.OnClick(func(ev ui.Event) {
		DiscordToken = tmpDiscordToken
		CommandPrefix = tmpCommandPrefix
	})
	//ui.CreateLabel(saveSettings, ui.AutoSize, ui.AutoSize, TxtSaveDesc, ui.Fixed)

}
