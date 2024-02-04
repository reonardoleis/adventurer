package generator

import (
	"fmt"
	"strings"
)

type Prompt struct {
	Content    string
	fieldCount int
}

func (v Prompt) Fill(values ...any) (Prompt, error) {
	if len(values) != v.fieldCount {
		return v, fmt.Errorf(
			"expected %d values, got %d",
			v.fieldCount,
			len(values),
		)
	}

	v.Content = fmt.Sprintf(v.Content, values...)
	return v, nil
}

func (v Prompt) Chain(p Prompt) Prompt {
	v.Content += "\n" + p.Content
	v.fieldCount += p.fieldCount
	return v
}

func (v Prompt) String() string {
	return v.Content
}

func SetBasePromptLanguage(language string) {
	Base.Content = strings.Replace(
		Base.Content,
		"English",
		language,
		-1,
	)
}

var (
	Base Prompt = Prompt{
		Content: `You are generating an environment for a "%s" RPG.
		The RPG is set in a "%s" world.
		You do not need to always be totally focused on the theme, but it should be the main focus.
		Example: if the theme is on a Cyberpunk world with rogue Androids, you can have situation that the rogue Androids are not present, preferring the focus on the Cyberpunk world.
		Always be focusing on maintaining the theme, but do not be afraid to be creative.
		This is a solo RPG, so never mention other players nor let other characters join the main character.
		The USER has preference to decide things about the main character, so if the player says that the main character do not like charity, then you can't make the main character do charity and you must consider that when generating the content.
		Generate in the following language: English.
		Always do the shortest you can, as if it were messages in a chat.`,
		fieldCount: 2,
	}

	NumberOfCharacters Prompt = Prompt{
		Content:    `Number of characters on the environment other than the player (never in a group with the player): %d`,
		fieldCount: 1,
	}

	MainCharacterInformation Prompt = Prompt{
		Content: `Main character name and story: "%s"
		Main character race and class: "%s"`,
		fieldCount: 2,
	}

	LastSituation Prompt = Prompt{
		Content: `Last situation: "%s"
		Did the player failed last situation: %t`,
		fieldCount: 2,
	}

	CurrentEnvironment Prompt = Prompt{
		Content:    `Current environment: "%s"`,
		fieldCount: 1,
	}

	LastPlayerDecision Prompt = Prompt{
		Content: `Last player decision: "%s"
		Did it start a battle: %t
		Was it a success: %t`,
		fieldCount: 3,
	}

	OtherCharactersInformation Prompt = Prompt{
		Content:    `Other characters on the environment are: %s`,
		fieldCount: 1,
	}

	PromptGenerateEnvironment = Prompt{
		Content: `Generate the name of the environment, the description of the environment, a list of character that are in the environment.
		You should answer in the following format, with it only:
	
		{
			"name": "Environment 1",
			"description": "Desc 1",
			"characters": [
				{
					"name": "Character 1",
					"story": "Story 1",
					"is_hostile": true
				},
				...add more characters here until you reach the number of characters...
			]
		}`,
		fieldCount: 0,
	}

	PromptGenerateSituation = Prompt{
		Content: `The new situation must be related to the environment and have continuity with the last situation.
		Move on from the last situation to a new one, or create a new one if there was none.
		Be creative, do not repeat the same situation twice.
		The situation does not have to be one hundred percent related to the environment, but it should be possible to happen in it.
		If the last situation is empty, consider that you should generate the first situation, the RPG campaign is starting, so make it interesting.
		The situation do not have to include the characters, be creative. It can be as simple as the characters finding a treasure or something on the floor, or as complex as a battle with all the characters.
		Generate a situation that the characters can find themselves in.
		
		Consider the STORY_LOG content to help you create the situation, avoid repeating the same situation and avoiding making unrelated situations.
		Should it start a battle? %t
		`,
		fieldCount: 1,
	}

	PromptGenerateDecision = Prompt{
		Content: `List of things to generate:
		NUMBER: The number of dice to roll you find appropriate, from 1 to 20.
			ALWAYS_START_BATTLE: true or false, if the decision always starts a battle. Be evil here.
			FAIL_START_BATTLE: true or false, if the decision starts a battle when failed. Be evil here.

			Examples on "difficulty of the decision":
			1. If the decision considering the situation is easy, the number of dice to roll should be low (0 to 8)
			2. If the decision considering the situation is medium, the number of dice to roll should be medium (8 to 15)
			3. If the decision considering the situation is hard, the number of dice to roll should be high (15 to 20)

			Example of easy decision as if the situation were "The character is walking in a safe forest": "I decide to smile!"
			Example of medium decision as if the situation were "The character is walking in a forest at night": "I decide to explore the area!"
			Example of hard decision as if the situation were "The character is walking in a forest at night and hears a noise": "I yell at the noise, come at me!"

			IMPORTANT: if it is a "pathetic" decision without any chance of failure (based on the decision + situation), roll 0 and no battle.
			Prioritize situations with low or 0 rolls, and no battle. If the situation is hard, prioritize situations with high rolls and battle.
			Your return format should follow the following EXAMPLE format, answering only with the values, in the order they are listed above, separated by two semicolons, like this:
		
			NUMBER;;ALWAYS_START_BATTLE;;FAIL_START_BATTLE
			
			Replace the values with the ones you generated. Never label the values, only the values separated by two semicolons.`,
		fieldCount: 0,
	}

	PromptGenerateDecisionOutcome = Prompt{
		Content:    `Generate the outcome text for the given decision.`,
		fieldCount: 0,
	}
)
