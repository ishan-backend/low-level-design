SDE-3 or above: experience with designing systems + knowing the behind knowledge of systems like cassandra etc. Grokking high level/advance + know about some papers.

SDE-1/2: machine coding, design patterns, coding and hld, experience questions.

**Problem Statement:**

Design a load balancer, where client comes and ask factory to give certain kind of load balancer.

* There can be 5-6 or more types of algorithms involved like:
round-robin,
least connection,
least response time, etc.
Read - https://medium.com/@leonardo5621_66451/building-a-load-balancer-in-go-1c68131dc0ef#:~:text=A%20load%20balancer%20is%20a,strategies%20to%20handle%20inbound%20requests.

* Algorithms can be mentioned by client.

* Can we change midway? For a request, has to be weighted robin, other request needs least response time.
 -> We can create different load balancers for different services.

**Approach :**

* What happens in load balancer? -> does not manipulate data, reroutes the request.
  * For e.g. Input (Request) with route -> Node (server to be routed to)

* 

