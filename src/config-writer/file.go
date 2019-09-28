package config_writer

import (
	"config-writer/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// 读取 虚拟网卡个数
type VFInfo struct {
	count  int
	master string
}

func (vf *VFInfo) ReadVFNum() (int, error) {
	sriovFile := fmt.Sprintf("/sys/class/net/%s/device/sriov_numvfs", vf.master)
	if _, err := os.Lstat(sriovFile); err != nil {
		return -1, fmt.Errorf("failed to open the sriov_numfs of device %q: %v", vf.master, err)
	}
	data, err := ioutil.ReadFile(sriovFile)
	if err != nil {
		return -1, fmt.Errorf("failed to read the sriov_numfs of device %q: %v", vf.master, err)
	}

	if len(data) == 0 {
		return -1, fmt.Errorf("no data in the file %q", sriovFile)
	}
	sriovNumfs := strings.TrimSpace(string(data))
	vfTotal, err := strconv.Atoi(sriovNumfs)
	if err != nil {
		return -1, fmt.Errorf("format num failed %v", err)
	}
	return vfTotal, nil
}

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
