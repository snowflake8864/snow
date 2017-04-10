package etcdclient

import (
    etcdclient "github.com/coreos/etcd/client"
    log "github.com/Sirupsen/logrus"
    "string"
    "os" 
)


const (
    DEFAULT_ETCD = "http://10.211.55.84"
)
var client etcdclient.Client
var machines []string
func init() {
    //get machines id
    machines = []string{DEFAULT_ETCD}
    if env := os.Getenv("ETCD_HOST"); env != "" {
        machines = strings.Split(env, ";")
    }
    // init config
    cfg := etcdclient.Config{
        Endpoints: machines,
        Transport: etcdclient.DefaultTransport,
    } 
    //create client
    c, err := etcdclient.New(cfg)
    if err != nil {
        log.Panic(err)
        return
    }
    client = c
}
//client API for etcd
func KeysAPI() etcdclient.KeysAPI {
    return etcdclient.NewKeysAPI(client)
}
