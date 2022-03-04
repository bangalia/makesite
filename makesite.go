package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	md "github.com/gomarkdown/markdown"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

type MdPage struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      template.HTML
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	fileFlag := flag.String("file", "first-post.txt", "Enter the name of the file to be converted")
	dirFlag := flag.String("dir", ".", "Enter the directory to find .txt files")
	mdFlag := flag.String("md", "test-1.md", "Enter the name of the .md file to be converted")
	flag.Parse()

	if isFlagPassed("file") {
		fileContents, _ := ioutil.ReadFile(*fileFlag)
		fileName := (*fileFlag)[:len(*fileFlag)-4]
		page := Page{
			TextFilePath: "./",
			TextFileName: fileName,
			HTMLPagePath: fileName + ".html",
			Content:      string(fileContents),
		}
		t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
		newFile, _ := os.Create(page.HTMLPagePath)
		t.Execute(newFile, page)
	}

	if isFlagPassed("dir") {
		files, _ := ioutil.ReadDir(*dirFlag)
		for _, file := range files {
			dirFileName := file.Name()

			if len(dirFileName) > 4 && dirFileName[len(dirFileName)-4:] == ".txt" {
				fmt.Println(dirFileName)
				fileName := dirFileName[:len(dirFileName)-4]
				fileContents, _ := ioutil.ReadFile(*dirFlag + "/" + dirFileName)
				page := Page{
					TextFilePath: *dirFlag,
					TextFileName: fileName,
					HTMLPagePath: *dirFlag + "/" + fileName + ".html",
					Content:      string(fileContents),
				}
				t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
				newFile, _ := os.Create(page.HTMLPagePath)
				t.Execute(newFile, page)
			}
		}
	}

	if isFlagPassed("md") {
		fileContents, _ := ioutil.ReadFile(*mdFlag)
		fileName := (*mdFlag)[:len(*mdFlag)-3]
		output := md.ToHTML(fileContents, nil, nil)

		page := MdPage{
			TextFilePath: "./",
			TextFileName: fileName,
			HTMLPagePath: fileName + ".html",
			Content:      template.HTML(output),
		}

		t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
		newFile, _ := os.Create(page.HTMLPagePath)
		t.Execute(newFile, page)
	}

}
