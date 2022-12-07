package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dirSizes map[string]int
var subDirs map[string][]string

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	currentPath := ""
	currentDir := ""

	dirSizes = make(map[string]int)
	subDirs = make(map[string][]string)

	for r.Scan() {
		line := r.Text()

		if strings.HasPrefix(line, "$ cd") {

			if line == "$ cd /" {
				currentPath = "/"
				currentDir = "/"
			} else if line == "$ cd .." {
				currentPath = strings.TrimSuffix(currentPath, "/"+currentDir)
				currentDirtemp := strings.Split(currentPath, "/")
				currentDir = currentDirtemp[len(currentDirtemp)-1]
			} else { //cd to another dir
				var cd string
				fmt.Sscanf(line, "$ cd %s\n", &cd)
				currentDir = cd
				if currentPath != "/" {
					currentPath = currentPath + "/" + cd
				} else {
					currentPath = "/" + cd
				}
			}

		} else if strings.HasPrefix(line, "$ ls") {
			//do nothing i guess
		} else if strings.HasPrefix(line, "dir ") {
			var dir string
			fmt.Sscanf(line, "dir %s\n", &dir)
			if currentPath != "/" {
				subDirs[currentPath] = append(subDirs[currentPath], currentPath+"/"+dir)
			} else {
				subDirs[currentPath] = append(subDirs[currentPath], "/"+dir)
			}

		} else {
			var filename string
			var filesize int
			fmt.Sscanf(line, "%d %s\n", &filesize, &filename)
			dirSizes[currentPath] += filesize
		}

	}

	totalSize("/")

	total := 0
	for _, v := range dirSizes {
		if v <= 100000 {
			total += v
		}
	}

	fmt.Println("Part 1: ", total)

	spaceNeeded := 70000000 - dirSizes["/"]
	spaceNeeded = 30000000 - spaceNeeded
	temp := 999999999999
	for _, v := range dirSizes {
		if v >= spaceNeeded && v < temp {
			temp = v
		}
	}

	fmt.Println("Part 2: ", temp)

}

func totalSize(dir string) int {
	total := dirSizes[dir]
	for _, v := range subDirs[dir] {
		total += totalSize(v)
	}

	dirSizes[dir] = total
	return total
}
