package domain

type ShopSearchRequest struct {
	Source interface{} `json:"_source,omitempty"`
	Query  Query       `json:"query,omitempty"`
	From   int         `json:"from,omitempty"`
	Size   int         `json:"size,omitempty"`
}

type ShopSearch struct {
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Id               int64     `json:"id"`
				Name             string    `json:"name"`
				Property         string    `json:"property,omitempty"`
				Alphabet         string    `json:"alphabet,omitempty"`
				NameKana         string    `json:"name_kana"`
				PrefId           string    `json:"pref_id,omitempty"`
				AreaId           string    `json:"area_id,omitempty"`
				StationId1       string    `json:"station_id1,omitempty"`
				StationTime1     string    `json:"station_time1,omitempty"`
				StationDistance1 string    `json:"station_distance1,omitempty"`
				StationId2       string    `json:"station_id2,omitempty"`
				StationTime2     string    `json:"station_time2,omitempty"`
				StationDistance2 string    `json:"station_distance2,omitempty"`
				StationId3       string    `json:"station_id3,omitempty"`
				StationTime3     string    `json:"station_time3,omitempty"`
				StationDistance3 string    `json:"station_distance3,omitempty"`
				CategoryId1      string    `json:"category_id1,omitempty"`
				CategoryId2      string    `json:"category_id2,omitempty"`
				CategoryId3      string    `json:"category_id3,omitempty"`
				CategoryId4      string    `json:"category_id4,omitempty"`
				CategoryId5      string    `json:"category_id5,omitempty"`
				Zip              string    `json:"zip,omitempty"`
				Address          string    `json:"address"`
				NorthLatitude    string    `json:"north_latitude,omitempty"`
				EastLongitude    string    `json:"east_longitude,omitempty"`
				Description      string    `json:"description"`
				Purpose          string    `json:"purpose,omitempty"`
				OpenMorning      int       `json:"open_morning,omitempty"`
				OpenLunch        int       `json:"open_lunch,omitempty"`
				OpenLate         int       `json:"open_late,omitempty"`
				PhotoCount       int       `json:"photo_count,omitempty"`
				SpecialCount     int       `json:"special_count,omitempty"`
				MenuCount        int       `json:"menu_count,omitempty"`
				FanCount         int       `json:"fan_count,omitempty"`
				AccessCount      int       `json:"access_count"`
				CreatedOn        string    `json:"created_on,omitempty"`
				ModifiedOn       string    `json:"modified_on,omitempty"`
				Closed           int       `json:"closed,omitempty"`
				AreaName         string    `json:"area_name"`
				PrefName         string    `json:"pref_name"`
				Pref             string    `json:"pref,omitempty"`
				Location         []float64 `json:"location"`
				Stas             []string  `json:"stas"`
				Cates            []string  `json:"cates"`
				Kuchikomi        []string  `json:"kuchikomi,omitempty"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
