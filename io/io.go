// Program that makes use of most of the utility functions from io/ioutil package
package io

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func RunIoProgram() {
	// using ioutil.ReadAll func
	fmt.Println("Running ReadAll func ----------------->")
	if err := runReadAll(); err != nil {
		fmt.Println("Error while running ioutil.ReadAll:", err)
	}

	// using iotuil.ReadDir func
	fmt.Println("Running ReadDir func ----------------->")
	if err := runReadDir(); err != nil {
		fmt.Println("Error while running ioutil.ReadDir:", err)
	}

	// using iotuil.ReadFile func
	fmt.Println("Running ReadFile func ----------------->")
	if err := runReadFile(); err != nil {
		fmt.Println("Error while running ioutil.ReadFile:", err)
	}

	// using iotuil.WriteFile func
	fmt.Println("Running WriteFile func ---------------->")
	if err := runWriteFile(); err != nil {
		fmt.Println("Error while running ioutil.WriteFile:", err)
	}
}

func runReadAll() error {
	inputStream := strings.NewReader("Printing input stream for iouti.ReadAll func")
	output, err := ioutil.ReadAll(inputStream)
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func runReadDir() error {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}

func runReadFile() error {
	content, err := ioutil.ReadFile("io/test")
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	return nil
}

func runWriteFile() error {
	content := []byte("** This content was added by iotuil.WriteFile func **")
	if err := ioutil.WriteFile("io/testWriteFile", content, 0644); err != nil {
		return err
	}
	return nil
}
