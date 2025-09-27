package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"
)

func main() {
	resp, err := http.Get("http://localhost:4444/hello")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	_, _ = fmt.Fprintln(tw, "StatusCode\tStatusText\tBody\tServerTime")
	_, _ = fmt.Fprintf(
		tw,
		"%d\t%s\t%s\t%s",
		resp.StatusCode,
		resp.Status,
		body,
		time.Now().Format(time.DateTime),
	)

	_ = tw.Flush()
}
