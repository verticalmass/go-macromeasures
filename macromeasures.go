package macromeasures

import "fmt"

// SharedResponse includes the Complete and Error booleans
type SharedResponse struct {
	Complete     bool   `json:"complete"`
	Error        bool   `json:"error"`
	ErrorMessage string `json:"message"`
}

// UserResponse is the user response from Macromeasures
type UserResponse struct {
	SharedResponse
	Labels map[string]User `json:"labels"`
}

func (u UserResponse) Users() ([]MacroUser, error) {
	if u.Error {
		return nil, fmt.Errorf("%s: api error: %v", libraryName, u.ErrorMessage)
	}
	if u.Labels == nil || len(u.Labels) == 0 {
		return nil, fmt.Errorf("%s: no users returned", libraryName)
	}
	users := make([]MacroUser, 0)
	for id, user := range u.Labels {
		users = append(users, MacroUser{id, user.Gender, user.All(), user.Language, user.Location, user.Platform, user.Type})
	}
	return users, nil
}

type MacroUser struct {
	ID        string              `json:"id"`
	Gender    *Gender             `json:"gender"`
	Interests []InterestsResponse `json:"interests"`
	Language  *Language           `json:"language"`
	Location  *Location           `json:"location"`
	Platform  *Platform           `json:"platform"`
	Type      *Type               `json:"type"`
}

type User struct {
	Valid     bool       `json:"valid"`
	Gender    *Gender    `json:"gender"`
	Interests *Interests `json:"interests"`
	Language  *Language  `json:"language"`
	Location  *Location  `json:"location"`
	Platform  *Platform  `json:"platform"`
	Type      *Type      `json:"type"`
}

func (u User) All() []InterestsResponse {
	all := make([]InterestsResponse, 0)
	for id, resp := range u.Interests.All {
		all = append(all, InterestsResponse{id, resp})
	}
	return all
}

type InterestsResponse struct {
	ID string `json:"id"`
	UserInterests
}

// Interests
// This attribute contains a structured list of psychographic characteristics that we've inferred about the user, organised in a tree structure in which branches are of arbitrary depths and each interest can have multiple parents.
type Interests struct {
	All       map[string]UserInterests `json:"all"`
	Confirmed bool                     `json:"confirmed"`
	Updated   Time                     `json:"updated"`
}

// UserInterests
// This structure contains over 4000 distinct interests, ranging from likely life stage characters (e.g., "Student", or "Parent", or "Married") to past and present TV shows (e.g., "Scream Queens" or "Breaking Bad") to brands (e.g., "Whole Foods" or "Urban Outfitters") to celebrities of all sorts (e.g., "Tiger Woods" or "Alex from Target") to print and digital publications (e.g., "PopSugar" or "Scientific American") to general hobbies and interests (e.g., "Golf", or "Korean pop music", or "Frozen yogurt").
type UserInterests struct {
	Category string   `json:"category"`
	Display  string   `json:"display"`
	Level    string   `json:"level"`
	Name     string   `json:"name"`
	Parents  []string `json:"parents"`
	Score    int      `json:"score"`
	Useful   bool     `json:"useful"`
}

// Type
// This attribute indicates whether or not we this account is classified as a personal account.
// The "personal" field will be true if we classify this account as a personal account and false otherwise.
// We define an account as being "personal" if:
// 1) there is one primary user managing the account;
// 2) it makes sense to assign a gender to the account;
// 3) the account is a reasonable representation of the person managing it;
// 4) the account is not clearly spam or a bot.
// Note that we do not return gender, location, or interest data for accounts that have been classified as not personal.
type Type struct {
	Confirmed bool `json:"confirmed"`
	Personal  bool `json:"personal"`
	Updated   Time `json:"updated"`
}

// Gender
// This attribute contains the inferred gender of the user.
// This attribute has only one field, "label", whose value must be "M" for male, "F" for female, or "?".
// We use a variety of signals to infer gender, including the name, username, bio, profile URL, tweets/posts, and network of a user.
// The label "?" is used for all non-personal accounts. It is also used for any user whose gender we cannot infer with reasonable confidence as well as for users who explicitly do not identify as either male or female.
// Confidence scores for gender are currently not available. If confidence scores would be useful for your application, let us know and we'll make it a priority.
type Gender struct {
	Confirmed bool   `json:"confirmed"`
	Label     string `json:"label"`
	Updated   Time   `json:"updated"`
}

// Platform
// This attribute -- only available for Twitter -- contains a list of the devices that the user uses to tweet, as inferred by the "source" field of recent tweets.
// The "primary" field contains the name of the user's primary tweeting device, while the "recent" field contains a sorted list of all detected tweeting devices, each identified by name and associated with a UNIX timestamp indicating the most recent usage (with most recent devices listed first).
type Platform struct {
	Confirmed bool           `json:"confirmed"`
	Updated   int64          `json:"updated"`
	Primary   *UserPlatform  `json:"primary"`
	Recent    []UserPlatform `json:"recent"`
}

// UserPlatform
// Possible device names include: Android, Android Tablet, Blackberry, iOS, iPad, iPhone, Mac, Mobile, Windows Phone, Computer.
type UserPlatform struct {
	Name      string `json:"name"`
	Timestamp Time   `json:"timestamp"`
}

// Location
// This attribute contains the user's inferred primary location.
type Location struct {
	Confirmed bool          `json:"confirmed"`
	Updated   Time          `json:"updated"`
	Primary   *UserLocation `json:"primary"`
}

// UserLocation
// This is inferred using the user's profile content and network to the most precise degree of granularity possible:
// either "city", or "subdivision", or "country", or "unknown".
type UserLocation struct {
	City        string        `json:"city"`
	Country     *LocationMeta `json:"country"`
	Display     string        `json:"display"`
	Granularity string        `json:"granularity"`
	Latitude    float64       `json:"latitude"`
	Longitude   float64       `json:"longitude"`
	Subdivision *LocationMeta `json:"subdivision"`
}

// LocationMeta
type LocationMeta struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Language
// This attribute contains the inferred languages used by the user.
// The "recent" field contains a list of languages detected in the user's recent tweets, for Twitter users, or post captions, for Instagram users. These recent languages are associated with the UNIX timestamp of their last detected usage, and are sorted with the most recent timestamps first. If the user is protected or does not have enough tweets or posts, the "recent" field will be empty.
// The "primary" field contains a single language that we have identified as the user's primary language, based on the proportion of tweets/posts identified as using that language. If we are unable to identify the language of the user's tweets/posts, we will look at the accounts that the user is following as well as the content in the user's profile.
type Language struct {
	Confirmed bool           `json:"confirmed"`
	Updated   Time           `json:"updated"`
	Primary   *UserLanguage  `json:"primary"`
	Recent    []UserLanguage `json:"recent"`
}

// UserLanguage
// Each language comes with a human-readable name and its corresponding ISO 639-1 code. Language locales are not supported because they are better represented by the location field.
type UserLanguage struct {
	ISOCode   string `json:"iso_code"`
	Name      string `json:"name"`
	Timestamp Time   `json:"timestamp"`
}
