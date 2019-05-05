package ping

import (
	"github.com/sparrc/go-ping"
	"fmt"
)

func Ping() *ping.Statistics {
	pinger, err := ping.NewPinger("61.135.169.121")
	if err != nil {
		panic(err)
	}
	pinger.SetPrivileged(false)

	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	return stats
}

func Ping2() *ping.Statistics {
	pinger, err := ping.NewPinger("www.baidu.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return nil
	}
	pinger.SetPrivileged(true)

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	pinger.Run()
	stats := pinger.Statistics() // get send/receive/rtt stats
	return stats
}
