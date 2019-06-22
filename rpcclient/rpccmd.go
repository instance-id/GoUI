package rpcclient

const (
	RPC_IP_ADRESS = "127.0.0.1"
	RPC_PORT      = ":14555"
)

const (
	RPC_START     = "-start"
	RPC_RESTART   = "-restart"
	RPC_STOP      = "-stop"
	RPC_STATUS    = "-status"
	RPC_UPDATE_DB = "-update" //TODO: implement update on demand functionality

	//administrative db operations
	RPC_TRUNCATE_DB    = "-truncatedb"    //TODO: TRUNCATE all tables
	RPC_TRUNCATE_TABLE = "-truncatetable" //TODO: TRUNCATE or DROP a specific table
	RPC_INITIALIZE_DB  = "-initializedb"  //TODO:	CREATE all tables that dont exist yet
	RPC_DESTROY_DB     = "-destroydb"     //TODO: DROP all tables
)
