package worker

import (
	"context"
	"errors"
	"fmt"
	"github.com/cavaliercoder/grab"
	"gopkg.in/go-playground/pool.v3"
	"time"
)

var p pool.Pool

func init() {
	p = pool.NewLimited(10)
	//ali = client.NewAli()
}

//download 下载文件
func download(url string, path string) error {
	// create client
	c := grab.NewClient()
	req, _ := grab.NewRequest(path, url)

	//限速8m/s
	req.RateLimiter = NewLimiter(8 * 1024 * 1024)
	// start download
	//fmt.Printf("下载 %v中...\n", req.URL())
	resp := c.Do(req)

	//fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()

	// 最大下载时间限制
	tLimit := time.NewTimer(time.Duration(resp.Size/1024/1024*10*10)*time.Second + time.Minute)

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("%s , %.2f m/s(%.2f%%) \n",
				req.URL(),
				resp.BytesPerSecond()/1024/1024,
				100*resp.Progress(),
			)
		case <-tLimit.C:
			return errors.New("下载超时 " + req.URL().String())

		case <-resp.Done:
			// download is complete
			break Loop
		}

	}

	// check for errors
	if err := resp.Err(); err != nil {
		return err
	}

	fmt.Printf("下载完成 ,saved to %v \n", resp.Filename)
	return nil
}

// testRateLimiter is a naive rate limiter that limits throughput to r tokens
// per second. The total number of tokens issued is tracked as n.
type testRateLimiter struct {
	r, n int
}

func NewLimiter(r int) grab.RateLimiter {
	return &testRateLimiter{r: r}
}

func (c *testRateLimiter) WaitN(ctx context.Context, n int) (err error) {
	c.n += n
	time.Sleep(
		time.Duration(1.00 / float64(c.r) * float64(n) * float64(time.Second)))
	return
}
