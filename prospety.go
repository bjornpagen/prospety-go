package prospety

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"

	"go.uber.org/ratelimit"
)

type Option func(option *options) error

type options struct {
	host       string
	rateLimit  *ratelimit.Limiter
	httpClient *http.Client
}

func WithHost(host string) Option {
	return func(option *options) error {
		// Check if host is valid.
		_, err := http.NewRequest("GET", fmt.Sprintf("https://%s", host), nil)
		if err != nil {
			return fmt.Errorf("invalid host: %w", err)
		}

		option.host = host
		return nil
	}
}

func WithRateLimit(rl ratelimit.Limiter) Option {
	return func(option *options) error {
		option.rateLimit = &rl
		return nil
	}
}

func WithHttpClient(hc http.Client) Option {
	return func(option *options) error {
		option.httpClient = &hc
		return nil
	}
}

type Client struct {
	apiKey  string
	options *options
}

func New(apiKey string, opts ...Option) (*Client, error) {
	o := &options{}
	for _, opt := range opts {
		err := opt(o)
		if err != nil {
			return nil, fmt.Errorf("bad option: %w", err)
		}
	}

	if o.host == "" {
		o.host = "app.prospety.com/api"
	}

	if o.rateLimit == nil {
		o.rateLimit = new(ratelimit.Limiter)
		*o.rateLimit = ratelimit.NewUnlimited()
	}

	if o.httpClient == nil {
		o.httpClient = http.DefaultClient
	}

	return &Client{
		apiKey:  apiKey,
		options: o,
	}, nil
}

type param struct {
	key   string
	value string
}

func (c *Client) buildUrl(p []string) string {
	return fmt.Sprintf("https://%s/%s", c.options.host, path.Join(p...))
}

func (c *Client) buildUrlWithParameters(path []string, params []param) string {
	url := c.buildUrl(path)
	for i, p := range params {
		// If it's the first parameter, use a question mark; otherwise, use an ampersand
		separator := "&"
		if i == 0 {
			separator = "?"
		}
		url = fmt.Sprintf("%s%s%s=%s", url, separator, p.key, p.value)
	}
	return url
}

func (c *Client) do(req *http.Request) (data []byte, err error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	(*c.options.rateLimit).Take()
	resp, err := c.options.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return data, nil
}

func (c *Client) get(path []string, params []param) (data []byte, err error) {
	url := c.buildUrlWithParameters(path, params)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return c.do(req)
}

func (c *Client) post(path []string, body any) (data []byte, err error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	url := c.buildUrl(path)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	return c.do(req)
}

func (c *Client) put(path []string, body any) (data []byte, err error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	url := c.buildUrl(path)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	return c.do(req)
}

func (c *Client) delete(path []string) (data []byte, err error) {
	url := c.buildUrl(path)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return c.do(req)
}

type getChannelsResponse struct {
	Total int       `json:"total"`
	Data  []Channel `json:"data"`
}

