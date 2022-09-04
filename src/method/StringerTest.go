package method

import "fmt"

type IPAddr [4]byte

func (p IPAddr) String() string {
	var ans string
	for n, i := range p {
		ans += string(i)
		if n < len(p)-1 {
			ans += "."
		}
	}
	return fmt.Sprintln(ans)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
