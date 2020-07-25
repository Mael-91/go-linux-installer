package app

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Application Application
}

type Application struct {
	Name string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description string `json:"description"`
	Version string `json:"version"`
	Github Github `json:"github"`
}

type Github struct {
	Repo             string `json:"repo"`
	Release          string `json:"release"`
	Issue            string `json:"issue"`
	PullRequest      string `json:"pr"`
	ContributorsList ContributorsList `json:"contributors_list"`
}

type ContributorsList struct {
	Authors []Author `json:"authors"`
}

type Author struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Github string `json:"github"`
	Website string `json:"website"`
}

type Contributors struct {
	Contributors []Author `json:"contributors"`
}

func LoadApplicationConf() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file. %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode into a struct. %s", err)
	}

	return config
}
