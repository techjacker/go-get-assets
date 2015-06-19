package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	// "path/filepath"
)

// $ go-get-assets <input-json> <download-dir> <output-json> <rewrite-urls-stem>
// make into command line arguments
// var (
// 	in          = filepath.Join(Cwd(), "fixtures", "cms.json")
// 	out         = filepath.Join(Cwd(), "fixtures", "cms.downloaded.json")
// 	downloadDir = filepath.Join(Cwd(), "src", "images")
// 	rel         = "/images"
// 	needle      = gDriveURL
// )

func main() {

	app := cli.NewApp()
	app.Name = "greet"
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

		goGetter := New(in, downloadDir, out, rel, needle)
		fmt.Printf("%q", goGetter)
		goGetter.Run()
	}

	app.Run(os.Args)
}
