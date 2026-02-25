package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	сmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "recursive size of directories (default: false)",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit) (default: false)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "include hidden files and directories (default: false)",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			n := cmd.NArg()
			// проверка на ввод аргументов функции
			if n == 0 {
				return fmt.Errorf("path not entered, operation not possible")
			}
			p := cmd.Args().Get(0)
			// проверка на корректный ввод пути
			_, err := os.Stat(p)
			if err != nil {
				return fmt.Errorf("invalid path entered")
			}
			size, err := code.GetPathSize(p, cmd.Bool("recursive"), cmd.Bool("human"), cmd.Bool("all"))
			if err != nil {
				return fmt.Errorf("error: %w", err)
			}
			fmt.Printf("%s\t%s\n", size, p)
			return nil
		},
	}

	if err := сmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
