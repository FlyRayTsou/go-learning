# 18

- タイマー：決まった時間の後に通知を受ける。決まった時間にコールバックをアプリケーションに返すもの。
- カウンター：現在時刻相当を取得する。現在時刻やそれ相当（プロセスやシステムの起動からの経過時間なども含む）を返すもの。
- タイマーやカウンターは時間を扱うサブシステムにとって不可欠な構成要素です。
- ハードウェアタイマーを効率よく利用するために複雑仕組みになっています。

# 18.1

## 18.1.go
```
rayt@RaydeMacBook-Pro go-learning % go run ch18_timer/18.1.go 
2023-11-19 08:48:06.562902 +0900 JST m=+0.000114080
```

## リアルタイムクロック(Real-time clock, RTC)
- ハードウェア
- パソコンを起動する時にシステムクロックはリアルタイムクロックを使って、初期化する
  - 現在の時刻を保持する。電源を切っても消えない。
  - 携帯なども同じか
    - https://www.reddit.com/r/apple/comments/7bassg/how_does_the_iphones_internal_clock_system_work/
- https://en.wikipedia.org/wiki/Real-time_clock

## システムクロック
- ソフトウェア
- OSが管理しているカウンター
- OSはハードウェアのタイマーを設定し、 一定間隔で割り込みがかかるようにします。
  - The operating system sets up the hardware timer to generate interrupts at regular intervals.
- この割り込みを受けて、システムクロックを更新したり、現在実行中のプ ロセスが持つ残りのタイムスライスを減らしたり、必要に応じてタスクの切り替えを 行ったりします。
  - In response to this interrupt, we update the system clock, reduce the remaining time slice of the currently running process, and perform task switching as needed.
  - リアルタイムクロックを使って、システムを更新することか？
- 電源が切れると消えてしまう
- 初期化する時にリアルタイムクロックを使う

## タイムスタンプカウンター(TSC、Time Stamp Counter)
- ハードウェア
- CPU内蔵のカウンター
- ナノ秒単位でカウントできる
- マルチコアCPUの場合、コアごろに結果にずれが生じることがある

## 各種タイマーデバイス
- ハードウェア
- 指定した間隔でハードウェア割り込みを発生するタイマーデバイス
- プログラマブルインターバルタイマー(programmable interval timer)
  - 別の電路のようみたい？
  - プログラマブル・インターバル・タイマー（PIT）チップ（Intel 8253/8254）は、基本的に発振子、プリスケーラー、および3つの独立した周波数分割器から構成されています。各周波数分割器には出力があり、これを使用してタイマーが外部回路（たとえば、IRQ 0）を制御できます。
  - https://en.wikipedia.org/wiki/Programmable_interval_timer
- APIC Timer(Advanced Programmable Interrupt Controller)
  - CPU内蔵の割り込みコントローラを使う
  - ローカルAPICタイマーの大きな利点は、それが各CPUコアにハードウェアで接続されていることです。これは、プログラマブルインターバルタイマーとは異なり、別の回路ではないです。そのため、リソース管理の必要がなく、事が簡単になります。欠点は、それがCPUの周波数（複数の中からの1つ）で振動していることで、マシンごとに異なります。一方、PITは標準の周波数（1,193,182 Hz）を使用しています。これを利用するには、それが何回の割り込み/秒に対応しているかを知っている必要があります。
  - https://wiki.osdev.org/APIC_Timer

## Linux
- タイマー割り込みの間隔は1秒あたり250回となっています。
- 割り込みのたびに、jiffiesというカウンターの変数が増えます。
  - 1が増加することを「1 Tick」と呼びます。
  - 1 Tickは固定時間ではなく、コンピューター上で観察可能な最小の時間間隔
  - ゲームのアプリケーション、1/60sは1 Tick

## 分解能
- 分解能より小さい間隔で分解できない。

## Japanese
- 抱える（かかえる）
- 割り込み：interrupt

# 18.2

## NTP(Network Time Protocol)
- NTPで正しい時間を取得になる。リアルタイム時刻。
- パケット交換による遅延時間が変動するネットワーク上のコンピュータシステム間で時刻同期させるための通信プロトコルである。
- OSI参照モデル(第7層 - アプリケーション層)

