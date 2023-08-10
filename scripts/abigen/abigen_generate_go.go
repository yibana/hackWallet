package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	// 设置contracts目录的路径
	contractsDir := "./contracts"
	abisDir := path.Join(contractsDir, "/abis")

	// 遍历contracts目录中的所有abi文件
	files, err := ioutil.ReadDir(abisDir)
	if err != nil {
		fmt.Println("无法读取contracts目录：", err)
		return
	}

	for _, file := range files {
		// 检查是否为ABI文件
		if strings.HasSuffix(file.Name(), ".abi") {
			// 获取文件名（不带扩展名）作为目录名
			filename := strings.TrimSuffix(file.Name(), ".abi")
			dirPath := filepath.Join(contractsDir, filename)

			// 检查目录是否已存在
			if _, err := os.Stat(dirPath); os.IsNotExist(err) {
				// 创建目录
				err := os.Mkdir(dirPath, 0755)
				if err != nil {
					fmt.Printf("无法创建目录 %s: %s\n", dirPath, err)
					continue
				}

				abigen_type := fmt.Sprintf("%sInterface", strings.Title(filename))
				// 使用abigen生成Go文件到目录中
				cmd := exec.Command("abigen", "--abi", filepath.Join(abisDir, file.Name()), "--pkg", abigen_type, "--type", abigen_type, "--out", filepath.Join(dirPath, abigen_type+".go"))
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					fmt.Printf("无法生成Go文件：%s\n", err)
					// 将错误信息写入到目录
					ioutil.WriteFile(filepath.Join(dirPath, abigen_type+".err"), []byte(err.Error()), 0644)
				} else {
					fmt.Printf("Go文件生成成功：%s\n", filepath.Join(dirPath, abigen_type+".go"))
				}
			} else {
				fmt.Printf("目录 %s 已存在，跳过生成Go文件。如果需要重新生成，可删除目录重新运行脚本\n", dirPath)
			}
		}
	}
}
