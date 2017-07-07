// +build linux

package net

import (
	"errors"

	"github.com/ContinuumLLC/platform-api-model/clients/model/Golang/resourceModel/asset"
	"github.com/ContinuumLLC/platform-asset-plugin/src/model"
)

// Info returns network information for Linux
func Info() ([]asset.AssetNetwork, error) {
	return nil, errors.New(model.ErrNotImplemented)
}
