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
	AssetData     DataAssets
	DiscordData   DataDiscord
)
