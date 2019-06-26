package fileio

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

func ReadFileWithOS(filepath string) (string, error) {

	/*os.OpenFile(name string, flag int, perm FileMode)打开文件并进行文件内容更改，需要注意flag相关的参数以及含义。
	const (
        O_RDONLY int = syscall.O_RDONLY // 只读打开文件和os.Open()同义
        O_WRONLY int = syscall.O_WRONLY // 只写打开文件
        O_RDWR   int = syscall.O_RDWR   // 读写方式打开文件
        O_APPEND int = syscall.O_APPEND // 当写的时候使用追加模式到文件末尾
        O_CREATE int = syscall.O_CREAT  // 如果文件不存在，此案创建
        O_EXCL   int = syscall.O_EXCL   // 和O_CREATE一起使用, 只有当文件不存在时才创建
        O_SYNC   int = syscall.O_SYNC   // 以同步I/O方式打开文件，直接写入硬盘.
        O_TRUNC  int = syscall.O_TRUNC  // 如果可以的话，当打开文件时先清空文件
	)*/

	fileObj, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	
	if err != nil {
		fmt.Println("ReadFileWithOS error at os.OpenFile():", err)
		panic(err)
	}

	defer fileObj.Close()

	contents, err := ioutil.ReadAll(fileObj)
	if err != nil {
		fmt.Println("ReadFileWithOS error at ioutil.ReadAll(): ", err)
		panic(err)
	} else {
		result := strings.Replace(string(contents),"\n","",1)
		return result, err
	}
}

func WriteFileWithOS(filepath string, content string) error {
	fileObject, err := os.OpenFile(filepath, os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("WriteFileWithOS error at os.OpenFile(): ", err)
		panic(err)
	}
	
	defer fileObject.Close()

	// 写入字符串
	// if _,err := fileObj.WriteString(content);err == nil {
    //     fmt.Println("Successful writing to the file with os.OpenFile and *File.WriteString method.",content)
	// }
	// 写入二进制
	contents := []byte(content)
	_, err = fileObject.Write(contents)
    if err != nil {
		fmt.Println("WriteFileWithOS error at fileObject(): ", err)
		panic(err)
	}
	return err
}

func ReadFileWithIoutil(filepath string) (string, error) {
	fileContentByte, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("ReadFileWithIoutil error at ioutil.ReadFile(): ", nil)
		panic(err)
	}
	fileContent := strings.Replace(string(fileContentByte), "\n", "", 1)

	return fileContent, err
}

func WriteFileWithIoutil(filepath string, content string) error {
	data := []byte(content)
	err := ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		fmt.Println("WriteFileWithIoutil error at ioutil.WriteFile(): ", err)
		panic(err)
	}
	return err
}