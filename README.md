# Go Image Converter

Go で作成した画像変換ツールです．JPG と PNG フォーマット間の画像変換を簡単に行うことができます．

## 機能

- JPG ファイルを PNG 形式に変換
- PNG ファイルを JPG 形式に変換
- ファイルが存在しない場合や無効な拡張子が指定された場合にエラーメッセージを出力

## 使用方法

### 1. プロジェクトのクローンまたはダウンロード

まず，このリポジトリをクローンまたはダウンロードしてください．

### 2. 依存関係の取得

プロジェクトに依存するパッケージは go.mod ファイルに記載されています．以下のコマンドでモジュールをダウンロードします．

```bash
go mod tidy
```

### 3. 実行

ターミナルで以下のコマンドを実行します．

```bash
go run .
```

プログラムが起動し，画像ファイルのパスを入力するように求められます．適切な JPG または PNG ファイルを指定してください．

### 4. ビルド

バイナリをビルドする場合は，以下のコマンドを実行します．

```bash
go build
```

## 入力例

実行中に画像ファイルのパスを入力すると，指定された画像のフォーマットが変換されます．

```bash
$ go run .

入力する画像ファイルのパスを入力してください: ./sample.jpg
PNGファイルを作成しました: ./sample.png
```

## ディレクトリ構成

```bash
.
├── go.mod             # Goモジュールファイル
├── main.go            # メインプログラム
├── mypkg              # カスタムパッケージ
│   ├── image.go       # 画像変換機能
│   └── input.go       # 入力処理とファイル確認機能
```

## 関数説明

### Input() string

- ユーザーから画像ファイルのパスを取得し，ファイルの存在確認と拡張子の検証を行います．

### LoadImage(imagePath string) GoImg

- 画像ファイルを読み込み，GoImg 構造体に格納します．

### GoImg.ConvertToPNG(file string)

- 読み込んだ画像を PNG 形式に変換して保存します．

### GoImg.ConvertToJPG(file string)

- 読み込んだ画像を JPG 形式に変換して保存します．
