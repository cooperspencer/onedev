package onedev

import "time"

type Client struct {
	Username string
	Password string
	Url      string
}

type Project struct {
	ID                     int       `json:"id"`
	ForkedFromID           int       `json:"forkedFromId"`
	Name                   string    `json:"name"`
	Description            string    `json:"description"`
	CreateDate             time.Time `json:"createDate"`
	UpdateDate             time.Time `json:"updateDate"`
	DefaultRoleID          int       `json:"defaultRoleId"`
	IssueManagementEnabled bool      `json:"issueManagementEnabled"`
}

type GroupAuthorization struct {
	ID        int `json:"id"`
	ProjectID int `json:"projectId"`
	GroupID   int `json:"groupId"`
	RoleID    int `json:"roleId"`
}

type UserAuthorization struct {
	ID        int `json:"id"`
	ProjectID int `json:"projectId"`
	UserID    int `json:"userId"`
	RoleID    int `json:"roleId"`
}

type ProjectSettings struct {
	BranchProtections []struct {
		Enabled           bool     `json:"enabled"`
		Branches          string   `json:"branches"`
		UserMatch         string   `json:"userMatch"`
		PreventForcedPush bool     `json:"preventForcedPush"`
		PreventDeletion   bool     `json:"preventDeletion"`
		PreventCreation   bool     `json:"preventCreation"`
		ReviewRequirement string   `json:"reviewRequirement"`
		JobNames          []string `json:"jobNames"`
		FileProtections   []struct {
			Paths             string   `json:"paths"`
			ReviewRequirement string   `json:"reviewRequirement"`
			JobNames          []string `json:"jobNames"`
		} `json:"fileProtections"`
	} `json:"branchProtections"`
	TagProtections []struct {
		Enabled         bool   `json:"enabled"`
		Tags            string `json:"tags"`
		UserMatch       string `json:"userMatch"`
		PreventUpdate   bool   `json:"preventUpdate"`
		PreventDeletion bool   `json:"preventDeletion"`
		PreventCreation bool   `json:"preventCreation"`
	} `json:"tagProtections"`
	IssueSetting struct {
		ListFields []string `json:"listFields"`
		BoardSpecs []struct {
			Name             string   `json:"name"`
			BaseQuery        string   `json:"baseQuery"`
			BacklogBaseQuery string   `json:"backlogBaseQuery"`
			IdentifyField    string   `json:"identifyField"`
			Columns          []string `json:"columns"`
			DisplayFields    []string `json:"displayFields"`
			EditColumns      []string `json:"editColumns"`
		} `json:"boardSpecs"`
		NamedQueries []struct {
			Name  string `json:"name"`
			Query string `json:"query"`
		} `json:"namedQueries"`
	} `json:"issueSetting"`
	BuildSetting struct {
		ListParams   []string `json:"listParams"`
		NamedQueries []struct {
			Name  string `json:"name"`
			Query string `json:"query"`
		} `json:"namedQueries"`
		JobSecrets []struct {
			Name               string `json:"name"`
			Value              string `json:"value"`
			AuthorizedBranches string `json:"authorizedBranches"`
		} `json:"jobSecrets"`
		BuildPreservations []struct {
			Condition string `json:"condition"`
			Count     int    `json:"count"`
		} `json:"buildPreservations"`
		ActionAuthorizations []struct {
			Type               string `json:"@type"`
			MilestoneNames     string `json:"milestoneNames"`
			AuthorizedBranches string `json:"authorizedBranches"`
		} `json:"actionAuthorizations"`
		DefaultFixedIssueFilters []struct {
			JobNames   string `json:"jobNames"`
			IssueQuery string `json:"issueQuery"`
		} `json:"defaultFixedIssueFilters"`
	} `json:"buildSetting"`
	PullRequestSetting struct {
		NamedQueries []struct {
			Name  string `json:"name"`
			Query string `json:"query"`
		} `json:"namedQueries"`
	} `json:"pullRequestSetting"`
	NamedCommitQueries []struct {
		Name  string `json:"name"`
		Query string `json:"query"`
	} `json:"namedCommitQueries"`
	NamedCodeCommentQueries []struct {
		Name  string `json:"name"`
		Query string `json:"query"`
	} `json:"namedCodeCommentQueries"`
	WebHooks []struct {
		PostURL    string   `json:"postUrl"`
		EventTypes []string `json:"eventTypes"`
		Secret     string   `json:"secret"`
	} `json:"webHooks"`
	ContributedSettings struct {
		String struct {
			Type            string `json:"@type"`
			ExampleProperty string `json:"exampleProperty"`
		} `json:"string"`
	} `json:"contributedSettings"`
}

type Milestone struct {
	ID          int       `json:"id"`
	ProjectID   int       `json:"projectId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	Closed      bool      `json:"closed"`
}

type ProjectQueryOptions struct {
	Query  string
	Offset int
	Count  int
}

type MilestoneQueryOptions struct {
	Name      string
	DueBefore time.Time
	DueAfter  time.Time
	Closed    bool
	Offset    int
	Count     int
}

type Contribution struct {
	Author struct {
		Name         string `json:"name"`
		EmailAddress string `json:"emailAddress"`
		When         int    `json:"when"`
		TzOffset     int    `json:"tzOffset"`
	} `json:"author"`
	TotalContribution struct {
		Commits   int `json:"commits"`
		Additions int `json:"additions"`
		Deletions int `json:"deletions"`
	} `json:"totalContribution"`
	DailyContributions map[string]int `json:"dailyContributions"`
}

type ContributorOptions struct {
	Type      string
	SinceDate time.Time
	UntilDate time.Time
	Count     int
}

type CreateProjectOptions struct {
	Name                   string `json:"name"`
	ForkedFromID           int    `json:"forkedFromId,omitempty"`
	Description            string `json:"description"`
	DefaultRoleID          int    `json:"defaultRoleId"`
	IssueManagementEnabled bool   `json:"issueManagementEnabled"`
}
