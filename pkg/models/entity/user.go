package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId     string
	Username   string
	Fname      string
	Lname      string
	Country    string
	State      string
	City       string
	Password   string
	Email      string
	CreateTime time.Time
	//	bool banned
	//	Timestamp bannedTime;
	//	private String city;
	//	private String code;
	//	private String createTime;
	//	private String registeredTime;
	//	private String activatedTime;
	//    private String email;
	//	private Integer lastAction;
	//	private Timestamp lastLoginTime;
	//	private Timestamp lastSessionTime;
	//	private float lat;
	//	private float lon;
	//	private Timestamp prevLastLoginTime;
	//	private Timestamp prevLastSessionTime;
	//	private String salt;
	//	private String sessionId;
	//	private String region;
	//	private Timestamp terminationTime;
	//  private UUID registrationToken;
	//    private Timestamp registrationExpirationTime;
	//    private UnitLocaleTypeMapEnum unitLocale;

	// private UserStatusCodeTypeMapEnum userStatusCode;
}

func NewUser(email string, fname string, lname string, country string, state string, city string) (*User, error) {

	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	user := &User{
		UserId:     uuid.NewString(),
		Username:   email,
		Fname:      fname,
		Lname:      lname,
		Country:    country,
		State:      state,
		City:       city,
		Email:      email,
		CreateTime: time.Now(),
	}

	return user, nil
}
