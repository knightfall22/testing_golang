package benchmark

import "os"

func FileLen(f string, size int) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}

	defer file.Close()
	buf := make([]byte, size)
	count := 0
	for {
		num, err := file.Read(buf)
		count += num
		if err != nil {
			break
		}
	}

	return count, nil
}
