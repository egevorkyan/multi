package sources

import(
	"fmt"
	"bufio"
	"os"
)

type Source struct{
	Path string
}

type FileReader interface{
	OpenFile(path string) ([]byte, error)
	readFile(f *os.File) ()
}

func (s Source) OpenFile() ([]string, error){
	file, err := os.Open(s.Path)
	if err != nil {
		return nil, fmt.Errorf("erro reading file : %s", err.Error())
	}
	lines := s.readFile(file)
	return lines, err
}

func (s Source) readFile(f *os.File) ([]string){
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	return lines
}