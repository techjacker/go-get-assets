package main

import (
	"fmt"
)

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

	if err = g.r.Run(); err != nil {
		return fmt.Errorf("%q", err)
	}

	return err
}
