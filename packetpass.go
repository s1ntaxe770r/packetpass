package main

import (
	"flag"

	"fmt"
	"os"

	"github.com/iovisor/gobpf/bcc"
)

import "C"

/*
#cgo CFLAGS: -I/usr/include/bcc/compat
#cgo LDFLAGS: -lbcc
#include <bcc/bcc_common.h>
#include <bcc/libbpf.h>
void perf_reader_free(void *ptr);
*/

const module_source string = `
#define KBUILD_MODNAME "packetpass"
#include <linux/if_ether.h>
#include <linux/if_packet.h>
#include <linux/in.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <uapi/linux/bpf.h>
int tcpfilter(struct CTXTYPE *ctx) {
	void *data = (void *)(long)ctx->data;
	void *data_end = (void *)(long)ctx->data_end;
	struct ethhdr *eth = data;
	if ((void*)eth + sizeof(*eth) <= data_end) {
	  	struct iphdr *ip = data + sizeof(*eth);
	  	if ((void*)ip + sizeof(*ip) <= data_end) {
		
			//Checking if the protocol with TCP Protocal
			if (ip->protocol == IPPROTO_TCP) {
				struct tcphdr *tcp = (void*)ip + sizeof(*ip);
				if ((void*)tcp + sizeof(*tcp) <= data_end) {
					//Checking if the destination port matches with the specified port
					if (tcp->dest != ntohs(PORT)) {
						//Drop Packet
						return RETURNCODE;
					}
				}
			}
	  	}
	}
	return XDP_PASS;
}`

func main() {
	var port string
	var iface string
	flag.StringVar(&iface, "interface", "lo", "interface to listen on ")
	flag.StringVar(&port, "port", "4040", "port to drop traffic on")
	flag.Parse()
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	retcode := "XDP_DROP"
	ctxtype := "xdp_md"

	module := bcc.NewModule(module_source, []string{
		"-w",
		"-DRETURNCODE=" + retcode,
		"-DCTXTYPE=" + ctxtype,
		"-DPORT=" + port,
	})
	defer module.Close()

	fn, err := module.Load("packetpass", C.BPF_PROG_TYPE_XDP, 1, 65536)

	if err != nil {
		fmt.Errorf("unable to load xdp progtam %s", err.Error())
		os.Exit(1)
	}
	err = module.AttachXDP(iface, fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to attach xdp prog: %v\n", err)
		os.Exit(1)
	}

	defer func() {
		if err := module.RemoveXDP(iface); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to remove XDP from %s: %v\n", iface, err)
		}
	}()

	fmt.Printf("Accepting connections on %s", port)

}
