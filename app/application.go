/*
Copyright © 2020 Mael Constantin <mael.constantin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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
