package mypkg

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ユーザーからJPGファイルのパスを入力
func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("入力する画像ファイルのパスを入力してください: ")
	scanner.Scan()
	filePath := scanner.Text()

	// ファイルパスが空の場合
	if filePath == "" {
		fmt.Println("ファイルパスが空です")
		os.Exit(1)
	}

	// ファイルが存在するか確認
	_, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("ファイルが存在しません")
		os.Exit(1)
	}

	// 拡張子がJPGか確認（大文字・小文字対応）
	ext := strings.ToLower(filepath.Ext(filePath))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		fmt.Println("JPGまたはPNGファイルを指定してください")
		os.Exit(1)
	}

	return filePath
}

// 拡張子から変換のフラグを設定
func Select(flag *bool, filePath string) {
	ext := strings.ToLower(filepath.Ext(filePath))
	if ext == ".jpg" || ext == ".jpeg" {
		*flag = false
	} else {
		*flag = true
	}
}
