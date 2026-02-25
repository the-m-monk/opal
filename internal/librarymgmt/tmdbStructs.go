package librarymgmt

import "time"

type tmdbImage struct {
	AspectRatio float64 `json:"aspect_ratio"`
	Height      int     `json:"height"`
	Iso31661    string  `json:"iso_3166_1"`
	Iso6391     string  `json:"iso_639_1"`
	FilePath    string  `json:"file_path"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Width       int     `json:"width"`
}

type tmdbMovie struct {
	Adult               bool   `json:"adult"`
	BackdropPath        string `json:"backdrop_path"`
	BelongsToCollection struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		PosterPath   string `json:"poster_path"`
		BackdropPath string `json:"backdrop_path"`
	} `json:"belongs_to_collection"`
	Budget int `json:"budget"`
	Genres []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage            string   `json:"homepage"`
	ID                  int      `json:"id"`
	ImdbID              string   `json:"imdb_id"`
	OriginCountry       []string `json:"origin_country"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalTitle       string   `json:"original_title"`
	Overview            string   `json:"overview"`
	Popularity          float64  `json:"popularity"`
	PosterPath          string   `json:"poster_path"`
	ProductionCompanies []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso31661 string `json:"iso_3166_1"`
		Name     string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date"`
	Revenue         int    `json:"revenue"`
	Runtime         int    `json:"runtime"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name"`
		Iso6391     string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Credits     struct {
		Cast []struct {
			Adult              bool    `json:"adult"`
			Gender             int     `json:"gender"`
			ID                 int     `json:"id"`
			KnownForDepartment string  `json:"known_for_department"`
			Name               string  `json:"name"`
			OriginalName       string  `json:"original_name"`
			Popularity         float64 `json:"popularity"`
			ProfilePath        string  `json:"profile_path"`
			CastID             int     `json:"cast_id"`
			Character          string  `json:"character"`
			CreditID           string  `json:"credit_id"`
			Order              int     `json:"order"`
		} `json:"cast"`
		Crew []struct {
			Adult              bool    `json:"adult"`
			Gender             int     `json:"gender"`
			ID                 int     `json:"id"`
			KnownForDepartment string  `json:"known_for_department"`
			Name               string  `json:"name"`
			OriginalName       string  `json:"original_name"`
			Popularity         float64 `json:"popularity"`
			ProfilePath        string  `json:"profile_path"`
			CreditID           string  `json:"credit_id"`
			Department         string  `json:"department"`
			Job                string  `json:"job"`
		} `json:"crew"`
	} `json:"credits"`
	Images struct {
		Backdrops []tmdbImage `json:"backdrops"`
		Logos     []tmdbImage `json:"logos"`
		Posters   []tmdbImage `json:"posters"`
	} `json:"images"`
	Videos struct {
		Results []struct {
			Iso6391     string    `json:"iso_639_1"`
			Iso31661    string    `json:"iso_3166_1"`
			Name        string    `json:"name"`
			Key         string    `json:"key"`
			Site        string    `json:"site"`
			Size        int       `json:"size"`
			Type        string    `json:"type"`
			Official    bool      `json:"official"`
			PublishedAt time.Time `json:"published_at"`
			ID          string    `json:"id"`
		} `json:"results"`
	} `json:"videos"`
	ReleaseDates struct {
		Results []struct {
			Iso31661     string `json:"iso_3166_1"`
			ReleaseDates []struct {
				Certification string    `json:"certification"`
				Descriptors   []any     `json:"descriptors"`
				Iso6391       string    `json:"iso_639_1"`
				Note          string    `json:"note"`
				ReleaseDate   time.Time `json:"release_date"`
				Type          int       `json:"type"`
			} `json:"release_dates"`
		} `json:"results"`
	} `json:"release_dates"`
	Keywords struct {
		Keywords []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"keywords"`
	} `json:"keywords"`
	Recommendations struct {
		Page    int `json:"page"`
		Results []struct {
			Adult            bool    `json:"adult"`
			BackdropPath     string  `json:"backdrop_path"`
			ID               int     `json:"id"`
			Title            string  `json:"title"`
			OriginalTitle    string  `json:"original_title"`
			Overview         string  `json:"overview"`
			PosterPath       string  `json:"poster_path"`
			MediaType        string  `json:"media_type"`
			OriginalLanguage string  `json:"original_language"`
			GenreIds         []int   `json:"genre_ids"`
			Popularity       float64 `json:"popularity"`
			ReleaseDate      string  `json:"release_date"`
			Video            bool    `json:"video"`
			VoteAverage      float64 `json:"vote_average"`
			VoteCount        int     `json:"vote_count"`
		} `json:"results"`
		TotalPages   int `json:"total_pages"`
		TotalResults int `json:"total_results"`
	} `json:"recommendations"`
	Similar struct {
		Page    int `json:"page"`
		Results []struct {
			Adult            bool    `json:"adult"`
			BackdropPath     string  `json:"backdrop_path"`
			GenreIds         []int   `json:"genre_ids"`
			ID               int     `json:"id"`
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title"`
			Overview         string  `json:"overview"`
			Popularity       float64 `json:"popularity"`
			PosterPath       string  `json:"poster_path"`
			ReleaseDate      string  `json:"release_date"`
			Title            string  `json:"title"`
			Video            bool    `json:"video"`
			VoteAverage      float64 `json:"vote_average"`
			VoteCount        int     `json:"vote_count"`
		} `json:"results"`
		TotalPages   int `json:"total_pages"`
		TotalResults int `json:"total_results"`
	} `json:"similar"`
	Translations struct {
		Translations []struct {
			Iso31661    string `json:"iso_3166_1"`
			Iso6391     string `json:"iso_639_1"`
			Name        string `json:"name"`
			EnglishName string `json:"english_name"`
			Data        struct {
				Homepage string `json:"homepage"`
				Overview string `json:"overview"`
				Runtime  int    `json:"runtime"`
				Tagline  string `json:"tagline"`
				Title    string `json:"title"`
			} `json:"data"`
		} `json:"translations"`
	} `json:"translations"`
	ExternalIds struct {
		ImdbID      string `json:"imdb_id"`
		WikidataID  string `json:"wikidata_id"`
		FacebookID  string `json:"facebook_id"`
		InstagramID string `json:"instagram_id"`
		TwitterID   string `json:"twitter_id"`
	} `json:"external_ids"`

	PosterPrimaryAspectRatio   float64 `json:"PosterPrimaryAspectRatio"`
	LogoPrimaryAspectRatio     float64 `json:"LogoPrimaryAspectRatio"`
	BackdropPrimaryAspectRatio float64 `json:"BackdropPrimaryAspectRatio"`
}

