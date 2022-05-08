from bcc import BPF 

device = "lo"
b = BPF(src_file="filter.c")
print("loading tcp filter")
fn = b.load_func("tcpfilter", BPF.XDP)
print("loaded tcp filter, attaching to device...")
b.attach_xdp(device, fn, 0)
try:
  b.trace_print()
except KeyboardInterrupt:
  pass

b.remove_xdp(device, 0) 