func (c *Client) GetChannels(limit, page int) ([]Channel, error) {
	data, err := c.get([]string{"channels"}, []param{
		{
			key:   "limit",
			value: strconv.Itoa(limit),
		},
		{
			key:   "page",
			value: strconv.Itoa(page),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get channels: %w", err)
	}

	res := &getChannelsResponse{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res.Data, nil
}

func (c *Client) GetChannel(id int) (*Channel, error) {
	data, err := c.get([]string{"channels", strconv.Itoa(id)}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel: %w", err)
	}

	res := &Channel{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res, nil
}

type getQuickSearchesResponse struct {
	Total int              `json:"total"`
	Data  []apiQuickSearch `json:"data"`
}

func (c *Client) GetQuickSearches(limit, page int) ([]QuickSearch, error) {
	data, err := c.get([]string{"quick_searches"}, []param{
		{
			key:   "limit",
			value: strconv.Itoa(limit),
		},
		{
			key:   "page",
			value: strconv.Itoa(page),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get quick searches: %w", err)
	}

	res := &getQuickSearchesResponse{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	quickSearches := make([]QuickSearch, len(res.Data))
	for i, quickSearch := range res.Data {
		s, err := quickSearch.toQuickSearch()
		if err != nil {
			return nil, fmt.Errorf("failed to convert quick search: %w", err)
		}

		quickSearches[i] = *s
	}

	return quickSearches, nil
}

type createQuickSearchPayload struct {
	ChannelID int    `json:"channel_id"`
	Url       string `json:"url"`
}

func (c *Client) CreateQuickSearch(channel ChannelType, url string) error {
	payload := createQuickSearchPayload{
		ChannelID: channel,
		Url:       url,
	}

	_, err := c.post([]string{"quick_searches"}, payload)
	if err != nil {
		return fmt.Errorf("failed to create quick search: %w", err)
	}

	return nil
}

func (c *Client) GetQuickSearch(id int) (*QuickSearch, error) {
	data, err := c.get([]string{"quick_searches", strconv.Itoa(id)}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get quick search: %w", err)
	}

	res := &apiQuickSearch{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res.toQuickSearch()
}

func (c *Client) DeleteQuickSearch(id int) error {
	_, err := c.delete([]string{"quick_searches", strconv.Itoa(id)})
	if err != nil {
		return fmt.Errorf("failed to delete quick search: %w", err)
	}

	return nil
}

type getPotentialProspectsCountResponse struct {
	Count int `json:"count"`
}

func (c *Client) GetPotentialProspectsCount(criteria any) (int, error) {
	switch v := criteria.(type) {
	case YouTubeStandardSearchCriteria:
		return getPotentialProspectsCountYouTubeStandard(c, v)
	case YouTubeSimilarSearchCriteria:
		return getPotentialProspectsCountYouTubeSimilar(c, v)
	default:
		return 0, fmt.Errorf("unknown search criteria data type: %T", v)
	}
}

type getPotentialProspectsCountPayload struct {
	Type      string      `json:"type"`
	ChannelId ChannelType `json:"channel_id"`
	Data      any         `json:"data"`
}

func getPotentialProspectsCountYouTubeStandard[T YouTubeStandardSearchCriteria](c *Client, criteria T) (int, error) {
	payload := getPotentialProspectsCountPayload{
		Type:      SearchTypeStandard,
		ChannelId: ChannelYouTube,
		Data:      criteria,
	}

	data, err := c.put([]string{"searches", "potential-prospects", "count"}, payload)
	if err != nil {
		return 0, fmt.Errorf("failed to get potential prospects count: %w", err)
	}

	res := &getPotentialProspectsCountResponse{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res.Count, nil
}

func getPotentialProspectsCountYouTubeSimilar[T YouTubeSimilarSearchCriteria](c *Client, criteria T) (int, error) {
	payload := getPotentialProspectsCountPayload{
		Type:      SearchTypeSimilar,
		ChannelId: ChannelYouTube,
		Data:      criteria,
	}

	data, err := c.put([]string{"searches", "potential-prospects", "count"}, payload)
	if err != nil {
		return 0, fmt.Errorf("failed to get potential prospects count: %w", err)
	}

	res := &getPotentialProspectsCountResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res.Count, nil
}

func (c *Client) GetPotentialProspects(criteria any) (any, error) {
	switch v := criteria.(type) {
	case YouTubeStandardSearchCriteria:
		return getPotentialProspectsYouTubeStandard(c, v)
	case YouTubeSimilarSearchCriteria:
		return getPotentialProspectsYouTubeSimilar(c, v)
	default:
		return nil, fmt.Errorf("unknown search criteria type: %T", v)
	}
}

type getPotentialProspectsPayload = getPotentialProspectsCountPayload

func getPotentialProspectsYouTubeStandard[T YouTubeStandardSearchCriteria](c *Client, criteria T) ([]YouTubeProspectPreview, error) {
	payload := getPotentialProspectsPayload{
		Type:      SearchTypeStandard,
		ChannelId: ChannelYouTube,
		Data:      criteria,
	}

	data, err := c.put([]string{"searches", "potential-prospects", "preview"}, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to get potential prospects: %w", err)
	}

	var res []YouTubeProspectPreview
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res, nil
}

func getPotentialProspectsYouTubeSimilar[T YouTubeSimilarSearchCriteria](c *Client, criteria T) ([]YouTubeProspectPreview, error) {
	payload := getPotentialProspectsPayload{
		Type:      SearchTypeSimilar,
		ChannelId: ChannelYouTube,
		Data:      criteria,
	}

	data, err := c.put([]string{"searches", "potential-prospects", "preview"}, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to get potential prospects: %w", err)
	}

	var res []YouTubeProspectPreview
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res, nil
}

type getSearchesResponse struct {
	Total int      `json:"total"`
	Data  []Search `json:"data"`
}

func (c *Client) GetSearches(limit, page int) ([]Search, error) {
	data, err := c.get([]string{"searches"}, []param{
		{
			key:   "limit",
			value: strconv.Itoa(limit),
		},
		{
			key:   "page",
			value: strconv.Itoa(page),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get searches: %w", err)
	}

	res := &getSearchesResponse{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res.Data, nil
}

type createSearchPayload struct {
	Title      string `json:"title"`
	Type       string `json:"type"`
	ChannelId  int    `json:"channel_id"`
	Limit      int    `json:"limit"`
	Data       any/*SearchData*/ `json:"data"`
	ImportFile string `json:"import_file,omitempty"` // not using this field. unsupported
	Method     string `json:"_method"`
}

func (c *Client) CreateSearch(title string, limit int, searchData any) error {
	switch v := searchData.(type) {
	case YouTubeStandardSearch:
		return createSearchYouTubeStandard(c, title, limit, v)
	default:
		return fmt.Errorf("unknown search data type: %T", v)
	}
}

func createSearchYouTubeStandard[T YouTubeStandardSearch](c *Client, title string, limit int, data T) error {
	payload := createSearchPayload{
		Title:     title,
		Type:      SearchTypeStandard,
		ChannelId: ChannelYouTube,
		Limit:     limit,
		Data:      data,
		Method:    "PUT",
	}

	_, err := c.put([]string{"searches"}, payload)
	if err != nil {
		return fmt.Errorf("failed to create search: %w", err)
	}

	return nil
}

func (c *Client) GetSearch(id int) (*Search, error) {
	data, err := c.get([]string{"searches", strconv.Itoa(id)}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get search: %w", err)
	}

	res := &Search{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Unmarshal the data field into one of our types. Currently res.Data is a map[string]interface{}.
	// Need to switch both on res.Data.Type and res.ChannelID.
	channelId := res.ChannelID
	var dataType string

	// first check if it's nil
	if res.Data == nil {
		return nil, fmt.Errorf("data not set")
	}
	// then check if it's a map, then check if it's a map[string]string
	var mapData = make(map[string]interface{})
	if mapData, ok := res.Data.(map[string]interface{}); ok {
		if mapData["type"] == nil {
			return nil, fmt.Errorf("data type not set")
		}
		dataType = mapData["type"].(string)
	} else {
		return nil, fmt.Errorf("data not set")
	}

	switch channelId {
	case ChannelYouTube:
		switch dataType {
		case SearchTypeStandard:
			newData, err := mapToYouTubeStandardSearchData(mapData)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal data: %w", err)
			}
			res.Data = newData
		case SearchTypeImport:
			return nil, fmt.Errorf("import search type not implemented")
		case SearchTypeSimilar:
			return nil, fmt.Errorf("similar search type not implemented")
		case SearchTypeFollower:
			return nil, fmt.Errorf("follower search type not implemented")
		case SearchTypeHashtag:
			return nil, fmt.Errorf("hashtag search type not implemented")
		default:
			return nil, fmt.Errorf("unknown search type: %s", dataType)
		}
	case ChannelInstagram:
		return nil, fmt.Errorf("instagram not implemented")
	default:
		return nil, fmt.Errorf("unknown channel id: %d", channelId)
	}

	return res, nil
}

type updateSearchPayload = createSearchPayload

func (c *Client) UpdateSearch(id int, title string, limit int, data any) error {
	switch v := data.(type) {
	case YouTubeStandardSearch:
		return updateSearchYouTubeStandard(c, id, title, SearchTypeStandard, ChannelYouTube, limit, v)
	default:
		return fmt.Errorf("unknown search data type: %T", v)
	}
}

func updateSearchYouTubeStandard[T YouTubeStandardSearch](c *Client, id int, title, dataType string, channelId, limit int, data T) error {
	payload := updateSearchPayload{
		Title:     title,
		Type:      dataType,
		ChannelId: channelId,
		Limit:     limit,
		Data:      data,
		Method:    "PUT",
	}

	_, err := c.put([]string{"searches", strconv.Itoa(id)}, payload)
	if err != nil {
		return fmt.Errorf("failed to update search: %w", err)
	}

	return nil
}

func (c *Client) DeleteSearch(id int) error {
	_, err := c.delete([]string{"searches", strconv.Itoa(id)})
	if err != nil {
		return fmt.Errorf("failed to delete search: %w", err)
	}

	return nil
}

func (c *Client) StartSearch(id int) error {
	_, err := c.put([]string{"searches", strconv.Itoa(id), "start"}, nil)
	if err != nil {
		return fmt.Errorf("failed to start search: %w", err)
	}

	return nil
}

func (c *Client) PauseSearch(id int) error {
	_, err := c.put([]string{"searches", strconv.Itoa(id), "pause"}, nil)
	if err != nil {
		return fmt.Errorf("failed to pause search: %w", err)
	}

	return nil
}

func (c *Client) FinishSearch(id int) error {
	_, err := c.put([]string{"searches", strconv.Itoa(id), "finish"}, nil)
	if err != nil {
		return fmt.Errorf("failed to finish search: %w", err)
	}

	return nil
}

type getProspectsResponse struct {
	Total int `json:"total"`
	Data  any `json:"data"`
}

func (c *Client) GetProspects(id int) (any, error) {
	data, err := c.get([]string{"searches", strconv.Itoa(id), "prospects"}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get prospects: %w", err)
	}

	res := &getProspectsResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Unmarshal the data field into one of our types. Currently res.Data is a []map[string]interface{}.
	// It's either a []YouTubeProspect or []InstagramProspect.
	// We can do this by checking for the Data["video_keywords"]. If it's set, it's a YouTubeProspect.
	// If it's not set, it's an InstagramProspect.
	// ONLY CHECK for video_keywords, because it's the only field that's unique to YouTubeProspect.
	// Literally don't check any other fields, because they're not unique to InstagramProspect.

	// first check if it's nil
	if res.Data == nil {
		return nil, fmt.Errorf("data not set")
	}
	// then check if it's a slice, then check if it's a []map[string]string
	if sliceData, ok := res.Data.([]map[string]interface{}); ok {
		if len(sliceData) == 0 {
			return nil, fmt.Errorf("data not set")
		}
		if sliceData[0]["video_keywords"] != nil {
			newData, err := mapToYouTubeProspects(sliceData, res.Total)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal data: %w", err)
			}
			res.Data = newData
		} else {
			newData, err := mapToInstagramProspects(sliceData, res.Total)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal data: %w", err)
			}
			res.Data = newData
		}
	} else {
		return nil, fmt.Errorf("data not set")
	}

	return res, nil
}

func (c *Client) ExportProspects(id int, fileType string) (string, error) {
	data, err := c.get([]string{"searches", strconv.Itoa(id), "prospects", "export"},
		[]param{
			{
				key:   "type",
				value: fileType,
			},
		})
	if err != nil {
		return "", fmt.Errorf("failed to export prospects: %w", err)
	}

	// just give us the csv string
	return string(data), nil
}

// theres a shit ton more methods but we're not using them! goodbye
