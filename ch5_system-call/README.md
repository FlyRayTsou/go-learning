# Chapter 5 system call

## 5.1 システムコールとは何か

### 5.1.1 CPUの動作モード

  - [Process Vs. Thread | Difference Between Process and Thread](https://www.javatpoint.com/process-vs-thread)
- 動作モード
  - [Intel系CPUはリングプロテクション](https://ja.wikipedia.org/wiki/リングプロテクション)

### 5.1.2 システムコールでモードの壁を越える

- arm は RISC(Reduced Instruction Set Computer)、x86とx64 は CISC(Complex Instruction Set Computer)
  - https://www.wakuwakubank.com/posts/809-it-cpu-x86-x64-arm/

### 5.1.3 システムコールがないとどうなるか

- WebAssembly
  - https://zenn.dev/highgrenade/articles/3238e4b6c00de5
- WebAssembly System Interface
  - https://hacks.mozilla.org/2019/03/standardizing-wasi-a-webassembly-system-interface
- https://wasmer.io/
- https://docs.docker.com/desktop/wasm/

## 5.2 Go言語におけるシステムコールの実装

### 5.2.1

### 5.2.2

- open /usr/local/go/src/syscall/asm_darwin_amd64.s
- open /usr/local/go/src/syscall/asm_darwin_arm64.s
- entersyscall()は現在実行中のOSスレッドが時間のかかるシステムコールでブロックされていることを示すマークを付ける関数です
  - entersyscall() is a function that marks the currently running OS thread as blocked by a time-consuming system call.
- これらの関数を使うのは、スレッド作成という重い処理を必要になるまで行わないためです
  - The reason for using these functions is that they do not do the heavy lifting of thread creation until it is needed
- Go言語では、システムコールのブロックなどが原因で、実行しなければならないタスクが多くあるのに動けるスレッドが不足すると、OSに依頼して新しい作業用スレッドを作成します
  - In the Go language, when there are too many tasks to perform but not enough threads to run them, due to blocked system calls, etc., it asks the OS to create a new working thread.

### 5.2.3

- open /usr/local/go/src/syscall/syscall_linux.go
- open /usr/local/go/src/syscall/zsyscall_linux_amd64.go

### 5.2.4

- https://www.codeguru.com/windows/how-do-windows-nt-system-calls-really-work/
- DLL：Dynamic-link library
- POSIX系OS：Portable Operating System Interface
- open /usr/local/go/src/syscall/zsyscall_windows.go

## 5.3

- IEEE：Institute of Electrical and Electronics Engineers

## 5.4

### 5.4.1

- https://github.com/torvalds/linux/blob/v5.16/fs/read_write.c#L652:L656

### 5.4.2

- https://github.com/torvalds/linux/blob/v5.16/arch/x86/entry/common.c#L40:L54
- https://github.com/torvalds/linux/blob/v5.16/arch/x86/entry/entry_64.S#L113
- https://github.com/torvalds/linux/blob/v5.16/arch/x86/kernel/cpu/common.c#L1794

## 5.5
- https://zenn.dev/hajimehoshi/articles/72f027db464280

## 5.6

- [Disabling and Enabling System Integrity Protection](https://developer.apple.com/documentation/security/disabling_and_enabling_system_integrity_protection)

- System Integrity Protectionをdisableしても、Permission deniedの状態

```
sudo dtruss ./go-project/ch5_system-call/5.2.1.go 
Password:
dtrace: system integrity protection is on, some features will not be available
dtrace: failed to execute ./go-project/ch5_system-call/5.2.1.go: Permission denied

## After disable System Integrity Protection
rayt@RaydeMacBook-Pro ~ % csrutil status
System Integrity Protection status: disabled.
rayt@RaydeMacBook-Pro ~ % sudo dtruss ./go-project/ch5_system-call/5.2.1.go
Password:
dtrace: failed to execute ./go-project/ch5_system-call/5.2.1.go: Permission denied
```


## 5.7

- https://linuxjm.osdn.jp/html/LDP_man-pages/man2/write.2.html
- https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-formatmessage

## 5.8
- [ユーザーランド](https://www.weblio.jp/content/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E3%83%A9%E3%83%B3%E3%83%89)：OSが動作するのに必要な、カーネル以外の部分。ファイルシステムやファイル操作コマンド、シェルなどの基本的なソフトウェア群を指す。

# 日本語

- 5.0
  - 探索（たんさく）
  - 座学（ざがく）
  - 紙面（しめん）
  - まれ：rare
- 5.1.1
  - 特権（とっけん）
  - 覗く（のぞく）
  - 行儀（ぎょうぎ）
  - 節制（せっせい）
  - 入出力（にゅうしゅつりょく）
  - 暴れ（あばれ）
  - 主流（しゅりゅう）
  - 異なる（ことなる）
  - 制約（せいやく）
  - 領域（りょういき）
- 5.1.2
  - 多々（たた）
  - 介する（かいする）
  - 割り当て：assign
  - なぞらえる：liken
  - プロシージャ：procedure
- 5.1.3
  - 無駄（むだ）
  - 出力（しゅつりょく）
  - 備える（そなえる）
  - まとも：proper
- 5.2.1
  - 下り（くだり）
- 5.2.2
  - 発生しうる：can occur
  - 見込まれる：expected
- 5.3
  - 先頭（せんとう）
  - 与える（あたえる）
  - 順序（じゅんじょ）
- 5.4.1
  - 隣接（りんせつ）
- 5.5
  - 眺める（ながめる）
- 5.6
  - 冒頭（ぼうとう）
  - 中途半端（ちゅうとはんぱ）
- 5.7
  - 単なる（たんなる）
  - 察し（さっし）
  - すばやい：fast
- 5.8
  - もっとも：most
  - またがず：
  - 擬似：ぎじ
  - 防ぐ（ふせぐ）


# Other

- press command+shift+V in the visual studio code to preview.