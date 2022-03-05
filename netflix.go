package netspeed

import (
	"context"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"time"
)

type netflixProvider struct {
	start        time.Time
	upload       string
	uploadUnit   string
	download     string
	downloadUnit string
}

// Measure perform network speed measurements based on Netflix's fast.com
func (n *netflixProvider) measure() (Result, error) {

	// setup the context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// setup timeout
	ctx, cancel = context.WithTimeout(ctx, 180*time.Second)
	defer cancel()

	cmds := []chromedp.Action{
		emulation.SetUserAgentOverride(`chromedp/chromedp v0.6.10`),
		chromedp.Navigate(`https://fast.com`),
		chromedp.ScrollIntoView(`footer`),
		chromedp.WaitVisible(`#speed-value.succeeded`),
		chromedp.Text(`#speed-value.succeeded`, &n.download, chromedp.NodeVisible, chromedp.ByQuery),
		//chromedp.Text(`#speed-units.succeeded`, &n.downloadUnit, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Click(`#show-more-details-link`),
		chromedp.WaitVisible(`#upload-value.succeeded`),
		chromedp.Text(`#upload-value.succeeded`, &n.upload, chromedp.NodeVisible, chromedp.ByQuery),
		//chromedp.Text(`#upload-units.succeeded`, &n.uploadUnit, chromedp.NodeVisible, chromedp.ByQuery),
	}

	err := chromedp.Run(ctx, cmds...)

	download, err := strconv.ParseFloat(n.download, 64)
	if err != nil {
		return nil, errors.Wrap(err, "unable convert download")
	}

	upload, err := strconv.ParseFloat(n.upload, 64)
	if err != nil {
		return nil, errors.Wrap(err, "unable convert upload")
	}

	return result{provider: KindNetflix, download: download, upload: upload, latency: time.Since(n.start)}, nil
}
