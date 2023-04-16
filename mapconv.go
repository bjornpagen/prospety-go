package prospety

import "fmt"

// boilerplate map functions

func mapToYouTubeStandardSearchData(data map[string]any) (*YouTubeStandardSearch, error) {
	res := &YouTubeStandardSearch{}

	// Keywords
	if v, ok := data["keywords"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(string); ok {
					res.Keywords = append(res.Keywords, v4)
				} else {
					return nil, fmt.Errorf("failed to cast keywords to string")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast keywords to []interface{}")
		}
	} else {
		return nil, fmt.Errorf("keywords not set")
	}

	// KeywordsMode
	if v, ok := data["keywords_mode"]; ok {
		if v2, ok := v.(string); ok {
			res.KeywordsMode = v2
		} else {
			return nil, fmt.Errorf("failed to cast keywords_mode to string")
		}
	} else {
		return nil, fmt.Errorf("keywords_mode not set")
	}

	// ExcludedKeywords
	if v, ok := data["excluded_keywords"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(string); ok {
					res.ExcludedKeywords = append(res.ExcludedKeywords, v4)
				} else {
					return nil, fmt.Errorf("failed to cast excluded_keywords to string")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast excluded_keywords to []interface{}")
		}
	} else {
		return nil, fmt.Errorf("excluded_keywords not set")
	}

	// ExcludedKeywordsMode
	if v, ok := data["excluded_keywords_mode"]; ok {
		if v2, ok := v.(string); ok {
			res.ExcludedKeywordsMode = v2
		} else {
			return nil, fmt.Errorf("failed to cast excluded_keywords_mode to string")
		}
	} else {
		return nil, fmt.Errorf("excluded_keywords_mode not set")
	}

	// VideoKeywords
	if v, ok := data["video_keywords"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(string); ok {
					res.VideoKeywords = append(res.VideoKeywords, v4)
				} else {
					return nil, fmt.Errorf("failed to cast video_keywords to string")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast video_keywords to []interface{}")
		}
	} else {
		return nil, fmt.Errorf("video_keywords not set")
	}

	// VideoKeywordsMode
	if v, ok := data["video_keywords_mode"]; ok {
		if v2, ok := v.(string); ok {
			res.VideoKeywordsMode = v2
		} else {
			return nil, fmt.Errorf("failed to cast video_keywords_mode to string")
		}
	} else {
		return nil, fmt.Errorf("video_keywords_mode not set")
	}

	// ExcludedVideoKeywords
	if v, ok := data["excluded_video_keywords"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(string); ok {
					res.ExcludedVideoKeywords = append(res.ExcludedVideoKeywords, v4)
				} else {
					return nil, fmt.Errorf("failed to cast excluded_video_keywords to string")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast excluded_video_keywords to []interface{}")
		}
	} else {
		return nil, fmt.Errorf("excluded_video_keywords not set")
	}

	// ExcludedVideoKeywordsMode
	if v, ok := data["excluded_video_keywords_mode"]; ok {
		if v2, ok := v.(string); ok {
			res.ExcludedVideoKeywordsMode = v2
		} else {
			return nil, fmt.Errorf("failed to cast excluded_video_keywords_mode to string")
		}
	} else {
		return nil, fmt.Errorf("excluded_video_keywords_mode not set")
	}

	// Category
	if v, ok := data["category"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(string); ok {
					res.Category = append(res.Category, v4)
				} else {
					return nil, fmt.Errorf("failed to cast category to string")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast category to []interface{}")
		}
	} else {
		return nil, fmt.Errorf("category not set")
	}

	// Country
	if v, ok := data["country"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(string); ok {
					res.Country = append(res.Country, v4)
				} else {
					return nil, fmt.Errorf("failed to cast country to string")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast country to []interface{}")
		}
	} else {
		return nil, fmt.Errorf("country not set")
	}

	// SubscribersRange
	if v, ok := data["subscribers_range"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(float64); ok {
					res.SubscribersRange = append(res.SubscribersRange, int64(v4))
				} else {
					return nil, fmt.Errorf("failed to cast subscribers_range to int64")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast subscribers_range to []interface{}")
		}
		if len(res.SubscribersRange) != 2 {
			return nil, fmt.Errorf("subscribers_range must have 2 elements")
		}
	} else {
		return nil, fmt.Errorf("subscribers_range not set")
	}

	// TotalViewsRange
	if v, ok := data["total_views_range"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(float64); ok {
					res.TotalViewsRange = append(res.TotalViewsRange, int64(v4))
				} else {
					return nil, fmt.Errorf("failed to cast total_views_range to int64")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast total_views_range to []interface{}")
		}
		if len(res.TotalViewsRange) != 2 {
			return nil, fmt.Errorf("total_views_range must have 2 elements")
		}
	} else {
		return nil, fmt.Errorf("total_views_range not set")
	}

	// AverageViewsPerVideoRange
	if v, ok := data["average_views_per_video_range"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(float64); ok {
					res.AverageViewsPerVideoRange = append(res.AverageViewsPerVideoRange, int64(v4))
				} else {
					return nil, fmt.Errorf("failed to cast average_views_per_video_range to int64")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast average_views_per_video_range to []interface{}")
		}
		if len(res.AverageViewsPerVideoRange) != 2 {
			return nil, fmt.Errorf("average_views_per_video_range must have 2 elements")
		}
	} else {
		return nil, fmt.Errorf("average_views_per_video_range not set")
	}

	// TotalVideosRange
	if v, ok := data["total_videos_range"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(float64); ok {
					res.TotalVideosRange = append(res.TotalVideosRange, int(v4))
				} else {
					return nil, fmt.Errorf("failed to cast total_videos_range to int")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast total_videos_range to []interface{}")
		}
		if len(res.TotalVideosRange) != 2 {
			return nil, fmt.Errorf("total_videos_range must have 2 elements")
		}
	} else {
		return nil, fmt.Errorf("total_videos_range not set")
	}

	// LatestVideoRange
	if v, ok := data["latest_video_range"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(float64); ok {
					res.LatestVideoRange = append(res.LatestVideoRange, int(v4))
				} else {
					return nil, fmt.Errorf("failed to cast latest_video_range to int")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast latest_video_range to []interface{}")
		}
		if len(res.LatestVideoRange) != 2 {
			return nil, fmt.Errorf("latest_video_range must have 2 elements")
		}
	} else {
		return nil, fmt.Errorf("latest_video_range not set")
	}

	// CreatedRange
	if v, ok := data["created_range"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(float64); ok {
					res.CreatedRange = append(res.CreatedRange, int(v4))
				} else {
					return nil, fmt.Errorf("failed to cast created_range to int")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast created_range to []interface{}")
		}
		if len(res.CreatedRange) != 2 {
			return nil, fmt.Errorf("created_range must have 2 elements")
		}
	} else {
		return nil, fmt.Errorf("created_range not set")
	}

	// PricingMethod
	if v, ok := data["pricing_method"]; ok {
		if v2, ok := v.(string); ok {
			res.PricingMethod = v2
		} else {
			return nil, fmt.Errorf("failed to cast pricing_method to string")
		}
	} else {
		return nil, fmt.Errorf("pricing_method not set")
	}

	// EmailVerificationMethod
	if v, ok := data["email_verification_method"]; ok {
		if v2, ok := v.(string); ok {
			res.EmailVerificationMethod = v2
		} else {
			return nil, fmt.Errorf("failed to cast email_verification_method to string")
		}
	} else {
		return nil, fmt.Errorf("email_verification_method not set")
	}

	return res, nil
}

