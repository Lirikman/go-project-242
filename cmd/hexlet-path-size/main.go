package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	сmd := &cli.Command{
		Name:  "hexlet-path-size - print size of a file or directory",
		Usage: "hexlet-path-size [global options]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			n := cmd.NArg()
			if n == 0 {
				return fmt.Errorf("path not entered, operation not possible")
			}
			p := cmd.Args().Get(0)
			size, _ := GetSize(p)
			fSize := FormatSize(float64(size), cmd.Bool("human"))
			fmt.Printf("%s\t%s\n", fSize, p)
			return nil
		},
	}

	if err := сmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

// функция для подсчёта размера файлов в папке
func GetSize(path string) (int64, error) {
	// получаем информацию о пути
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, errors.New("invalid file path")
	}
	// проверяем, что путь это директория
	if fileInfo.IsDir() == true {
		files, err := os.ReadDir(path)
		if err != nil {
			return 0, fmt.Errorf("Error: %w", err)
		}
		var totalSize int64
		for _, file := range files {
			fileInfo, _ := file.Info() // проверяем что объект это файл
			if fileInfo.IsDir() == false {
				totalSize += fileInfo.Size()
			}
		}
		return totalSize, nil
	}

	// иначе, путь это файл
	size := fileInfo.Size()
	return size, nil
}

// функция конвертирования размера в человекочитаемый формат
func FormatSize(size float64, human bool) string {
	if human == false {
		return fmt.Sprintf("%.0fB", size)
	}
	sizeMb := size / 1000000.0
	return fmt.Sprintf("%.1fMB", sizeMb)
}
