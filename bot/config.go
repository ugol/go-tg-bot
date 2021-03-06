package bot

type Token string

type Config struct {
	Bot struct {
		AccessControl struct {
			Groups []struct {
				Group struct {
					ID             int64    `yaml:"id"`
					URL            string   `yaml:"url"`
					WelcomeMessage string   `yaml:"welcome_message"`
					CrossPostTags  []string `yaml:"crosspost_tags"`
					Admin          struct {
						IDs         []int    `yaml:"ids"`
						Permissions []string `yaml:"permissions"`
					} `yaml:"admin"`
					Moderator struct {
						IDs         []int    `yaml:"ids"`
						Permissions []string `yaml:"permissions"`
					} `yaml:"moderator"`
				} `yaml:"group"`
			} `yaml:"groups"`
		} `yaml:"access_control"`
	} `yaml:"bot"`
}

func (c *Config) CheckAdminPermissions(chatID int64, userID int, operation string) bool {
	var (
		userIDFound    = false
		operationFound = false
	)

	for _, g := range c.Bot.AccessControl.Groups {
		if g.Group.ID == chatID {
			for _, id := range g.Group.Admin.IDs {
				if id == userID {
					userIDFound = true
				}
			}
			for _, perm := range g.Group.Admin.Permissions {
				if perm == operation {
					operationFound = true
				}
			}
		}
	}
	return userIDFound && operationFound
}

func (c *Config) isAdmin(chatID int64, userID int) bool {
	for _, g := range c.Bot.AccessControl.Groups {
		if g.Group.ID == chatID {
			for _, id := range g.Group.Admin.IDs {
				if id == userID {
					return true
				}
			}
		}
	}
	return false
}

func (c *Config) isModerator(chatID int64, userID int) bool {
	for _, g := range c.Bot.AccessControl.Groups {
		if g.Group.ID == chatID {
			for _, id := range g.Group.Moderator.IDs {
				if id == userID {
					return true
				}
			}
		}
	}
	return false
}
