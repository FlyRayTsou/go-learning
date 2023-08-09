# 11
この章は主にコマンドシェルの紹介です。
 
# 11.1

シェルを聞いたら、何を思い浮かぶ。

- 「キーボードで命令を入れて、ファイルの処理やコマンドの実行をする画面」：bash、zsh、fish
- 「Windowsのファイルやアイコンが置かれるデスクトップ」
- 「DOSシェル」
    - https://news.mynavi.jp/article/20130408-dosshell/

CUI, GUI, TUI
- CUI(character user interface, Command-line interface)
    - https://en.wikipedia.org/wiki/Command-line_interface
- GUI(Graphical user interface)
- TUI(Text Based User Interface)
    - https://en.wikipedia.org/wiki/Text-based_user_interface


ユーザーがコンピューターを操作するために使う接点となるシステムにおいてコンピューターシステムの外側の「殻(Shell)」です。
- Windows：Explorer
- MaxOS：Finder

「命令をキーボードで入力して実行するモード」

node
```
rayt@RaydeMacBook-Pro ~ % node
Welcome to Node.js v19.1.0.
Type ".help" for more information.
> exit()
```

python3
```
rayt@RaydeMacBook-Pro ~ % python3
Python 3.7.8 (v3.7.8:4b47a5b6ba, Jun 27 2020, 04:47:50)
[Clang 6.0 (clang-600.0.57)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> exit()
```

## 11.1.1

コマンドシェルはファイルからソフトウェアを起動したり、ファイルを整理したりする役割を担います。

- シェルを/bin/falseにして、ログインできないシステムユーザーを作る
    - https://www.express.nec.co.jp/linux/distributions/knowledge/system/false.html


## 11.1.2

- シェルガないシステム
    - 例：かっぱ寿司（実際はキオスクモードかもしれないので）
        - https://jouhou-kan.net/archives/6149
    - Distroless
        - 攻撃される可能性が少なくなり、目的対応できれば、積極的に使うのは良いかも。イメージサイズも小さいです。
        - https://devops-blog.virtualtech.jp/entry/20220713/1657683666
        - https://qiita.com/tkusumi/items/dbda4abadd24d3db0981
    - unikernel
        - ハードウェア上でOSいらずでアプリケーションを起動できるように
        - http://unikernel.org/projects/

# 11.2
シェル自体が持つ機能（内部コマンド）や外部コマンド（プログラムなど）を端末エミュレータと呼ばれるソフトウェアから起動し、それによりファイルの整理や加工といったコンピュータ上での作業を実行

## 11.2.1 端末エミュレータ上でコマンドを入力

iTerm、terminalなど: 擬似端末(Pseudo Terminal)、端末アプリケーション、端末端末エミュレータ

## 11.2.2 シェルスクリプト

```
rayt@RaydeMacBook-Pro ch11_command_shell_101 % echo $SHELL
/bin/zsh

rayt@RaydeMacBook-Pro ch11_command_shell_101 % zsh 11.2.2.zsh
+11.2.2.zsh:4> set -e
+11.2.2.zsh:6> echo 'hello world'
hello world

rayt@RaydeMacBook-Pro ch11_command_shell_101 % bash ./11.2.2.bash
+ set -e
+ echo 'hello world'
hello world

rayt@RaydeMacBook-Pro ch11_command_shell_101 % ./11.2.2.bash
zsh: permission denied: ./11.2.2.bash

rayt@RaydeMacBook-Pro ch11_command_shell_101 % ls -l
total 16
-rw-r--r--@ 1 rayt  staff    47  7 23 13:13 11.2.2.bash

rayt@RaydeMacBook-Pro ch11_command_shell_101 % chmod +x 11.2.2.bash

rayt@RaydeMacBook-Pro ch11_command_shell_101 % ls -l
total 16
-rwxr-xr-x@ 1 rayt  staff    47  7 23 13:13 11.2.2.bash

rayt@RaydeMacBook-Pro ch11_command_shell_101 % ./11.2.2.bash
+ set -e
+ echo 'hello world'
hello world

rayt@RaydeMacBook-Pro ch11_command_shell_101 % ./11.2.2.zsh
zsh: permission denied: ./11.2.2.zsh

rayt@RaydeMacBook-Pro ch11_command_shell_101 % chmod +x 11.2.2.zsh

rayt@RaydeMacBook-Pro ch11_command_shell_101 % ./11.2.2.zsh
+./11.2.2.zsh:4> set -e
+./11.2.2.zsh:6> echo 'hello world'
hello world

rayt@RaydeMacBook-Pro ch11_command_shell_101 % cat /etc/shells
# List of acceptable shells for chpass(1).
# Ftpd will not allow users to connect who are not using
# one of these shells.

/bin/bash
/bin/csh
/bin/dash
/bin/ksh
/bin/sh
/bin/tcsh
/bin/zsh

```

