# golang_pub_sub

 - [x] https://www.youtube.com/watch?v=blaYXmmA8zw

## 文章參考

[simple-publish-subscribe-pattern-in-golang](https://blog.wu-boy.com/2022/04/simple-publish-subscribe-pattern-in-golang/) 

## pub/sub 架構

```mermaid
graph RL
    hub(Hub)
    subscriber1(subscriber)
    subscriber2(subscriber)
    subscriber3(subscriber)
    subscriber1 -- subscribe --> hub
    subscriber2 -- subscribe --> hub
    subscriber3 -- subscribe --> hub
```

## message flow

```mermaid
graph LR
    message(message)
    hub(Hub)
    message -- publish --> hub
    subscriber1(subscriber)
    subscriber2(subscriber)
    subscriber3(subscriber)
    hub -- dispatch --> subscriber1
    hub -- dispatch --> subscriber2
    hub -- dispatch --> subscriber3
```

## unsubscribe flow

```mermaid
graph LR
    message(message)
    hub(Hub)
    message -- publish --> hub
    subscriber1(subscriber)
    subscriber2(subscriber)
    subscriber3(subscriber)
    hub -. unsubscribe .- subscriber1
    hub -- dispatch --> subscriber2
    hub -- dispatch --> subscriber3
```