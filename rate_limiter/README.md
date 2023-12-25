https://blog.logrocket.com/rate-limiting-go-application/

* The Rate Limit design pattern is a strategy used in software design to *control the number of requests a user or service
  can make to a particular system or API within a given time frame*.

* This pattern is crucial for maintaining the stability and reliability of a system by preventing overuse or abuse,
  such as *denial-of-service attacks or resource hogging*.

* **Pros of using this pattern**:
  Enhanced Stability, Fair Resource Allocation, Increased Security (mitigate certain types of cyber-attacks, reduces latency and increases reliability)
* **Cons of using this pattern**:
  Potential for Blocking Legitimate Requests, Complex to implement - especially in distributed systems, Requires additional resources for tracking and enforcing limits.

Below article summarises 4 types of rate limiting algorithms:
https://levelup.gitconnected.com/rate-limit-design-pattern-in-golang-with-unit-tests-1-overview-four-algorithms-b6065c19a816

**Steps to implement**
```text
The following are steps on how to implement the rate limit pattern in Go.

1. Define Rate Limits
Decide on the number of requests allowed per user/IP and the time window, for example, 100 requests per hour.

2. Choose a Rate Limiting Algorithm
As I mentioned, there are four rate-limiting algorithms, please select one from Fixed Window, Sliding Window Log, Token Bucket, or Leaky Bucket.

3. Implement Rate Limiting Logic
Use middleware for HTTP servers and store request counts or timestamps in a fast, memory-based storage system.

4. Track Requests
Identify users or IPs for each incoming request and check the current count or timestamp against the defined limits.

5. Enforce Limits
If within the limit, proceed and update the count or timestamp. If exceeded, return an error response, typically HTTP 429, which means “Too Many Requests”. You can define a friendly response by yourself.


```

**Problem 1:** 
```text
A social media platform wants to limit the number of posts or comments a user can make per hour to prevent spam and 
encourage more thoughtful interactions. They decided to implement a fixed window rate-limiting algorithm 
where each user is allowed, for example, 10 comments per minute. Once a user reaches this limit, they cannot comment until the next minute begins.
```
- use an in-memory map to track user activity. If you like, you can use a more scalable and persistent storage solution like a database or Redis.