Mac のデフォルトシェルが bash から zsh に変更された機に乗じて乗り換える手順
- https://qiita.com/normalsalt/items/5b71fbd03721800de4a0

PowerShell
- https://learn.microsoft.com/en-us/powershell/module/activedirectory/?view=windowsserver2022-ps

## 11.2.3 プログラムからの外部プロセス起動


```
Docker
myprogはmy programの意味ですか。
EXTRYPOINT "./myprog hello world"
```

Goのos/execは, シェルを経由しない実行のみをサポートしています。シェル経由の実行をエミュレーションするには次のようにします。
```
rayt@RaydeMacBook-Pro ch11_command_shell_101 % go run 11.2.3.go
darwin
/bin/zsh
rayt@RaydeMacBook-Pro ch11_command_shell_101 % sleep 5
```

# 11.3

[Shell & Utilities] > [Utilities]
- https://pubs.opengroup.org/onlinepubs/9699919799/


完全に準拠したOSは「UNIX」を名乗れます。
- Uniplexed Information and Computing Service
- ex: maxOS 10.5 Leopard ~ 11 Big Sur
- https://ja.wikipedia.org/wiki/UNIX

POSIXやLSBは巨大な規格であることから、それらの小さいサブセットとして、BusyBoxと呼ばれるソフトウェアもあります。多くのPOSIX系OSではコマンドが特立したプログラムとして提供されていますが、BusyBoxでは1つの実行プログラム中に複数のUNIXのコマンドが実装されています。
https://www.busybox.net/

