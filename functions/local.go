package functions

import (
	"os"

	"github.com/gravitl/netmaker/logger"
	"github.com/gravitl/netmaker/logic"
)

// FileExists - checks if file exists
func FileExists(f string) bool {
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// SetDNSDir - sets the dns directory of the system
func SetDNSDir() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = os.Stat(dir + "/config/dnsconfig")
	if os.IsNotExist(err) {
		os.Mkdir(dir+"/config/dnsconfig", 0744)
	} else if err != nil {
		logger.Log(0, "couldnt find or create /config/dnsconfig")
		return err
	}
	_, err = os.Stat(dir + "/config/dnsconfig/Corefile")
	if os.IsNotExist(err) {
		err = logic.SetCorefile(".")
		if err != nil {
			logger.Log(0, err.Error())
		}
	}
	_, err = os.Stat(dir + "/config/dnsconfig/netmaker.hosts")
	if os.IsNotExist(err) {
		_, err = os.Create(dir + "/config/dnsconfig/netmaker.hosts")
		if err != nil {
			logger.Log(0, err.Error())
		}
	}
	return nil
}
