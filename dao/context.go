package dao

import (
	"sync"
	"gopkg.in/mgo.v2"
	"github.com/shvetsovau/gohttpserver/config"
)

type context struct {
	Session *mgo.Session
	DB      *mgo.Database
}

var instance *context
var once sync.Once

func (c *context) init() error {
	session, err := mgo.Dial(config.AppArgs.Host)
	if err != nil {
		return err
	}

	c.Session = session
	c.Session.SetMode(mgo.Monotonic, true)
	c.DB = c.Session.DB(config.AppArgs.DB)

	return nil
}

func getContext() *context {
	once.Do(func() {
		instance = &context{}
		if err := instance.init(); err != nil {
			panic(err)
		}
	})

	if err := instance.Session.Ping(); err != nil {
		instance.Session.Refresh()
	}
	return instance
}