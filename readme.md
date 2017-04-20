#Benchmark

#### Go will auto use goroutines to improve performance.

    BruceMacAir:webbench-1.5 bruce$ ./webbench -t 10 -c 50 'http://127.0.0.1:8787/qrcode?token=4yPqRtS3KHJyF9o5XEiM93FGJqafpd55vEiYOa&message=%E5%B0%BC%E7%8E%9B%E7%9A%84%E5%82%BB%E9%80%BC%E6%88%BF%E4%B8%9C%E5%95%8A%E6%88%BF%E4%B8%9C%E9%98%BF%E9%A3%9E%E6%87%82%E5%95%8A%E6%88%BF%E4%B8%9C%E9%98%BF%E9%A3%9E%E6%87%82%E5%95%8A%E6%88%BF%E4%B8%9C%E5%95%8A%E6%88%BF%E4%B8%9C%E5%95%8A'
    Webbench - Simple Web Benchmark 1.5
    Copyright (c) Radim Kolar 1997-2004, GPL Open Source Software.
    
    Benchmarking: GET http://127.0.0.1:8787/qrcode?token=4yPqRtS3KHJyF9o5XEiM93FGJqafpd55vEiYOa&message=%E5%B0%BC%E7%8E%9B%E7%9A%84%E5%82%BB%E9%80%BC%E6%88%BF%E4%B8%9C%E5%95%8A%E6%88%BF%E4%B8%9C%E9%98%BF%E9%A3%9E%E6%87%82%E5%95%8A%E6%88%BF%E4%B8%9C%E9%98%BF%E9%A3%9E%E6%87%82%E5%95%8A%E6%88%BF%E4%B8%9C%E5%95%8A%E6%88%BF%E4%B8%9C%E5%95%8A
    50 clients, running 10 sec.
    
    Speed=6336 pages/min, 120172 bytes/sec.
    Requests: 1056 susceed, 0 failed.
    BruceMacAir:webbench-1.5 bruce$ ./webbench -t 10 -c 50 'http://127.0.0.1:8787/qrcode?token=4yPqRtS3KHJyF9o5XEiM93FGJqafpd55vEiYOa&message=%E5%B0%BC%E7%'
    Webbench - Simple Web Benchmark 1.5
    Copyright (c) Radim Kolar 1997-2004, GPL Open Source Software.
    
    Benchmarking: GET http://127.0.0.1:8787/qrcode?token=4yPqRtS3KHJyF9o5XEiM93FGJqafpd55vEiYOa&message=%E5%B0%BC%E7%
    50 clients, running 10 sec.
    
    Speed=97950 pages/min, 272627 bytes/sec.
    Requests: 16325 susceed, 0 failed.
    
    
   