package view

import (
	"fmt"
	. "github.com/instance-id/GoUI/dicontainer"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
)

func CreateViewDiscordSettings() {

	var tmpGuidId = DiCon.Cnt.Dac.Discord.Guild
	var tmpBotUsers = DiCon.Cnt.Dac.Discord.BotUsers
	var tmpRoles = DiCon.Cnt.Dac.Discord.Roles

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
	ui.CreateEditField(botUsersFrame, ui.AutoSize, DiCon.Cnt.Dac.Discord.BotUsers[0], ui.Fixed)
	ui.CreateLabel(botUsersFrame, ui.AutoSize, ui.AutoSize, TxtBotUsersDesc, ui.Fixed)

	// --- Asset Details -------------------------------------------------
	logLevelFrame := NewFramedInput(settingsFrame, TxtAssetDetails, nil)
	BtnAssetDetails = ui.CreateButton(logLevelFrame, ui.AutoSize, ui.AutoSize, TxtAssetDetailsBtn, ui.Fixed)
	BtnAssetDetails.SetAlign(ui.AlignLeft)
	BtnAssetDetails.SetShadowType(ui.ShadowHalf)
	BtnAssetDetails.OnClick(func(ev ui.Event) {
		FrmDiscordSettings.SetActive(false)
		td := CreateTableDialog(BtnAssetDetails)
		td.SetActive(true)

	})
	ui.CreateLabel(logLevelFrame, ui.AutoSize, ui.AutoSize, TxtAssetDetailsDesc, ui.Fixed)

	// --- Save Settings -------------------------------------------------
	btnFrame := ui.CreateFrame(settingsFrame, 10, 1, ui.BorderNone, ui.Fixed)
	btnFrame.SetPaddings(2, 2)
	var params = FramedInputParams{Orientation: ui.Vertical, Width: 25, Height: 4}
	saveSettings := NewFramedInput(btnFrame, TxtSaveDesc, &params)
	BtnMainSettingsSave = ui.CreateButton(saveSettings, 25, ui.AutoSize, TxtSaveBtn, ui.Fixed)
	BtnMainSettingsSave.SetAlign(ui.AlignLeft)
	BtnMainSettingsSave.SetShadowType(ui.ShadowHalf)
	BtnMainSettingsSave.OnClick(func(ev ui.Event) {
		DiCon.Cnt.Dac.Discord.Guild = tmpGuidId
		DiCon.Cnt.Dac.Discord.BotUsers = tmpBotUsers
		DiCon.Cnt.Dac.Discord.Roles = tmpRoles
		_, err := DiCon.Cnt.Wtr.SetConfig()
		if err != nil {
			ui.CreateAlertDialog(ErrCouldNotSaveCfg, fmt.Sprintf("Error Could not save config: &s", err), TxtCloseBtn)
		}
	})
	FrmDiscordSettings.SetVisible(false)
}
