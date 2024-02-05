package generator

import (
	"fmt"
)

type Prompt struct {
	Content          string
	neededValueCount int
	valueCount       int
}

func (v Prompt) Fill(values ...any) Prompt {
	v.Content = fmt.Sprintf(v.Content, values...)
	v.valueCount += len(values)
	return v
}

func (v Prompt) Chain(p Prompt) Prompt {
	v.Content += "\n" + p.Content
	v.neededValueCount += p.neededValueCount
	return v
}

func (v Prompt) Finish() (string, error) {
	if v.neededValueCount != v.valueCount {
		return "", fmt.Errorf("prompt has %d values, but %d are needed", v.valueCount, v.neededValueCount)
	}

	return v.Content, nil
}

var (
	Base Prompt = Prompt{
		Content: `You are generating RPG content for a world callled "%s".
		The world is described as: "%s".
		Story log to avoid generating similar content: "%s".
		Try to not repeat things from the story log. Example: if the story log contains information that John Doe went to Bar X, you should not generate a situation where John Doe goes to Bar X.
		The narrative is focused around the main character: "%s" who is a "%s".
		It is a solo RPG, no one never joins the main character.`,
		neededValueCount: 5,
	}

	CreateEnvironment Prompt = Prompt{
		Content: `You should generate environment data that will be used in the story.
		Avoid generating environment data that is too to these previous environments: "%s".,
		You should answer with a JSON only, with the following format:
		{
			"name": "The name of the environment",
			"description": "The description of the environment",
			"characters": [
				{ 
					"name": "The name of the character",
					"story": "The story of the character",
					"race": "The race of the character",
					"class": "The class of the character"
				}, ...
			]
		}`,
		neededValueCount: 1,
	}

	CreateSituation Prompt = Prompt{
		Content: `You should generate a situation that will be used in the story.
		Previous situation information: "%s".
		The current environment is: "%s".
		Maintain continuity with the environment, previous situations and the story log.
		IMPORTANT: the outcome of the situation should always be the starting point of the next situation.
		NEVER repeat things from the story log.
		IMPORTANT: do it highly immersive and detailed. Example: if in the previous outcome the character went to his home, you should describe the home in the next situation, the player will choose what to do in the home, but if you think it is important to describe a thing that happened out of player's control, you should do it.
		If the previous situation is not empty, you should generate a situation that is a consequence of the last one, for example: if the last situation of the list was "John Doe escaped" then the new situation could start with "After John Doe escaped, he ...".
		You should answer with a JSON only, with the following format:
		{
			"situation": "The situation content"
		}`,
		neededValueCount: 2,
	}

	CreateDecisionOutcome Prompt = Prompt{
		Content: `You should generate the outcome of the situation.
		Besides the outcome content, you should fill the field "to_remember" if you detect that the outcome should be remembered for future situations. This field should be really short, like "John Doe is a liar".
		The situation is: "%s".
		The player's decision was: "%s".
		The player succeded? %t.
		Should you start a battle if the player failed? %t.
		Should you always start a battle? %t.
		If the player failed, the outcome should be a consequence of the failure.
		If the player succeeded, the outcome should be a consequence of the success.
		You should answer with a JSON only, with the following format:
		{
			"outcome": "The outcome content",
			"battle_started": false // or true,
		}`,
		neededValueCount: 5,
	}

	CreateRollForDecision Prompt = Prompt{
		Content: `You should generate a roll for the player's decision.
		The situation is: "%s".
		The player's decision was: "%s".
		Generate "level" for the roll: low (0-5), medium (6-10), high (11-15), very_high (16-20).
		Generate "fail_starts_battle" for the roll: true or false. Consider the situation as if it was in real life.
		Generate "always_starts_battle" for the roll: true or false. Consider the situation as if it was in real life.
		The "level" should consider the factibility of the decision given the situation.
		You should answer with a JSON only, with the following format:
		{
			"level": "low, medium, high, very_high",
			"fail_starts_battle": true, // or false
			"always_starts_battle": false // or true
		}`,
		neededValueCount: 2,
	}
)

func Language(language string) Prompt {
	return Prompt{
		Content:          fmt.Sprintf("Generate the content in %s. JSON keys must remain in the original language.", language),
		neededValueCount: 0,
	}
}
