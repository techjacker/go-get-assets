package main

import (
	"fmt"
)

const gDriveURL string = "https://drive.google.com/file/d"

func New(in, downloadDir, out, rel, needle string) *GoGetAsseter {
	return &GoGetAsseter{
		l: NewLister(needle, in),
		d: NewDownloader(downloadDir, rel),
		r: NewRenamer(needle, in, out, rel),
	}
}

type GoGetAsseter struct {
	l *Lister
	d *Downloader
	r *Renamer
}

func (g *GoGetAsseter) Run() error {

	var err error

	if err = g.l.Run(); err != nil {
		return fmt.Errorf("%v", err)
	}

	if err = g.d.Run(g.l.Assets); err != nil {
		return fmt.Errorf("%q", err)
	}
	fmt.Printf("%q", g.l.Assets)

	if g.r.Out != "" {
		if err = g.r.Run(); err != nil {
			return fmt.Errorf("%q", err)
		}
	}

	return err
}
