package response

import system "Miyo_Assignment/model/User"

type ResigterResponse struct {
	User system.User `json:"user"`
}
