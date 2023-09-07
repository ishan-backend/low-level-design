SDE-3 or above: experience with designing systems + knowing the behind knowledge of systems like cassandra etc. Grokking high level/advance + know about some papers.

SDE-1/2: machine coding, design patterns, coding and hld, experience questions.

**Problem Statement:**

Design a load balancer, where client comes and ask factory to give certain kind of load balancer.

* There can be 5-6 or more types of algorithms involved like:
round-robin,
least connection,
least response time, etc.
Read - 
  * [Article 1](https://medium.com/@leonardo5621_66451/building-a-load-balancer-in-go-1c68131dc0ef#:~:text=A%20load%20balancer%20is%20a,strategies%20to%20handle%20inbound%20requests.)
  * [Article2](https://dev.to/bmf_san/implement-a-load-balancer-in-golang-8gj#:~:text=There%20are%20different%20types%20of,least%20number%20of%20unprocessed%20requests.)

* Algorithms can be mentioned by client.

* Can we change midway? For a request, has to be weighted robin, other request needs least response time.
 -> We can create different load balancers for different services.

**Approach :**

* LB is an application of reverseProxy, using existing wheel.
* Thinking about backend server, what responsibility individual backend server has.
* What things would be configurable?
  * FE would pass strategy to use for a particular request in headers
  * Max attempts to connect to a backend host should be configurable.
  * Increasing backend servers can be optional
* Server Pool :- 
* Single Host Reverse Proxy (via stdlib): https://fideloper.com/golang-single-host-reverse-proxy
* To make your own reverse proxy (with load balancing): https://fideloper.com/go-http

