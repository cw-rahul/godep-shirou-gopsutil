package main

import (
	"github.com/ContinuumLLC/platform-asset-plugin/src/dal"
	"github.com/ContinuumLLC/platform-asset-plugin/src/msgl"
	"github.com/ContinuumLLC/platform-asset-plugin/src/services"
	"github.com/ContinuumLLC/platform-common-lib/src/clar"
	"github.com/ContinuumLLC/platform-common-lib/src/env"
	cjson "github.com/ContinuumLLC/platform-common-lib/src/json"
	"github.com/ContinuumLLC/platform-common-lib/src/plugin/protocol/http"
	"github.com/ContinuumLLC/platform-common-lib/src/pluginUtils"
	"github.com/ContinuumLLC/platform-common-lib/src/procParser"
)

type factory struct {
	clar.ServiceInitFactoryImpl
	dal.AssetDalFactoryImpl
	dal.ConfigDalFactoryImpl

	services.AssetServiceFactoryImpl
	services.ConfigServiceFactoryImpl

	pluginUtils.StandardIOReaderImpl
	pluginUtils.StandardIOWriterImpl
	cjson.FactoryJSONImpl
	env.FactoryEnvImpl
	procParser.ParserFactoryImpl

	msgl.AssetListenerFactory
	msgl.ProcessAssetFactoryImpl

	http.ServerHTTPFactory
}
