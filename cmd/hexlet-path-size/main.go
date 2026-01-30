package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

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
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "include hidden files and directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			n := cmd.NArg()
			if n == 0 {
				return fmt.Errorf("path not entered, operation not possible")
			}
			p := cmd.Args().Get(0)
			size, _ := GetSize(p, cmd.Bool("all"))
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
func GetSize(path string, all bool) (int64, error) {
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
			fileInfo, _ := file.Info()
			// проверяем что объект это файл
			if fileInfo.IsDir() == false {
				// проверяем что флаг all == false
				if all == false {
					// проверяем что файл скрыт и пропускаем его
					if strings.HasPrefix(fileInfo.Name(), ".") == true {
						continue
						// если не скрыт, то прибавляем к общему размеру
					} else {
						totalSize += fileInfo.Size()
					}
					// если флаг all == true учитываем все файлы
				} else {
					totalSize += fileInfo.Size()
				}
			}
		}
		return totalSize, nil
	}

	// иначе, путь это файл
	// проверяем что флаг all == false
	if all == false {
		// проверяем что файл скрытый и возвращаем нулевой размер
		if strings.HasPrefix(fileInfo.Name(), ".") == true {
			return 0, nil
		}
	}
	// если флаг all == true возвращаем размер любого файла
	return fileInfo.Size(), nil
}

// функция конвертирования размера в человекочитаемый формат
func FormatSize(size float64, human bool) string {
	if human == false {
		return fmt.Sprintf("%.0fB", size)
	}
	sizeMb := size / 1000000.0
	return fmt.Sprintf("%.1fMB", sizeMb)
}
