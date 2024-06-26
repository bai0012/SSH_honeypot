// data.go
package main

const CpuInfo = `processor   : 0
vendor_id   : GenuineIntel
cpu family  : 6
model       : 85
model name  : Intel(R) Xeon(R) Platinum 8175M CPU @ 2.50GHz
stepping    : 4
microcode   : 0x2007006
cpu MHz     : 2499.998
cache size  : 33792 KB
physical id : 0
siblings    : 2
core id     : 0
cpu cores   : 1
apicid      : 0
initial apicid  : 0
fpu     : yes
fpu_exception   : yes
cpuid level : 13
wp      : yes
flags       : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ss ht syscall nx pdpe1gb rdtscp lm constant_tsc rep_good nopl xtopology nonstop_tsc cpuid tsc_known_freq pni pclmulqdq ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand hypervisor lahf_lm abm 3dnowprefetch invpcid_single pti fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid mpx avx512f avx512dq rdseed adx smap clflushopt clwb avx512cd avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves ida arat pku ospke
bugs        : cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit mmio_stale_data retbleed gds
bogomips    : 4999.99
clflush size    : 64
cache_alignment : 64
address sizes   : 46 bits physical, 48 bits virtual
power management:

processor   : 1
vendor_id   : GenuineIntel
cpu family  : 6
model       : 85
model name  : Intel(R) Xeon(R) Platinum 8175M CPU @ 2.50GHz
stepping    : 4
microcode   : 0x2007006
cpu MHz     : 2499.998
cache size  : 33792 KB
physical id : 0
siblings    : 2
core id     : 0
cpu cores   : 1
apicid      : 1
initial apicid  : 1
fpu     : yes
fpu_exception   : yes
cpuid level : 13
wp      : yes
flags       : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ss ht syscall nx pdpe1gb rdtscp lm constant_tsc rep_good nopl xtopology nonstop_tsc cpuid tsc_known_freq pni pclmulqdq ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand hypervisor lahf_lm abm 3dnowprefetch invpcid_single pti fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid mpx avx512f avx512dq rdseed adx smap clflushopt clwb avx512cd avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves ida arat pku ospke
bugs        : cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit mmio_stale_data retbleed gds
bogomips    : 4999.99
clflush size    : 64
cache_alignment : 64
address sizes   : 46 bits physical, 48 bits virtual
power management:
`

const IfconfigOutput = `lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
	inet 127.0.0.1  netmask 255.0.0.0
	inet6 ::1  prefixlen 128  scopeid 0x10<host>
	loop  txqueuelen 1000  (本地环回)
	RX packets 102383991  bytes 42288672418 (42.2 GB)
	RX errors 0  dropped 0  overruns 0  frame 0
	TX packets 102383991  bytes 42288672418 (42.2 GB)
	TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0`