func mapToYouTubeProspects(data []map[string]any, length int) ([]YouTubeProspect, error) {
	res := make([]YouTubeProspect, length)

	for i, v := range data {
		prospect, err := mapToYouTubeProspect(v)
		if err != nil {
			return nil, err
		}
		res[i] = *prospect
	}

	return res, nil
}

func mapToInstagramProspects(data []map[string]any, length int) ([]InstagramProspect, error) {
	res := make([]InstagramProspect, length)

	for i, v := range data {
		prospect, err := mapToInstagramProspect(v)
		if err != nil {
			return nil, err
		}
		res[i] = *prospect
	}

	return res, nil
}

func mapToYouTubeProspect(data map[string]any) (*YouTubeProspect, error) {
	res := &YouTubeProspect{}

	// Photo
	if v, ok := data["photo"]; ok {
		if v2, ok := v.(string); ok {
			res.Photo = v2
		} else {
			return nil, fmt.Errorf("failed to cast photo to string")
		}
	} else {
		return nil, fmt.Errorf("photo not set")
	}

	// Name
	if v, ok := data["name"]; ok {
		if v2, ok := v.(string); ok {
			res.Name = v2
		} else {
			return nil, fmt.Errorf("failed to cast name to string")
		}
	} else {
		return nil, fmt.Errorf("name not set")
	}

	// URL
	if v, ok := data["url"]; ok {
		if v2, ok := v.(string); ok {
			res.URL = v2
		} else {
			return nil, fmt.Errorf("failed to cast url to string")
		}
	} else {
		return nil, fmt.Errorf("url not set")
	}

	// Email
	if v, ok := data["email"]; ok {
		if v2, ok := v.(string); ok {
			res.Email = v2
		} else {
			return nil, fmt.Errorf("failed to cast email to string")
		}
	} else {
		return nil, fmt.Errorf("email not set")
	}

	// Phone
	if v, ok := data["phone"]; ok {
		if v2, ok := v.(string); ok {
			res.Phone = v2
		} else {
			return nil, fmt.Errorf("failed to cast phone to string")
		}
	} else {
		return nil, fmt.Errorf("phone not set")
	}

	// Keywords
	if v, ok := data["keywords"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(string); ok {
					res.Keywords = append(res.Keywords, v4)
				} else {
					return nil, fmt.Errorf("failed to cast keywords to string")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast keywords to []interface{}")
		}
	} else {
		return nil, fmt.Errorf("keywords not set")
	}

	// VideoKeywords
	if v, ok := data["video_keywords"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(string); ok {
					res.VideoKeywords = append(res.VideoKeywords, v4)
				} else {
					return nil, fmt.Errorf("failed to cast video_keywords to string")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast video_keywords to []interface{}")
		}
	} else {
		return nil, fmt.Errorf("video_keywords not set")
	}

	// Category
	if v, ok := data["category"]; ok {
		if v2, ok := v.(string); ok {
			res.Category = v2
		} else {
			return nil, fmt.Errorf("failed to cast category to string")
		}
	} else {
		return nil, fmt.Errorf("category not set")
	}

	// Country
	if v, ok := data["country"]; ok {
		if v2, ok := v.(string); ok {
			res.Country = v2
		} else {
			return nil, fmt.Errorf("failed to cast country to string")
		}
	} else {
		return nil, fmt.Errorf("country not set")
	}

	// Links
	if v, ok := data["links"]; ok {
		if v2, ok := v.([]interface{}); ok {
			for _, v3 := range v2 {
				if v4, ok := v3.(string); ok {
					res.Links = append(res.Links, v4)
				} else {
					return nil, fmt.Errorf("failed to cast links to string")
				}
			}
		} else {
			return nil, fmt.Errorf("failed to cast links to []interface{}")
		}
	} else {
		return nil, fmt.Errorf("links not set")
	}

	// CreatedAt
	if v, ok := data["created_at"]; ok {
		if v2, ok := v.(string); ok {
			res.CreatedAt = v2
		} else {
			return nil, fmt.Errorf("failed to cast created_at to string")
		}
	} else {
		return nil, fmt.Errorf("created_at not set")
	}

	// Subscribers
	if v, ok := data["subscribers"]; ok {
		if v2, ok := v.(float64); ok {
			res.Subscribers = int64(v2)
		} else {
			return nil, fmt.Errorf("failed to cast subscribers to float64")
		}
	} else {
		return nil, fmt.Errorf("subscribers not set")
	}

	// TotalViews
	if v, ok := data["total_views"]; ok {
		if v2, ok := v.(float64); ok {
			res.TotalViews = int64(v2)
		} else {
			return nil, fmt.Errorf("failed to cast total_views to float64")
		}
	} else {
		return nil, fmt.Errorf("total_views not set")
	}

	// TotalVideos
	if v, ok := data["total_videos"]; ok {
		if v2, ok := v.(float64); ok {
			res.TotalVideos = int(v2)
		} else {
			return nil, fmt.Errorf("failed to cast total_videos to float64")
		}
	} else {
		return nil, fmt.Errorf("total_videos not set")
	}

	// LastVideo
	if v, ok := data["last_video"]; ok {
		if v2, ok := v.(string); ok {
			res.LastVideo = v2
		} else {
			return nil, fmt.Errorf("failed to cast last_video to string")
		}
	} else {
		return nil, fmt.Errorf("last_video not set")
	}

	return res, nil
}

func mapToInstagramProspect(data map[string]any) (*InstagramProspect, error) {
	return nil, fmt.Errorf("not implemented")
}
