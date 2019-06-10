package view

import (
	"fmt"
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

	var tmpProviders = DatabaseData.Providers
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
	providerFrame := NewFramedInput(settingsFrame, TxtDbProvider, nil)
	providerFrame.SetPaddings(2, 2)
	BtnDatabaseProvider = ui.CreateButton_NoShadow(providerFrame, ui.AutoSize, ui.AutoSize, fmt.Sprintf(" | %s", DatabaseData.Providers[tmpProvider]), ui.Fixed)
	BtnDatabaseProvider.SetAlign(ui.AlignLeft)
	BtnDatabaseProvider.OnClick(func(ev ui.Event) {
		dbProvider := ui.CreateSelectDialog(TxtDbProvider, tmpProviders, ui.AutoSize, ui.SelectDialogList)
		dbProvider.OnClose(func() {
			tmpProvider = dbProvider.Value()
			BtnDatabaseProvider.SetTitle(fmt.Sprintf(" | %s", DatabaseData.Providers[tmpProvider]))
		})
	})

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

	var saveParams = FramedInputParams{Orientation: ui.Vertical, Width: 25, Height: 4, Scale: ui.Fixed}
	saveSettings := NewFramedInput(btnFrame, TxtSaveDesc, &saveParams)
	BtnMainSettingsSave = ui.CreateButton_NoShadow(saveSettings, ui.AutoSize, ui.AutoSize, TxtSaveBtn, ui.Fixed)
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
