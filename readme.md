## make все поднимет прокачает и запустит бенчмарк
# без внешних запросов
```data_received..................: 2.7 GB   2.9 MB/s
data_sent......................: 1.8 GB   2.0 MB/s
dropped_iterations.............: 899969   967.223988/s
http_req_blocked...............: avg=2.57µs   min=0s         med=1µs   max=1.2s   p(90)=1µs   p(95)=3µs  
http_req_connecting............: avg=0ns      min=0s         med=0s    max=1.65ms p(90)=0s    p(95)=0s   
http_req_duration..............: avg=376.78µs min=15µs       med=159µs max=10.73s p(90)=547µs p(95)=931µs
{ expected_response:true }...: avg=376.78µs min=15µs       med=159µs max=10.73s p(90)=547µs p(95)=931µs
http_req_failed................: 0.00%    ✓ 0            ✗ 19337836
http_req_receiving.............: avg=70.89µs  min=-1074000ns med=5µs   max=10.71s p(90)=16µs  p(95)=33µs 
http_req_sending...............: avg=16.03µs  min=1µs        med=2µs   max=9.59s  p(90)=7µs   p(95)=14µs 
http_req_tls_handshaking.......: avg=0s       min=0s         med=0s    max=0s     p(90)=0s    p(95)=0s   
http_req_waiting...............: avg=289.85µs min=0s         med=146µs max=4.23s  p(90)=496µs p(95)=827µs
http_reqs......................: 19337836 20782.959021/s
iteration_duration.............: avg=9m3s     min=9m3s       med=9m3s  max=9m3s   p(90)=9m3s  p(95)=9m3s 
iterations.....................: 15       0.016121/s
vus............................: 9        min=9          max=15    
vus_max........................: 15       min=15         max=15
```
# с внешними запросами
```data_received..................: 576 kB 619 B/s
data_sent......................: 405 kB 436 B/s
dropped_iterations.............: 899986 967.72806/s
http_req_blocked...............: avg=13.82µs  min=0s    med=7µs   max=3.31ms  p(90)=14µs  p(95)=16µs    
http_req_connecting............: avg=1.08µs   min=0s    med=0s    max=734µs   p(90)=0s    p(95)=0s      
http_req_duration..............: avg=3.24s    min=1.15s med=2.37s max=16.79s  p(90)=5.35s p(95)=10.09s  
{ expected_response:true }...: avg=3.24s    min=1.15s med=2.37s max=16.79s  p(90)=5.35s p(95)=10.09s  
http_req_failed................: 0.00%  ✓ 0         ✗ 4295
http_req_receiving.............: avg=141.13µs min=8µs   med=94µs  max=8.63ms  p(90)=201µs p(95)=266.29µs
http_req_sending...............: avg=50.93µs  min=3µs   med=31µs  max=14.75ms p(90)=58µs  p(95)=69µs    
http_req_tls_handshaking.......: avg=0s       min=0s    med=0s    max=0s      p(90)=0s    p(95)=0s      
http_req_waiting...............: avg=3.24s    min=1.15s med=2.37s max=16.79s  p(90)=5.35s p(95)=10.09s  
http_reqs......................: 4295   4.618285/s
vus............................: 15     min=15      max=15
vus_max........................: 15     min=15      max=15
```