// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
)

// UsersService handles communication with the user related
// methods of the GitHub API.
//
// GitHub API docs: https://docs.github.com/en/rest/users/
type UsersService service

// User represents a GitHub user.
type User struct {
	Login                   *string    `json:"login,omitempty"`
	ID                      *int64     `json:"id,omitempty"`
	NodeID                  *string    `json:"node_id,omitempty"`
	AvatarURL               *string    `json:"avatar_url,omitempty"`
	HTMLURL                 *string    `json:"html_url,omitempty"`
	GravatarID              *string    `json:"gravatar_id,omitempty"`
	Name                    *string    `json:"name,omitempty"`
	Company                 *string    `json:"company,omitempty"`
	Blog                    *string    `json:"blog,omitempty"`
	Location                *string    `json:"location,omitempty"`
	Email                   *string    `json:"email,omitempty"`
	Hireable                *bool      `json:"hireable,omitempty"`
	Bio                     *string    `json:"bio,omitempty"`
	TwitterUsername         *string    `json:"twitter_username,omitempty"`
	PublicRepos             *int       `json:"public_repos,omitempty"`
	PublicGists             *int       `json:"public_gists,omitempty"`
	Followers               *int       `json:"followers,omitempty"`
	Following               *int       `json:"following,omitempty"`
	CreatedAt               *Timestamp `json:"created_at,omitempty"`
	UpdatedAt               *Timestamp `json:"updated_at,omitempty"`
	SuspendedAt             *Timestamp `json:"suspended_at,omitempty"`
	Type                    *string    `json:"type,omitempty"`
	SiteAdmin               *bool      `json:"site_admin,omitempty"`
	TotalPrivateRepos       *int64     `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos       *int64     `json:"owned_private_repos,omitempty"`
	PrivateGists            *int       `json:"private_gists,omitempty"`
	DiskUsage               *int       `json:"disk_usage,omitempty"`
	Collaborators           *int       `json:"collaborators,omitempty"`
	TwoFactorAuthentication *bool      `json:"two_factor_authentication,omitempty"`
	Plan                    *Plan      `json:"plan,omitempty"`
	LdapDn                  *string    `json:"ldap_dn,omitempty"`

	// API URLs
	URL               *string `json:"url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`

	// TextMatches is only populated from search results that request text matches
	// See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata
	TextMatches []*TextMatch `json:"text_matches,omitempty"`

	// Permissions and RoleName identify the permissions and role that a user has on a given
	// repository. These are only populated when calling Repositories.ListCollaborators.
	Permissions map[string]bool `json:"permissions,omitempty"`
	RoleName    *string         `json:"role_name,omitempty"`
}

func (u User) String() string {
	return Stringify(u)
}

// Get fetches a user. Passing the empty string will fetch the authenticated
// user.
//
// GitHub API docs: https://docs.github.com/en/rest/users/users#get-a-user
func (s *UsersService) Get(ctx context.Context, user string) (*User, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v", user)
	} else {
		u = "user"
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(User)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// GetByID fetches a user.
//
// Note: GetByID uses the undocumented GitHub API endpoint /user/:id.
func (s *UsersService) GetByID(ctx context.Context, id int64) (*User, *Response, error) {
	u := fmt.Sprintf("user/%d", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	resp, err := s.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// HovercardOptions specifies optional parameters to the UsersService.GetHovercard
// method.
type HovercardOptions struct {
	// SubjectType specifies the additional information to be received about the hovercard.
	// Possible values are: organization, repository, issue, pull_request. (Required when using subject_id.)
	SubjectType string `url:"subject_type"`

	// SubjectID specifies the ID for the SubjectType. (Required when using subject_type.)
	SubjectID string `url:"subject_id"`
}

// Hovercard represents hovercard information about a user.
type Hovercard struct {
	Contexts []*UserContext `json:"contexts,omitempty"`
}

// UserContext represents the contextual information about user.
type UserContext struct {
	Message *string `json:"message,omitempty"`
	Octicon *string `json:"octicon,omitempty"`
}

// GetHovercard fetches contextual information about user. It requires authentication
// via Basic Auth or via OAuth with the repo scope.
//
// GitHub API docs: https://docs.github.com/en/rest/users/users#get-contextual-information-for-a-user
func (s *UsersService) GetHovercard(ctx context.Context, user string, opts *HovercardOptions) (*Hovercard, *Response, error) {
	u := fmt.Sprintf("users/%v/hovercard", user)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	hc := new(Hovercard)
	resp, err := s.client.Do(ctx, req, hc)
	if err != nil {
		return nil, resp, err
	}

	return hc, resp, nil
}

// UserListOptions specifies optional parameters to the UsersService.ListAll
// method.
type UserListOptions struct {
	// ID of the last user seen
	Since int64 `url:"since,omitempty"`

	// Note: Pagination is powered exclusively by the Since parameter,
	// ListOptions.Page has no effect.
	// ListOptions.PerPage controls an undocumented GitHub API parameter.
	ListOptions
}

// ListAll lists all GitHub users.
//
// To paginate through all users, populate 'Since' with the ID of the last user.
//
// GitHub API docs: https://docs.github.com/en/rest/users/users#list-users
func (s *UsersService) ListAll(ctx context.Context, opts *UserListOptions) ([]*User, *Response, error) {
	u, err := addOptions("users", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*User
	resp, err := s.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}
