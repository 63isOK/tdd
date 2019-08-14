package loop

func repeat(char string, count int) (ret string) {
	for i := 0; i < count; i++ {
		ret = ret + char
	}

	return
}
