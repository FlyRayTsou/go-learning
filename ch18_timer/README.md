# 18

# 18.1

## 18.1.go
```
rayt@RaydeMacBook-Pro go-learning % go run ch18_timer/18.1.go 
2023-11-19 08:48:06.562902 +0900 JST m=+0.000114080
```

## リアルタイムクロック(Real-time clock, RTC)

- パソコンを起動する時にシステムクロックはリアルタイムクロックを使って、初期化する

## Japanese
- 抱える（かかえる）

## Ref
- https://en.wikipedia.org/wiki/Real-time_clock
- https://www.reddit.com/r/apple/comments/7bassg/how_does_the_iphones_internal_clock_system_work/


# 18.2

## NTP(Network Time Protocol)
- パケット交換による遅延時間が変動するネットワーク上のコンピュータシステム間で時刻同期させるための通信プロトコルである。
- OSI参照モデル(第7層 - アプリケーション層)
- モノトニック(monotonic time)
  - Monotonic time is a measurement of the time since the system was powered on and is maintained by the kernel. Note that the monotonic time may not always reset to zero on reboot and the behavior during sleep/suspend has not yet been defined. (Fuchsia)

## ウォールクロック時間(wall clock time)
- 理解：本当の世界の時間

## CPU時間
- 1 core, 10%, 5mins(300s)。 1 * 300s * 10% = 30s
- 8 cores, 100%, 5mins。 8 * 5 * 100% = 40mins

## Japanese
- うるう秒: 閏秒
- 尺度（しゃくど）

## Ref
- https://ja.wikipedia.org/wiki/Network_Time_Protocol
- https://ja.wikipedia.org/wiki/OSI%E5%8F%82%E7%85%A7%E3%83%A2%E3%83%87%E3%83%AB
- https://fuchsia.dev/fuchsia-src/concepts/kernel/time/monotonic
- https://www.techtarget.com/whatis/definition/wall-time-real-world-time-or-wall-clock-time