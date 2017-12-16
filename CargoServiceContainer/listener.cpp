#include "listener.hpp"
#include <QDebug>

Listener::Listener(QObject *parent) : QThread(parent)
{

}

Listener::~Listener(){
    this->stop(); // Stop the loop
    this->terminate(); // Terminate the thread
    this->wait(); // Wait to terminate.
}

bool  Listener::hasEventNext(){
    QMutexLocker ml(&this->mutex);
    return !this->events.isEmpty();
}

Event  Listener::getNextEvent(){
    QMutexLocker ml(&this->mutex);
    Event evt = this->events.top();
    return evt;
}

void Listener::popEvent(){
    QMutexLocker ml(&this->mutex);
    this->events.pop();
}

void Listener::run(){
    while(this->isRuning){
        // Process event here.
        if(this->hasEventNext()){
            Event evt = getNextEvent();
            // Try to process the event.
            this->processEvent(evt);
            this->popEvent();
        }else{
            // wait for more event to cames in
            stackNotEmpty.wait(&this->mutex_);
        }
    }

    // Clear the stack.
    this->events.clear();
}

void Listener::onEvent(const Event& evt){
    QMutexLocker ml(&this->mutex);
    this->events.push(evt);
    this->stackNotEmpty.wakeAll();
}

void Listener::stop(){
    this->isRuning = false;
}
