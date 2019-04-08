package main

import (
	"os"
	"path/filepath"

	"github.com/TGSAN/freedom-routes/routes"
	"github.com/codegangsta/cli"
)

const VERSION = "1.1.0"

func genRoutes(templateNames []string, outputDir string) {
	ips := routes.FetchIps()

	if len(templateNames) == 1 {
		os.MkdirAll(outputDir, 0755)
		routes.Generate(templateNames[0], ips, outputDir)
	} else {
		for _, templateName := range templateNames {
			output := filepath.Join(outputDir, templateName)
			os.MkdirAll(output, 0755)
			routes.Generate(templateName, ips, output)
		}
	}
}

func main() {
	cli.AppHelpTemplate = `{{.Name}} v{{.Version}} - {{.Usage}}

USAGE:
   {{.Name}} {{if .Flags}}[options] {{end}}<template ..>

COMMANDS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`

	app := cli.NewApp()
	app.Name = "freedom-routes"
	app.Version = VERSION
	app.Usage = "generate routes-up.sh and route-down.sh"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: ".",
			Usage: "output directory",
		},
	}

	app.Action = func(c *cli.Context) {
		if c.Args().Present() {
			genRoutes(c.Args(), c.String("output"))
		} else {
			cli.ShowAppHelp(c)
			routes.PrintTemplatesPath()
		}
	}

	app.Run(os.Args)
}
