package elements

type LogLevels struct {
	LogLevel        []string
	CurrentLogLevel string
	DefaultLogLevel string
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

var (
	Log           LogLevels
	DiscordToken  string
	CommandPrefix string
)

var DiscordData = DataDiscord{
	GuildId:  "123123123123",
	BotUsers: []string{"12312312344444", "32132132188888"},
	Roles:    map[string]string{"ABC": "44444444444444", "BCD": "55555555555555"},
}

var AssetData = DataAssets{
	AssetCodes:    []string{"ABC", "BCD"},
	AssetPackages: map[string]string{"ABC": "ABC - Best Asset", "BCD": "BCD : Also Best Asset"},
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
