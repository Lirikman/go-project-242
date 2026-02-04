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

// функция для подсчёта размера файлов в папке
func GetPathSize(path string, human, all, recursive bool) (string, error) {
	// получаем информацию о пути
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", errors.New("invalid file path")
	}
	// проверяем, что путь это директория
	if fileInfo.IsDir() {
		var totalSize int64
		// проверяем что флаг recursive == false
		if !recursive {
			// получаем список файлов и папок указанной папки
			files, err := os.ReadDir(path)
			if err != nil {
				return "", fmt.Errorf("error: %w", err)
			}
			// получаем размер всех файлов
			for _, file := range files {
				fileInfo, _ := file.Info()
				// проверяем что объект это файл
				if !fileInfo.IsDir() {
					// проверяем что флаг all == false и пропускаем скрытые файлы
					if !all && strings.HasPrefix(fileInfo.Name(), ".") {
						continue
						// иначе учитываем размер файла
					} else {
						totalSize += fileInfo.Size()
					}
				}
			}
			return FormatSize(totalSize, human), nil
			// если флаг recursive == true
		} else {
			// получаем размер файлов во всех вложенных папках
			err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
				// проверяем, что объект это файл
				if !d.IsDir() {
					info, _ := d.Info()
					// проверяем что флаг all == false и пропускаем скрытые файлы
					if !all && strings.HasPrefix(info.Name(), ".") {
						totalSize += 0
					} else {
						totalSize += info.Size()
					}
				}
				return nil
			})
			if err != nil {
				return "", fmt.Errorf("error: %w", err)
			}
			return FormatSize(totalSize, human), nil
		}
	}
	// иначе, путь это файл
	// проверяем что флаг all == false и что файл скрытый и возвращаем нулевой размер
	if !all && strings.HasPrefix(fileInfo.Name(), ".") {
		return "", nil
	}
	// если флаг all == true возвращаем размер любого файла
	return FormatSize(fileInfo.Size(), human), nil
}

// функция конвертирования размера в человекочитаемый формат
func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}
	var msg_size string
	switch {
	case size < 1000:
		msg_size = fmt.Sprintf("%dB", size)
	case size >= 1000 && size < 1000000:
		msg_size = fmt.Sprintf("%.1fKB", float64(size)/1000.0)
	case size >= 1000000 && size < 1000000000:
		msg_size = fmt.Sprintf("%.1fMB", float64(size)/1000.0/1000.0)
	default:
		msg_size = fmt.Sprintf("%.1fGB", float64(size)/1000.0/1000.0/1000.0)
	}
	return msg_size
}

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
			size, _ := GetPathSize(p, cmd.Bool("human"), cmd.Bool("all"), cmd.Bool("recursive"))
			fmt.Printf("%s\t%s\n", size, p)
			return nil
		},
	}

	if err := сmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
