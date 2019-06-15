package components

var Asset *AssetContainer

//var Providers = []string{"MySQL", "Postgres", "MSSQL", "SqLite"} // 0=MySQL, 1=Postgres, 2=MSSQL, 3=SqLite

//var DatabaseData = &DatabaseDetails{
//	Providers: []string{"mysql", "postgres", "mssql", "sqlite"}, // 0=MySQL, 1=Postgres, 2=MSSQL, 3=SqLite
//}

func InitAssetData() {

}

type AssetContainer struct {
	AD []AssetDetails
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

type DatabaseDetails struct {
	Providers []string
}
