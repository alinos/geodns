package main

import (
	"log"
	"net"
	"strings"
)

func getInterfaces() []string {

	var inter []string
	uniq := make(map[string]bool)

	for _, host := range strings.Split(*flaginter, ",") {
		ip, port, err := net.SplitHostPort(host)
		if err != nil {
			switch {
			case strings.Contains(err.Error(), "missing port in address"):
				// 127.0.0.1
				ip = host
			case strings.Contains(err.Error(), "too many colons in address") &&
				// [a:b::c]
				strings.LastIndex(host, "]") == len(host)-1:
				ip = host[1 : len(host)-1]
				port = ""
			case strings.Contains(err.Error(), "too many colons in address"):
				// a:b::c
				ip = host
				port = ""
			default:
				log.Fatalf("Could not parse %s: %s\n", host, err)
			}
		}
		if len(port) == 0 {
			port = *flagport
		}
		host = net.JoinHostPort(ip, port)
		if uniq[host] {
			continue
		}
		uniq[host] = true

		// default to the first interfaces
		// todo: skip 127.0.0.1 and ::1 ?

		if ip != "127.0.0.1" {
			if len(serverInfo.ID) == 0 {
				serverInfo.ID = ip
			}
			if len(serverInfo.IP) == 0 {
				serverInfo.IP = ip
			}
		}

		inter = append(inter, host)

	}

	return inter
}
