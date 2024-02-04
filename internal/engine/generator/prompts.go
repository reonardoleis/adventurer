package generator

const (
	PromptGenerateEnvironment = `You are generating an environment for a "%s" RPG.

	The RPG is set in a "%s" world.

	Generate the name of the environment, the description of the environment, a list of character that are in the environment.

	Number of characters: %d

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
	}
	`

	PromptGenerateSituation = `You are generating a situation for a "%s" RPG.

	The RPG is set in a "%s" world.

	This is a solo RPG, so the narrative is focused on the main character.
	Main character name and story: "%s"
	Main character race and class: "%s"

	Last situation was (ignore if empty): "%s"
	Failed last situation (ignore if last situation is empty): %t

	The new situation must be related to the environment and have continuity with the last situation.
	Move on from the last situation to a new one, or create a new one if there was none.
	Be creative, do not repeat the same situation twice.
	The situation does not have to be one hundred percent related to the environment, but it should be possible to happen in it.
	If the last situation is empty, consider that you should generate the first situation, the RPG campaign is starting, so make it interesting.
	The situation do not have to include the characters, be creative. It can be as simple as the characters finding a treasure or something on the floor, or as complex as a battle with all the characters.
	Current environment: "%s"
	Characters in the environment other than the main character: %s
	Starts battle: %t

	Generate a situation that the characters can find themselves in.
	`

	PromptGenerateDecision = `You are generating a dice roll for a "%s" RPG.
		The RPG is set in a "%s" world.

		Current environment: "%s"
		Situation: "%s"
		Player decision: "%s"

		List of things to generate:

		NUMBER: The number of dice to roll you find appropriate, from 1 to 20.
		ALWAYS_START_BATTLE: true or false, if the decision always starts a battle. Be evil here.
		FAIL_START_BATTLE: true or false, if the decision starts a battle when failed. Be evil here.

		Your return format should follow the following EXAMPLE format, answering only with the values, in the order they are listed above, separated by two semicolons, like this:

		NUMBER;;ALWAYS_START_BATTLE;;FAIL_START_BATTLE`

	PromptGenerateDecisionOutcome = `You are generating the outcome of a decision for a "%s" RPG.
		The RPG is set in a "%s" world.

		The situation the player was in: "%s"
		Their decision: "%s"
		Was it a success? %t
		Did it start a battle? %t

		Generate the outcome text for the decision.
		`
)
