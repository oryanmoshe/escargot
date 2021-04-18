package store

import (
	"github.com/oryanmoshe/escargot/internal/config"
	logrus "github.com/sirupsen/logrus"
)

type Store interface {
	GetName() string
	Whoop(string)
}

type store struct {
	id   string
	name string
	host string
	port int
	pass string
	log  *logrus.Entry
}

func FromID(id string) Store {
	params := config.GetConfigByStoreID(id)

	return New(params)
}

func New(config config.StoreConfig) Store {
	if config.Port == 0 {
		config.Port = 6379
	}

	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.TextFormatter{})

	log := logger.WithFields(logrus.Fields{
		"component": "store",
		"storeId":   config.ID,
	})

	s := store{
		id:   config.ID,
		name: config.Name,
		host: config.Host,
		port: config.Port,
		pass: config.Pass,
		log:  log,
	}

	return s
}

func (s store) GetName() string {
	return s.name
}

func (s store) Whoop(msg string) {
	s.log.Info(msg)
}
