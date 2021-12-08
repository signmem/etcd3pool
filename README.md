# etcd3pool  

> 包含 client, server 端程序  
> 可用于对 etcd API=3 进行压测   

# client  
> 用于对 etcd 进行压测  
> 根据配置文件生成 "testline" 个随机主机名字，并循环写入 etcd  
> 每个主机名 value 为当前写入的 timestamp     

# server  

> 每分钟获取 "/" 所有数据  
> 当 value (timestamp) 延后当前 3 分钟， 则删除 servername  


# metrics   

> 客户端可以通过下面 api 获取一些监控数据   
> curl http://localhost:7071/metrics (上一分钟总共成功写入 etcd 的量)   
> curl http://localhost:7071/now  (上一分钟到现在，成功写入 etcd 的量)  
> curl http://localhost:7071/health (返回 ok 即程序正常)  