```
rayt@RaydeMacBook-Pro ch11_command_shell_101 % docker run --rm -it alpine busybox
Unable to find image 'alpine:latest' locally
latest: Pulling from library/alpine
31e352740f53: Pull complete
Digest: sha256:82d1e9d7ed48a7523bdebc18cf6290bdb97b82302a8a9c27d4fe885949ea94d1
Status: Downloaded newer image for alpine:latest
BusyBox v1.36.1 (2023-06-02 00:42:02 UTC) multi-call binary.
BusyBox is copyrighted by many authors between 1998-2015.
Licensed under GPLv2. See source distribution for detailed
copyright notices.

Usage: busybox [function [arguments]...]
   or: busybox --list[-full]
   or: busybox --install [-s] [DIR]
   or: function [arguments]...

	BusyBox is a multi-call binary that combines many common Unix
	utilities into a single executable.  Most people will create a
	link to busybox for each function they wish to use and BusyBox
	will act like whatever it was invoked as.

Currently defined functions:
	[, [[, acpid, add-shell, addgroup, adduser, adjtimex, arch, arp,
	arping, ash, awk, base64, basename, bbconfig, bc, beep, blkdiscard,
	blkid, blockdev, brctl, bunzip2, bzcat, bzip2, cal, cat, chattr, chgrp,
	chmod, chown, chpasswd, chroot, chvt, cksum, clear, cmp, comm, cp,
	cpio, crond, crontab, cryptpw, cut, date, dc, dd, deallocvt, delgroup,
	deluser, depmod, df, diff, dirname, dmesg, dnsdomainname, dos2unix, du,
	dumpkmap, echo, ed, egrep, eject, env, ether-wake, expand, expr,
	factor, fallocate, false, fatattr, fbset, fbsplash, fdflush, fdisk,
	fgrep, find, findfs, flock, fold, free, fsck, fstrim, fsync, fuser,
	getopt, getty, grep, groups, gunzip, gzip, halt, hd, head, hexdump,
	hostid, hostname, hwclock, id, ifconfig, ifdown, ifenslave, ifup, init,
	inotifyd, insmod, install, ionice, iostat, ip, ipaddr, ipcalc, ipcrm,
	ipcs, iplink, ipneigh, iproute, iprule, iptunnel, kbd_mode, kill,
	killall, killall5, klogd, last, less, link, linux32, linux64, ln,
	loadfont, loadkmap, logger, login, logread, losetup, ls, lsattr, lsmod,
	lsof, lsusb, lzcat, lzma, lzop, lzopcat, makemime, md5sum, mdev, mesg,
	microcom, mkdir, mkdosfs, mkfifo, mkfs.vfat, mknod, mkpasswd, mkswap,
	mktemp, modinfo, modprobe, more, mount, mountpoint, mpstat, mv, nameif,
	nanddump, nandwrite, nbd-client, nc, netstat, nice, nl, nmeter, nohup,
	nologin, nproc, nsenter, nslookup, ntpd, od, openvt, partprobe, passwd,
	paste, pgrep, pidof, ping, ping6, pipe_progress, pivot_root, pkill,
	pmap, poweroff, printenv, printf, ps, pscan, pstree, pwd, pwdx,
	raidautorun, rdate, rdev, readahead, readlink, realpath, reboot,
	reformime, remove-shell, renice, reset, resize, rev, rfkill, rm, rmdir,
	rmmod, route, run-parts, sed, sendmail, seq, setconsole, setfont,
	setkeycodes, setlogcons, setpriv, setserial, setsid, sh, sha1sum,
	sha256sum, sha3sum, sha512sum, showkey, shred, shuf, slattach, sleep,
	sort, split, stat, strings, stty, su, sum, swapoff, swapon,
	switch_root, sync, sysctl, syslogd, tac, tail, tar, tee, test, time,
	timeout, top, touch, tr, traceroute, traceroute6, tree, true, truncate,
	tty, ttysize, tunctl, udhcpc, udhcpc6, umount, uname, unexpand, uniq,
	unix2dos, unlink, unlzma, unlzop, unshare, unxz, unzip, uptime, usleep,
	uudecode, uuencode, vconfig, vi, vlock, volname, watch, watchdog, wc,
	wget, which, who, whoami, whois, xargs, xxd, xzcat, yes, zcat, zcip
```

内部コマンド
```
cd, pwd
```

内部コマンド以外
```
cat, chmod
```

# 11.4

環境変数は、「ユーザーごとの固有の設定」が含まれた配列です。

Q：（ただし、設定したプロセスの親には伝搬しないので、その点ではグローバル変数と異なります。）

```
rayt@RaydeMacBook-Pro ch11_command_shell_101 % env
SSH_AUTH_SOCK=/private/tmp/com.apple.launchd.mXQfocblUh/Listeners
LC_TERMINAL_VERSION=3.4.17
(略)
```

HTTP_PROXY
- https://qiita.com/toshyss/items/f229f2595f6ea20ebd70


The Twelve Factor App (クラウド時代のアプリケーション設計の方法論として知られる)
- https://12factor.net/
- 「設定の変更はすべて環境変数で行えるようにすべき」

Docker command -> Daemon -> container

## 11.4.1

11.4.1.go

## 11.4.2

.env idea was create by Ruby on Rails

```
rayt@RaydeMacBook-Pro ch11_command_shell_101 % chmod +x my-printenv.bash
rayt@RaydeMacBook-Pro ch11_command_shell_101 % go run dotenv.go ./my-printenv.bash
NVM_INC=/Users/rayt/.nvm/versions/node/v19.1.0/include/node
TERM_PROGRAM=iTerm.app
(略)
HELLO=WORLD
```

- https://github.com/joho/godotenv

## 11.4.3
WSL(Windows Subsystem for linux)
WSL2(WindowsとLinuxがシームレスに提携できるようになっています。)：まるで一つのOS環境である

