package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "go-get-assets"
	app.Usage = "download the google drive assets contained in a JSON file"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "output, o", Value: "", Usage: "path to write rewritten JSON file"},
		cli.StringFlag{Name: "rewrite, r", Value: "", Usage: "stem to replace downloaded URLs in rewritten JSON file"},
		cli.StringFlag{Name: "needle, n", Value: gDriveURL, Usage: "needle to use to look for download URLs (defaults to Google drive stem)"},
	}

	app.Action = func(c *cli.Context) {
		var (
			in          = c.Args().Get(0)
			downloadDir = c.Args().Get(1)
			out         = c.String("output")
			rel         = c.String("rewrite")
			needle      = c.String("needle")
		)

		// println("in", in)
		// println("downloadDir", downloadDir)
		// println("out", out)
		// println("rel", rel)
		// println("needle", needle)

		goGetter := New(in, downloadDir, out, rel, needle)
		goGetter.Run()
	}

	app.Run(os.Args)
}