## モノトニック(monotonic time)
- 時刻調整をせずに起動後の経過時間などの尺度で管理され、決して巻き戻らない時刻が使われる。
- タイマー待ちで正確に「100ミリ秒測定したい」の場合、モノトニックを使う必要がある。
- Monotonic time is a measurement of the time since the system was powered on and is maintained by the kernel. Note that the monotonic time may not always reset to zero on reboot and the behavior during sleep/suspend has not yet been defined. (Fuchsia)
  - モノトニックタイムは、システムが起動してからの時間の測定で、カーネルによって維持されています。再起動時にモノトニックタイムが常にゼロにリセットされない場合があり、また、スリープ/サスペンド中の挙動はまだ定義されていません。

## ウォールクロック時間(wall clock time)
- 理解：本当の世界の時間

## CPU時間
- CPUが消費した時間です。
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

# 18.3

- OSの中の仕組みは複雑ですが、Go言語のランタイムと時間関係の接点はシンプルです。

# 18.3.1 runtime.now()

- 現在時刻を取得してくる関数
- ソフトウェア世界一番細かい粒度はナノ秒の単位です。10^−9 s
- 1GHzのパルス幅は1ナノ秒
- 現状のCPUのピークのクロック周波数が4GHzや5GHz
- オーバーヘッドを考慮すると、ナノよりも良い精度が出ないのは想像に難くないでしょうか。
  - Considering the overhead, it's not hard to imagine that it won't have better accuracy than nano.

## Windows

- 7ffe0000番地から1キロバイト:SharedUserDataという読み込み専用のデータ領域
  - 先頭から20バイト：SystemTimeというシステム時間が保存される。Windowsカーネルが100ナノ秒で更新。Goはこのアドレスを参照し、現在時刻を取得。
- +0x014 SystemTime       : _KSYSTEM_TIME
  - http://uninformed.org/index.cgi?v=2&a=2&p=15

## Linux

- clock_gettime()
- gettimeofday()
- Go --最初--> clock_gettime() --利用できない場合--> gettimeofday()
- vDSO(仮想ELF動的の共有オブジェクト)：システムオーバーヘッドがなく。
- Note
  - Regist hardware timer's event handler first
  - update_vsyscall() could be called regularly
    - Stores the current timestamp in a static variable inside the vDSO library

## Mac

- First, RDTSC (CPU Assembly)でCPU Timestamp Counter
- Second, gettimeofday()システムコール

# 18.3.2 runtime.semasleep()

- タイマー処理を使用した時に最終的に呼び出される関数
- マルチスレッドの共有資源の管理。セマフォの仕組み
- 通常、セマフォを使うときは. 正常ケースで資源管理を獲得し、期待どおりにいかなかったケースでタイムアウトさせますが、Goのタイマーでは逆にタイムアウトの機構を使って処理待ちを行っています。
  - Usually, when using semaphores, you acquire resource management in normal cases and timeout in cases where things don't go as expected. However, with Go timers, processing waits are performed using the timeout mechanism in the opposite way.
  - タイムアウトの仕組みでタイマーを作ること？
- すべてのタイマーがruntime.timersというグローバル変数に格納される。
  - いつ呼ばれるか、時間の順番にソートされて格納。
  - [50s, 124s, 40s] --- 現在時刻との差をタイムアウトに設定 ---> [124s, 40s, 45s] ---> [40s, 45s, 119s] --> [45s, 119s, 35s]
  - Queueみたい？
- 現在もっとも早く起動するタイマーが10秒後に設定されているとして、新たに5 秒後にタイマーを設定するとします。すでに10秒後のタイマーによるスリープが実 行中です。このときは、セマフォを通常どおりruntime .semawai tで 獲得して早 く待ちを終わらせます。その後、タイマーのリストに新しい5秒後のタイマーを挿入 して、この先頭になったタイマーを開始します。
  - The current timer is set to go off in 10 seconds as the fastest one. Now, let's say we want to set a new timer for 5 seconds. Meanwhile, the sleep caused by the 10-second timer is already in progress. In this situation, we acquire the semaphore as usual with runtime.semawait to quickly finish waiting. After that, we insert the new 5-second timer into the timer list and start the timer that now becomes the head.

# 18.4
# 18.4.1
- 時間(time.Duration)
- 時刻(time.Time)は日時情報
- show time.kitchen and other format
  - Click time.Kitchen
- モノトニック時刻による補正機能
  - t.Add(), t.Sub()
  - ウォールクロック：t.AddDate(), t.Round(), t.Truncate()
- 18.4.1.go

# 18.4.2.go

# 18.4.3-1.go

# 18.4.3-2.go

# 18.5.go

# Epoch

- 32bit -> 2038years
- 64bit -> 1400億years