Design a messaging queue supporting publisher/subscriber model. It should also support following operations:

1. It should support *multiple topics* where message can be published.
2. Publisher should be able to publish message to particular topic.
3. Subscribers should be able to subscribe to a topic.

4. Whenever a message is published to a topic, all subscribers subscribed to that topic should receive message.

5. We should be able to reset offset for a subscriber. -> meaning, subscriber would start reading again from that offset.
	- It is like replaying the messages.

6. Subscribers should be able to run in parallel.

Video links:
https://www.youtube.com/watch?v=4BEzgPlLKTo

Java solution: https://github.com/anomaly2104/low-level-design-messaging-queue-pub-sub
