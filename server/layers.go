package server

import (
	"emperror.dev/errors"
	"github.com/apex/log"
	"os"

	"github.com/LittleBigBug/ptero-wings/config"
)

func EnsureLayerDataDirectoryExists() error {
	uid := config.Get().System.User.Uid
	gid := config.Get().System.User.Gid
	layerDir := config.Get().System.LayerDirectory

	if _, err := os.Lstat(layerDir); err != nil {
		if os.IsNotExist(err) {
			log.Debug("layers: creating root directory and setting permissions")
			if err := os.MkdirAll(layerDir, 0o700); err != nil {
				return errors.WithStack(err)
			}
			if err := os.Chown(layerDir, uid, gid); err != nil {
				log.WithField("error", err).Warn("layers: failed to chown layer data directory")
			}
		} else {
			return errors.WrapIf(err, "server: failed to stat server root directory")
		}
	}

	return nil
}
