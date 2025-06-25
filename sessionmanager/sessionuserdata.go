package sessionmanager

type SessionUserData struct {
	Id        *int    `json:"id"`
	Mobile    *string `json:"mobile"`
	Email     *string `json:"email"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}
