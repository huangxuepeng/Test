package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

/**
* 代码描述:创建文件
* 作者:黄雪朋
* 创建时间:2024/09/10 19:58:38
 */
func createFile() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("create fail")
		return
	}

	defer file.Close()
	fmt.Println("create successfully")
}

/**
* 代码描述:删除文件
* 作者:黄雪朋
* 创建时间:2024/09/10 19:58:55
 */
func removeFile() {
	err := os.Remove("example.txt")
	if err != nil {
		fmt.Println("error deleting file: ", err)
		return
	}

	fmt.Println("File remove successfully")
}

/**
* 代码描述:文件读写操作
* 作者:黄雪朋
* 创建时间:2024/09/10 20:01:40
 */
func readAndWrite() {
	// createFile()
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		fmt.Println(n)
		if err != nil {
			fmt.Println("read fail: ", err)
			return
		}
		if n == 0 {
			break
		}
		fmt.Println(string(buffer))
	}
}

/**
* 代码描述:文件写入数据
* 作者:黄雪朋
* 创建时间:2024/09/10 20:12:38
 */
func writeData() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("error creatign file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("hello word"))
	if err != nil {
		fmt.Println("error writing to file: ", err)
		return
	}
	fmt.Println("data written to file successfully")
}

func writeDataPro() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("error opening file: ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error reading file: ", err)
			return
		}
		fmt.Print(line)
		if err != nil {
			break
		}
	}
}

/**
* 代码描述: 文件信息获取与修改
* 作者:黄雪朋
* 创建时间:2024/09/10 20:37:36
 */
func FileStat() {
	info, err := os.Stat("example.txt")
	if err != nil {
		fmt.Println("error stating file: ", err)
		return
	}
	fmt.Println(info.Name(), info.IsDir())

}

/**
* 代码描述: 读取变量
* 作者:黄雪朋
* 创建时间:2024/09/10 21:42:07
 */
func GetEnv() {
	// 读取环境变量
	value := os.Getenv("PATH")
	// fmt.Println(value)

	// 查找环境变量
	value, ok := os.LookupEnv("PATH")
	if ok {
		fmt.Println("PATH exists:", value)
	} else {
		fmt.Println("PATH does not exist")
	}
}

/**
* 代码描述:设置环境变量
* 作者:黄雪朋
* 创建时间:2024/09/10 21:53:23
 */
func setEnv() {
	// 设置环境
	err := os.Setenv("MY_VAR", "Hello word")
	if err != nil {
		fmt.Println("error setting var:", err)
		return
	}

	// 读取环境变量
	value := os.Getenv("MY_VAR")
	fmt.Println(value)
}

func main() {
	path, err := exec.LookPath("prog")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)

}
