package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultOutputDirectory = "out"
)

func main() {
	xmlDir := flag.String("xmlDir", "dictionaries", "The path for xml dictionary files")
	outDir := flag.String("out", defaultOutputDirectory, "The path for output SQL files")
	flag.Parse()

	if *xmlDir == "" {
		slog.Error("xml directory should not be empty")
		return
	}

	if *outDir == "" {
		*outDir = defaultOutputDirectory
	}

	err := filepath.Walk(*xmlDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		slog.Info(fmt.Sprintf("processing %s", path))
		data, err := os.ReadFile(path)
		if err != nil {
			slog.Error("read file error: %s", err)
			return err
		}

		isList := true
		dict := Dictionary{}
		dicts := Dictionaries{}

		err = xml.Unmarshal(data, &dicts)
		if err != nil {
			slog.Error("could not unmarshal to list of dictionaries. Trying single dictionary", err)

			err = xml.Unmarshal(data, &dict)
			if err != nil {
				slog.Error("could not unmarshal to single dictionary", err)
				return err
			}

			isList = false
		}
		slog.Info("marshalled successfully")

		sql := ""
		if !isList {
			sql, err = dict.ToSQL()
			if err != nil {
				slog.Error("could not convert dictionary to sql", err)
				return err
			}
		}

		for _, dictionary := range dicts.Dictionaries {
			dictSQL, err := dictionary.ToSQL()
			if err != nil {
				slog.Error("could not convert dictionary to sql", err)
				return err
			}

			sql += dictSQL + ";\n"
		}

		err = os.MkdirAll(*outDir, os.ModePerm)
		if err != nil {
			slog.Error("could not create output directory", err)
			return err
		}

		fileName := fmt.Sprintf("%s/%s", *outDir, strings.ReplaceAll(info.Name(), "xml", "sql"))
		slog.Info("file name: %s", fileName)
		file, err := os.Create(fileName)
		if err != nil {
			slog.Error("could not create output file for "+path, err)
			return err
		}

		_, err = file.Write([]byte(sql))
		if err != nil {
			slog.Error("could not write to output file for "+path, err)
			return err
		}

		return nil
	})
	if err != nil {
		slog.Error("walk error: %s", err)
		return
	}

	slog.Info("done")

}
