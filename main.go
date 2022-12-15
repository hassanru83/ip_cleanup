package main

import (
	"strings"
)

func main() {
	//fmt.Println("Hello Bangladesh")
	ips := []string{"192.168.1.1", "192.168.2", "192.168.3.1", "192.168.4"}
	ip_cleanup(ips)
}

func ip_cleanup(ips []string) ([]string, []string) {

	cleanips := make([]string, 0)
	uncleanips := make([]string, 0)

	for _, v := range ips {

		v = strings.TrimSpace(v)
		k := strings.Split(v, ".")
		if len(k) < 4 {
			v = strings.Join(k, ".")
			uncleanips = append(uncleanips, v)
		}

		if len(k) == 4 {
			v = strings.Join(k, ".")
			cleanips = append(cleanips, v)
		}
	}
	return cleanips, uncleanips
}
