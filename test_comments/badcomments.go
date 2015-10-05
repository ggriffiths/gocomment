package comments

// LargeFunction is a function with a comment that is only a little bit too long for us.
func LargeFunction() string {
	return nil
}

// MultipleLinesCommentNoPunctuationFirstLine is a function with a comment that is too long
// on the first line and has no punctuation on the first or second line
func MultipleLinesCommentNoPunctuationFirstLine() string {
	return nil
}

// MultipleLinesCommentNoPunctuationAnywhere is a function with no punctuation on
// the first line and not on the second line either
func MultipleLinesCommentNoPunctuationAnywhere() string {
	return nil
}

// MultipleLinesCommentNoPunctuationOnOnlySecondLine is a function with punctuation.
// on the first line and not on the second line
func MultipleLinesCommentNoPunctuationOnOnlySecondLine() string {
	return nil
}

// HugeFunction is a funcion with a comment that is really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really long.
func HugeFunction() string {
	return nil
}

// HugeFunctionAgain is another funcion with a comment that is really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really really long.
func HugeFunctionAgain() string {
	return nil
}

// GoodFunction is a function without a period at the end :(
func GoodFunction() string {
	return nil
}

// GoodFunction is a function without a period at the end and is way too long. Get it together :(
func GoodFunction() string {
	return nil
}
