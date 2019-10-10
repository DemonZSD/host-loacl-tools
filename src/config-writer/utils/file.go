package utils

import (
	"config-writer/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/sirupsen/logrus"
	"config-writer/log"
)


var log *logrus.Logger

func init() {
	log = logger.GetLog()
}


// load json file to HostLocal
func ReadJsonFile(path string) (hostlocal *types.HostLocal, err error) {

	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteValue, &hostlocal)
	if err != nil {

		return nil, err
	}
	return hostlocal, nil
}
// write data to file
func WriteJsonToFile(path string, data interface{}) error{
	preData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Errorln(fmt.Sprintf("parse data failed: %v", err))
		return err
	}
	err = ioutil.WriteFile(path, preData, 0644)
	if err != nil {
		log.Errorln(fmt.Sprintf("write data to file %s failed: %v", path, err))
		return err
	}
	return nil
}
