package onedev

import "time"

type Client struct {
	Username string
	Password string
	Token    string
	Url      string
}

type Project struct {
	ID                  int                 `json:"id,omitempty"`
	ForkedFromID        int                 `json:"forkedFromId,omitempty"`
	ParentID            int                 `json:"parentId,omitempty"`
	Description         string              `json:"description,omitempty"`
	CreateDate          time.Time           `json:"createDate,omitempty"`
	DefaultRoleID       int                 `json:"defaultRoleId,omitempty"`
	Name                string              `json:"name,omitempty"`
	CodeManagement      bool                `json:"codeManagement,omitempty"`
	IssueManagement     bool                `json:"issueManagement,omitempty"`
	GitPackConfig       GitPackConfig       `json:"gitPackConfig,omitempty"`
	CodeAnalysisSetting CodeAnalysisSetting `json:"codeAnalysisSetting,omitempty"`
}
type GitPackConfig struct {
	WindowMemory  string `json:"windowMemory,omitempty"`
	PackSizeLimit string `json:"packSizeLimit,omitempty"`
	Threads       string `json:"threads,omitempty"`
	Window        string `json:"window,omitempty"`
}
type CodeAnalysisSetting struct {
	AnalysisFiles string `json:"analysisFiles,omitempty"`
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

type IssueQueryOptions struct {
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
	ForkedFromID        int                 `json:"forkedFromId,omitempty"`
	ParentID            int                 `json:"parentId,omitempty"`
	Description         string              `json:"description,omitempty"`
	DefaultRoleID       int                 `json:"defaultRoleId,omitempty"`
	Name                string              `json:"name,omitempty"`
	CodeManagement      bool                `json:"codeManagement,omitempty"`
	IssueManagement     bool                `json:"issueManagement,omitempty"`
	GitPackConfig       GitPackConfig       `json:"gitPackConfig,omitempty"`
	CodeAnalysisSetting CodeAnalysisSetting `json:"codeAnalysisSetting,omitempty"`
}

type Issue struct {
	ID            int       `json:"id"`
	State         string    `json:"state"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	NumberScopeID int       `json:"numberScopeId"`
	ProjectID     int       `json:"projectId"`
	MilestoneID   int       `json:"milestoneId"`
	SubmitterID   int       `json:"submitterId"`
	SubmitterName string    `json:"submitterName"`
	SubmitDate    time.Time `json:"submitDate"`
	VoteCount     int       `json:"voteCount"`
	CommentCount  int       `json:"commentCount"`
	UUID          string    `json:"uuid"`
	Number        int       `json:"number"`
	LastUpdate    struct {
		UserID   int       `json:"userId"`
		UserName string    `json:"userName"`
		Date     time.Time `json:"date"`
		Activity string    `json:"activity"`
	} `json:"lastUpdate"`
}

type IssueField struct {
	ID      int    `json:"id"`
	IssueID int    `json:"issueId"`
	Name    string `json:"name"`
	Value   string `json:"value"`
	Type    string `json:"type"`
	Ordinal int    `json:"ordinal"`
}

type IssueChange struct {
	ID       int       `json:"id"`
	IssueID  int       `json:"issueId"`
	UserID   int       `json:"userId"`
	UserName string    `json:"userName"`
	Date     time.Time `json:"date"`
	Data     struct {
		Type         string `json:"@type"`
		OldState     string `json:"oldState"`
		NewState     string `json:"newState"`
		OldMilestone string `json:"oldMilestone"`
		NewMilestone string `json:"newMilestone"`
		Comment      string `json:"comment"`
		OldFields    struct {
			String struct {
				Name   string   `json:"name"`
				Type   string   `json:"type"`
				Values []string `json:"values"`
			} `json:"string"`
		} `json:"oldFields"`
		NewFields struct {
			String struct {
				Name   string   `json:"name"`
				Type   string   `json:"type"`
				Values []string `json:"values"`
			} `json:"string"`
		} `json:"newFields"`
	} `json:"data"`
}

type Comment struct {
	ID       int       `json:"id"`
	IssueID  int       `json:"issueId"`
	UserID   int       `json:"userId"`
	UserName string    `json:"userName"`
	Date     time.Time `json:"date"`
	Content  string    `json:"content"`
}

type Vote struct {
	ID      int       `json:"id"`
	IssueID int       `json:"issueId"`
	UserID  int       `json:"userId"`
	Date    time.Time `json:"date"`
}

type Watch struct {
	ID       int  `json:"id"`
	IssueID  int  `json:"issueId"`
	UserID   int  `json:"userId"`
	Watching bool `json:"watching"`
}

type PullRequest struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	CloseInfo struct {
		UserID   int       `json:"userId"`
		UserName string    `json:"userName"`
		Date     time.Time `json:"date"`
		Status   string    `json:"status"`
	} `json:"closeInfo"`
	Description                 string    `json:"description"`
	SubmitterID                 int       `json:"submitterId"`
	SubmitterName               string    `json:"submitterName"`
	NumberScopeID               int       `json:"numberScopeId"`
	TargetProjectID             int       `json:"targetProjectId"`
	TargetBranch                string    `json:"targetBranch"`
	SourceProjectID             int       `json:"sourceProjectId"`
	SourceBranch                string    `json:"sourceBranch"`
	BaseCommitHash              string    `json:"baseCommitHash"`
	LastCodeCommentActivityDate time.Time `json:"lastCodeCommentActivityDate"`
	SubmitDate                  time.Time `json:"submitDate"`
	MergeStrategy               string    `json:"mergeStrategy"`
	UUID                        string    `json:"uuid"`
	Number                      int       `json:"number"`
	CommentCount                int       `json:"commentCount"`
	LastUpdate                  struct {
		UserID   int       `json:"userId"`
		UserName string    `json:"userName"`
		Date     time.Time `json:"date"`
		Activity string    `json:"activity"`
	} `json:"lastUpdate"`
	CheckError string `json:"checkError"`
}

type CreateIssueOptions struct {
	ProjectID   int               `json:"projectId"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	MilestoneID int               `json:"milestoneId"`
	Fields      map[string]string `json:"fields"`
}

type StateTransitionOptions struct {
	State        string            `json:"state"`
	Fields       map[string]string `json:"fields"`
	RemoveFields []string          `json:"removeFields"`
	Comment      string            `json:"comment"`
}

type MergePreview struct {
	TargetHeadCommitHash string `json:"targetHeadCommitHash"`
	HeadCommitHash       string `json:"headCommitHash"`
	MergeStrategy        string `json:"mergeStrategy"`
	MergeCommitHash      string `json:"mergeCommitHash"`
}

type Assignment struct {
	ID        int `json:"id"`
	UserID    int `json:"userId"`
	RequestID int `json:"requestId"`
}

type Review struct {
	ID        int `json:"id"`
	UserID    int `json:"userId"`
	RequestID int `json:"requestId"`
	Result    struct {
		Commit   string `json:"commit"`
		Approved bool   `json:"approved"`
		Comment  string `json:"comment"`
	} `json:"result"`
}

type Update struct {
	ID                   int       `json:"id"`
	RequestID            int       `json:"requestId"`
	HeadCommitHash       string    `json:"headCommitHash"`
	TargetHeadCommitHash string    `json:"targetHeadCommitHash"`
	Date                 time.Time `json:"date"`
}

type Build struct {
	ID            int       `json:"id"`
	NumberScopeID int       `json:"numberScopeId"`
	ProjectID     int       `json:"projectId"`
	SubmitterID   int       `json:"submitterId"`
	SubmitterName string    `json:"submitterName"`
	CancellerID   int       `json:"cancellerId"`
	CancellerName string    `json:"cancellerName"`
	JobName       string    `json:"jobName"`
	JobWorkspace  string    `json:"jobWorkspace"`
	RefName       string    `json:"refName"`
	Version       string    `json:"version"`
	Number        int       `json:"number"`
	CommitHash    string    `json:"commitHash"`
	Status        string    `json:"status"`
	SubmitDate    time.Time `json:"submitDate"`
	PendingDate   time.Time `json:"pendingDate"`
	RunningDate   time.Time `json:"runningDate"`
	FinishDate    time.Time `json:"finishDate"`
	RetryDate     time.Time `json:"retryDate"`
	SubmitReason  string    `json:"submitReason"`
	ErrorMessage  string    `json:"errorMessage"`
	RequestID     int       `json:"requestId"`
}

type Change struct {
	ID        int       `json:"id"`
	RequestID int       `json:"requestId"`
	UserID    int       `json:"userId"`
	UserName  string    `json:"userName"`
	Date      time.Time `json:"date"`
	Data      struct {
		Type    string `json:"@type"`
		Comment string `json:"comment"`
	} `json:"data"`
}

type PullRequestOptions struct {
	TargetProjectID int    `json:"targetProjectId"`
	SourceProjectID int    `json:"sourceProjectId"`
	TargetBranch    string `json:"targetBranch"`
	SourceBranch    string `json:"sourceBranch"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	MergeStrategy   string `json:"mergeStrategy"`
	ReviewerIds     []int  `json:"reviewerIds"`
	AssigneeIds     []int  `json:"assigneeIds"`
}

type CreateUserOptions struct {
	Name         string `json:"name,omitempty"`
	Password     string `json:"password,omitempty"`
	FullName     string `json:"fullName,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
}

type CloneUrl struct {
	HTTP string `json:"http,omitempty"`
	SSH  string `json:"ssh,omitempty"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
}

type UserMemebership struct {
	ID      int `json:"id"`
	UserID  int `json:"userId"`
	GroupID int `json:"groupId"`
}

type Group struct {
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	Description        interface{} `json:"description"`
	Administrator      bool        `json:"administrator"`
	CreateRootProjects bool        `json:"createRootProjects"`
	Enforce2FA         bool        `json:"enforce2FA"`
}

type Artifact struct {
	Type     string `json:"@type"`
	Children []struct {
		Type         string      `json:"@type"`
		Children     interface{} `json:"children,omitempty"`
		Path         string      `json:"path"`
		LastModified int         `json:"lastModified"`
		Length       int         `json:"length,omitempty"`
		MediaType    interface{} `json:"mediaType,omitempty"`
	} `json:"children"`
	Path         string `json:"path"`
	LastModified int    `json:"lastModified"`
}

type CommitQueryOptions struct {
	Query string
	Count int
}

type Commit struct {
	CommitHash string `json:"commitHash"`
	Author     struct {
		Name         string `json:"name"`
		EmailAddress string `json:"emailAddress"`
		When         int64  `json:"when"`
		TzOffset     int    `json:"tzOffset"`
	} `json:"author"`
	Committer struct {
		Name         string `json:"name"`
		EmailAddress string `json:"emailAddress"`
		When         int64  `json:"when"`
		TzOffset     int    `json:"tzOffset"`
	} `json:"committer"`
	CommitMessage string `json:"commitMessage"`
}
