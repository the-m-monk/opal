package jfstructs

import "time"

type ResponseUsersItemsByUuidMovieMediaSources struct {
	Protocol                            string         `json:"Protocol"`
	ID                                  string         `json:"Id"`
	Path                                string         `json:"Path"`
	Type                                string         `json:"Type"`
	Container                           string         `json:"Container"`
	Size                                int            `json:"Size"`
	Name                                string         `json:"Name"`
	IsRemote                            bool           `json:"IsRemote"`
	ETag                                string         `json:"ETag"`
	RunTimeTicks                        int64          `json:"RunTimeTicks"`
	ReadAtNativeFramerate               bool           `json:"ReadAtNativeFramerate"`
	IgnoreDts                           bool           `json:"IgnoreDts"`
	IgnoreIndex                         bool           `json:"IgnoreIndex"`
	GenPtsInput                         bool           `json:"GenPtsInput"`
	SupportsTranscoding                 bool           `json:"SupportsTranscoding"`
	SupportsDirectStream                bool           `json:"SupportsDirectStream"`
	SupportsDirectPlay                  bool           `json:"SupportsDirectPlay"`
	IsInfiniteStream                    bool           `json:"IsInfiniteStream"`
	UseMostCompatibleTranscodingProfile bool           `json:"UseMostCompatibleTranscodingProfile"`
	RequiresOpening                     bool           `json:"RequiresOpening"`
	RequiresClosing                     bool           `json:"RequiresClosing"`
	RequiresLooping                     bool           `json:"RequiresLooping"`
	SupportsProbing                     bool           `json:"SupportsProbing"`
	VideoType                           string         `json:"VideoType"`
	MediaStreams                        []CommonStream `json:"MediaStreams"`
	MediaAttachments                    []any          `json:"MediaAttachments"`
	Formats                             []any          `json:"Formats"`
	Bitrate                             int            `json:"Bitrate"`
	RequiredHTTPHeaders                 struct {
	} `json:"RequiredHttpHeaders"`
	TranscodingSubProtocol  string `json:"TranscodingSubProtocol"`
	DefaultAudioStreamIndex int    `json:"DefaultAudioStreamIndex"`
	HasSegments             bool   `json:"HasSegments"`
}

type ResponseUsersItemsByUuidMovieExternalUrls struct {
	Name string `json:"Name"`
	URL  string `json:"Url"`
}

