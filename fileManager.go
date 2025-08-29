package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

type FileManager struct {
	appDirInfo    os.FileInfo
	dataDirName   string
	dataPath      string
	configDirName string
	configPath    string
	logger        Logger
}

/* MAXWELL. This code is broken. You probably don't remember how, but it is. Fix it. Don't use AI. You already did that once and you learned nothing.
You'll likely need to make a log system. Your only hint is that you're better off logging to a file rather than trying to get terminal logs.
Bonus points if you can get the logs to readout from the file to a terminal live.

- PM */

func NewFileManager() *FileManager {
	fm := &FileManager{}

	fm.logger = *NewLogger("FileManager")

	fm.dataDirName = "p2p-chat"
	fm.dataPath = filepath.Join(xdg.DataHome, fm.dataDirName)
	fm.configDirName = "p2p-chat"
	fm.configPath = filepath.Join(xdg.ConfigHome, fm.configDirName)

	err := fm.initializeAppDirs()
	if err != nil {
		log.Printf("Error creating app directories: %v", err)
	}

	dirInf, err := fm.checkAppDir(fm.dataPath)
	if err != nil {
		log.Printf("Error checking app directory: %v", err)
	}
	if dirInf != nil {
		fm.appDirInfo = dirInf
	}

	return fm
}

func (fm *FileManager) checkAppDir(path string) (os.FileInfo, error) {
	appDirInfo, err := os.Stat(path)
	if err == nil {
		return appDirInfo, nil
	}
	if os.IsNotExist(err) {
		return nil, nil
	}
	return nil, err
}

// Create the folders/files necessary for app to run then return the directory
func (fm *FileManager) initializeAppDirs() error {
	for _, path := range []string{fm.dataPath, fm.configPath} {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fm *FileManager) RequestDir() (result string) {
	result = "err"
	if fm.dataPath != "" {
		result = fm.dataPath
	}
	return
}

func (fm *FileManager) ReadAppFiles() string {
	log.Println("Home data directory:", xdg.DataHome)
	return xdg.DataHome
}

func (fm *FileManager) ReadAppConfig() string {
	log.Println("Config data directory:", xdg.ConfigHome)
	return xdg.ConfigHome
}
