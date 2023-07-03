Design a **messaging queue** supporting publisher/subscriber model. It should also support following operations:

1. It should support *multiple topics* where message can be published.
2. Publisher should be able to publish message to particular topic.
3. Subscribers should be able to subscribe to a topic.

4. Whenever a message is published to a topic, all subscribers subscribed to that topic should receive message.

5. We should be able to reset offset for a subscriber. -> meaning, subscriber would start reading again from that offset.
	- It is like replaying the messages.

6. Subscribers should be able to run in parallel.



*Idea* - Design the whole system in such a way that you're handling the threading model correctly. Threads and inter-thread communication.

A queue has multiple topics. Publisher publishes a message to a particular topic.
All subscribers subscribed to a topic, will get the message.
Allow reset the offset for a subcriber(only), not other subscribers from same topic. In meantime, if new message gets published to topic, this subsriber won't receive the message until offset reaches end.
Make system performant, using threads, multiple threads should be able to run in parallel.

Video links:
https://www.youtube.com/watch?v=4BEzgPlLKTo

Java solution: https://github.com/anomaly2104/low-level-design-messaging-queue-pub-sub

Solution formulation:
- Create Queue.h
- Create ISubscriber.h, Subscriber.h
- Create implementations of above interfaces
- Create ThreadPool interface