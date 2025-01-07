package datasets

import (
	"fmt"
	"io"
	"time"

	units "github.com/docker/go-units"
	"github.com/gosuri/uiprogress"
)

func byteCountSI(b int64) string {
	return units.HumanSizeWithPrecision(float64(b), 3)
}

func byteCountPerSecSI(b int64, seconds float64) string {
	return units.HumanSizeWithPrecision(float64(b)/seconds, 3) + "/s"
}

type copyProgressBar struct {
	io.Writer
	total     int64 // Total # of bytes written
	bar       *uiprogress.Bar
	startTime time.Time
	filename  string
	status    string
}

func (progressBar *copyProgressBar) Write(p []byte) (n int, err error) {
	n, err = progressBar.Writer.Write(p)
	progressBar.total += int64(n)
	if err == nil {
		progressBar.status = byteCountPerSecSI(progressBar.total, time.Since(progressBar.startTime).Seconds())
	}
	return
}

func (progressBar *copyProgressBar) Copy(dest io.Writer, src io.Reader) (n int64, err error) {
	if !argNoProgress && progressBar.bar == nil {
		progressBar.bar = progress.AddBar(1)
		progressBar.bar.LeftEnd = ' '
		progressBar.bar.RightEnd = ' '
		progressBar.bar.Width = 2
		progressBar.bar.Width = 2
		progressBar.bar.PrependFunc(func(b *uiprogress.Bar) string {
			return "Downloading: " + progressBar.filename
		})
		progressBar.bar.AppendFunc(func(b *uiprogress.Bar) string {
			return fmt.Sprintf("%s %s", byteCountSI(progressBar.total), progressBar.status)
		})
	}
	progressBar.status = "connecting"
	progressBar.total = 0
	progressBar.startTime = time.Now()
	progressBar.Writer = dest
	n, err = io.Copy(progressBar, src)
	progressBar.status = "done"
	return
}
