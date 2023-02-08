package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/lufia/iostat"
	"github.com/urfave/cli"
	"golang.org/x/term"
)

var (
	Revision              = "beta"
	headerIsPrinted bool  = false
	linePageCount   int64 = 0
	pageSize        int64 = 3
	disk            string
	count           int64 = -1
	wait            int64 = -1
	curr            int   = -1
	xtended         bool
	short           bool
	err             error
)
var MB_read_s, MB_wrtn_s, read_s, wrtn_s, T_read_s, T_wrtn_s, R_lat_ms, W_lat_ms, r_err, w_err, r_retr, w_retr float64

var previousDriveStats map[string]*iostat.DriveStats = make(map[string]*iostat.DriveStats)
var timeOfLastUpdate time.Time

func main() {
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("version=%s revision=%s\n", cCtx.App.Version, Revision)
	}

	// get height of terminal for header periodic printing
	if term.IsTerminal(0) {
		_, height, err := term.GetSize(0)
		if err != nil {
			return
		}
		// println("width:", width, "height:", height)
		pageSize = int64(height) / 2
	} else {
		println("not in a term")
	}

	app := &cli.App{
		Name:    "iiostat",
		Version: "v0.0.1",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "x",
				Usage:       "extended information",
				Destination: &xtended,
			},
			&cli.BoolFlag{
				Name:        "d",
				Usage:       "short",
				Destination: &short,
			},
		},
		ArgsUsage: string("wait count"),
		Action: func(cCtx *cli.Context) error {
			//print OS kernel version by different OS platform
			//for darwin
			if runtime.GOOS == "darwin" {
				fmt.Printf("OS kernel version is %s\n\n", runtime.GOOS)
			}
			//for linux
			if runtime.GOOS == "linux" {
				fmt.Printf("OS kernel version is %s\n\n", runtime.GOOS)
			}

			//check if arguments are numbers
			if cCtx.Args().Get(0) != "" {
				wait, err = strconv.ParseInt(cCtx.Args().Get(0), 10, 64)
				if err != nil {
					wait = 0
					count = 0
					disk = cCtx.Args().Get(0)
				}
			}

			//Parse the int64 argument
			if cCtx.Args().Get(1) != "" {
				count, err = strconv.ParseInt(cCtx.Args().Get(1), 10, 8)
				if err != nil {
					count = -1
					disk = cCtx.Args().Get(1)
				} else {
					count--
				}
			}

			//Filter by disk name in the argument 3
			if cCtx.Args().Get(2) != "" {
				disk = cCtx.Args().Get(2)
			}
			return nil
		},
	}
	app.Run(os.Args)

	if wait == -1 && count == -1 {
		count = 0
		wait = 1
	}

	for i := 0; i <= int(count)+1; i++ {
		iostatGetInfo()
		if wait > 0 && count == -1 {
			i = -1
		}
		//exclude the last iteration
		if i < int(count)+1 {
			time.Sleep(time.Duration(wait) * time.Second)
		}

		curr = i

	}

}

func iostatGetInfo() {

	dstats, err := iostat.ReadDriveStats()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	for i, dstat := range dstats {
		if previousDriveStats[dstat.Name] != nil {

			if disk != "" && disk != dstat.Name {
				continue
			}
			
			
			//calculate the difference between current and previous and speed change
			calculateDriveStats(dstat)

			if !headerIsPrinted {
				//Print header
				if short {
					fmt.Printf("\nDevice:            MB_read/s MB_wrtn/s  #_read/s  #_wrtn/s  T_read/ms  T_wrtn/ms\n\n")
				} else {
					fmt.Printf("\nDevice:            MB_read/s MB_wrtn/s  #_read/s  #_wrtn/s  T_read/ms  T_wrtn/ms  R_lat(ms)  W_lat(ms)  #_r_err    #_w_err    #_r_retr    #_w_retr\n\n")
				}

				headerIsPrinted = true
			}
		
			//make header is printed every pageSize lines
			if linePageCount > pageSize {
				headerIsPrinted = false
				linePageCount = 0
			}
			linePageCount++

			if short {
				fmt.Printf("%s\t         %9.2f %9.2f %9.2f %9.2f %9.2f %9.2f\n", dstat.Name, MB_read_s, MB_wrtn_s, read_s, wrtn_s, T_read_s, T_wrtn_s)
			} else {
				fmt.Printf("%s\t         %9.2f %9.2f %9.2f %9.2f %9.2f %9.2f %10.2f %10.2f %10.2f %10.2f %11.2f %11.2f\n", dstat.Name, MB_read_s, MB_wrtn_s, read_s, wrtn_s, T_read_s, T_wrtn_s, R_lat_ms, W_lat_ms, r_err, w_err, r_retr, w_retr)
			}
		}

		//save current dsstat to previousDriveStats
		previousDriveStats[dstat.Name] = dstat

		//print new line after last drive
		if i == len(dstats)-1 && curr < int(count)+1 {
			fmt.Printf("\n")
		}

	}

	//save current time
	timeOfLastUpdate = time.Now()
	dstats = nil

}

func calculateDriveStats(dstat *iostat.DriveStats) {

	MB_read_s = (float64(dstat.BytesRead-previousDriveStats[dstat.Name].BytesRead) / 1024 / 1024) / time.Since(timeOfLastUpdate).Seconds()
	MB_wrtn_s = (float64(dstat.BytesWritten-previousDriveStats[dstat.Name].BytesWritten) / 1024 / 1024) / time.Since(timeOfLastUpdate).Seconds()
	read_s = float64(dstat.NumRead-previousDriveStats[dstat.Name].NumRead) / time.Since(timeOfLastUpdate).Seconds()
	wrtn_s = float64(dstat.NumWrite-previousDriveStats[dstat.Name].NumWrite) / time.Since(timeOfLastUpdate).Seconds()
	T_read_s = float64(dstat.TotalReadTime-previousDriveStats[dstat.Name].TotalReadTime) / 1000000000
	T_wrtn_s = float64(dstat.TotalWriteTime-previousDriveStats[dstat.Name].TotalWriteTime) / 1000000000
	R_lat_ms = float64(dstat.ReadLatency - previousDriveStats[dstat.Name].ReadLatency)
	W_lat_ms = float64(dstat.WriteLatency - previousDriveStats[dstat.Name].WriteLatency)
	r_err = float64(dstat.ReadErrors - previousDriveStats[dstat.Name].ReadErrors)
	w_err = float64(dstat.WriteErrors - previousDriveStats[dstat.Name].WriteErrors)
	r_retr = float64(dstat.ReadRetries - previousDriveStats[dstat.Name].ReadRetries)
	w_retr = float64(dstat.WriteRetries - previousDriveStats[dstat.Name].WriteRetries)

}
