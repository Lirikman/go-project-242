package main

import ("fmt"
	"log"
	"os"
	"context"
	"github.com/urfave/cli/v3"
	"errors"
	)

func GetSize(path string) (string, error) {
	// получаем информацию о пути
	fileInfo, err := os.Lstat(path)
		if err != nil {
			return "", errors.New("Invalid file path")
			}

	// проверяем, что путь это директория
	if fileInfo.IsDir() == true {
		files, err := os.ReadDir(path)
		if err != nil {
			return "", fmt.Errorf("Error: %w", err)
			}
		var totalSize int64
		for _, file := range files {
			fileInfo, _ := file.Info()
				totalSize += fileInfo.Size()
				return fmt.Sprintf("%d	%s", totalSize, path), nil
				}
	}
	
	// иначе, путь это файл
	size := fileInfo.Size()
	return fmt.Sprintf("%d	%s", size, path), nil
}



func main() {
    cmd := &cli.Command{
        Name:  "hexlet-path-size",
        Usage: "print size of a file or directory",
        Action: func(context.Context, *cli.Command) error {
            fmt.Println("Hello from Hexlet!")
            return nil
        },
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
    
    // Проверка расчёта размера файла или папки 
    var path string
    fmt.Print("Введите путь к файлу или папке: ")
    fmt.Scanln(&path)
    fmt.Println(GetSize(path))
}