type ResponseUsersItemsByUuidMovie struct {
	Name            string                                      `json:"Name"`
	OriginalTitle   string                                      `json:"OriginalTitle"`
	ServerID        string                                      `json:"ServerId"`
	ID              string                                      `json:"Id"`
	Container       string                                      `json:"Container"`
	Overview        string                                      `json:"Overview"`
	Path            string                                      `json:"Path"`
	RunTimeTicks    int64                                       `json:"RunTimeTicks"`
	PlayAccess      string                                      `json:"PlayAccess"`
	MediaType       string                                      `json:"MediaType"`
	LocationType    string                                      `json:"LocationType"`
	MediaStreams    []CommonStream                              `json:"MediaStreams"`
	MediaSources    []ResponseUsersItemsByUuidMovieMediaSources `json:"MediaSources"`
	VideoType       string                                      `json:"VideoType"`
	CanDownload     bool                                        `json:"CanDownload"`
	CanDelete       bool                                        `json:"CanDelete"`
	Taglines        []string                                    `json:"Taglines"`
	PremiereDate    time.Time                                   `json:"PremiereDate"`
	OfficialRating  string                                      `json:"OfficialRating"`
	CommunityRating float64                                     `json:"CommunityRating"`
	ExternalUrls    []ResponseUsersItemsByUuidMovieExternalUrls `json:"ExternalUrls"`

	ImageTags struct {
		Logo    string `json:"Logo"`
		Thumb   string `json:"Thumb"`
		Primary string `json:"Primary"`
	} `json:"ImageTags"`
	PrimaryImageAspectRatio float64  `json:"PrimaryImageAspectRatio"`
	BackdropImageTags       []string `json:"BackdropImageTags"`

	//Etag          string    `json:"Etag"`
	//DateCreated  time.Time `json:"DateCreated"`

	//HasSubtitles bool      `json:"HasSubtitles"`

	//SortName     string    `json:"SortName"`

	//LockedFields []any     `json:"LockedFields"`
	//LockData     bool      `json:"LockData"`
	//Width        int       `json:"Width"`
	//Height       int       `json:"Height"`

	/*

	 */
	/*

			MediaAttachments []struct {
				Codec    string `json:"Codec"`
				CodecTag string `json:"CodecTag"`
				Index    int    `json:"Index"`
				FileName string `json:"FileName"`
				MimeType string `json:"MimeType"`
			} `json:"MediaAttachments"`
			Formats             []any `json:"Formats"`
			Bitrate             int   `json:"Bitrate"`
			RequiredHTTPHeaders struct {
			} `json:"RequiredHttpHeaders"`
			TranscodingSubProtocol     string `json:"TranscodingSubProtocol"`
			DefaultAudioStreamIndex    int    `json:"DefaultAudioStreamIndex"`
			DefaultSubtitleStreamIndex int    `json:"DefaultSubtitleStreamIndex"`
			HasSegments                bool   `json:"HasSegments"`
		} `json:"MediaSources"`
	*/
	/*
		CriticRating             int      `json:"CriticRating"`
		ProductionLocations      []string `json:"ProductionLocations"`

		EnableMediaSourceDisplay bool     `json:"EnableMediaSourceDisplay"`

		ChannelID                any      `json:"ChannelId"`


		Genres                   []string `json:"Genres"`


		ProductionYear           int      `json:"ProductionYear"`
		RemoteTrailers           []struct {
			URL  string `json:"Url"`
			Name string `json:"Name"`
		} `json:"RemoteTrailers"`
		ProviderIds struct {
			Imdb           string `json:"Imdb"`
			Tmdb           string `json:"Tmdb"`
			TmdbCollection string `json:"TmdbCollection"`
		} `json:"ProviderIds"`
		IsHD     bool   `json:"IsHD"`
		IsFolder bool   `json:"IsFolder"`
		ParentID string `json:"ParentId"`
		Type     string `json:"Type"`
	*/
	/*

	 */
	/*
		Studios []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"Studios"`
		GenreItems []struct {
			Name string `json:"Name"`
			ID   string `json:"Id"`
		} `json:"GenreItems"`
		LocalTrailerCount int `json:"LocalTrailerCount"`
	*/
	/*
		UserData struct {
			PlayedPercentage      float64   `json:"PlayedPercentage"`
			PlaybackPositionTicks int64     `json:"PlaybackPositionTicks"`
			PlayCount             int       `json:"PlayCount"`
			IsFavorite            bool      `json:"IsFavorite"`
			LastPlayedDate        time.Time `json:"LastPlayedDate"`
			Played                bool      `json:"Played"`
			Key                   string    `json:"Key"`
			ItemID                string    `json:"ItemId"`
		} `json:"UserData"`
		SpecialFeatureCount     int      `json:"SpecialFeatureCount"`
		DisplayPreferencesID    string   `json:"DisplayPreferencesId"`
		Tags                    []string `json:"Tags"`

		MediaStreams            []struct {
			Codec                    string  `json:"Codec"`
			TimeBase                 string  `json:"TimeBase"`
			VideoRange               string  `json:"VideoRange"`
			VideoRangeType           string  `json:"VideoRangeType"`
			AudioSpatialFormat       string  `json:"AudioSpatialFormat"`
			DisplayTitle             string  `json:"DisplayTitle,omitempty"`
			IsInterlaced             bool    `json:"IsInterlaced"`
			IsAVC                    bool    `json:"IsAVC"`
			BitRate                  int     `json:"BitRate,omitempty"`
			BitDepth                 int     `json:"BitDepth,omitempty"`
			RefFrames                int     `json:"RefFrames,omitempty"`
			IsDefault                bool    `json:"IsDefault"`
			IsForced                 bool    `json:"IsForced"`
			IsHearingImpaired        bool    `json:"IsHearingImpaired"`
			Height                   int     `json:"Height,omitempty"`
			Width                    int     `json:"Width,omitempty"`
			AverageFrameRate         float64 `json:"AverageFrameRate,omitempty"`
			RealFrameRate            float64 `json:"RealFrameRate,omitempty"`
			ReferenceFrameRate       float64 `json:"ReferenceFrameRate,omitempty"`
			Profile                  string  `json:"Profile,omitempty"`
			Type                     string  `json:"Type"`
			AspectRatio              string  `json:"AspectRatio,omitempty"`
			Index                    int     `json:"Index"`
			IsExternal               bool    `json:"IsExternal"`
			IsTextSubtitleStream     bool    `json:"IsTextSubtitleStream"`
			SupportsExternalStream   bool    `json:"SupportsExternalStream"`
			PixelFormat              string  `json:"PixelFormat,omitempty"`
			Level                    int     `json:"Level"`
			IsAnamorphic             bool    `json:"IsAnamorphic,omitempty"`
			Language                 string  `json:"Language,omitempty"`
			LocalizedDefault         string  `json:"LocalizedDefault,omitempty"`
			LocalizedExternal        string  `json:"LocalizedExternal,omitempty"`
			ChannelLayout            string  `json:"ChannelLayout,omitempty"`
			Channels                 int     `json:"Channels,omitempty"`
			SampleRate               int     `json:"SampleRate,omitempty"`
			LocalizedUndefined       string  `json:"LocalizedUndefined,omitempty"`
			LocalizedForced          string  `json:"LocalizedForced,omitempty"`
			LocalizedHearingImpaired string  `json:"LocalizedHearingImpaired,omitempty"`
			ColorSpace               string  `json:"ColorSpace,omitempty"`
		} `json:"MediaStreams"`


		BackdropImageTags []string `json:"BackdropImageTags"`
		ImageBlurHashes   struct {
			Backdrop struct {
				SixD10Dfc7Cea2613F76324108A2518334 string `json:"6d10dfc7cea2613f76324108a2518334"`
			} `json:"Backdrop"`
			Logo struct {
				Ecf86C54317Ef8676E98Ca49E688Df29 string `json:"ecf86c54317ef8676e98ca49e688df29"`
			} `json:"Logo"`
			Thumb struct {
				D12457F1F322Ff757Bdfcfae0E839Bdf string `json:"d12457f1f322ff757bdfcfae0e839bdf"`
			} `json:"Thumb"`
			Primary struct {
				NineC3152Facacd7C88A6Db5E4C5B041051 string `json:"9c3152facacd7c88a6db5e4c5b041051"`
			} `json:"Primary"`
		} `json:"ImageBlurHashes"`
		Chapters []struct {
			StartPositionTicks int       `json:"StartPositionTicks"`
			Name               string    `json:"Name"`
			ImageDateModified  time.Time `json:"ImageDateModified"`
		} `json:"Chapters"`
		Trickplay struct {
		} `json:"Trickplay"`
	*/
}

