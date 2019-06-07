package view

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
)

func CreateViewDiscordSettings() {

	var discordData = DataDiscord{
		GuildId:  "123123123123",
		BotUsers: []string{"12312312344444", "32132132188888"},
		Roles:    map[string]string{"ABC": "44444444444444", "BCD": "55555555555555"},
	}

	var assetData = DataAssets{
		AssetCodes:    []string{"ABC", "BCD"},
		AssetPackages: map[string]string{"ABC": "ABC - Best Asset", "BCD": "BCD : Also Best Asset"},
		AssetApiKeys:  map[string]string{"ABC": "1231232123123123", "BCD": "3453453453453645"},
	}

	var tmpGuidId = discordData.GuildId
	var tmpBotUsers = discordData.BotUsers
	var tmpRoles = discordData.Roles
	var tmpAssetCodes = assetData.AssetCodes

	// --- Discord Settings Frame ------------------------------------------
	FrmDiscordSettings = ui.CreateFrame(FrameContent, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	FrmDiscordSettings.SetTitle("FrameTop")
	FrmDiscordSettings.SetPack(ui.Horizontal)

	// --- Discord Settings Content ----------------------------------------
	settingsFrame := ui.CreateFrame(FrmDiscordSettings, 100, ui.AutoSize, ui.BorderThin, ui.Fixed)
	settingsFrame.SetPaddings(2, 2)
	settingsFrame.SetTitle(TxtDiscordSettings)
	settingsFrame.SetPack(ui.Vertical)

	// --- GuildId -------------------------------------------------------
	guildIdFrame := NewFramedInput(settingsFrame, TxtGuildId, nil)
	ui.CreateEditField(guildIdFrame, 70, tmpGuidId, ui.Fixed)
	ui.CreateLabel(guildIdFrame, ui.AutoSize, ui.AutoSize, TxtGuildIdDesc, ui.Fixed)

	// --- Bot Users -----------------------------------------------------
	botUsersFrame := NewFramedInput(settingsFrame, TxtBotUsers, nil)
	ui.CreateEditField(botUsersFrame, 50, discordData.BotUsers[0], ui.Fixed)
	ui.CreateLabel(botUsersFrame, ui.AutoSize, ui.AutoSize, TxtBotUsersDesc, ui.Fixed)

	// --- Asset Codes ---------------------------------------------------
	assetCodesFrame := NewFramedInput(settingsFrame, TxtAssetCodes, nil)
	BtnAssetCodes = ui.CreateButton(assetCodesFrame, ui.AutoSize, ui.AutoSize, TxtAssetCodesBtn, ui.Fixed)
	BtnAssetCodes.OnClick(func(ev ui.Event) {
		BtnAssetCodes.SetEnabled(false)
		ChangeAssetCodes(BtnAssetCodes)
	})
	ui.CreateLabel(assetCodesFrame, ui.AutoSize, ui.AutoSize, TxtAssetCodesDesc, ui.Fixed)

	// --- Select Log Level ----------------------------------------------
	logLevelFrame := NewFramedInput(settingsFrame, TxtLogLevel, nil)
	BtnLogLevel := ui.CreateButton(logLevelFrame, ui.AutoSize, ui.AutoSize, TxtLogLevel, ui.Fixed)
	BtnLogLevel.OnClick(func(ev ui.Event) {
		BtnLogLevel.SetEnabled(false)
		SelectLogLevel(BtnLogLevel)
	})
	ui.CreateLabel(logLevelFrame, ui.AutoSize, ui.AutoSize, TxtLogLevelDesc, ui.Fixed)

	// --- Save Settings ------------------------------------------------
	var params = FramedInputParams{Orientation: ui.Vertical, Width: 4, Height: 4}
	saveSettings := NewFramedInput(settingsFrame, TxtSaveDesc, &params)
	BtnMainSettingsSave = ui.CreateButton(saveSettings, ui.AutoSize, ui.AutoSize, TxtSave, ui.Fixed)
	BtnMainSettingsSave.OnClick(func(ev ui.Event) {
		discordData.GuildId = tmpGuidId
		discordData.BotUsers = tmpBotUsers
		discordData.Roles = tmpRoles
		assetData.AssetCodes = tmpAssetCodes
	})
	FrmDiscordSettings.SetVisible(false)
}
