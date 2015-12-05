package app

import (
	"errors"

	"github.com/Alienero/IamServer/config"

	"github.com/golang/glog"
)

func InitServer() error {
	if len(config.Config.Apps) == 0 {
		return errors.New("empty app list.")
	}
	for n, application := range config.Config.Apps {
		if application.RTMP != nil {
			glog.Infof("Load RTMP serve:%v", n)
			// TODO: start rtmp publisher server.s
		} else {
			// should throws a panic.
		}
		// IM & HTTP use one port, by default.
		if application.HTTP_FLV != nil {
			glog.Infof("Load HTTP-FLV serve:%v", n)
			// TODO
		}
		if application.IM != nil {
			glog.Infof("Load IM serve:%v", n)
			// TODO
		}
	}
	return nil
}
