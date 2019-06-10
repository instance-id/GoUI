package elements

import (
	ui "github.com/instance-id/clui"
)

type LogLevels struct {
	LogLevel        []string
	CurrentLogLevel string
	DefaultLogLevel string
}

type TableEdit struct {
	Row    int
	Col    int
	OldVal string
	NewVal string
}

type AssetDetails struct {
	AssetCode     string
	AssetName     string
	AssetApiKey   string
	AssetRole     string
	AssetVersion  string
	AssetReplaced string
	ReplaceDate   string
}

type DataAssets struct {
	AssetCodes    []string
	AssetPackages map[string]string
	AssetApiKeys  map[string]string
}

type DataDiscord struct {
	GuildId  string
	BotUsers []string
	Roles    map[string]string
}

type DatabaseDetails struct {
	Providers   []string
	Provider    int
	Address     string
	Username    string
	Password    string
	Database    string
	TablePrefix string
}

type TableDialog struct {
	View       *ui.Window
	Frame      *ui.Frame
	result     int
	value      int
	edtResult  string
	table      *ui.TableView
	edit       *ui.EditField
	editDialog *ui.SelectDialog
	confirm    *ui.ConfirmationDialog
	onClose    func()
}

type ListDialog struct {
	View       *ui.Window
	Frame      *ui.Frame
	result     int
	value      int
	edtResult  string
	list       *ui.ListBox
	edit       *ui.EditField
	editDialog *ui.SelectDialog
	confirm    *ui.ConfirmationDialog
	onClose    func()
}

type LogDialog struct {
	View       *ui.Window
	Frame      *ui.Frame
	result     int
	value      int
	edtResult  string
	log        *ui.TextView
	edit       *ui.EditField
	editDialog *ui.SelectDialog
	confirm    *ui.ConfirmationDialog
	onClose    func()
}

var (
	Log           LogLevels
	DiscordToken  string
	CommandPrefix string
)

var AssetDetail = []*AssetDetails{{AssetCode: "SCT", AssetName: "SCT - Scriptable Text", AssetApiKey: "123123112312312323123123123", AssetRole: "123123123123123123", AssetReplaced: "No", AssetVersion: "1", ReplaceDate: ""},
	{AssetCode: "UFPS1", AssetName: "UFPS : Ultimate FPS", AssetApiKey: "123123112312312323123123123", AssetRole: "123123123123123123", AssetReplaced: "Yes", AssetVersion: "1", ReplaceDate: "2018-06-06"},
	{AssetCode: "UCC", AssetName: "Ultimate Character Controller", AssetApiKey: "123123112312312323123123123", AssetRole: "123123123123123123", AssetReplaced: "Yes", AssetVersion: "1", ReplaceDate: "2018-06-06"},
	{AssetCode: "TPC", AssetName: "Third Person Controller", AssetApiKey: "123123112312312323123123123", AssetRole: "123123123123123123", AssetReplaced: "Yes", AssetVersion: "1", ReplaceDate: "2018-06-06"},
	{AssetCode: "UTPS", AssetName: "UTPS: Ultimate Third Person Shooter", AssetApiKey: "123123112312312323123123123", AssetRole: "123123123123123123", AssetReplaced: "Yes", AssetVersion: "1", ReplaceDate: "2018-06-06"},
	{AssetCode: "UTPM", AssetName: "UTPM: Ultimate Third Person Melee", AssetApiKey: "123123112312312323123123123", AssetRole: "123123123123123123", AssetReplaced: "Yes", AssetVersion: "1", ReplaceDate: "2018-06-06"},
	{AssetCode: "TPC1", AssetName: "Third Person Controller", AssetApiKey: "123123112312312323123123123", AssetRole: "123123123123123123", AssetReplaced: "Yes", AssetVersion: "1", ReplaceDate: "2018-06-06"},
	{AssetCode: "FPC", AssetName: "First Person Controller", AssetApiKey: "123123112312312323123123123", AssetRole: "123123123123123123", AssetReplaced: "Yes", AssetVersion: "1", ReplaceDate: "2018-06-06"},
	{AssetCode: "BD", AssetName: "Behavior Designer - Behavior Trees for Everyone", AssetApiKey: "123123112312312323123123123", AssetRole: "123123123123123123", AssetReplaced: "No", AssetVersion: "1", ReplaceDate: "2018-06-06"},
	{AssetCode: "UFPS", AssetName: "UFPS : Ultimate FPS", AssetApiKey: "123123112312312323123123123", AssetRole: "123123123123123123", AssetReplaced: "Yes", AssetVersion: "1", ReplaceDate: "2018-06-06"},
	{AssetCode: "UFPM", AssetName: "UFPM: Ultimate First Person Melee", AssetApiKey: "123123112312312323123123123", AssetRole: "123123123123123123", AssetReplaced: "Yes", AssetVersion: "1", ReplaceDate: "2018-06-06"}}

var DiscordData = DataDiscord{
	GuildId:  "123123123123",
	BotUsers: []string{"12312312344444", "32132132188888"},
	Roles:    map[string]string{"ABC": "44444444444444", "BCD": "55555555555555"},
}

var DatabaseData = DatabaseDetails{
	Address:     "instance.id",
	Username:    "Username",
	Password:    "Password",
	Database:    "Verifier",
	TablePrefix: "verifier_",
	Providers:   []string{"MySQL", "Postgres", "MSSQL", "SqLite"}, // 0=MySQL, 1=Postgres, 2=MSSQL, 3=SqLite
	Provider:    0,                                                // 0=MySQL, 1=Postgres, 2=MSSQL, 3=SqLite
}

var AssetData = DataAssets{
	AssetCodes:    []string{"ABC", "BCD"},
	AssetPackages: map[string]string{"ABC": "ABC - Best Asset", "BCD": "BCD : Also Best Asset", AssetDetail[0].AssetCode: AssetDetail[0].AssetName},
	AssetApiKeys:  map[string]string{"ABC": "1231232123123123", "BCD": "3453453453453645"},
}

const (
	// DialogClosed - a user clicked close button on the dialog title
	DialogClosed = -1
	// DialogAlive - a user does not close the dialog yet, exit code is unavailable
	DialogAlive = 0
	// DialogButton1 - a user clicked the first button in the dialog (by default, it is 'Yes' or 'OK')
	DialogButton1 = 1
	// DialogButton2 - a user clicked the second button in the dialog
	DialogButton2 = 2
	// DialogButton3 - a user clicked the third button in the dialog
	DialogButton3 = 3
	// DialogButton4 - a user clicked the fourth button in the dialog
	DialogButton4 = 4
)