# 11.5
シェルがコマンドを実行するまでの次の各ステップについて説明


Tiny Shell
- https://github.com/shibukawa/tish
- タブ機能がない

```
rayt@RaydeMacBook-Pro tish-main % go run cmd/tish/main.go
🐸 tiny shell

  🕓23:18:19    👤rayt@RaydeMacBook-Pro.local    📁~/Downloads/tish-main    ✔︎
```

1. ユーザーの入力を受け付ける
2. 入力されたテキストの分解
3. コマンドと引数の全処理
4. 実行ファイルの探索
5. ワイルドカードの展開
6. プロセスの起動

## 11.5.1 ユーザーの入力を受け付ける
シェルは起動すると「ユーザーの入力受け付ける状態である」ことをプロンプトによって示します。そして「ユーザーがやりたいこと」を文字列の形で受け取ります。

### 11.5.1-1.go
- fmt.Scanf()など単一の文字列として受け取ればいいでしょう

```
rayt@RaydeMacBook-Pro ch11_command_shell_101 % go run 11.5.1-1.go
cd ..
// why % in the end?
// https://stackoverflow.com/questions/59091824/why-does-printf-leave-an-extra-after-my-output
// The '%' you are seeing is your bash / shell prompt for missing newlines.
// You input cmd : cd .. .%
// fmt.Printf("You input cmd : %s %s .\n", cmd, cmd2)
// You input cmd : cd .. .
// fmt.Printf("You input cmd : %s %s .\n", cmd, cmd2)
```

### 11.5.1-2.go
- https://github.com/peterh/liner
- SetCompleter

## 11.5.2 入力されたテキストの分解



### Shlex
- https://github.com/google/shlex
- This repository has been archived by the owner on Dec 30, 2022. It is now read-only.

```
// expected boolean or range expression, found assignment (missing parentheses around composite literal?)synta
for ; token, err := l.Next(); err != nil {
	args = append(args, token)
}
```

11.5.2-1.go

11.5.2-2.go

 | ; || && < > >> 2> 2>> &> &>>

## 11.5.3 コマンドと引数の全処理

- 引数やコマンドの中に変数があれば、それらを実際のデータと置き換えます。
```
rayt@RaydeMacBook-Pro ch11_command_shell_101 % echo "hello ${USER}"
hello rayt
```

- さらに実用的なシェルでは、あるコマンドの引数で別のコマンドをバッククオート(`)や$(...)でxむことにより、その部分を「別のコマンドを実行した結果の文字列」で置き換えコマンドを実行するという機能があります。

```
file_count=$(ls -l | wc -l)
echo "ファイル数は ${file_count} です。"
```

## 11.5.4 実行ファイルの探索

```
rayt@RaydeMacBook-Pro ch11_command_shell_101 % which echo
echo: shell built-in command
rayt@RaydeMacBook-Pro ch11_command_shell_101 % which cd
cd: shell built-in command
rayt@RaydeMacBook-Pro ch11_command_shell_101 % which cat
/bin/cat
```

## 11.5.5 ワイルドカードの展開

```
rayt@RaydeMacBook-Pro ch11_command_shell_101 % ls *.go
```

11.5.5.go

## 11.5.6 プロセスの起動

プロセス：Ch12
メモリ：Ch17

# 11.6 Unix哲学とシェル
- UNIXというシステムは小さなユーティリティを組み合わせて多くの仕事を行うように設計されています。その設計の背後（はいご）にある考え方を称して（しょうして）「Unix哲学」と呼ぶことがあります
- これがUnix哲学です。一つのことだけをうまく行うプログラムをいくつも作成しましょう。互いに強調して動作するプログラムをいくつも作成しましょう。テキストストリームを扱うプログラムをいくつも作りましょう。なぜならテキストストリームはユニバーサルなインタフェースだからです。

```
rayt@RaydeMacBook-Pro ch11_command_shell_101 % cat *.md | grep unix | wc -l
       2
rayt@RaydeMacBook-Pro ch11_command_shell_101 % cat *.rst | grep unix | wc -l
zsh: no matches found: *.rst
       0
```

# 11.7

