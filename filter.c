#define KBUILD_MODNAME "packetpass"
#include <linux/if_ether.h>
#include <linux/if_packet.h>
#include <linux/in.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <uapi/linux/bpf.h>
int tcpfilter(struct xdp_md *ctx)
{
    void *data = (void *)(long)ctx->data;
    void *data_end = (void *)(long)ctx->data_end;
    struct ethhdr *eth = data;
    if ((void *)eth + sizeof(*eth) <= data_end)
    {
        struct iphdr *ip = data + sizeof(*eth);
        if ((void *)ip + sizeof(*ip) <= data_end)
        {

            
            if (ip->protocol == IPPROTO_TCP)
            {
                struct tcphdr *tcp = (void *)ip + sizeof(*ip);
                if ((void *)tcp + sizeof(*tcp) <= data_end)
                {
                    // Checking if the destination port matches with the specified port
                    if (tcp->dest == ntohs(4040))
                    {
                        // Drop Packet
                        return XDP_DROP;
                    }
                }
            }
        }
    }
    return XDP_PASS;
}