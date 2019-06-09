package view

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
)

func CreateViewDatabaseSettings() {

	var tmpProvider = DatabaseData.Provider
	var tmpAddress = DatabaseData.Address
	var tmpUsername = DatabaseData.Username
	var tmpPassword = DatabaseData.Password
	var tmpDatabase = DatabaseData.Database
	var tmpTablePrefix = DatabaseData.TablePrefix

	// --- Database Settings Frame ---------------------------------------
	FrmDatabaseSettings = ui.CreateFrame(FrameContent, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	FrmDatabaseSettings.SetTitle("FrameTop")
	FrmDatabaseSettings.SetPack(ui.Horizontal)

	// --- Database Settings Content -------------------------------------
	settingsFrame := ui.CreateFrame(FrmDatabaseSettings, 130, ui.AutoSize, ui.BorderThin, ui.Fixed)
	settingsFrame.SetPaddings(2, 2)
	settingsFrame.SetTitle(TxtDatabaseSettings)
	settingsFrame.SetPack(ui.Vertical)

	// --- Database Provider ---------------------------------------------
	//var gIdparams = FramedInputParams{Border: ui.BorderNone}
	providerFrame := NewFramedInput(settingsFrame, TxtDbProvider, nil)
	providerFrame.SetPaddings(2, 2)

	mySQLRadio := ui.CreateRadio(providerFrame, ui.AutoSize, TxtMysql, ui.Fixed)
	postgresRadio := ui.CreateRadio(providerFrame, ui.AutoSize, TxtPostgres, ui.Fixed)
	mssqlRadio := ui.CreateRadio(providerFrame, ui.AutoSize, TxtMSSQL, ui.Fixed)
	sqliteRadio := ui.CreateRadio(providerFrame, ui.AutoSize, TxtSqlite, ui.Fixed)

	providerGroup := ui.CreateRadioGroup()
	providerGroup.AddItem(mySQLRadio)
	providerGroup.AddItem(postgresRadio)
	providerGroup.AddItem(mssqlRadio)
	providerGroup.AddItem(sqliteRadio)

	ui.CreateLabel(providerFrame, ui.AutoSize, ui.AutoSize, TxtGuildIdDesc, ui.Fixed)

	// --- Bot Users -----------------------------------------------------
	botUsersFrame := NewFramedInput(settingsFrame, TxtBotUsers, nil)
	botUsersFrame.SetPaddings(2, 2)
	ui.CreateEditField(botUsersFrame, ui.AutoSize, DatabaseData.TablePrefix, ui.Fixed)
	ui.CreateLabel(botUsersFrame, ui.AutoSize, ui.AutoSize, TxtBotUsersDesc, ui.Fixed)

	// --- Window Control ------------------------------------------------
	btnFrame := ui.CreateFrame(settingsFrame, 1, 1, ui.BorderNone, ui.Fixed)
	btnFrame.SetPaddings(1, 1)

	var params = FramedInputParams{Orientation: ui.Vertical, Width: 4, Height: 4}
	saveSettings := NewFramedInput(btnFrame, TxtSaveDesc, &params)
	BtnMainSettingsSave = ui.CreateButton(saveSettings, ui.AutoSize, ui.AutoSize, TxtSave, ui.Fixed)
	BtnMainSettingsSave.OnClick(func(ev ui.Event) {
		DatabaseData.Provider = tmpProvider
		DatabaseData.Address = tmpAddress
		DatabaseData.Username = tmpUsername
		DatabaseData.Password = tmpPassword
		DatabaseData.Database = tmpDatabase
		DatabaseData.TablePrefix = tmpTablePrefix
	})
	FrmDatabaseSettings.SetVisible(false)
}
