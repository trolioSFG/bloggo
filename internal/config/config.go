package blogconfig

import (
	"os"
	"encoding/json"
)

type Config struct {
	DbURL string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

func Read() Config {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(home + "/.gatorconfig.json")
	if err != nil {
		panic(err)
	}

	c := Config{}
	err = json.Unmarshal(data, &c)

	return c
}

func (c Config) SetUser(user string) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	c.CurrentUser = user
	data, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(home + "/.gatorconfig.json", data, 0644)
	if err != nil {
		panic(err)
	}
}

