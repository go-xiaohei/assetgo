package main

import (
	"errors"
	"os"

	"github.com/go-xiaohei/assetgo"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Assetgo"
	app.Version = "0.1.0"
	app.Usage = "Assetgo embed static assets to go file"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "pkg, p",
			Usage: "pkg set `package` name of output go file",
		},
		cli.StringFlag{
			Name:  "out, o",
			Value: "asset.go",
			Usage: "pkg set output go `file`",
		},
	}
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			return errors.New("no input directory")
		}

		var (
			w   = new(assetgo.Writer)
			err error
		)
		if err = w.WritePackage(c.String("pkg")); err != nil {
			return err
		}
		if err = w.WriteImport(); err != nil {
			return err
		}
		if err = w.WriteInitBegin(); err != nil {
			return err
		}

		for _, dir := range c.Args() {
			if err := assetgo.WalkDirectory(w, dir); err != nil {
				return err
			}
		}

		if err = w.WriteInitEnd(); err != nil {
			return err
		}
		return w.ToFile("asset.go")
	}
	app.Run(os.Args)
}
