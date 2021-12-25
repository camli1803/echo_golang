package config

import (
    "fmt"
    "log"
    "time"

    mgo "gopkg.in/mgo.v2"
)

type Session struct {
    session *mgo.Session
}

func NewSession() *Session {
    session, err := mgo.DialWithInfo(&mgo.DialInfo{
        Addrs:    []string{AppConfig.MongoDBHost},
        Username: AppConfig.MongoDBUser,
        Password: MongoDBPwd,
        Timeout:  60 * time.Second,,
    })

    if err != nil {
        log.Fatalf("[ConnectDB]: %s\n", err)
    }
    session.SetMode(mgo.Monotonic, true)

    return &Session{session}
}

func (s *Session) Copy() *mgo.Session {
    return s.session.Copy()
}

func (s *Session) Close() {
    if s.session != nil {
        s.session.Close()
    }
}