package api

var (
	MsgPlus429 = `
	{
		"detail": {
		  "clears_in": 252,
		  "code": "model_cap_exceeded",
		  "message": "You have sent too many messages to the model. Please try again later."
		}
	  }
	`

	MsgMod400 = `
	{
		"detail": {
		  "code": "flagged_by_moderation",
		  "message": "This content may violate [OpenAI Usage Policies](https://openai.com/policies/usage-policies)."
		}
	}
	`
)
