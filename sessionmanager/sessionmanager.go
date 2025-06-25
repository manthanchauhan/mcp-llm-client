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
		*sd.CustomerRequestCategory = rc.(string)
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
