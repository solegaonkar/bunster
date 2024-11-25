package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/yassinebenaid/godump"
	"github.com/yassinebenaid/ryuko/lexer"
	"github.com/yassinebenaid/ryuko/parser"
)

func main() {
	app := cli.Command{
		Name: "ryuko",
		Commands: []*cli.Command{
			{
				Name:        "tree",
				Description: "Print the script tree",
				Action:      treeCMD,
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "no-ansi", Aliases: []string{"n"}},
				},
			},
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func treeCMD(_ context.Context, cmd *cli.Command) error {
	filename := cmd.Args().Get(0)
	v, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	p := parser.New(
		lexer.New(v),
	)

	script := p.ParseScript()
	if p.Error != nil {
		return p.Error
	}

	var d godump.Dumper
	d.ShowPrimitiveNamedTypes = true

	if !cmd.Bool("no-ansi") {
		d.Theme = godump.DefaultTheme
	}

	return d.Println(script)
}
