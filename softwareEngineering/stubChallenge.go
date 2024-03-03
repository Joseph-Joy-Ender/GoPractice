package main

import "image"

type stubChallenger string

func (c stubChallenger) Challenge() (image.Image, string) {
	return image.NewRGBA(image.Rect(0, 0, 100, 100)), string(c)

}

type stubPrompter string

func (p stubPrompter) Prompt(_ image.Image) string {
	return string(p)

}
