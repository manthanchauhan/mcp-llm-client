package sessionmanager

import "time"

var SingletonSessionManager *SessionManager = nil

type SessionManager struct {
	SessionCollection map[int64]*SessionData
}

func (sm *SessionManager) CreateSession() *SessionData {
	id := time.Now().Unix()
	sd := SessionData{SessionId: id, UserData: nil}
	sm.SessionCollection[id] = &sd
	return &sd
}

func (sm *SessionManager) GetSession(sid int64) *SessionData {
	if sd, ok := sm.SessionCollection[sid]; !ok {
		return nil
	} else {
		return sd
	}
}

func (sm *SessionManager) EnrichSessionData(sid int64, info *map[string]any) error {
	sd := sm.GetSession(sid)

	if rc, ok := (*info)["request_category"]; ok {
		rcStr := rc.(string)
		(*sd).CustomerRequestCategory = &rcStr
	}

	if ud, ok := (*info)["user_data"]; ok {
		if (*sd).UserData == nil {
			(*sd).UserData = &SessionUserData{}
		}

		udMap := ud.(map[string]any)

		if mobile, ok2 := udMap["USER_MOBILE_NUMBER"]; ok2 {
			mobileStr := mobile.(string)
			(*sd).UserData.Mobile = &mobileStr
		}

		if email, ok2 := udMap["email"]; ok2 {
			emailStr := email.(string)
			(*sd).UserData.Email = &emailStr
		}

		if fn, ok2 := udMap["first_name"]; ok2 {
			fnStr := fn.(string)
			(*sd).UserData.FirstName = &fnStr
		}

		if ln, ok2 := udMap["last_name"]; ok2 {
			lnStr := ln.(string)
			(*sd).UserData.LastName = &lnStr
		}
	}

	return nil
}

func GetOrCreateSessionManager() *SessionManager {
	if SingletonSessionManager == nil {
		SingletonSessionManager = &SessionManager{
			SessionCollection: map[int64]*SessionData{},
		}
	}

	return SingletonSessionManager
}
