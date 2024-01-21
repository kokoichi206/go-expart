## sec 1

- `performance = (accuracy*efficiency*speed)`
- Optimization often results in less readable code
- Efficiency with Algorithms, Performance with Data Structures
  - https://www.youtube.com/watch?v=fHNmRkzxHWs&ab_channel=CppCon
- The Key to Pragmatic Code Performance
  - to stop focusing strictly on the speed and latency of our code
  - **we should generally focus on efficiency**
    - not speed
  - why?
    - It is much harder to make efficient software slow.
    - Speed is more fragile.
    - Speed is less portable.

## sec 2

- Why go
  - strong ecosystem
  - unused import or variable causes build error
    - 結構珍しいかも
  - table tests
  - backward compatibility
  - go runtime
    - hardware or os portability
      - like jvm, clr ...
      - just-in-time(jit) compilation
  - goroutine: CSP
- tips?
  - `// BUG(who)`

## sec 3

DEFINE RAER: **Resource-Aware Efficiency Requirements**

- optimizing our software was not easy.
  - We are really bad at guessing which part of the program consumes the most resources and how much.
  - Paret Principle
- Understand Your Goals
  - But what does "enough" mean for us?
  - sweet spot
- I would encourage you to try defining a number
  - it would make your software much easier to assess!
- Efficiency-aware development flow
  - TFBO:
    - **test, fix, benchmark, optimize**
  - It is far, far **easier to make a correct program fast** than it is to make a fast program correct.

## english

- compromises
  - 妥協
- Don't get me wrong
- mitigate
  - 和らげる
- tame
  - 飼い慣らす
- deliberately
  - わざと
- with a grain of salt
  - 話半分に
- First of all, don't panic!
- at first hand
  - 一次資料から、直接に
