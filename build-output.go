package nixpacks

import (
	"errors"
	"fmt"
	"strings"
)

type BuildOutput struct {
	Response      []byte
	IsBrokenImage bool
	BuildError    string
	ImageName     string
	Language      string
	Install       string
	Build         string
	Start         string
}

func (o *BuildOutput) Parse() error {
	//convert output to string
	if len(o.Response) == 0 {
		return errors.New("response is empty")
	}
	output := string(o.Response)
	r := strings.Split(output, "\n")

	var save bool
	var skipFirst bool

	for _, line := range r {
		fmt.Println("l:", line)
		if strings.HasPrefix(line, "║ setup") {
			fmt.Println("found setup: " + line)
			o.Language = strings.ReplaceAll(line, "║ setup", "")
			o.Language = strings.Split(o.Language, "│")[1]
			o.Language = strings.ReplaceAll(o.Language, "║", "")
			o.Language = strings.TrimSpace(o.Language)
		} else if strings.HasPrefix(line, "║ install") {
			o.Install = strings.ReplaceAll(line, "║ install", "")
			o.Install = strings.Split(o.Install, "│")[1]
			o.Install = strings.ReplaceAll(o.Install, "║", "")
			o.Install = strings.TrimSpace(o.Install)
		} else if strings.HasPrefix(line, "║ build") {
			o.Build = strings.ReplaceAll(line, "║ build", "")
			o.Build = strings.Split(o.Build, "│")[1]
			o.Build = strings.ReplaceAll(o.Build, "║", "")
			o.Build = strings.TrimSpace(o.Build)
		} else if strings.HasPrefix(line, "║ start") {
			o.Start = strings.ReplaceAll(line, "║ start", "")
			o.Start = strings.Split(o.Start, "│")[1]
			o.Start = strings.ReplaceAll(o.Start, "║", "")
			o.Start = strings.TrimSpace(o.Start)
		} else if strings.HasPrefix(line, "  docker run -it ") {
			o.ImageName = strings.ReplaceAll(line, "  docker run -it ", "")
		}

		if o.IsBrokenImage {
			if line == "------" {
				save = !save
				continue
			}
			if save {
				if !skipFirst {
					skipFirst = true
					continue
				}
				o.BuildError += line + "\n"
			}
		}
	}

	return nil
}
