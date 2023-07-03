#include "Topic.h"
#include <boost/asio/post.hpp>

Topic::Topic(std::string topic_name): name(topic_name) {

}

std::string Topic::getName() const
{
    return name;
}

void Topic::addSubscriber(std::shared_ptr<ISubscriber> sub_ptr) {
    auto ts = std::make_shared<TopicSubscriber>(sub_ptr);
    std::lock_guard<std::mutex> lck(sub_mtx);
    if (topic_subscribers.find(ts) == topic_subscribers.end()){
        topic_subscribers.insert(ts);
    }
}

void Topic::removeSubscriber(std::shared_ptr<ISubscriber> sub_ptr){
    auto ts = std::make_shared<TopicSubscriber>(sub_ptr);
    std::lock_guard<std::mutex> lck(sub_mtx);
    /*
    The line std::lock_guard<std::mutex> lck(sub_mtx); is used to create a std::lock_guard object named lck, which is a type of lock wrapper in C++ that automatically acquires a lock on a given std::mutex object when constructed and releases the lock when the std::lock_guard object goes out of scope (i.e., when it is destroyed).
    This part constructs the std::lock_guard object lck and acquires the lock on the sub_mtx mutex at the same time.
    */
    auto it = topic_subscribers.find(ts);
    if(it != topic_subscribers.end()){
        topic_subscribers.erase(it);
    }
}

// By using a mutex (msg_mtx) to protect the messages vector from concurrent access, the function ensures that adding messages to the topic is thread-safe. 
// Additionally, by publishing the message asynchronously through publishMessages(), the function may allow subscribers to process the new message without blocking the main thread that called addMessage().
void Topic::addMessage(Message message, boost::asio::thread_pool &thread_pool)
{
    // Step 1: Lock the msg_mtx mutex to protect the messages vector from concurrent access. The std::lock_guard constructor locks the mutex upon creation and automatically unlocks it when the lck object goes out of scope. 
    {
        std::lock_guard<std::mutex> lck(msg_mtx);

        // Step 2: Add the new message to the end of the messages vector. Inside the locked scope, function adds the message safely because the mutex is locked during this operation.
        messages.push_back(message);
    }

    // Step 3: Call publishMessages() to publish the added message to subscribers asynchronously
    publishMessages(thread_pool);
}

void Topic::publishMessages(boost::asio::thread_pool &thread_pool) {
    boost::asio::post(thread_pool, [this]() {
        /*
        This line is a Boost.Asio function call to post a new task (lambda function) to the specified thread_pool. The lambda function, [this]() { ... }, captures the current object (this) by reference and represents the task that will be executed asynchronously in one of the threads from the thread pool.
        */
        // Step 1: Lock the sub_mtx and msg_mtx mutexes to protect the subscribers and messages from concurrent access
        std::scoped_lock lck(sub_mtx, msg_mtx);
        /*
        - protecting the critical sections
        Inside the lambda function, a std::scoped_lock is created. The std::scoped_lock is a multiple-mutex lock that locks all the provided mutexes (sub_mtx and msg_mtx) at the same time. This ensures that both the topic_subscribers set and the messages vector are protected from concurrent access during the message publishing process.
        */

        // Step 2: Iterate through all topic subscribers and publish messages to each subscriber
        for (auto it = topic_subscribers.begin(); it != topic_subscribers.end(); it++) {
            int offset = (*it)->getOffset(); // This offset represents the index of the last consumed message for that particular subscriber.

            // Step 3: Consume all unread messages by the current subscriber
            while (offset < messages.size()) {
                (*it)->getSubscriber()->consumeMessage(messages[offset]); //  The current subscriber's getSubscriber() function is called to retrieve the associated ISubscriber object. The consumeMessage() function of that ISubscriber object is then called with the corresponding Message from the messages vector at the current offset. This allows the subscriber to process the message.
                offset++;
                (*it)->incrementOffset();
            }
        }
    });
}



