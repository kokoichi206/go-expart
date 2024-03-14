Event-Driven Architecture in Golang

## EDA: Event-Driven Architecture

- architecture that helps organizations to decouple microservices
- **3 uses**
  - Event notifications
  - Event-carried state transfer
    - asynchronous REST ?
  - Event sourcing
- **Core components**
  - Event
    - immutable fact
    - e.g.
      - signing up
      - payments
      - received for orders
  - Queue
    - also referred as
      - bus, channel, stream, topic, ...
    - Events are frequently (not always) organized in FIFO
    - **Message queues**
      - have a limited lifetime
        - lack of event retention
    - **Event streams**
      - with event retention
    - Event stores
      - typically NOT used for message communication
  - Producers
  - Consumers
- Benefits of EDA
  - compared to synchronous or point-to-point communication patterns
  - Resiliency
    - loosely coupled
    - event broker
  - Agility

## Patterns

- Domain-driven design (DDD)
  - about modeling a complex business idea
  - key components?
    - ubizuitous language
    - bounded contexts
      - Every bounded context has its own UL
- Domain-centric architectures
  - [The Onion Architecture](https://jeffreypalermo.com/2008/07/the-onion-architecture-part-1/)
  - [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
  - Hexagonal architecture
  - you might skip using them
    - if you keep your **services small** enough or **never have to deal with** migrating cloud providers or switching databases.
- CQRS:
  - Command and Query Responsibility Segregation
    - Command
      - Performs a mutation of the application state
    - Query
      - Returns application state to the caller
  - [CQRS Documents by Greg Young](https://cqrs.files.wordpress.com/2010/11/cqrs_documents.pdf)

## Designs

- [EventStorming](https://blog.kinto-technologies.com/posts/2022-11-01-eventstorming/)
  - システムの目的、前提、すでに決まっている事項
- BDD: Behavior-Driven Development
- ADR: 
  - examples:
    - https://github.com/PacktPublishing/Event-Driven-Architecture-in-Golang/blob/main/Chapter03/docs/ADL/0002-use-a-modular-monolith-architecture.md

## English

- Errata
  - 正誤表
- piracy
  - 海賊行為
  - 剽窃
    - plagiarism
- predominantly
  - 主に、大部分は
- conjunction
  - 接続詞
- fire-and-forget
  - 米国が誇る巡航ミサイルの性能の良さを説明する言葉として著名
  - 発射ボタンを押しさせえすれば、後は忘れていても標的に当たる
- Choreography
  - 振り付け
- interventions
  - 介入介入
- go south
  - 悪化する

sentences

- Our world is made up of events
- architectures can go south if they are followed rigidly

## Links

- https://github.com/PacktPublishing/Event-Driven-Architecture-in-Golang
