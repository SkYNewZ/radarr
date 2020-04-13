package main

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"
)

var versionCommand *cli.Command = &cli.Command{
	Name:     "version",
	Usage:    "Print version",
	HideHelp: true,
	Action:   showVersion,
}

type version struct {
	GOOS    string `json:"GOOS"`
	GOARCH  string `json:"GOARCH"`
	Version string `json:"version"`
	Runtime string `json:"runtime"`
}

func (v *version) String() string {
	return fmt.Sprintf("radarr %s compiled with %v on %v/%v", v.Version, v.Runtime, v.GOOS, v.GOARCH)
}

var v *version = &version{
	GOARCH:  runtime.GOARCH,
	GOOS:    runtime.GOOS,
	Runtime: runtime.Version(),
	Version: Version,
}

func init() {
	app.Commands = append(app.Commands, versionCommand)
}

func showVersion(c *cli.Context) error {
	if c.Bool("json") {
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		return nil
	}

	fmt.Println(v)
	return nil
}
