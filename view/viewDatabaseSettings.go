package view

import (
	"fmt"

	. "github.com/instance-id/GoUI/components"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
)

var pendingDB = struct {
	provider bool
	address  bool
	username bool
	password bool
	database bool
	prefix   bool
}{
	provider: false,
	address:  false,
	username: false,
	password: false,
	database: false,
	prefix:   false,
}

func CreateViewDatabaseSettings() {
	var tmpProviders = Cntnrs.Dbd.Providers
	var tmpProvider = Cntnrs.Dbd.Database
	var tmpAddress = Cntnrs.Dbd.Data.Address
	var tmpUsername = Cntnrs.Dbd.Data.Username
	var tmpPassword = Cntnrs.Dbd.Data.Password
	var tmpDatabase = Cntnrs.Dbd.Data.DbName
	var tmpTablePrefix = Cntnrs.Dbd.Data.TablePrefix

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
	BtnDatabaseProvider = ui.CreateButton(providerFrame, ui.AutoSize, ui.AutoSize, fmt.Sprintf(" | %s", Cntnrs.Dbd.Providers[tmpProvider]), ui.Fixed)
	BtnDatabaseProvider.SetAlign(ui.AlignLeft)
	BtnDatabaseProvider.SetShadowType(ui.ShadowHalf)
	BtnDatabaseProvider.OnClick(func(ev ui.Event) {
		dbProvider := ui.CreateSelectDialog(TxtDbProvider, tmpProviders, tmpProvider, ui.SelectDialogList)
		dbProvider.OnClose(func() {
			switch dbProvider.Result() {
			case ui.DialogButton1:
				tmpProvider = dbProvider.Value()
				BtnDatabaseProvider.SetTitle(fmt.Sprintf(" | %s", Cntnrs.Dbd.Providers[tmpProvider]))
				if tmpProvider != Cntnrs.Dbd.Database {
					pendingDB.provider = true
					SavePendingDatabase()
				}
				if tmpProvider == Cntnrs.Dbd.Database {
					pendingDB.provider = false
					SavePendingDatabase()
				}
			}
		})
	})

	// --- Database Details ----------------------------------------------
	dbDetailsFrame := NewFramedInput(settingsFrame, TxtDbDetails, nil)
	dbDetailsFrame.SetPaddings(2, 2)
	dbDetailsFrame.SetGaps(0, 0)

	// --- Address ------------------------------
	addressFrame := ui.CreateFrame(dbDetailsFrame, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	addressFrame.SetPack(ui.Horizontal)
	addressFrame.SetPaddings(0, 0)
	addressFrame.SetGaps(0, 0)
	ui.CreateLabel(addressFrame, ui.AutoSize, ui.AutoSize, TxtDbAddress, ui.Fixed)
	addressResult := ui.CreateEditField(addressFrame, 50, tmpAddress, ui.Fixed)
	addressResult.OnChange(func(event ui.Event) {
		tmpAddress = addressResult.Title()
		if tmpAddress != Cntnrs.Dbd.Data.Address {
			pendingDB.address = true
			SavePendingDatabase()
		}
		if tmpAddress == Cntnrs.Dbd.Data.Address {
			pendingDB.address = false
			SavePendingDatabase()
		}

	})

	// --- Username -----------------------------
	usernameFrame := ui.CreateFrame(dbDetailsFrame, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	usernameFrame.SetPack(ui.Horizontal)
	ui.CreateLabel(usernameFrame, ui.AutoSize, ui.AutoSize, TxtDbUsername, ui.Fixed)
	usernameResult := ui.CreateEditField(usernameFrame, 50, tmpUsername, ui.Fixed)
	usernameResult.OnChange(func(event ui.Event) {
		tmpUsername = usernameResult.Title()
		if tmpUsername != Cntnrs.Dbd.Data.Username {
			pendingDB.username = true
			SavePendingDatabase()
		}
		if tmpUsername == Cntnrs.Dbd.Data.Username {
			pendingDB.username = false
			SavePendingDatabase()
		}
	})

	// --- Password -----------------------------
	passwordFrame := ui.CreateFrame(dbDetailsFrame, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	passwordFrame.SetPack(ui.Horizontal)
	ui.CreateLabel(passwordFrame, ui.AutoSize, ui.AutoSize, TxtDbPassword, ui.Fixed)
	passwordResult := ui.CreateEditField(passwordFrame, 50, tmpPassword, ui.Fixed)
	passwordResult.OnChange(func(event ui.Event) {
		tmpPassword = passwordResult.Title()
		if tmpPassword != Cntnrs.Dbd.Data.Password {
			pendingDB.password = true
			SavePendingDatabase()
		}
		if tmpPassword == Cntnrs.Dbd.Data.Password {
			pendingDB.password = false
			SavePendingDatabase()
		}
	})

	// --- Database -----------------------------
	databaseFrame := ui.CreateFrame(dbDetailsFrame, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	databaseFrame.SetPack(ui.Horizontal)
	ui.CreateLabel(databaseFrame, ui.AutoSize, ui.AutoSize, TxtDbDatabase, ui.Fixed)
	databaseResult := ui.CreateEditField(databaseFrame, 50, tmpDatabase, ui.Fixed)
	databaseResult.OnChange(func(event ui.Event) {
		tmpDatabase = databaseResult.Title()
		if tmpDatabase != Cntnrs.Dbd.Data.DbName {
			pendingDB.database = true
			SavePendingDatabase()
		}
		if tmpDatabase == Cntnrs.Dbd.Data.DbName {
			pendingDB.database = false
			SavePendingDatabase()
		}
	})

	// --- Table Prefix -------------------------
	prefixFrame := ui.CreateFrame(dbDetailsFrame, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	prefixFrame.SetPack(ui.Horizontal)
	ui.CreateLabel(prefixFrame, ui.AutoSize, ui.AutoSize, TxtDbTablePrefix, ui.Fixed)
	tblPrefixResult := ui.CreateEditField(prefixFrame, 50, tmpTablePrefix, ui.Fixed)
	tblPrefixResult.OnChange(func(event ui.Event) {
		tmpTablePrefix = tblPrefixResult.Title()
		if tmpTablePrefix != Cntnrs.Dbd.Data.TablePrefix {
			pendingDB.prefix = true
			SavePendingDatabase()
		}
		if tmpTablePrefix == Cntnrs.Dbd.Data.TablePrefix {
			pendingDB.prefix = false
			SavePendingDatabase()
		}
	})

	// --- Window Control -----------------------
	btnFrame := ui.CreateFrame(settingsFrame, 10, 1, ui.BorderNone, ui.Fixed)
	btnFrame.SetPaddings(2, 2)

	// --- Save Button --------------------------
	var saveParams = FramedInputParams{Orientation: ui.Vertical, Width: 25, Height: 4, Scale: ui.Fixed}
	saveSettings := NewFramedInput(btnFrame, TxtSaveDesc, &saveParams)
	BtnDatabaseSettingsSave = ui.CreateButton(saveSettings, ui.AutoSize, ui.AutoSize, TxtSaveBtn, ui.Fixed)
	BtnDatabaseSettingsSave.SetAlign(ui.AlignLeft)
	BtnDatabaseSettingsSave.SetShadowType(ui.ShadowHalf)
	// --- Save settings back to container ------
	BtnDatabaseSettingsSave.OnClick(func(ev ui.Event) {
		Cntnrs.Dbd.Database = tmpProvider
		Cntnrs.Dbd.Data.Address = tmpAddress
		Cntnrs.Dbd.Data.Username = tmpUsername
		Cntnrs.Dbd.Data.Password = tmpPassword
		Cntnrs.Dbd.Data.DbName = tmpDatabase
		Cntnrs.Dbd.Data.TablePrefix = tmpTablePrefix
		_, err := Cntnrs.Wtr.SetDbConfig()
		if err != nil {
			ui.CreateAlertDialog(ErrCouldNotSaveDb, fmt.Sprintf("Error Could not save DB config: %s", err), TxtCloseBtn)
		}
		BtnDatabaseSettingsSave.SetTitle(TxtSaveBtn)
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	})

	BtnDatabaseSettingsSave.SetActive(false)
	FrmDatabaseSettings.SetVisible(false)
}

func SavePendingDatabase() {
	if pendingDB.provider || pendingDB.address || pendingDB.username || pendingDB.password || pendingDB.database || pendingDB.prefix == true {
		BtnDatabaseSettingsSave.SetTitle(TxtSavePendingBtn)
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	} else {
		BtnDatabaseSettingsSave.SetTitle(TxtSaveBtn)
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	}
}