/*
type AutoGenerated struct {
	Name          string    `json:"Name"`
	OriginalTitle string    `json:"OriginalTitle"`
	ServerID      string    `json:"ServerId"`
	ID            string    `json:"Id"`
	Etag          string    `json:"Etag"`
	DateCreated   time.Time `json:"DateCreated"`
	CanDelete     bool      `json:"CanDelete"`
	CanDownload   bool      `json:"CanDownload"`
	Container     string    `json:"Container"`
	SortName      string    `json:"SortName"`
	PremiereDate  time.Time `json:"PremiereDate"`
	ExternalUrls  []struct {
		Name string `json:"Name"`
		URL  string `json:"Url"`
	} `json:"ExternalUrls"`
	MediaSources []struct {
		Protocol                            string `json:"Protocol"`
		ID                                  string `json:"Id"`
		Path                                string `json:"Path"`
		Type                                string `json:"Type"`
		Container                           string `json:"Container"`
		Size                                int    `json:"Size"`
		Name                                string `json:"Name"`
		IsRemote                            bool   `json:"IsRemote"`
		ETag                                string `json:"ETag"`
		RunTimeTicks                        int64  `json:"RunTimeTicks"`
		ReadAtNativeFramerate               bool   `json:"ReadAtNativeFramerate"`
		IgnoreDts                           bool   `json:"IgnoreDts"`
		IgnoreIndex                         bool   `json:"IgnoreIndex"`
		GenPtsInput                         bool   `json:"GenPtsInput"`
		SupportsTranscoding                 bool   `json:"SupportsTranscoding"`
		SupportsDirectStream                bool   `json:"SupportsDirectStream"`
		SupportsDirectPlay                  bool   `json:"SupportsDirectPlay"`
		IsInfiniteStream                    bool   `json:"IsInfiniteStream"`
		UseMostCompatibleTranscodingProfile bool   `json:"UseMostCompatibleTranscodingProfile"`
		RequiresOpening                     bool   `json:"RequiresOpening"`
		RequiresClosing                     bool   `json:"RequiresClosing"`
		RequiresLooping                     bool   `json:"RequiresLooping"`
		SupportsProbing                     bool   `json:"SupportsProbing"`
		VideoType                           string `json:"VideoType"`
		MediaStreams                        []struct {
			Codec                  string `json:"Codec"`
			CodecTag               string `json:"CodecTag"`
			Language               string `json:"Language"`
			TimeBase               string `json:"TimeBase"`
			VideoRange             string `json:"VideoRange"`
			VideoRangeType         string `json:"VideoRangeType"`
			AudioSpatialFormat     string `json:"AudioSpatialFormat"`
			DisplayTitle           string `json:"DisplayTitle"`
			NalLengthSize          string `json:"NalLengthSize,omitempty"`
			IsInterlaced           bool   `json:"IsInterlaced"`
			IsAVC                  bool   `json:"IsAVC"`
			BitRate                int    `json:"BitRate"`
			BitDepth               int    `json:"BitDepth,omitempty"`
			RefFrames              int    `json:"RefFrames,omitempty"`
			IsDefault              bool   `json:"IsDefault"`
			IsForced               bool   `json:"IsForced"`
			IsHearingImpaired      bool   `json:"IsHearingImpaired"`
			Height                 int    `json:"Height,omitempty"`
			Width                  int    `json:"Width,omitempty"`
			AverageFrameRate       int    `json:"AverageFrameRate,omitempty"`
			RealFrameRate          int    `json:"RealFrameRate,omitempty"`
			ReferenceFrameRate     int    `json:"ReferenceFrameRate,omitempty"`
			Profile                string `json:"Profile"`
			Type                   string `json:"Type"`
			AspectRatio            string `json:"AspectRatio,omitempty"`
			Index                  int    `json:"Index"`
			IsExternal             bool   `json:"IsExternal"`
			IsTextSubtitleStream   bool   `json:"IsTextSubtitleStream"`
			SupportsExternalStream bool   `json:"SupportsExternalStream"`
			PixelFormat            string `json:"PixelFormat,omitempty"`
			Level                  int    `json:"Level"`
			IsAnamorphic           bool   `json:"IsAnamorphic,omitempty"`
			Title                  string `json:"Title,omitempty"`
			LocalizedDefault       string `json:"LocalizedDefault,omitempty"`
			LocalizedExternal      string `json:"LocalizedExternal,omitempty"`
			ChannelLayout          string `json:"ChannelLayout,omitempty"`
			Channels               int    `json:"Channels,omitempty"`
			SampleRate             int    `json:"SampleRate,omitempty"`
		} `json:"MediaStreams"`
		MediaAttachments    []any `json:"MediaAttachments"`
		Formats             []any `json:"Formats"`
		Bitrate             int   `json:"Bitrate"`
		RequiredHTTPHeaders struct {
		} `json:"RequiredHttpHeaders"`
		TranscodingSubProtocol  string `json:"TranscodingSubProtocol"`
		DefaultAudioStreamIndex int    `json:"DefaultAudioStreamIndex"`
		HasSegments             bool   `json:"HasSegments"`
	} `json:"MediaSources"`
	CriticRating             int      `json:"CriticRating"`
	ProductionLocations      []string `json:"ProductionLocations"`
	Path                     string   `json:"Path"`
	EnableMediaSourceDisplay bool     `json:"EnableMediaSourceDisplay"`

	ChannelID                any      `json:"ChannelId"`
	Overview                 string   `json:"Overview"`
	Taglines                 []string `json:"Taglines"`
	Genres                   []string `json:"Genres"`
	CommunityRating          float64  `json:"CommunityRating"`
	RunTimeTicks             int64    `json:"RunTimeTicks"`
	PlayAccess               string   `json:"PlayAccess"`
	ProductionYear           int      `json:"ProductionYear"`
	RemoteTrailers           []struct {
		URL  string `json:"Url"`
		Name string `json:"Name"`
	} `json:"RemoteTrailers"`
	ProviderIds struct {
		Imdb string `json:"Imdb"`
		Tmdb string `json:"Tmdb"`
	} `json:"ProviderIds"`
	IsHD     bool   `json:"IsHD"`
	IsFolder bool   `json:"IsFolder"`
	ParentID string `json:"ParentId"`
	Type     string `json:"Type"`
	People   []struct {
		Name            string `json:"Name"`
		ID              string `json:"Id"`
		Role            string `json:"Role"`
		Type            string `json:"Type"`
		PrimaryImageTag string `json:"PrimaryImageTag,omitempty"`
		ImageBlurHashes struct {
			Primary struct {
				ThreeE2Fa050E881C894D264B770E1579E52 string `json:"3e2fa050e881c894d264b770e1579e52"`
			} `json:"Primary"`
		} `json:"ImageBlurHashes"`
	} `json:"People"`
	Studios []struct {
		Name string `json:"Name"`
		ID   string `json:"Id"`
	} `json:"Studios"`
	GenreItems []struct {
		Name string `json:"Name"`
		ID   string `json:"Id"`
	} `json:"GenreItems"`
	LocalTrailerCount int `json:"LocalTrailerCount"`
	UserData          struct {
		PlaybackPositionTicks int    `json:"PlaybackPositionTicks"`
		PlayCount             int    `json:"PlayCount"`
		IsFavorite            bool   `json:"IsFavorite"`
		Played                bool   `json:"Played"`
		Key                   string `json:"Key"`
		ItemID                string `json:"ItemId"`
	} `json:"UserData"`
	SpecialFeatureCount     int      `json:"SpecialFeatureCount"`
	DisplayPreferencesID    string   `json:"DisplayPreferencesId"`
	Tags                    []string `json:"Tags"`
	PrimaryImageAspectRatio float64  `json:"PrimaryImageAspectRatio"`
	MediaStreams            []struct {
		Codec                  string `json:"Codec"`
		CodecTag               string `json:"CodecTag"`
		Language               string `json:"Language"`
		TimeBase               string `json:"TimeBase"`
		VideoRange             string `json:"VideoRange"`
		VideoRangeType         string `json:"VideoRangeType"`
		AudioSpatialFormat     string `json:"AudioSpatialFormat"`
		DisplayTitle           string `json:"DisplayTitle"`
		NalLengthSize          string `json:"NalLengthSize,omitempty"`
		IsInterlaced           bool   `json:"IsInterlaced"`
		IsAVC                  bool   `json:"IsAVC"`
		BitRate                int    `json:"BitRate"`
		BitDepth               int    `json:"BitDepth,omitempty"`
		RefFrames              int    `json:"RefFrames,omitempty"`
		IsDefault              bool   `json:"IsDefault"`
		IsForced               bool   `json:"IsForced"`
		IsHearingImpaired      bool   `json:"IsHearingImpaired"`
		Height                 int    `json:"Height,omitempty"`
		Width                  int    `json:"Width,omitempty"`
		AverageFrameRate       int    `json:"AverageFrameRate,omitempty"`
		RealFrameRate          int    `json:"RealFrameRate,omitempty"`
		ReferenceFrameRate     int    `json:"ReferenceFrameRate,omitempty"`
		Profile                string `json:"Profile"`
		Type                   string `json:"Type"`
		AspectRatio            string `json:"AspectRatio,omitempty"`
		Index                  int    `json:"Index"`
		IsExternal             bool   `json:"IsExternal"`
		IsTextSubtitleStream   bool   `json:"IsTextSubtitleStream"`
		SupportsExternalStream bool   `json:"SupportsExternalStream"`
		PixelFormat            string `json:"PixelFormat,omitempty"`
		Level                  int    `json:"Level"`
		IsAnamorphic           bool   `json:"IsAnamorphic,omitempty"`
		Title                  string `json:"Title,omitempty"`
		LocalizedDefault       string `json:"LocalizedDefault,omitempty"`
		LocalizedExternal      string `json:"LocalizedExternal,omitempty"`
		ChannelLayout          string `json:"ChannelLayout,omitempty"`
		Channels               int    `json:"Channels,omitempty"`
		SampleRate             int    `json:"SampleRate,omitempty"`
	} `json:"MediaStreams"`
	VideoType string `json:"VideoType"`
	ImageTags struct {
		Primary string `json:"Primary"`
		Logo    string `json:"Logo"`
		Thumb   string `json:"Thumb"`
	} `json:"ImageTags"`
	BackdropImageTags []string `json:"BackdropImageTags"`
	ImageBlurHashes   struct {
		Backdrop struct {
			ThreeA7Fde3D3509249054Fd591B77468A69 string `json:"3a7fde3d3509249054fd591b77468a69"`
		} `json:"Backdrop"`
		Primary struct {
			Cf44375D10118801C4540811B46502B3 string `json:"cf44375d10118801c4540811b46502b3"`
		} `json:"Primary"`
		Logo struct {
			Nine9C0B052454Be28857C73Bcf22D12065 string `json:"99c0b052454be28857c73bcf22d12065"`
		} `json:"Logo"`
		Thumb struct {
			D203D5B15Ec7A905C9Cc0E415E34Ca81 string `json:"d203d5b15ec7a905c9cc0e415e34ca81"`
		} `json:"Thumb"`
	} `json:"ImageBlurHashes"`
	Chapters  []any `json:"Chapters"`
	Trickplay struct {
	} `json:"Trickplay"`
	LocationType string `json:"LocationType"`
	MediaType    string `json:"MediaType"`
	LockedFields []any  `json:"LockedFields"`
	LockData     bool   `json:"LockData"`
	Width        int    `json:"Width"`
	Height       int    `json:"Height"`
}
*/
