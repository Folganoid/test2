package config

func Setup() map[string]string {

	setup := make(map[string]string)

	// mysql
	//setup["db_name"] = "go"
	//setup["db_user"] = "fg"
	//setup["db_pass"] = "56195619"
	//setup["db_prefix"] = ""

	//postgres
	//setup["db_host"] = "postgres"
	//setup["db_port"] = "5432"
	//setup["db_user"] = "postgres"
	//setup["db_pass"] = "postgres"
	//setup["db_name"] = "d8ha2ongbrinnm"
	//setup["db_prefix"] = "test."
	//setup["frontPath"] = "https://localhost:3000"

	setup["db_host"] = "mouse.db.elephantsql.com"
	setup["db_port"] = "5432"
	setup["db_user"] = "ppiogato"
	setup["db_pass"] = "5ghas9TvM1HjUejfRPsGNBgpC6ewY2EN"
	setup["db_name"] = "ppiogato"
	setup["db_prefix"] = "test."
	setup["frontPath"] = "https://velofg.herokuapp.com"

	return setup
}
