package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {

	inputPath := ""
	out, _ := os.Getwd()
	flag.StringVar(&inputPath, "input", "", "path to the source")
	flag.StringVar(&out, "output", out, "output file")
	flag.Parse()

	filename := getFilename(inputPath)
	out += "\\" + filename + ".wav"
	err := ffmpeg.Input(inputPath).OverWriteOutput().Output(out, ffmpeg.KwArgs{
		"format": "wav",
	}).Audio().Run()
	if err != nil {
		log.Fatalln(err.Error())
	}

}

func getFilename(path string) string {
	filename := strings.Split(filepath.Base(path), ".")[0]
	return strings.ReplaceAll(filename, " ", "")
}
