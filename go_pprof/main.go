package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	cpuFile, err := os.Create("mem.profile")
	if err != nil {
		log.Printf("os file error:%s", err.Error())
		panic(err)
	}
	//pprof.StartCPUProfile(cpuFile)
	profile := pprof.Lookup("heap")
	//一直sleep
	i := 0
	strs := make([]string, 1000)
	for ; ; {
		//go func() {
		//	fmt.Println("hello")
		//}()
		strs = append(strs, make([]string, 1000)...)
		time.Sleep(time.Millisecond * 10)
		i++
		if i > 500 {
			break
		}
		profile.WriteTo(cpuFile, 1)
		//pprof.WriteHeapProfile(cpuFile)
	}
	//pprof.WriteHeapProfile(cpuFile)
	//defer pprof.StopCPUProfile()
	 cpuFile.Close()
}
