package soldier_strategy

import (
	"fmt"
)

type Adapter struct {
}

func (c *Adapter) TranslateTwoLanguages(com GreekSpeech) {
	fmt.Println("Soldier start to speak on his language.")
	com.SpeakingGreek()
}