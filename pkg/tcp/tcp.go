package tcp

import (
	"context"
	"net"
	"time"
)

const keepaliveTime = 30 * time.Second // 初始：7,200,000 milliseconds

//func controlFunc(network, address string, c syscall.RawConn) error {
//	_, _ = fmt.Fprintf(os.Stdout, "create network:%s address:%s, server fd:%v\n", network, address, c)
//	return nil
//}

func Listen(network, address string) (net.Listener, error) {
	// ////////////////////////////////////////////////////////////////////////
	// @see:
	//   func (ln *TCPListener) accept() (*TCPConn, error) 中调用了 SetKeepAlive
	var lc net.ListenConfig
	lc.KeepAlive = keepaliveTime
	// lc.Control = controlFunc
	return lc.Listen(context.Background(), network, address)
	// ////////////////////////////////////////////////////////////////////////
	/*
		第二种方式
		ln, err := net.Listen(network, address)
		if err != nil {
			return nil, err
		}
		ln = &tcpKeepaliveListener{TCPListener: ln.(*net.TCPListener), KeepalivePeriod: keepaliveTime}
		return ln, nil
	*/
}

/*
int val = 1;
if (setsockopt(fd, SOL_SOCKET, SO_KEEPALIVE, &val, sizeof(val)) == -1) {
	printf("setsockopt SO_KEEPALIVE: %s", strerror(errno));
	return -1;
}

// Send first probe after `interval' seconds.
val = interval;
if (setsockopt(fd, IPPROTO_TCP, TCP_KEEPIDLE, &val, sizeof(val)) < 0) {
	printf("setsockopt TCP_KEEPIDLE: %s\n", strerror(errno));
	return -1;
}

// Send next probes after the specified interval. Note that we set the
// delay as interval / 3, as we send three probes before detecting
// an error (see the next setsockopt call).
val = interval/3;
if (setsockopt(fd, IPPROTO_TCP, TCP_KEEPINTVL, &val, sizeof(val)) < 0) {
	printf("setsockopt TCP_KEEPINTVL: %s\n", strerror(errno));
	return -1;
}

// Consider the socket in error state after three we send three ACK
// probes without getting a reply.
val = 3;
if (setsockopt(fd, IPPROTO_TCP, TCP_KEEPCNT, &val, sizeof(val)) < 0) {
	printf("setsockopt TCP_KEEPCNT: %s\n", strerror(errno));
	return -1;
}
*/
/*
// 第二种实现方式
type tcpKeepaliveListener struct {
	*net.TCPListener
	KeepalivePeriod time.Duration
}

func (ln *tcpKeepaliveListener) Accept() (net.Conn, error) {
	tc, err := ln.TCPListener.AcceptTCP()
	if err != nil {
		return nil, err
	}

	if err := tc.SetKeepAlive(true); err != nil {
		_ = tc.Close() // nolint:errcheck
		return nil, err
	}
	if ln.KeepalivePeriod > 0 {
		if err := tc.SetKeepAlivePeriod(ln.KeepalivePeriod); err != nil {
			_ = tc.Close() // nolint:errcheck
			return nil, err
		}
	}
	return tc, nil
}
*/
