package main

import (
	"awesomeProject1/YALIB"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	loc := "C:\\Program Files (x86)\\Steam\\steamapps\\common\\BorderlandsGOTYEnhanced"
	//loc := "C:\\Users\\yasha\\GolandProjects\\awesomeProject1\\test"
	var zap string
	var i int
	write, _ := os.Create("log.txt")
	defer write.Close()

	fmt.Println("что мы ищем? для выхода напишите \"z\" без кавычек ")
	fmt.Scanln(&zap)
	zap = strings.TrimSpace(zap)
	fmt.Printf("%q\n", zap)

	if zap == "z" {
		return
	}

	filepath.Walk(loc, func(path string, info os.FileInfo, err error) error { // тут Я написал анонимную функцию для бега по файлам, анонимную, чтобы была возможность передавать ей данные из вне
		if !info.IsDir() { // игнорируем папки
			file, errOp := os.Open(path) // открываем все файлы
			if errOp != nil {            // проверка ошибок и запись в лог
				fmt.Fprintln(write, errOp, path)
				fmt.Println(errOp, path)
			}
			datafile, errRed := ioutil.ReadAll(file) // читаем все файлы
			if errRed != nil {                       // проверка ошибок и запись в лог
				fmt.Fprintln(write, errRed, path)
				fmt.Println(errRed, path)
			}
			fmt.Fprintln(write, path)
			fmt.Println(path)
			kolVo := YALIB.IndaxRune(string(datafile), zap, write)
			if kolVo > -1 {
				i++
			}
			fmt.Println("количество нужных подстрок в файле =", kolVo)
			fmt.Fprintln(write, "количество нужных подстрок в файле =", kolVo)
		}
		return nil // не отдаем никаких ошибок в функцию выше, потому что она прекратит прохождение по файлам, а нам это не нужно
	})
	fmt.Println("найдено файлов с нужными данными:", i) // выводим в консольку количество найденных файлов, и если оно больше 0 то мы что-то нашли

}
