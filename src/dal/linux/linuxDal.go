package linux

import (
	"os/exec"
	"strings"
	"time"

	amodel "github.com/ContinuumLLC/platform-api-model/clients/model/Golang/resourceModel/asset"
	"github.com/ContinuumLLC/platform-asset-plugin/src/model"
	"github.com/ContinuumLLC/platform-common-lib/src/logging"
	"github.com/ContinuumLLC/platform-common-lib/src/procParser"
)

// AssetCollection Proc related constants
const (
	cAssetCreatedBy string = "/continuum/agent/plugin/asset"
	cAssetDataType  string = "assetCollection"
	cAssetProcPath  string = "/proc/meminfo"
)

//Error Codes
const (
// INVALIDAssetCollectionMEASURE = "Invalid measure :"
)

// AssetCollectionDalLinux ...
type AssetCollectionDalLinux struct {
	Factory model.AssetCollectionDalDependencies
	Logger  logging.Logger
}

//GetAssetData ...
func (dal *AssetCollectionDalLinux) GetAssetData() (*amodel.AssetCollection, error) {
	reader, err := dal.Factory.GetEnv().GetFileReader(cAssetProcPath)
	if err != nil {
		dal.Logger.Logf(logging.DEBUG, "Error in reading file %v", err)
		return nil, err
	}
	defer reader.Close()
	parser := dal.Factory.GetParser()
	cfg := procParser.Config{
		ParserMode:    procParser.ModeKeyValue,
		IgnoreNewLine: true,
	}
	data, err := parser.Parse(cfg, reader)
	if err != nil {
		dal.Logger.Logf(logging.DEBUG, "Error in parsing config %v", err)
		return nil, err
	}
	return translateAssetCollection{logger: dal.Logger}.translateAssetCollectionProcToModel(data), nil
}

type translateAssetCollection struct {
	logger logging.Logger
}

func (t translateAssetCollection) translateAssetCollectionProcToModel(data *procParser.Data) *amodel.AssetCollection {
	assetCollection := new(amodel.AssetCollection)
	assetCollection.CreateTimeUTC = time.Now().UTC()
	assetCollection.Type = cAssetDataType
	assetCollection.CreatedBy = cAssetCreatedBy
	assetCollection.BaseBoard = *(getBaseBoardInfo())
	assetCollection.Bios = *(getBiosInfo())
	assetCollection.Memory = *(getMemoryInfo())
	assetCollection.Os = *(getOsInfo())
	assetCollection.System = *(getSystemInfo())
	return assetCollection
}

func (t translateAssetCollection) getDataFromMap(key string, data *procParser.Data) int64 {
	if _, exists := data.Map[key]; !exists {
		return int64(0)
	}
	val, err := procParser.GetInt64(data.Map[key].Values[1])
	if err != nil {
		t.logger.Logf(logging.DEBUG, "Error in GetInt64 for key %s, Error : %v", key, err)
		return 0
	}
	return procParser.GetBytes(val, data.Map[key].Values[2])
}

type dalUtil struct{}

func (util dalUtil) execCommand(cmdName string) (string, error) {
	out, err := exec.Command("bash", "-c", cmdName).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
