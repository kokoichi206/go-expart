## 準備

``` sh
sudo ip netns add router1
sudo ip netns ls

# Network Namespace でコマンドを実行する。
$ sudo ip netns exec router1 ip a
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00

# Network Interface 間を接続するために veth インタフェースを作成！
# veth: 仮想的な LAN ケーブル！
sudo ip link add name host1-router1 type veth peer name router1-host1
sudo ip link add name host2-router1 type veth peer name router1-host2

# ホストから確認可能
$ ip a
49: router1-host2@host2-router1: <BROADCAST,MULTICAST,M-DOWN> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 8a:78:7b:ce:96:2d brd ff:ff:ff:ff:ff:ff
50: host2-router1@router1-host2: <BROADCAST,MULTICAST,M-DOWN> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether a6:93:0f:ec:10:5b brd ff:ff:ff:ff:ff:**ff**

# Network Namespace への接続
sudo ip link set host1-router1 netns host1
sudo ip link set router1-host1 netns router1
sudo ip link set host2-router1 netns host2
sudo ip link set router1-host2 netns router1

# ホストからの確認
# → 先程まで見えていたのが見えなくなっていることがわかる（Network Namespace に接続されたため）
ip a

# こっちには増えてる
$ sudo ip netns exec host1 ip a
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
48: host1-router1@if47: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 02:01:1f:89:b3:67 brd ff:ff:ff:ff:ff:ff link-netns router1

# インタフェースがダウンしてるので UP させる。
sudo ip netns exec host1 ip link set host1-router1 up
sudo ip netns exec router1 ip link set router1-host1 up
sudo ip netns exec host2 ip link set host2-router1 up
sudo ip netns exec router1 ip link set router1-host2 up

# LinkUP 状態になってることが確認できる。
sudo ip netns exec host1 ip a

# Linkup したインタフェースに IP アドレスを設定。
sudo ip netns exec host1 ip addr add 192.168.1.2/24 dev host1-router1
sudo ip netns exec host1 ip route add default via 192.168.1.1
sudo ip netns exec router1 ip addr add 192.168.1.1/24 dev router1-host1
# sudo ip netns exec host2 ip addr del 192.168.2.2/24 dev host2-router1
# sudo ip netns exec host2 ip route del default via 192.168.2.1
# sudo ip netns exec router1 ip addr del 192.168.2.2/24 dev router1-host2
sudo ip netns exec host2 ip addr add 192.168.0.2/24 dev host2-router1
sudo ip netns exec host2 ip route add default via 192.168.0.1
sudo ip netns exec router1 ip addr add 192.168.0.1/24 dev router1-host2

# ip アドレス, デフォルトゲート
$ sudo ip netns exec host1 ip a
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
48: host1-router1@if47: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 02:01:1f:89:b3:67 brd ff:ff:ff:ff:ff:ff link-netns router1
    inet 192.168.1.2/24 scope global host1-router1
       valid_lft forever preferred_lft forever
    inet6 fe80::1:1fff:fe89:b367/64 scope link 
       valid_lft forever preferred_lft forever

# host1 から host2 に、router 経由で ping が飛んでる
# (network namespace では IP フォワードのカーネルパラメータが有効になっているので)
sudo ip netns exec host1 ping -c 2 192.168.0.2
$ sudo ip netns exec host1 traceroute 192.168.0.2
traceroute to 192.168.0.2 (192.168.0.2), 64 hops max
  1   192.168.1.1  0.022ms  0.015ms  0.016ms 
  2   192.168.0.2  0.016ms  0.017ms  0.016ms 
```

## ソケット

| Layer | プロトコル | go で扱うなら？ |
| :---: | :---: | :---: |
| アプリケーション層 | HTTP | net/http |
| トランスポート層 | TCP, UDP | net.TCPConn<br />net.UDPConn |
| インターネット層 | IP, ARP, ICMP | net.IPConn |
| ネットワーク<br />インタフェース層 | Ethernet | syscall |

- syscall を開いで取得したディスクリプタに read, write を行う
  - 通常パッケージが OS ごとに異なるソケットの利用を抽象化している
- デフォルトではブロッキングモードになる
  - 本当？ubuntu 上のラズパイでは違ったかも？
    - epollWait でエラー発生したから、blocking mode にしたら解決
  - TCP サーバーでは1つのソケットを見るだけなので、ブロッキングモードでも良い
  - そこで epoll
    - poll や select と同じように、複数の fd の入出力を監視してくれるシステムコール
    - ソケットにパケットが受信したら教えてね！

## memo

- epoll
- オリジナルの curo との差分
  - ソケットをノンブロキングもどで開いてパケットを受信（オリジナル）
    - epoll に変更
  - IP アドレスをスクリプトで設定（今回）
  - パケットバッファ構造体を定義（今回）
    - オリジナルでは、パケットデータを作成するときに、メモリコピーが発生しないように工夫してそう

## Links

- Go から学ぶ I/O
  - https://zenn.dev/hsaki/books/golang-io-package/viewer/intro
- socket の man ページ
  - https://linuxjm.osdn.jp/html/LDP_man-pages/man2/socket.2.html
