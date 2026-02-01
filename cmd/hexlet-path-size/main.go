package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v3"
)

func main() {
	сmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
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
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "recursive size of directories (default: false)",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			n := cmd.NArg()
			if n == 0 {
				return fmt.Errorf("path not entered, operation not possible")
			}
			p := cmd.Args().Get(0)
			size, _ := GetSize(p, cmd.Bool("all"), cmd.Bool("recursive"))
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
func GetSize(path string, all bool, recursive bool) (int64, error) {
	// получаем информацию о пути
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, errors.New("invalid file path")
	}
	// проверяем, что путь это директория
	if fileInfo.IsDir() {
		var totalSize int64
		// проверяем что флаг recursive == false
		if !recursive {
			// получаем список файлов и папок указанной папки
			files, err := os.ReadDir(path)
			if err != nil {
				return 0, fmt.Errorf("error: %w", err)
			}
			// получаем размер всех файлов
			for _, file := range files {
				fileInfo, _ := file.Info()
				// проверяем что объект это файл
				if !fileInfo.IsDir() {
					// проверяем что флаг all == false
					if !all {
						// проверяем что файл скрыт и пропускаем его
						if strings.HasPrefix(fileInfo.Name(), ".") {
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
			// если флаг recursive == true
		} else {
			// получаем размер файлов во всех вложенных папках
			err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
				// проверяем, что объект это файл
				if !info.IsDir() {
					// проверяем что флаг all == false
					if !all {
						// проверяем что файл не является скрытым
						if !strings.HasPrefix(info.Name(), ".") {
							totalSize += info.Size()
						}
						// если флаг all == true учитываем все файлы
					} else {
						// если флаг all == true учитываем все файлы
						totalSize += info.Size()
					}
				}
				return nil
			})
			if err != nil {
				return 0, fmt.Errorf("error: %w", err)
			}
			return totalSize, nil
		}
	}
	// иначе, путь это файл
	// проверяем что флаг all == false
	if !all {
		// проверяем что файл скрытый и возвращаем нулевой размер
		if strings.HasPrefix(fileInfo.Name(), ".") {
			return 0, nil
		}
	}
	// если флаг all == true возвращаем размер любого файла
	return fileInfo.Size(), nil
}

// функция конвертирования размера в человекочитаемый формат
func FormatSize(size float64, human bool) string {
	if !human {
		return fmt.Sprintf("%.0fB", size)
	}
	sizeMb := size / 1000000.0
	return fmt.Sprintf("%.1fMB", sizeMb)
}
