package view

import (
	"fmt"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
)

func CreateViewDiscordSettings() {

	var tmpGuidId = DiscordData.GuildId
	var tmpBotUsers = DiscordData.BotUsers
	var tmpRoles = DiscordData.Roles

	// --- Discord Settings Frame ------------------------------------------
	FrmDiscordSettings = ui.CreateFrame(FrameContent, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	FrmDiscordSettings.SetTitle("FrameTop")
	FrmDiscordSettings.SetPack(ui.Horizontal)

	// --- Discord Settings Content ----------------------------------------
	settingsFrame := ui.CreateFrame(FrmDiscordSettings, 130, ui.AutoSize, ui.BorderThin, ui.Fixed)
	settingsFrame.SetPaddings(2, 2)
	settingsFrame.SetTitle(TxtDiscordSettings)
	settingsFrame.SetPack(ui.Vertical)

	// --- GuildId -------------------------------------------------------
	//var gIdparams = FramedInputParams{Border: ui.BorderNone}
	guildIdFrame := NewFramedInput(settingsFrame, TxtGuildId, nil)
	guildIdFrame.SetPaddings(2, 2)
	ui.CreateEditField(guildIdFrame, ui.AutoSize, tmpGuidId, ui.Fixed)
	ui.CreateLabel(guildIdFrame, ui.AutoSize, ui.AutoSize, TxtGuildIdDesc, ui.Fixed)

	// --- Bot Users -----------------------------------------------------
	botUsersFrame := NewFramedInput(settingsFrame, TxtBotUsers, nil)
	botUsersFrame.SetPaddings(2, 2)
	ui.CreateEditField(botUsersFrame, ui.AutoSize, DiscordData.BotUsers[0], ui.Fixed)
	ui.CreateLabel(botUsersFrame, ui.AutoSize, ui.AutoSize, TxtBotUsersDesc, ui.Fixed)

	// --- Asset Codes ---------------------------------------------------
	assetCodesFrame := NewFramedInput(settingsFrame, fmt.Sprintf("%s - %s", TxtAssetCodes, TxtAssetCodesDesc), nil)
	BtnAssetCodes = ui.CreateButton_NoShadow(assetCodesFrame, 22, ui.AutoSize, TxtAssetCodesBtn, ui.Fixed)
	BtnAssetCodes.SetAlign(ui.AlignLeft)
	BtnAssetCodes.OnClick(func(ev ui.Event) {
		BtnAssetCodes.SetEnabled(false)
		assetCodes := CreateListDialog(TxtAssetCodes)
		assetCodes.OnClose(func() {
			BtnAssetCodes.SetEnabled(true)
		})
	})

	// --- Asset Details -------------------------------------------------
	logLevelFrame := NewFramedInput(settingsFrame, TxtAssetDetails, nil)
	BtnAssetDetails = ui.CreateButton_NoShadow(logLevelFrame, ui.AutoSize, ui.AutoSize, TxtAssetDetailsBtn, ui.Fixed)
	BtnAssetDetails.SetAlign(ui.AlignLeft)
	BtnAssetDetails.OnClick(func(ev ui.Event) {
		FrmDiscordSettings.SetActive(false)
		CreateTableDialog(BtnAssetDetails, TxtAssetDetails)
	})
	ui.CreateLabel(logLevelFrame, ui.AutoSize, ui.AutoSize, TxtAssetDetailsDesc, ui.Fixed)

	// --- Save Settings -------------------------------------------------
	btnFrame := ui.CreateFrame(settingsFrame, 10, 1, ui.BorderNone, ui.Fixed)
	btnFrame.SetPaddings(2, 2)
	var params = FramedInputParams{Orientation: ui.Vertical, Width: 25, Height: 4}
	saveSettings := NewFramedInput(btnFrame, TxtSaveDesc, &params)
	BtnMainSettingsSave = ui.CreateButton_NoShadow(saveSettings, 25, ui.AutoSize, TxtSave, ui.Fixed)
	BtnMainSettingsSave.SetAlign(ui.AlignLeft)
	BtnMainSettingsSave.OnClick(func(ev ui.Event) {
		DiscordData.GuildId = tmpGuidId
		DiscordData.BotUsers = tmpBotUsers
		DiscordData.Roles = tmpRoles
	})
	FrmDiscordSettings.SetVisible(false)
}
