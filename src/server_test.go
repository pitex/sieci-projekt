package test

import "./joinservice"
import "testing"
import "os"

func TestRewriteFile(t *testing.T) {
	file, _ := os.Create("test1")

	defer os.Remove("test1")

	text := []byte("Just some testing text,\nAnd some more,\nAnd again some more.\n")
	file.Write(text)

	file.Close()
	file, _ = os.Create("test2")
	defer os.Remove("test2")
	file.Close()
	joinservice.RewriteFile("test1", "test2")

	var file2 *os.File

	file, _ = os.Open("test1")
	defer file.Close()
	file2, _ = os.Open("test2")
	defer file2.Close()

	bytes := make([]byte, 1024)
	bytes2 := make([]byte, 1024)
	var n1,n2 int

	for ;; {
		n1, _ = file.Read(bytes)
		n2, _ = file2.Read(bytes2)
		t.Log(string(bytes))
		t.Log(string(bytes2))

		if n1 != n2 {
			t.FailNow()
		}

		for i := range bytes {
			if bytes[i] != bytes2[i] {
				t.FailNow()
			}
		}

		if n1 < 1024 {
			break
		}
	}
}