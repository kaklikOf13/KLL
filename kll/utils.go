package kll

type Callstack struct {
	Line uint
	Col  uint
	Show string
}

func containsString(slice []string, target string) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

func splitWithSeparators(input []Token, separators []string) []Block {
	result := make([]Block, 0, 13)
	currentSegment := make(Block, 0, 6)

	for _, item := range input {
		containsSeparator := false

		for _, separator := range separators {
			if item.Type == separator {
				containsSeparator = true
				break
			}
		}

		if containsSeparator {
			if len(currentSegment) > 0 {
				result = append(result, currentSegment)
				currentSegment = make([]Token, 0)
				continue
			}
		}

		currentSegment = append(currentSegment, item)
	}

	if len(currentSegment) > 0 {
		result = append(result, currentSegment)
	}

	return result
}
