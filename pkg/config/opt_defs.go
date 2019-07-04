package config

var flagsOpts = []flagOpt{
	{
		optName:         FLAG_KEY_SERVER_HOST,
		optDefaultValue: "0.0.0.0",
		optUsage:        "server listen host",
	},
	{
		optName:         FLAG_KEY_SERVER_PORT,
		optDefaultValue: 8081,
		optUsage:        "server listen port",
	},
	{
		optName:         FLAG_KEY_GIN_MODE,
		optDefaultValue: "debug",
		optUsage:        "gin mode",
	},
	{
		optName:         FLAG_KEY_LOG_LEVEL,
		optDefaultValue: "info",
		optUsage:        "log level",
	},
}
