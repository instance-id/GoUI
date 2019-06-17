package components

import "fmt"

var Asset AssetContainer

type AssetContainer struct {
	AD []*AssetDetails
}

type AssetDetails struct {
	AssetCode        string
	AssetName        string
	AssetApiKey      string
	AssetRole        string
	AssetReplaced    string
	AssetReplacement string
	ReplaceDate      string
}

type ConfigPath struct {
	FolderName string
	FileName   string
	DbFileName string
	Path       string
	DbPath     string
}

var configPath = new(ConfigPath)

func (a *AssetContainer) AddNewAsset(asset *AssetDetails) []*AssetDetails {
	a.AD = append(a.AD, asset)
	return a.AD
}

func LoadTableData() AssetContainer {
	ad := Cntnrs.Dac
	//name := ad.Assets.AssetCodes[0]
	//fmt.Printf(" - AssetCode: %s \n", name)
	//fmt.Printf(" - AssetName: %s \n", ad.Assets.Packages[name])
	//fmt.Printf(" - AssetApikey: %s \n", ad.Assets.ApiKeys[name])
	//fmt.Printf(" - AssetRoles: %s \n", ad.Discord.Roles[name])
	AssetD := func(settings *MainSettings) AssetContainer {
		Ac := new(AssetContainer)
		for a := range ad.Assets.AssetCodes {
			Ac.AD = append(Ac.AD, &AssetDetails{
				ad.Assets.AssetCodes[a],
				ad.Assets.Packages[ad.Assets.AssetCodes[a]],
				ad.Assets.ApiKeys[ad.Assets.AssetCodes[a]],
				ad.Discord.Roles[ad.Assets.AssetCodes[a]],
				ad.Assets.AssetReplaced[ad.Assets.AssetCodes[a]],
				ad.Assets.AssetReplacement[ad.Assets.AssetCodes[a]],
				ad.Assets.ReplaceDate[ad.Assets.AssetCodes[a]]})
		}
		return *Ac
	}(ad)

	return AssetD

}

func init() {

	fmt.Printf("Loading Init in data.go \n")
	configPath.FolderName = "config"
	configPath.FileName = "config.yml"
	configPath.DbFileName = "dbconfig.yml"
	configPath.Path = fmt.Sprintf("./%s/%s", configPath.FolderName, configPath.FileName)
	configPath.DbPath = fmt.Sprintf("./%s/%s", configPath.FolderName, configPath.DbFileName)

}
