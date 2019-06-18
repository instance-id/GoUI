package view

import (
	"fmt"

	. "github.com/instance-id/GoUI/components"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
)

var pendingDiscord = struct {
	guildId  bool
	botUsers bool
}{
	guildId:  false,
	botUsers: false,
}

func CreateViewDiscordSettings() {

	var tmpGuidId = Cntnrs.Dac.Discord.Guild
	var tmpBotUsers = Cntnrs.Dac.Discord.BotUsers[0]
	var tmpRoles = Cntnrs.Dac.Discord.Roles

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
	guildIdFrame := NewFramedInput(settingsFrame, TxtGuildId, nil)
	guildIdFrame.SetPaddings(2, 2)
	guildIdResult := ui.CreateEditField(guildIdFrame, ui.AutoSize, tmpGuidId, ui.Fixed)
	guildIdResult.OnChange(func(event ui.Event) {
		tmpGuidId = guildIdResult.Title()
		if tmpGuidId != Cntnrs.Dac.Discord.Guild {
			pendingDiscord.guildId = true
			SavePendingDiscord()
		} else {
			pendingDiscord.guildId = false
			SavePendingDiscord()
		}

	})
	ui.CreateLabel(guildIdFrame, ui.AutoSize, ui.AutoSize, TxtGuildIdDesc, ui.Fixed)

	// --- Bot Users -----------------------------------------------------
	botUsersFrame := NewFramedInput(settingsFrame, TxtBotUsers, nil)
	botUsersFrame.SetPaddings(2, 2)
	botUsersResult := ui.CreateEditField(botUsersFrame, ui.AutoSize, tmpBotUsers, ui.Fixed)
	botUsersResult.OnChange(func(event ui.Event) {
		tmpBotUsers = botUsersResult.Title()
		if tmpBotUsers != Cntnrs.Dac.Discord.BotUsers[0] {
			pendingDiscord.botUsers = true
			SavePendingDiscord()
		} else {
			pendingDiscord.botUsers = false
			SavePendingDiscord()
		}
	})
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
	BtnDiscordSettingsSave = ui.CreateButton(saveSettings, 25, ui.AutoSize, TxtSaveBtn, ui.Fixed)
	BtnDiscordSettingsSave.SetAlign(ui.AlignLeft)
	BtnDiscordSettingsSave.SetShadowType(ui.ShadowHalf)
	BtnDiscordSettingsSave.OnClick(func(ev ui.Event) {
		Cntnrs.Dac.Discord.Guild = tmpGuidId
		Cntnrs.Dac.Discord.BotUsers[0] = tmpBotUsers
		Cntnrs.Dac.Discord.Roles = tmpRoles
		_, err := Cntnrs.Wtr.SetConfig()
		if err != nil {
			ui.CreateAlertDialog(ErrCouldNotSaveCfg, fmt.Sprintf("Error Could not save config: %s", err), TxtCloseBtn)
		}
		BtnDiscordSettingsSave.SetTitle(TxtSaveBtn)
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	})
}

func SavePendingDiscord() {
	if pendingDiscord.guildId || pendingDiscord.botUsers == true {
		BtnDiscordSettingsSave.SetTitle(TxtSavePendingBtn)
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	} else {
		BtnDiscordSettingsSave.SetTitle(TxtSaveBtn)
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	}
}
