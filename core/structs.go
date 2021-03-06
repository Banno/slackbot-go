package core

type Config struct {
	Token       string    `json:"token,omitempty"`
	Last_update string    `json:"last-update"`
	Webhooks    []Hook    `json:"webhooks,omitempty"`
	Channels    []Channel `json:"channels,omitempty"`
	Team        []User    `json:"team,omitempty"`
}

type Hook struct {
	Name  string `json:"channel_name"`
	Url   string `json:"url"`
	Token string `json:"token"`
	Team  string `json:"team_id"`
}

type Channel struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type User struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

// Returns a specific channels Id
func (c *Config) GetChannelId(name string) string {
	for _, v := range c.Channels {
		if v.Name == name {
			return v.Id
		}
	}
	return ""
}

// Takes a list of channel names and returns their Slack Ids
func (c *Config) GetChannelIds(s ...string) []string {
	var r []string

	for _, v := range s {
		for _, n := range c.Channels {
			if n.Name == v {
				r = append(r, n.Id)
			}
		}
	}

	return r
}

//Returns a map of all channel names and ids
func (c *Config) GetChannels() map[string]string {
	m := make(map[string]string)
	for _, v := range c.Channels {
		m[v.Name] = v.Id
	}
	return m
}

// Returns a specific user's id
func (c *Config) GetUserId(name string) string {
	for _, v := range c.Team {
		if v.Name == name {
			return v.Id
		}
	}
	return ""
}

// Takes a list of people and returns their Slack Ids
func (c *Config) GetUserIds(s ...string) []string {
	var r []string

	for _, v := range s {
		for _, n := range c.Team {
			if n.Name == v {
				r = append(r, n.Id)
			}
		}
	}

	return r
}

// Returns a map of all names/user Ids
func (c *Config) GetUsers() map[string]string {
	m := make(map[string]string)
	for _, v := range c.Team {
		m[v.Name] = v.Id
	}
	return m
}

func (c *Config) GetToken() string {
	return c.Token
}

// Pretty sure this stuff below isn't currently used either

type AttachmentMessage struct {
	Attachments []AttachmentObject `json:"attachments,omitempty"`
}

type AttachmentObject struct {
	Fallback    string        `json:"fallback,omitempty"`
	Color       string        `json:"color,omitempty"`
	Pretext     string        `json:"pretext,omitempty"`
	Author_Name string        `json:"author_name,omitempty"`
	Author_Link string        `json:"author_link,omitempty"`
	Author_Icon string        `json:"author_icon,omitempty"`
	Title       string        `json:"title,omitempty"`
	Title_Link  string        `json:"title_link,omitempty"`
	Text        string        `json:"text,omitempty"`
	Fields      []FieldObject `json:"fields,omitempty"`
	Image_Url   string        `json:"image_url,omitempty"`
	Thumb_Url   string        `json:"thumb_url,omitempty"`
	Footer      string        `json:"footer,omitempty"`
	Footer_Icon string        `json:"footer_icon,omitempty"`
	Ts          int           `json:"ts,omitempty"`
}

type FieldObject struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

//Not currently used
type SlashRequest struct {
	Channel_Id   string `json:"channel_id,omitempty"`
	Channel_Name string `json:"channel_name,omitempty"`
	Command      string `json:"command,omitempty"`
	Response_Url string `json:"response_url,omitempty"`
	Team_Domain  string `json:"team_domain,omitempty"`
	Team_Id      string `json:"team_id,omitempty"`
	Text         string `json:"text,omitempty"`
	Token        string `json:"token,omitempty"`
	Trigger_Id   string `json:"trigger_id,omitempty"`
	User_Id      string `json:"user_id,omitempty"`
	User_Name    string `json:"user_name,omitempty"`
}

type ApprovalObject struct {
	Question string     `json:"text"`
	Buttons  Attachment `json:"attachments"`
}

type Attachment struct {
	Fallback    string   `json:"fallback"`
	Callback_Id string   `json:"callback_id"`
	Color       string   `json:"color,omitempty"`
	Choices     []Button `json:"actions"`
}

type Button struct {
	Name  string `json:"name"`
	Text  string `json:"text"`
	Type  string `json:"type"`
	Value string `json:"value"`
	Style string `json:"style,omitempty"`
}

/*
func (o *ApprovalObject) Init() {
	Question = ""
	Buttons := Attachment.Init()
}

func (o *Attachment) Init() {
	Fallback := "Unable to issue approval at this time."
	Choices := [2]Button(
		Button{text: "Approve", Type: "button", Value: "true"},
		Button{text: "Deny", Type: "button", Value: "false", Style: "danger"})
}
*/
/*
    "attachments":
        {
            "fallback": "You are unable to choose a game",
            "callback_id": "wopr_game",
            "color": "#3AA3E3",
            "actions": [
                {
                    "name": "zipapproval",
                    "text": "Approve",
                    "type": "button",
                    "value": "true"
                },
                {
                    "name": "game",
                    "text": "Deny",
                    "style": "danger",
                    "type": "button",
                    "value": "false",
                    "confirm": {
                        "title": "Are you sure?",
                        "text": "Are you sure?",
                        "ok_text": "Yes",
                        "dismiss_text": "No"
                    }
                }
            ]
        }
    ]
}
*/
