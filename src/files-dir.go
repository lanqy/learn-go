package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// Create an empty file
func main() {
	emptyFile, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()
}

// Create directory
func main1() {
	_, err := os.Stat("test")

	if os.IsNotExist(err) {
		errDir := os.MKdirAll("test", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}

// Rename a file
func main2() {
	oldName := "test.txt"
	newName := "testing.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

// Move a file from one location to another
func main3() {
	oldLocation := "/var/www/html/test.txt"
	newLocation := "/var/www/html/src/test.txt"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

// Copy file at specified location
func main4() {
	sourceFile, err := os.Open("/var/www/html/src/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()
	// create new file
	newFile, err := os.Create("/var/www/html/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesCopied)

}

// Getting metadata of file
func main5() {
	fileStat, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileStat.Name())
	fmt.Println("Size:", fileStat.Size())
	fmt.Println("Permissions:", fileStat.Mode())
	fmt.Println("Last Modifiled:", fileStat.ModTime())
	fmt.Println("IS Directory: ", fileStat.ISDir())
}

//Deleting a specific file
func main6() {
	err := os.Remove("/var/www/html/test.txt")
	if err != nil {
		log.Fatal(err)
	}
}

// Read a text file character by character
func main7() {
	filename := "test.txt"

	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	data := buffio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)
	for data.Scan() {
		fmt.Print(data.Text())
	}
}

// Truncate file content
func main8() {
	err := os.Truncate("test.txt", 100)

	if err != nil {
		log.Fatal(err)
	}
}

// Append content in text file
func main9() {
	message := "add this content at end"
	filename := "test.txt"

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()

	fmt.Fprintf(f, "%s\n", message)
}

// Change file permission, ownership, and timestamps
func main10() {
	_, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File exist.")

	// change permissions Linux.
	err = os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	// Change file ownership.
	err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}

	// Change file timestamps.
	addoneDayFromNow := time.Now().Add(24 * time.Hour)
	lastAccessTime := addOneDayFromNow
	lastModifyTime := addOneDayFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModityTime)
	if err != nil {
		log.Println(err)
	}
}

//Compress list of files into ZIP file
func main11() {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile("test.zip", flags, 0644)
	if err != nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}
	defer file.Close()

	var files = []string{"test1.txt", "test2.txt", "test3.txt"}

	zipw := zip.NewWriter(file)
	defer zipw.Close()

	for _, filename := range files {
		if err := appendFiles(filename, zipw); err != nil {
			log.Fatalf("Failed to add file %s to zip: %s", filename, err)
		}
	}
}

func appendFiles(filename string, zipw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Failed to open %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zipw.Create(filename)
	if err != nil {
		msg := "Failed to create entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("Failed to write %s to zip: %s", filename, err)
	}

	return nil
}

// Read list of files inside ZIP file
func main12() {
	read, err := zip.OpenReader("test.zip")
	if err != nil {
		msg := "Failed to open: %s"
		log.Fatalf(msg, err)
	}
	defer read.Close()

	for _, file := range read.File {
		if err := listFiles(file); err != nil {
			log.Fatalf("Failed to read %s from zip: %s", file.Name, err)
		}
	}
}

func listFiles(file *zip.File) error {
	fileread, err := file.Open()
	if err != nil {
		msg := "Failed to open zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}
	defer fileread.Close()

	fmt.Fprintf(os.Stdout, "%s:", file.Name)

	if err != nil {
		msg := "Failed to read zip %s for reading: %s"
		return fmt.Errorf(msg, file.Name, err)
	}

	fmt.Println()

	return nil
}

// Extract or Unzip
func main13() {
	zipReader, _ := zip.OpenReader("test.zip")
	for _, file := range zipReader.Reader.File {
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		targetDir := "./"
		extractedFilePath := filepath.Join(targetDir, fileName)
		if file.FileInfo().IsDir() {
			log.Println("Directory Created:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			log.Println("File extracted:", file.Name)
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
