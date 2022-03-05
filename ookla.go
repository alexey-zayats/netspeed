package netspeed

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/showwin/speedtest-go/speedtest"
	"time"
)

var (
	ErrOoklaFetchUser    = fmt.Errorf("unable fetch user info")
	ErrOoklaFetchServers = fmt.Errorf("unable fetch servers")
	ErrOoklaTestUpload   = fmt.Errorf("unable measure upload speed")
	ErrOoklaTestDownload = fmt.Errorf("unable measure download speed")
)

type ooklaProvider struct{}

// Measure perform network speed measurements based on Speedtest.net
func (n *ooklaProvider) measure() (Result, error) {

	var (
		err     error
		sp      *speedtest.Speedtest
		user    *speedtest.User
		servers speedtest.Servers
		t       time.Time
	)

	sp = speedtest.New()

	// fetch user info
	user, err = sp.FetchUserInfo()
	if err != nil {
		return nil, errors.Wrap(err, ErrOoklaFetchUser.Error())
	}

	// fetch servers list
	servers, err = sp.FetchServers(user)
	if err != nil {
		return nil, errors.Wrap(err, ErrOoklaFetchServers.Error())
	}

	t = time.Now()

	r := result{provider: KindOokla}

	for _, server := range servers {

		// perform download test
		if err = server.DownloadTest(false); err != nil {
			return nil, errors.Wrap(err, ErrOoklaTestDownload.Error())
		}

		// perform upload test
		if err = server.UploadTest(false); err != nil {
			return nil, errors.Wrap(err, ErrOoklaTestUpload.Error())
		}

		// fill in result
		r.download = server.DLSpeed
		r.upload = server.ULSpeed

		// break after the first server
		break
	}

	// calculate measurement latency
	r.latency = time.Since(t)

	return r, nil
}
