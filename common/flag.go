package common

import (
	"flag"
)

func Banner() {
	banner := `
   ___                              _    
  / _ \     ___  ___ _ __ __ _  ___| | __ 
 / /_\/____/ __|/ __| '__/ _` + "`" + ` |/ __| |/ /
/ /_\\_____\__ \ (__| | | (_| | (__|   <    
\____/     |___/\___|_|  \__,_|\___|_|\_\   
                     fscan version: 1.4.2
`
	print(banner)
}

func Flag(Info *HostInfo) {
	Banner()
	flag.StringVar(&Info.Host, "h", "", "IP address of the host you want to scan,for example: 192.168.11.11 | 192.168.11.11-255 | 192.168.11.11,192.168.11.12")
	flag.StringVar(&Info.HostFile, "hf", "", "host file, -hs ip.txt")
	flag.StringVar(&Info.Ports, "p", DefaultPorts, "Select a port,for example: 22 | 1-65535 | 22,80,3306")
	flag.StringVar(&Info.Command, "c", "", "exec command (ssh)")
	flag.IntVar(&Info.Threads, "t", 200, "Thread nums")
	flag.IntVar(&Info.IcmpThreads, "it", 11000, "Icmp Threads nums")
	flag.BoolVar(&Info.Isping, "np", false, "not to ping")
	flag.BoolVar(&Info.Ping, "ping", false, "using ping replace icmp")
	flag.BoolVar(&Info.IsSave, "no", false, "not to save output log")
	flag.StringVar(&Info.Domain, "domain", "", "smb domain")
	flag.StringVar(&Info.Username, "user", "", "username")
	flag.StringVar(&Info.Userfile, "userf", "", "username file")
	flag.StringVar(&Info.Password, "pwd", "", "password")
	flag.StringVar(&Info.Passfile, "pwdf", "", "password file")
	flag.StringVar(&Info.Outputfile, "o", "result.txt", "Outputfile")
	flag.Int64Var(&Info.Timeout, "time", 3, "Set timeout")
	flag.BoolVar(&Info.Debug, "debug", false, "debug mode will print more error info")
	flag.Int64Var(&Info.WebTimeout, "wt", 3, "Set web timeout")
	flag.StringVar(&Info.Scantype, "m", "all", "Select scan type ,as: -m ssh")
	flag.StringVar(&Info.RedisFile, "rf", "", "redis file to write sshkey file (as: -rf id_rsa.pub) ")
	flag.StringVar(&Info.RedisShell, "rs", "", "redis shell to write cron file (as: -rs 192.168.1.1:6666) ")

	flag.BoolVar(&Info.IsWebCan, "nopoc", false, "not to scan web vul")
	flag.StringVar(&Info.PocInfo.PocName, "pocname", "", "use the pocs these contain pocname, -pocname weblogic")
	flag.StringVar(&Info.PocInfo.Proxy, "proxy", "", "set poc proxy, -proxy http://127.0.0.1:8080")
	flag.IntVar(&Info.PocInfo.Num, "Num", 20, "poc rate")
	flag.Parse()
}
