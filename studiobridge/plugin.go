package studiobridge

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/vinegarhq/vinegar/wine"
)

//go:generate go run genplugin.go

//go:embed vinegarbridge.rbxm
var plugin []byte

var portLoaderCode = "plugin:SetSetting(\"VinegarBridge_ServerPort\", %d)"

var pluginDirPath = "Local/Roblox/Plugins"

func getPluginDir(pfx *wine.Prefix) (string, error) {
	appdataDir, err := pfx.AppDataDir()

	if (err != nil) {
		return "", err
	}

	target := path.Join(appdataDir, pluginDirPath)

	if err = os.Mkdir(target, os.ModePerm); err != nil && !errors.Is(err, fs.ErrExist) {
		return "", err
	}

	return target, nil
}

func Install(pfx *wine.Prefix, port int) error {
	target, err := getPluginDir(pfx)

	if err != nil {
		return err
	}

	if err := Remove(pfx); err != nil {
		return err
	}



	f, err := os.OpenFile(path.Join(target, "vinegarbridgeportshim.lua"), os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}

	if _, err := io.WriteString(f, fmt.Sprintf(portLoaderCode, port)); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	f, err = os.OpenFile(path.Join(target, "vinegarbridge.rbxm"), os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}

	if _, err := f.Write(plugin); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func Remove(pfx *wine.Prefix) error {
	target, err := getPluginDir(pfx)

	if err != nil {
		return err
	}

	if err = os.Remove(path.Join(target, "vinegarbridgeportshim.lua")); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}

	if err = os.Remove(path.Join(target, "vinegarbridge.rbxm")); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}

	return nil
}