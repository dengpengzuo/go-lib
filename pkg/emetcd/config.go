package emetcd

/*
在go.mod中替换,才能编译通过,主要是etcd使用的grpc是v1.26.0

replace github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5
replace github.com/coreos/etcd v3.3.25+incompatible => go.etcd.io/etcd v3.3.25+incompatible
replace google.golang.org/grpc v1.33.1 => google.golang.org/grpc v1.26.0
*/

/*
import (
	"github.com/coreos/etcd/embed"
)

func GenEmbedEtcd() (e *embed.Etcd, err error) {
	cfg := embed.NewConfig()
	cfg.Dir = "data"
	cfg.WalDir = "wal"
	e, err = embed.StartEtcd(cfg)
	return e, err
}
*/
