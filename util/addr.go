package util

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

func GetAddrByHostAndPort(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func GetOutBoundIP() (ip string, err error) {
	// UDP 是无连接的，发送数据之前不需要建立连接
	conn, err := net.Dial("udp", "114.114.114.114:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println("GetOutBoundIP:" + localAddr.String())
	ipList := strings.Split(localAddr.String(), ":")
	if len(ipList) > 0 {
		ip = ipList[0]
		return
	}
	err = errors.New("no ip")
	return
}

// GetMacByIp 取的mac，全小写
func GetMacByIp(localIp string) string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("无法获取接口列表：", err)
		return ""
	}
	for _, iface := range interfaces {
		// 忽略回环接口和虚拟接口
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("无法获取接口地址：", err)
			continue
		}

		for _, addr := range addrs {
			// 检查地址类型是否为IP地址
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				//fmt.Println("MAC地址：", iface.HardwareAddr, addr)
				// 检查IP地址是否为IPv4
				if ipnet.IP.To4() != nil && strings.Contains(addr.String(), localIp) {
					return iface.HardwareAddr.String()
				}
			}
		}
	}
	return ""

}
