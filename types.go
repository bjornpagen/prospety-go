package prospety

import (
	"time"
)

type ChannelType = int

const (
	ChannelYouTube   = ChannelType(1)
	ChannelInstagram = ChannelType(2)
)

type Channel struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type SearchStatus = string

const (
	SearchStatusPending  = SearchStatus("pending")
	SearchStatusFinished = SearchStatus("finished")
)

type YouTubeCategory = int

/*
YouTube: Autos & Vehicles: 2, Comedy: 23, Education: 27, Entertainment: 24, Film & Animation: 1, Gaming: 20, Howto & Style: 26, Music: 10, News & Politics: 25, Nonprofits & Activism: 29, People & Blogs: 22, Pets & Animals: 15, Science & Technology: 28, Sports: 17, Travel & Events: 19, Shows: 43, Trailers: 44.
*/

const (
	YouTubeCategoryAutosVehicles = YouTubeCategory(2)
	YouTubeCategoryComedy        = YouTubeCategory(23)
	YouTubeCategoryEducation     = YouTubeCategory(27)
	YouTubeCategoryEntertainment = YouTubeCategory(24)
	YouTubeCategoryFilmAnimation = YouTubeCategory(1)
	YouTubeCategoryGaming        = YouTubeCategory(20)
	YouTubeCategoryHowtoStyle    = YouTubeCategory(26)
	YouTubeCategoryMusic         = YouTubeCategory(10)
	YouTubeCategoryNewsPolitics  = YouTubeCategory(25)
	YouTubeCategoryNonprofits    = YouTubeCategory(29)
	YouTubeCategoryPeopleBlogs   = YouTubeCategory(22)
	YouTubeCategoryPetsAnimals   = YouTubeCategory(15)
	YouTubeCategoryScienceTech   = YouTubeCategory(28)
	YouTubeCategorySports        = YouTubeCategory(17)
	YouTubeCategoryTravelEvents  = YouTubeCategory(19)
	YouTubeCategoryShows         = YouTubeCategory(43)
	YouTubeCategoryTrailers      = YouTubeCategory(44)
)

type QuickSearch struct {
	ID        int              `json:"id"`
	Status    SearchStatus     `json:"status"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	Prospect  *ProspectPreview `json:"prospect"`
}

type ProspectPreview struct {
	Photo string `json:"photo"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

const (
	SearchTypeStandard = "standard"
	SearchTypeImport   = "import"
	SearchTypeSimilar  = "similar"
	SearchTypeFollower = "follower"
	SearchTypeHashtag  = "hashtag"
)

type StandardSearchCriteria struct {
	Keywords                  []string `json:"keywords"`
	KeywordsMode              string   `json:"keywords_mode"`
	ExcludedKeywords          []string `json:"excluded_keywords"`
	ExcludedKeywordsMode      string   `json:"excluded_keywords_mode"`
	VideoKeywords             []string `json:"video_keywords"`
	VideoKeywordsMode         string   `json:"video_keywords_mode"`
	ExcludedVideoKeywords     []string `json:"excluded_video_keywords"`
	ExcludedVideoKeywordsMode string   `json:"excluded_video_keywords_mode"`
	Category                  []string `json:"category"`
	Country                   []string `json:"country"`
	SubscribersRange          []int64  `json:"subscribers_range"`
	TotalViewsRange           []int64  `json:"total_views_range"`
	AverageViewsPerVideoRange []int64  `json:"average_views_per_video_range"`
	TotalVideosRange          []int    `json:"total_videos_range"`
	LatestVideoRange          []int    `json:"latest_video_range"`
	CreatedRange              []int    `json:"created_range"`
}

type SimilarSearchCriteria struct {
	RequiredKeywords                    bool     `json:"required_keywords"`
	RequiredVideoKeywords               bool     `json:"required_video_keywords"`
	RequiredCategory                    bool     `json:"required_category"`
	RequiredCountry                     bool     `json:"required_country"`
	RequiredSubscribersRange            bool     `json:"required_subscribers_range"`
	SubscribersDifferenceRange          []int64  `json:"subscribers_difference_range"`
	RequiredTotalViewsRange             bool     `json:"required_total_views_range"`
	TotalViewsDifferenceRange           []int64  `json:"total_views_difference_range"`
	RequiredAverageViewsPerVideoRange   bool     `json:"required_average_views_per_video_range"`
	AverageViewsPerVideoDifferenceRange []int64  `json:"average_views_per_video_difference_range"`
	RequiredTotalVideosRange            bool     `json:"required_total_videos_range"`
	TotalVideosDifferenceRange          []int    `json:"total_videos_difference_range"`
	RequiredLatestVideoRange            bool     `json:"required_latest_video_range"`
	LatestVideoDifferenceRange          []int    `json:"latest_video_difference_range"`
	RequiredCreatedRange                bool     `json:"required_created_range"`
	CreatedDifferenceRange              []int    `json:"created_difference_range"`
	MinimumScore                        []int    `json:"minimum_score"`
	References                          []string `json:"references"`
}

type StandardSearch struct {
	StandardSearchCriteria

	PricingMethod           string `json:"pricing_method"`
	EmailVerificationMethod string `json:"email_verification_method"`
}

type Search struct {
	ID                 int            `json:"id"`
	Title              string         `json:"title"`
	Status             string         `json:"status"`
	Type               string         `json:"type"`
	IsTypeSet          bool           `json:"is_type_set"`
	ChannelID          int            `json:"channel_id"`
	ChannelTitle       string         `json:"channel_title"`
	Limit              int            `json:"limit"`
	CreatedAtFormatted string         `json:"created_at_formatted"`
	UpdatedAtFormatted string         `json:"updated_at_formatted"`
	CreatedAt          string         `json:"created_at"`
	UpdatedAt          string         `json:"updated_at"`
	Progress           SearchProgress `json:"progress"`
	Searched           bool           `json:"searched"`
	GatheringProspects bool           `json:"gathering_prospects"`
	Data               StandardSearch `json:"data"` // TODO: make it generic
}

type SearchProgress struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

type Prospect struct {
	ProspectPreview

	Email         string   `json:"email"`
	Phone         string   `json:"phone"`
	Keywords      []string `json:"keywords"`
	VideoKeywords []string `json:"video_keywords"`
	Category      string   `json:"category"`
	Country       string   `json:"country"`
	Links         []string `json:"links"`
	CreatedAt     string   `json:"created_at"`
	Subscribers   int64    `json:"subscribers"`
	TotalViews    int64    `json:"total_views"`
	TotalVideos   int      `json:"total_videos"`
	LastVideo     string   `json:"last_video"`
}
