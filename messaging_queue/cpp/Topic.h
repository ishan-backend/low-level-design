#pragma once
#include <string>
#include <mutex>
#include <vector>
#include "Message.h"
#include <boost/uuid/random_generator.hpp>
#include <set>
#include <memory>
#include "ISubscriber.h"
#include <boost/asio/thread_pool.hpp>
#include "TopicSubscriber.h"

class Topic {
    private:
        std::string name;
        std::mutex sub_mtx, msg_mtx;
        std::vector<Message> messages;
        static boost::uuids::random_generator rg;
        std::set<std::shared_ptr<TopicSubscriber>, TopicSubscriberCompare> topic_subscribers;

        void publishMessages(boost::asio::thread_pool &);
    public:
        Topic(std::string);
        std::string getName() const;
        void addSubscriber(std::shared_ptr<ISubscriber>);
        void removeSubscriber(std::shared_ptr<ISubscriber>);
        void addMessage(Message, boost::asio::thread_pool &);
};

struct TopicCompare {
    bool operator()(const std::shared_ptr<Topic> & tptr1, const std::shared_ptr<Topic> & tptr2) const {
        return tptr1->getName() < tptr2->getName();
    }
};

/*
The code provided defines a C++ class named `Topic`, which seems to represent a specific topic or channel in the context of a messaging system. The class has private member variables and public member functions to manage subscribers and messages associated with this topic.

Let's break down each part of the `Topic` class definition:

1. `std::string name;`: This member variable represents the name or identifier of the topic. Each `Topic` object is associated with a unique name.

2. `std::mutex sub_mtx, msg_mtx;`: These member variables are mutexes (mutual exclusions). They are used to synchronize access to shared data (in this case, the `topic_subscribers` and `messages` containers) among multiple threads. `sub_mtx` is likely used when adding or removing subscribers, and `msg_mtx` is used when adding messages.

3. `std::vector<Message> messages;`: This member variable is a vector that holds `Message` objects. It represents the collection of messages associated with this topic.

4. `std::set<std::shared_ptr<TopicSubscriber>, TopicSubscriberCompare> topic_subscribers;`: This member variable is a set of shared pointers to `TopicSubscriber` objects. It stores the subscribers that have subscribed to this topic.

5. `static boost::uuids::random_generator rg;`: This is a static member variable of type `boost::uuids::random_generator`. It is likely used to generate unique identifiers for messages or subscribers associated with this topic.

6. `void publishMessages(boost::asio::thread_pool &);`: This private member function is used to publish messages to the subscribers of this topic. It takes a reference to a `boost::asio::thread_pool` as a parameter, which is likely used to handle the message publishing asynchronously.

7. `Topic(std::string);`: This is the constructor of the `Topic` class. It takes a `std::string` parameter representing the name of the topic and initializes the `name` member variable.

8. `std::string getName() const;`: This member function is a getter method that returns the name of the topic.

9. `void addSubscriber(std::shared_ptr<ISubscriber>);`: This member function is used to add a subscriber to this topic. It takes a shared pointer to an `ISubscriber` object (representing a subscriber) as a parameter and adds it to the `topic_subscribers` set.

10. `void removeSubscriber(std::shared_ptr<ISubscriber>);`: This member function is used to remove a subscriber from this topic. It takes a shared pointer to an `ISubscriber` object as a parameter and removes it from the `topic_subscribers` set.

11. `void addMessage(Message, boost::asio::thread_pool &);`: This member function is used to add a `Message` to the `messages` vector of this topic. It takes a `Message` object as well as a reference to a `boost::asio::thread_pool` as parameters. The `boost::asio::thread_pool` is

 likely used to handle message publishing asynchronously.

12. `struct TopicCompare { ... };`: This is a custom comparison functor used for comparing shared pointers to `Topic` objects. It is likely used to sort or order topics based on their names.

In summary, the `Topic` class represents a messaging topic that can have multiple subscribers. It can store messages and notify its subscribers when new messages are published. The use of mutexes ensures thread safety when modifying the list of subscribers or adding messages. The `boost::asio::thread_pool` is likely used to manage asynchronous tasks, such as message publication and subscriber notification, in a separate thread pool to avoid blocking the main program's execution.

*/