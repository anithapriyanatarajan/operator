//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"net/http"
	"net/url"
)

type (
	// GitIgnoreTemplatesServiceInterface defines all the API methods for the GitIgnoreTemplatesService
	GitIgnoreTemplatesServiceInterface interface {
		ListTemplates(*ListTemplatesOptions, ...RequestOptionFunc) ([]*GitIgnoreTemplateListItem, *Response, error)
		GetTemplate(string, ...RequestOptionFunc) (*GitIgnoreTemplate, *Response, error)
	}

	// GitIgnoreTemplatesService handles communication with the gitignore
	// templates related methods of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/templates/gitignores/
	GitIgnoreTemplatesService struct {
		client *Client
	}
)

var _ GitIgnoreTemplatesServiceInterface = (*GitIgnoreTemplatesService)(nil)

// GitIgnoreTemplate represents a GitLab gitignore template.
//
// GitLab API docs: https://docs.gitlab.com/api/templates/gitignores/
type GitIgnoreTemplate struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// GitIgnoreTemplateListItem represents a GitLab gitignore template from the list.
//
// GitLab API docs: https://docs.gitlab.com/api/templates/gitignores/
type GitIgnoreTemplateListItem struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

// ListTemplatesOptions represents the available ListAllTemplates() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/templates/gitignores/#get-all-gitignore-templates
type ListTemplatesOptions ListOptions

// ListTemplates get a list of available git ignore templates
//
// GitLab API docs:
// https://docs.gitlab.com/api/templates/gitignores/#get-all-gitignore-templates
func (s *GitIgnoreTemplatesService) ListTemplates(opt *ListTemplatesOptions, options ...RequestOptionFunc) ([]*GitIgnoreTemplateListItem, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "templates/gitignores", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var gs []*GitIgnoreTemplateListItem
	resp, err := s.client.Do(req, &gs)
	if err != nil {
		return nil, resp, err
	}

	return gs, resp, nil
}

// GetTemplate get a git ignore template
//
// GitLab API docs:
// https://docs.gitlab.com/api/templates/gitignores/#get-a-single-gitignore-template
func (s *GitIgnoreTemplatesService) GetTemplate(key string, options ...RequestOptionFunc) (*GitIgnoreTemplate, *Response, error) {
	u := fmt.Sprintf("templates/gitignores/%s", url.PathEscape(key))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	g := new(GitIgnoreTemplate)
	resp, err := s.client.Do(req, g)
	if err != nil {
		return nil, resp, err
	}

	return g, resp, nil
}
