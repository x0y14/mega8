# mega8
ロックマン８の素材を使用して、ロックマンを自作を試みるもの。


## usage
#### unix
1. `go mod tidy`
2. `cd scripts`
3. `make build`
4. `./bin/mega8`
---
#### windows
[make](http://gnuwin32.sourceforge.net/packages/make.htm) をインストールし、pathを通す必要があります。
1. `go mod tidy`
2. `cd scripts`
3. `make build`
4. `./bin/mega8.exe`

---

### メモ
#### Entity
状態を持った全て, エフェクトだったり、キャラクタだったり
### Character
ステータスを持ったEntity
### Object
Entity

####Stage
背景の管理を担当

motionは、キャラクターにworkを付与するものにするのはどうだろうか。
workを司る役目がいて、そこに渡せば、一括管理ができる。
そうすれば、複数入力に対応できそう？
ステージからの効果とか、、、？



---
MegaMan and all related characters are copyright CAPCOM.