type tmdbTvshow struct {
	Adult          bool   `json:"adult"`
	BackdropPath   string `json:"backdrop_path"`
	CreatedBy      []any  `json:"created_by"`
	EpisodeRunTime []int  `json:"episode_run_time"`
	FirstAirDate   string `json:"first_air_date"`
	Genres         []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage         string   `json:"homepage"`
	ID               int      `json:"id"`
	InProduction     bool     `json:"in_production"`
	Languages        []string `json:"languages"`
	LastAirDate      string   `json:"last_air_date"`
	LastEpisodeToAir struct {
		ID             int     `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		VoteAverage    float64 `json:"vote_average"`
		VoteCount      int     `json:"vote_count"`
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		EpisodeType    string  `json:"episode_type"`
		ProductionCode string  `json:"production_code"`
		Runtime        int     `json:"runtime"`
		SeasonNumber   int     `json:"season_number"`
		ShowID         int     `json:"show_id"`
		StillPath      string  `json:"still_path"`
	} `json:"last_episode_to_air"`
	Name             string `json:"name"`
	NextEpisodeToAir any    `json:"next_episode_to_air"`
	Networks         []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"networks"`
	NumberOfEpisodes    int      `json:"number_of_episodes"`
	NumberOfSeasons     int      `json:"number_of_seasons"`
	OriginCountry       []string `json:"origin_country"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalName        string   `json:"original_name"`
	Overview            string   `json:"overview"`
	Popularity          float64  `json:"popularity"`
	PosterPath          string   `json:"poster_path"`
	ProductionCompanies []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso31661 string `json:"iso_3166_1"`
		Name     string `json:"name"`
	} `json:"production_countries"`
	Seasons []struct {
		AirDate      string  `json:"air_date"`
		EpisodeCount int     `json:"episode_count"`
		ID           int     `json:"id"`
		Name         string  `json:"name"`
		Overview     string  `json:"overview"`
		PosterPath   string  `json:"poster_path"`
		SeasonNumber int     `json:"season_number"`
		VoteAverage  float64 `json:"vote_average"`
	} `json:"seasons"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name"`
		Iso6391     string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spoken_languages"`
	Status         string  `json:"status"`
	Tagline        string  `json:"tagline"`
	Type           string  `json:"type"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
	ContentRatings struct {
		Results []struct {
			Descriptors []any  `json:"descriptors"`
			Iso31661    string `json:"iso_3166_1"`
			Rating      string `json:"rating"`
		} `json:"results"`
	} `json:"content_ratings"`
	Credits struct {
		Cast []struct {
			Adult              bool    `json:"adult"`
			Gender             int     `json:"gender"`
			ID                 int     `json:"id"`
			KnownForDepartment string  `json:"known_for_department"`
			Name               string  `json:"name"`
			OriginalName       string  `json:"original_name"`
			Popularity         float64 `json:"popularity"`
			ProfilePath        string  `json:"profile_path"`
			Character          string  `json:"character"`
			CreditID           string  `json:"credit_id"`
			Order              int     `json:"order"`
		} `json:"cast"`
		Crew []struct {
			Adult              bool    `json:"adult"`
			Gender             int     `json:"gender"`
			ID                 int     `json:"id"`
			KnownForDepartment string  `json:"known_for_department"`
			Name               string  `json:"name"`
			OriginalName       string  `json:"original_name"`
			Popularity         float64 `json:"popularity"`
			ProfilePath        string  `json:"profile_path"`
			CreditID           string  `json:"credit_id"`
			Department         string  `json:"department"`
			Job                string  `json:"job"`
		} `json:"crew"`
	} `json:"credits"`
	Videos struct {
		Results []struct {
			Iso6391     string    `json:"iso_639_1"`
			Iso31661    string    `json:"iso_3166_1"`
			Name        string    `json:"name"`
			Key         string    `json:"key"`
			Site        string    `json:"site"`
			Size        int       `json:"size"`
			Type        string    `json:"type"`
			Official    bool      `json:"official"`
			PublishedAt time.Time `json:"published_at"`
			ID          string    `json:"id"`
		} `json:"results"`
	} `json:"videos"`
	Images struct {
		Backdrops []tmdbImage `json:"backdrops"`
		Logos     []tmdbImage `json:"logos"`
		Posters   []tmdbImage `json:"posters"`
	} `json:"images"`
	Keywords struct {
		Results []struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"results"`
	} `json:"keywords"`
	ExternalIds struct {
		ImdbID      string `json:"imdb_id"`
		FreebaseMid string `json:"freebase_mid"`
		FreebaseID  string `json:"freebase_id"`
		TvdbID      int    `json:"tvdb_id"`
		TvrageID    int    `json:"tvrage_id"`
		WikidataID  string `json:"wikidata_id"`
		FacebookID  string `json:"facebook_id"`
		InstagramID any    `json:"instagram_id"`
		TwitterID   string `json:"twitter_id"`
	} `json:"external_ids"`
	AggregateCredits struct {
		Cast []struct {
			Adult              bool    `json:"adult"`
			Gender             int     `json:"gender"`
			ID                 int     `json:"id"`
			KnownForDepartment string  `json:"known_for_department"`
			Name               string  `json:"name"`
			OriginalName       string  `json:"original_name"`
			Popularity         float64 `json:"popularity"`
			ProfilePath        string  `json:"profile_path"`
			Roles              []struct {
				CreditID     string `json:"credit_id"`
				Character    string `json:"character"`
				EpisodeCount int    `json:"episode_count"`
			} `json:"roles"`
			TotalEpisodeCount int `json:"total_episode_count"`
			Order             int `json:"order"`
		} `json:"cast"`
		Crew []struct {
			Adult              bool    `json:"adult"`
			Gender             int     `json:"gender"`
			ID                 int     `json:"id"`
			KnownForDepartment string  `json:"known_for_department"`
			Name               string  `json:"name"`
			OriginalName       string  `json:"original_name"`
			Popularity         float64 `json:"popularity"`
			ProfilePath        string  `json:"profile_path"`
			Jobs               []struct {
				CreditID     string `json:"credit_id"`
				Job          string `json:"job"`
				EpisodeCount int    `json:"episode_count"`
			} `json:"jobs"`
			Department        string `json:"department"`
			TotalEpisodeCount int    `json:"total_episode_count"`
		} `json:"crew"`
	} `json:"aggregate_credits"`
	Recommendations struct {
		Page    int `json:"page"`
		Results []struct {
			Adult            bool     `json:"adult"`
			BackdropPath     string   `json:"backdrop_path"`
			ID               int      `json:"id"`
			Name             string   `json:"name"`
			OriginalName     string   `json:"original_name"`
			Overview         string   `json:"overview"`
			PosterPath       string   `json:"poster_path"`
			MediaType        string   `json:"media_type"`
			OriginalLanguage string   `json:"original_language"`
			GenreIds         []int    `json:"genre_ids"`
			Popularity       float64  `json:"popularity"`
			FirstAirDate     string   `json:"first_air_date"`
			VoteAverage      float64  `json:"vote_average"`
			VoteCount        int      `json:"vote_count"`
			OriginCountry    []string `json:"origin_country"`
		} `json:"results"`
		TotalPages   int `json:"total_pages"`
		TotalResults int `json:"total_results"`
	} `json:"recommendations"`

	PosterPrimaryAspectRatio   float64 `json:"PosterPrimaryAspectRatio"`
	LogoPrimaryAspectRatio     float64 `json:"LogoPrimaryAspectRatio"`
	BackdropPrimaryAspectRatio float64 `json:"BackdropPrimaryAspectRatio"`
}
