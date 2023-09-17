// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
)

// SelfService handles communication with the
// methods of the GitHub API which handle the currently logged in user
//
// GitHub API docs: https://docs.github.com/en/rest/user/
type SelfService struct {
	Self
	service
}

// Self is a [User] with some extra properties.
// It refers to the user making the API call.
type Self struct {
	User
	Collaborators           int
	DiskUsage               int
	Plan                    Plan
	OwnedPrivateRepos       int
	TotalPrivateRepos       int
	PrivateGists            int
	TwoFactorAuthentication bool
}

func (u Self) String() string {
	return Stringify(u)
}

// Get fetches the currently logged-in user.
//
// GitHub API docs: https://docs.github.com/en/rest/users/users#get-the-authenticated-user
func (s *SelfService) Get(ctx context.Context) (*Self, *Response, error) {
	req, err := s.client.NewRequest("GET", "/user", nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(Self)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	s.Self = *uResp

	return uResp, resp, nil
}

// Edit the authenticated user.
//
// GitHub API docs: https://docs.github.com/en/rest/users/users#update-the-authenticated-user
func (s *SelfService) Edit(ctx context.Context) (*Self, *Response, error) {
	req, err := s.client.NewRequest("PATCH", "/user", nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(Self)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	s.Self = *uResp

	return uResp, resp, nil
}

// ListInvitations lists invitations for the currently authenticated user
//
// GitHub API docs: https://docs.github.com/en/rest/collaborators/invitations#list-repository-invitations-for-the-authenticated-user
func (s *SelfService) ListInvitations(ctx context.Context, opts *ListOptions) ([]*RepositoryInvitation, *Response, error) {
	u, err := addOptions("user/repository_invitations", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	invites := []*RepositoryInvitation{}
	resp, err := s.client.Do(ctx, req, &invites)
	if err != nil {
		return nil, resp, err
	}

	return invites, resp, nil
}

// AcceptInvitation accepts the currently-open repository invitation for the
// authenticated user.
//
// GitHub API docs: https://docs.github.com/en/rest/collaborators/invitations#accept-a-repository-invitation
func (s *SelfService) AcceptInvitation(ctx context.Context, invitationID int64) (*Response, error) {
	u := fmt.Sprintf("user/repository_invitations/%v", invitationID)
	req, err := s.client.NewRequest("PATCH", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeclineInvitation declines the currently-open repository invitation for the
// authenticated user.
//
// GitHub API docs: https://docs.github.com/en/rest/collaborators/invitations#decline-a-repository-invitation
func (s *SelfService) DeclineInvitation(ctx context.Context, invitationID int64) (*Response, error) {
	u := fmt.Sprintf("user/repository_invitations/%v", invitationID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
