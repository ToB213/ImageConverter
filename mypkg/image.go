package mypkg

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type GoImg struct {
	Image  image.Image
	Path   string
	Height int
	Width  int
}

// JPG画像を読み込んでGoImg構造体に格納
func LoadImage(imagePath string) GoImg {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("画像ファイルを開くことができませんでした:", err)
		os.Exit(1)
	}
	defer file.Close()

	// 画像をデコード
	src, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("画像の読み込みに失敗しました:", err)
		os.Exit(1)
	}

	// GoImg構造体に画像情報を格納
	img := GoImg{
		Image: src,
		Path:  imagePath,
	}

	return img
}

// 画像をPNGに変換
func (img *GoImg) ConvertToPNG(file string) {
	// 出力ファイル名を生成
	pngFile := strings.TrimSuffix(file, filepath.Ext(file)) + ".png"

	// 出力ファイルを作成
	out, err := os.Create(pngFile)
	if err != nil {
		fmt.Println("PNGファイルを作成できませんでした:", err)
		os.Exit(1)
	}
	defer out.Close()

	// PNG形式でエンコード
	err = png.Encode(out, img.Image)
	if err != nil {
		fmt.Println("PNG形式でエンコードできませんでした:", err)
		os.Exit(1)
	}

	fmt.Printf("PNGファイルを作成しました: %s\n", pngFile)
}

// 画像をJPGに変換
func (img *GoImg) ConvertToJPG(file string) {
	// 出力ファイル名を生成
	jpgFile := strings.TrimSuffix(file, filepath.Ext(file)) + ".jpg"

	// 出力ファイルを作成
	out, err := os.Create(jpgFile)
	if err != nil {
		fmt.Println("JPGファイルを作成できませんでした:", err)
		os.Exit(1)
	}
	defer out.Close()

	// JPG形式でエンコード
	err = jpeg.Encode(out, img.Image, nil)
	if err != nil {
		fmt.Println("JPG形式でエンコードできませんでした:", err)
		os.Exit(1)
	}

	fmt.Printf("JPGファイルを作成しました: %s\n", jpgFile)
}
