package main

import (
	"JpgToPng/mypkg"
)

func main() {
	// ユーザーからファイルパスを取得
	filePath := mypkg.Input()

	// 変換フラグを判定
	var flag bool
	mypkg.Select(&flag, filePath)

	// 画像を読み込み
	img := mypkg.LoadImage(filePath)

	// JPG <=> PNG 変換を実行
	if flag {
		img.ConvertToJPG(filePath)
	} else {
		img.ConvertToPNG(filePath)
	}
}
