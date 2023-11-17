# Demo Asqnq Go Server
A demo of a Distributed Task Server using Go and Reddis. 

## High Level Architecture
* Client puts tasks on a queue
* Server pulls tasks off queues and starts a worker goroutine for each task
* Tasks are processed concurrently by multiple workers



