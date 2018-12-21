# RMCS_HUB-go
rewrite 'RMCS_HUB'(cpp originally) with golang

---
## VERSIONS
**date:&emsp;&emsp;&emsp;2018-12-01**  
**version:&emsp;&emsp;0.2.0**  
**description:**  
* Finish the framework of RMCS_HUB with golang.
* Use beego/orm in the operation of MYSQL.
* Use goroutine and channel instead of multi_thread and circular queue, which make the project more simple.
* Move **db.go** to **infra/** from **app/**.

**date:&emsp;&emsp;&emsp;2018-12-21**  
**version:&emsp;&emsp;1.0.0**  
**description:**  
* Use RegisterModel in beego/orm, hence the advanced query.
* Add mutex in rtu's connFd.
* Add **parser.go** in order to transfer between "float32", "uint32" and "byte".
* Take care of the difference between **copy()** and **append()**.