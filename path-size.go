package code

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// функция для подсчёта размера файлов в папке с учётом всех флагов
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	// получаем информацию о пути
	fileInfo, _ := os.Lstat(path)
	// проверяем, что путь это директория
	if fileInfo.IsDir() {
		if !all && strings.HasPrefix(fileInfo.Name(), ".") {
			return "", fmt.Errorf("the folder %s is hidden, enable the flag -a", fileInfo.Name())
		}
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
				totalSize += GetSize(file, all)
			}
			return FormatSize(totalSize, human), nil
			// если флаг recursive == true
		} else if recursive {
			// получаем размер файлов во всех вложенных папках
			err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
				// проверяем что флаг all == false и игнорируем все скрытые файлы и папки
				if !all {
					// проверяем что объект это папка и если скрытая то игнорируем
					if d.IsDir() && strings.HasPrefix(d.Name(), ".") {
						return fs.SkipDir
					} else {
						totalSize += GetSize(d, all)
					}
					// если флаг all == true учитываем все файлы во всех папках
				} else if all {
					// проверяем что объект это файл
					if !d.IsDir() {
						totalSize += GetSize(d, all)
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
		return "", fmt.Errorf("the file %s is hidden, enable the flag -a", fileInfo.Name())
	}
	// если флаг all == true возвращаем размер любого файла
	return FormatSize(fileInfo.Size(), human), nil
}

// функция для получения размера файла с учётом фильтра скрытых файлов (флаг -a)
func GetSize(filePath os.DirEntry, all bool) int64 {
	var fileSize int64
	// проверяем что переданный путь это файл
	if !filePath.IsDir() {
		// проверяем что флаг all == false и если файл скрытый возвращаем нулевой размер
		if !all && strings.HasPrefix(filePath.Name(), ".") {
			fileSize = 0
			// показывет список файлов и их размер, с выбранными фильтрами (использовать для отладки)
			// fmt.Println(filePath, fileSize)
		} else {
			info, _ := filePath.Info()
			fileSize = info.Size()
			// показывет список файлов и их размер, с выбранными фильтрами (использовать для отладки)
			// fmt.Println(filePath, fileSize)
		}
	}
	return fileSize
}

// функция конвертирования размера в человекочитаемый формат
func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}
	var msg_size string
	switch {
	case size < 1024:
		msg_size = fmt.Sprintf("%dB", size)
	case size >= 1024 && size < 1048576:
		msg_size = fmt.Sprintf("%.1fKB", float64(size)/1024.0)
	case size >= 1048576 && size < 1073741824:
		msg_size = fmt.Sprintf("%.1fMB", float64(size)/1024.0/1024.0)
	case size >= 1073741824 && size < 1099511627776:
		msg_size = fmt.Sprintf("%.1fGB", float64(size)/1024.0/1024.0/1024.0)
	default:
		msg_size = fmt.Sprintf("%.1fTB", float64(size)/1024.0/1024.0/1024.0/1024.0)
	}
	return msg_size
}
