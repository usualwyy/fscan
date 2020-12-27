package Plugins

import (
	"fmt"
	"github.com/shadow1ng/fscan/common"
	"net"
	"strconv"
	"sync"
	"time"
)

func ProbeHosts(host string, ports <-chan int, respondingHosts chan<- string, done chan<- bool, adjustedTimeout int64) {
	for port := range ports {
		con, err := net.DialTimeout("tcp4", fmt.Sprintf("%s:%d", host, port), time.Duration(adjustedTimeout)*time.Second)
		if err == nil {
			con.Close()
			address := host + ":" + strconv.Itoa(port)
			result := fmt.Sprintf("%s open", address)
			common.LogSuccess(result)
			respondingHosts <- address
		}
	}
	done <- true
}

func ScanAllports(address string, probePorts []int, threads int, adjustedTimeout int64) ([]string, error) {
	ports := make(chan int, 20)
	results := make(chan string)
	done := make(chan bool, threads)

	for worker := 0; worker < threads; worker++ {
		go ProbeHosts(address, ports, results, done, adjustedTimeout)
	}

	for _, port := range probePorts {
		ports <- port
	}
	close(ports)

	var responses = []string{}
	for {
		select {
		case found := <-results:
			responses = append(responses, found)
		case <-done:
			threads--
			if threads == 0 {
				return responses, nil
			}
		}
	}
}

func TCPportScan(hostslist []string, ports string, timeout int64) []string {
	var AliveAddress []string
	probePorts := common.ParsePort(ports)
	lm := 20
	if len(hostslist) > 5 && len(hostslist) <= 50 {
		lm = 40
	} else if len(hostslist) > 50 && len(hostslist) <= 100 {
		lm = 50
	} else if len(hostslist) > 100 && len(hostslist) <= 150 {
		lm = 60
	} else if len(hostslist) > 150 && len(hostslist) <= 200 {
		lm = 70
	} else if len(hostslist) > 200 {
		lm = 75
	}

	thread := 10
	if len(probePorts) > 500 && len(probePorts) <= 4000 {
		thread = len(probePorts) / 100
	} else if len(probePorts) > 4000 && len(probePorts) <= 6000 {
		thread = len(probePorts) / 200
	} else if len(probePorts) > 6000 && len(probePorts) <= 10000 {
		thread = len(probePorts) / 350
	} else if len(probePorts) > 10000 && len(probePorts) < 50000 {
		thread = len(probePorts) / 400
	} else if len(probePorts) >= 50000 && len(probePorts) <= 65535 {
		thread = len(probePorts) / 500
	}

	var wg sync.WaitGroup
	mutex := &sync.Mutex{}
	limiter := make(chan struct{}, lm)
	for _, host := range hostslist {
		wg.Add(1)
		limiter <- struct{}{}
		go func(host string) {
			defer wg.Done()
			if aliveAdd, err := ScanAllports(host, probePorts, thread, timeout); err == nil && len(aliveAdd) > 0 {
				mutex.Lock()
				AliveAddress = append(AliveAddress, aliveAdd...)
				mutex.Unlock()
			}
			<-limiter
		}(host)
	}
	wg.Wait()
	return AliveAddress
}
