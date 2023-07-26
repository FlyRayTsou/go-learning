package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Environ()
	fmt.Println("---------- os.Environ() start 文字列のリストで全部取得 ---------- ")
	fmt.Println(os.Environ())
	fmt.Println("---------- os.Environ() end 文字列のリストで全部取得 ---------- ")

	// os.ExpandEnv()
	fmt.Println("---------- os.ExpandEnv() start 環境変数が埋め込まれた文字列を展開（環境変数を使う） ---------- ")
	fmt.Println(os.ExpandEnv("${HOME}"))
	fmt.Println("---------- os.ExpandEnv() end 環境変数が埋め込まれた文字列を展開（環境変数を使う） ---------- ")

	// os.Expand()
	// https://pkg.go.dev/os#Expand
	fmt.Println("---------- os.Expand() start 環境変数が埋め込まれた文字列を展開（値のマップは自前） ---------- ")
	mapper := func(placeholderName string) string {
		switch placeholderName {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Gopher"
		}

		return ""
	}

	fmt.Println(os.Expand("Good ${DAY_PART}, $NAME!", mapper))
	fmt.Println("---------- os.Expand() end 環境変数が埋め込まれた文字列を展開（値のマップは自前） ---------- ")

	// os.Setenv()
	fmt.Println("---------- os.Setenv() start キーに対する値を設定 ---------- ")
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")
	fmt.Println("---------- os.Setenv() end キーに対する値を設定 ---------- ")

	// os.LookupEnv()
	// https://pkg.go.dev/fmt#hdr-Printing
	fmt.Println("---------- os.LookupEnv() start キーに対する値を取得(有無をboolで返す) ---------- ")
	show := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("status:%t, %s not set\n", ok, key)
		} else {
			fmt.Printf("status:%t, %s=%s\n", ok, key, val)
		}
	}
	show("BURROW")
	show("FAKE_BURROW")

	fmt.Println("---------- os.LookupEnv() end キーに対する値を取得(有無をboolで返す) ---------- ")

	// os.Getenv()
	fmt.Println("---------- os.Getenv() start キーに対する値を取得 ---------- ")
	fmt.Printf("%s lives in %s.\n", os.Getenv("NAME"), os.Getenv("BURROW"))
	fmt.Printf("FAKE %s lives in %s.\n", os.Getenv("FAKE_NAME"), os.Getenv("FAKE_BURROW"))
	fmt.Println("---------- os.Getenv() end キーに対する値を取得 ---------- ")

	// os.Unsetenv()
	fmt.Println("---------- os.Unsetenv() start 指定されたキーを削除する ---------- ")
	os.Unsetenv("NAME")
	fmt.Printf("%s lives in %s.\n", os.Getenv("NAME"), os.Getenv("BURROW"))
	fmt.Println("---------- os.Unsetenv() end 指定されたキーを削除する ---------- ")

	// os.Clearenv()
	fmt.Println("---------- os.Clearenv() start 全部クリアする ---------- ")
	os.Clearenv()
	fmt.Printf("%s lives in %s.\n", os.Getenv("NAME"), os.Getenv("BURROW"))
	fmt.Println("---------- os.Clearenv() end 全部クリアする ---------- ")

}
