package main

import "time"

type QiitaServices struct {
	SidebarProps struct {
		Ads struct {
			ID         string `json:"id"`
			AdUnitPath string `json:"adUnitPath"`
			Size       []int  `json:"size"`
		} `json:"ads"`
		Announcement struct {
			EncryptedID    string `json:"encryptedId"`
			Title          string `json:"title"`
			Body           string `json:"body"`
			DetailURL      string `json:"detailUrl"`
			IsStaffOnly    bool   `json:"isStaffOnly"`
			IsReadByViewer bool   `json:"isReadByViewer"`
		} `json:"announcement"`
		OfficialEvent interface{} `json:"officialEvent"`
		UserRanking   struct {
			Scope string `json:"scope"`
		} `json:"userRanking"`
		TagRanking struct {
			TagRanking struct {
				Edges []struct {
					Score int `json:"score"`
					Node  struct {
						MediumIconURL string `json:"mediumIconUrl"`
						Name          string `json:"name"`
						URLName       string `json:"urlName"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"tagRanking"`
			Scope string `json:"scope"`
		} `json:"tagRanking"`
		Sponsers []struct {
			Alt  string `json:"alt"`
			Src  string `json:"src"`
			Link string `json:"link"`
		} `json:"sponsers"`
		QiitaServices []struct {
			Alt  string `json:"alt"`
			Src  string `json:"src"`
			Link string `json:"link"`
		} `json:"qiitaServices"`
	} `json:"sidebarProps"`
	HeaderBanner struct {
		BannerDisplayKey string `json:"bannerDisplayKey"`
		BannerPath       string `json:"bannerPath"`
		DesktopBannerAlt string `json:"desktopBannerAlt"`
		DesktopBannerURL string `json:"desktopBannerUrl"`
		DisplayBanner    bool   `json:"displayBanner"`
		MobileBannerAlt  string `json:"mobileBannerAlt"`
		MobileBannerURL  string `json:"mobileBannerUrl"`
	} `json:"headerBanner"`
	ImageUrls struct {
		TwitterQiitaLogo          string `json:"twitterQiitaLogo"`
		TwitterQiitaMilestoneLogo string `json:"twitterQiitaMilestoneLogo"`
		TwitterQiitapoiLogo       string `json:"twitterQiitapoiLogo"`
		FacebookQiitaLogo         string `json:"facebookQiitaLogo"`
		Feed                      string `json:"feed"`
	} `json:"imageUrls"`
	Trend struct {
		Scope string `json:"scope"`
		Trend struct {
			Edges []struct {
				IsLikedByViewer bool          `json:"isLikedByViewer"`
				IsNewArrival    bool          `json:"isNewArrival"`
				FollowingLikers []interface{} `json:"followingLikers"`
				Node            struct {
					EncryptedID         string    `json:"encryptedId"`
					IsLikedByViewer     bool      `json:"isLikedByViewer"`
					IsStockableByViewer bool      `json:"isStockableByViewer"`
					IsStockedByViewer   bool      `json:"isStockedByViewer"`
					LikesCount          int       `json:"likesCount"`
					LinkURL             string    `json:"linkUrl"`
					PublishedAt         time.Time `json:"publishedAt"`
					Title               string    `json:"title"`
					UUID                string    `json:"uuid"`
					Author              struct {
						ProfileImageURL string `json:"profileImageUrl"`
						URLName         string `json:"urlName"`
					} `json:"author"`
					Organization            interface{}   `json:"organization"`
					RecentlyFollowingLikers []interface{} `json:"recentlyFollowingLikers"`
					Tags                    []struct {
						URLName string `json:"urlName"`
						Name    string `json:"name"`
					} `json:"tags"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"trend"`
		Type string `json:"type"`
	} `json:"trend"`
}
