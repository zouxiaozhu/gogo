package yiigo

func init() {
	loadEnv("env.toml")
}

func Bootstrap(mysql bool) error{
	if mysql {
		if err := initMySQL(); err != nil {
			return err
		}
	}

	return nil
}
