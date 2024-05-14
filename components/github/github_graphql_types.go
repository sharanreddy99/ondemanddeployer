package github

import "github.com/shurcooL/githubv4"

type AllStatsQuery struct {
	User struct {
		RepositoriesContributedTo struct {
			TotalCount githubv4.Int `json:"totalCount"`
			Nodes      []struct {
				Owner struct {
					Login     githubv4.String `json:"login"`
					AvatarURL githubv4.URI    `json:"avatarURL"`
					Typename  githubv4.String `graphql:"__typename" json:"__typename"`
				} `json:"owner"`
			} `json:"nodes"`
		} `graphql:"repositoriesContributedTo(last: 100)"`
		PinnedItems struct {
			TotalCount githubv4.Int `json:"totalCount"`
			Nodes      []struct {
				Repository struct {
					ID          githubv4.ID       `json:"id"`
					Name        githubv4.String   `json:"name"`
					CreatedAt   githubv4.DateTime `json:"createdAt"`
					URL         githubv4.URI      `json:"url"`
					Description githubv4.String   `json:"description"`
					IsFork      githubv4.Boolean  `json:"isFork"`
					Languages   struct {
						Nodes []struct {
							Name githubv4.String `json:"name"`
						} `json:"nodes"`
					} `graphql:"languages(first: 10)" json:"languages"`
				} `graphql:"... on Repository"`
			} `json:"nodes"`
		} `graphql:"pinnedItems(first: 6, types: REPOSITORY)"`
		Issues struct {
			TotalCount githubv4.Int `json:"totalCount"`
			Nodes      []struct {
				ID        githubv4.ID       `json:"id"`
				Closed    githubv4.Boolean  `json:"closed"`
				Title     githubv4.String   `json:"title"`
				CreatedAt githubv4.DateTime `json:"createdAt"`
				URL       githubv4.URI      `json:"url"`
				Number    githubv4.Int      `json:"number"`
				Assignees struct {
					Nodes []struct {
						AvatarURL githubv4.URI    `json:"avatarURL"`
						Name      githubv4.String `json:"name"`
						URL       githubv4.URI    `json:"url"`
					} `json:"nodes"`
				} `graphql:"assignees(first: 100)" json:"assignees"`
				Repository struct {
					Name  githubv4.String `json:"name"`
					URL   githubv4.URI    `json:"url"`
					Owner struct {
						Login     githubv4.String `json:"login"`
						AvatarURL githubv4.URI    `json:"avatarURL"`
						URL       githubv4.URI    `json:"url"`
					}
				} `json:"repository"`
			} `json:"nodes"`
		} `graphql:"issues(last: 100, orderBy: {field:CREATED_AT, direction: DESC})"`
		PullRequests struct {
			TotalCount githubv4.Int `json:"totalCount"`
			Nodes      []struct {
				ID       githubv4.ID     `json:"id"`
				Title    githubv4.String `json:"title"`
				URL      githubv4.URI    `json:"url"`
				State    githubv4.String `json:"state"`
				MergedBy *struct {
					AvatarURL githubv4.URI    `json:"avatarURL"`
					URL       githubv4.URI    `json:"url"`
					Login     githubv4.String `json:"login"`
				} `json:"mergedBy"`
				CreatedAt      githubv4.DateTime `json:"createdAt"`
				Number         githubv4.Int      `json:"number"`
				ChangedFiles   githubv4.Int      `json:"changedFiles"`
				Additions      githubv4.Int      `json:"additions"`
				Deletions      githubv4.Int      `json:"deletions"`
				BaseRepository struct {
					Name  githubv4.String `json:"name"`
					URL   githubv4.URI    `json:"url"`
					Owner struct {
						AvatarURL githubv4.URI    `json:"avatarURL"`
						Login     githubv4.String `json:"login"`
						URL       githubv4.URI    `json:"url"`
					} `json:"owner"`
				} `graphql:"baseRepository" json:"baseRepository"`
			} `json:"nodes"`
		} `graphql:"pullRequests(last: 100, orderBy: {field: CREATED_AT, direction: DESC})"`
	} `graphql:"user(login: $username)"`
}
