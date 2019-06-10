package view

import (
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
)

type ProviderGroup struct {
	Group      *ui.RadioGroup
	RadioGroup *ui.Radio
}

func GetProvider(db *DatabaseDetails, group *ProviderGroup) {
	switch db.Provider {
	case 0:
		//group.Group.SetSelected(group.RadioGroup[0])
	}

}

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
	settingsFrame.SetTitle(TxtDatabase)
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

	//var group = [mySQLRadio, postgresRadio, mssqlRadio, sqliteRadio]

	//GetProvider(DatabaseData, providerGroup)

	providerGroup.SelectItem(mySQLRadio)
	tmpProvider = providerGroup.Selected()

	// --- Database Details -----------------------------------------------------
	dbDetailsFrame := NewFramedInput(settingsFrame, TxtDbDetails, nil)
	dbDetailsFrame.SetPaddings(2, 2)
	dbDetailsFrame.SetGaps(0, 0)

	addressFrame := ui.CreateFrame(dbDetailsFrame, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	addressFrame.SetPack(ui.Horizontal)
	addressFrame.SetPaddings(0, 0)
	addressFrame.SetGaps(0, 0)
	ui.CreateLabel(addressFrame, ui.AutoSize, ui.AutoSize, TxtDbAddress, ui.Fixed)
	ui.CreateEditField(addressFrame, 50, tmpAddress, ui.Fixed)

	usernameFrame := ui.CreateFrame(dbDetailsFrame, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	usernameFrame.SetPack(ui.Horizontal)
	ui.CreateLabel(usernameFrame, ui.AutoSize, ui.AutoSize, TxtDbUsername, ui.Fixed)
	ui.CreateEditField(usernameFrame, 50, tmpUsername, ui.Fixed)

	passwordFrame := ui.CreateFrame(dbDetailsFrame, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	passwordFrame.SetPack(ui.Horizontal)
	ui.CreateLabel(passwordFrame, ui.AutoSize, ui.AutoSize, TxtDbPassword, ui.Fixed)
	ui.CreateEditField(passwordFrame, 50, tmpPassword, ui.Fixed)

	databaseFrame := ui.CreateFrame(dbDetailsFrame, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	databaseFrame.SetPack(ui.Horizontal)
	ui.CreateLabel(databaseFrame, ui.AutoSize, ui.AutoSize, TxtDbDatabase, ui.Fixed)
	ui.CreateEditField(databaseFrame, 50, tmpDatabase, ui.Fixed)

	prefixFrame := ui.CreateFrame(dbDetailsFrame, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	prefixFrame.SetPack(ui.Horizontal)
	ui.CreateLabel(prefixFrame, ui.AutoSize, ui.AutoSize, TxtDbTablePrefix, ui.Fixed)
	ui.CreateEditField(prefixFrame, 50, tmpTablePrefix, ui.Fixed)

	// --- Window Control ------------------------------------------------
	btnFrame := ui.CreateFrame(settingsFrame, 10, 1, ui.BorderNone, ui.Fixed)
	btnFrame.SetPaddings(2, 2)

	var params = FramedInputParams{Orientation: ui.Vertical, Width: 10, Height: 4, Scale: ui.Fixed}
	saveSettings := NewFramedInput(btnFrame, TxtSaveDesc, &params)
	BtnMainSettingsSave = ui.CreateButton_NoShadow(saveSettings, ui.AutoSize, ui.AutoSize, TxtSave, ui.Fixed)
	BtnMainSettingsSave.SetAlign(ui.AlignLeft)
	BtnMainSettingsSave.OnClick(func(ev ui.Event) {
		DatabaseData.Provider = tmpProvider
		DatabaseData.Address = tmpAddress
		DatabaseData.Username = tmpUsername
		DatabaseData.Password = tmpPassword
		DatabaseData.Database = tmpDatabase
		DatabaseData.TablePrefix = tmpTablePrefix
	})

	BtnMainSettingsSave.SetActive(false)
	FrmDatabaseSettings.SetVisible(false)
}
