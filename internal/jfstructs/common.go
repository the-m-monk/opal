package jfstructs

import "time"

type CommonUsersConfiguration struct {
	AudioLanguagePreference    string   `json:"AudioLanguagePreference,omitempty"` //usersPublic does not implement but usersAuthenticateByName does
	PlayDefaultAudioTrack      bool     `json:"PlayDefaultAudioTrack"`
	SubtitleLanguagePreference string   `json:"SubtitleLanguagePreference"`
	DisplayMissingEpisodes     bool     `json:"DisplayMissingEpisodes"`
	GroupedFolders             []any    `json:"GroupedFolders"`
	SubtitleMode               string   `json:"SubtitleMode"`
	DisplayCollectionsView     bool     `json:"DisplayCollectionsView"`
	EnableLocalPassword        bool     `json:"EnableLocalPassword"`
	OrderedViews               []string `json:"OrderedViews"`
	LatestItemsExcludes        []any    `json:"LatestItemsExcludes"`
	MyMediaExcludes            []any    `json:"MyMediaExcludes"`
	HidePlayedInLatest         bool     `json:"HidePlayedInLatest"`
	RememberAudioSelections    bool     `json:"RememberAudioSelections"`
	RememberSubtitleSelections bool     `json:"RememberSubtitleSelections"`
	EnableNextEpisodeAutoPlay  bool     `json:"EnableNextEpisodeAutoPlay"`
	CastReceiverID             string   `json:"CastReceiverId"`
}

type CommonUsersPolicy struct {
	IsAdministrator                  bool   `json:"IsAdministrator"`
	IsHidden                         bool   `json:"IsHidden"`
	EnableCollectionManagement       bool   `json:"EnableCollectionManagement"`
	EnableSubtitleManagement         bool   `json:"EnableSubtitleManagement"`
	EnableLyricManagement            bool   `json:"EnableLyricManagement"`
	IsDisabled                       bool   `json:"IsDisabled"`
	BlockedTags                      []any  `json:"BlockedTags"`
	AllowedTags                      []any  `json:"AllowedTags"`
	EnableUserPreferenceAccess       bool   `json:"EnableUserPreferenceAccess"`
	AccessSchedules                  []any  `json:"AccessSchedules"`
	BlockUnratedItems                []any  `json:"BlockUnratedItems"`
	EnableRemoteControlOfOtherUsers  bool   `json:"EnableRemoteControlOfOtherUsers"`
	EnableSharedDeviceControl        bool   `json:"EnableSharedDeviceControl"`
	EnableRemoteAccess               bool   `json:"EnableRemoteAccess"`
	EnableLiveTvManagement           bool   `json:"EnableLiveTvManagement"`
	EnableLiveTvAccess               bool   `json:"EnableLiveTvAccess"`
	EnableMediaPlayback              bool   `json:"EnableMediaPlayback"`
	EnableAudioPlaybackTranscoding   bool   `json:"EnableAudioPlaybackTranscoding"`
	EnableVideoPlaybackTranscoding   bool   `json:"EnableVideoPlaybackTranscoding"`
	EnablePlaybackRemuxing           bool   `json:"EnablePlaybackRemuxing"`
	ForceRemoteSourceTranscoding     bool   `json:"ForceRemoteSourceTranscoding"`
	EnableContentDeletion            bool   `json:"EnableContentDeletion"`
	EnableContentDeletionFromFolders []any  `json:"EnableContentDeletionFromFolders"`
	EnableContentDownloading         bool   `json:"EnableContentDownloading"`
	EnableSyncTranscoding            bool   `json:"EnableSyncTranscoding"`
	EnableMediaConversion            bool   `json:"EnableMediaConversion"`
	EnabledDevices                   []any  `json:"EnabledDevices"`
	EnableAllDevices                 bool   `json:"EnableAllDevices"`
	EnabledChannels                  []any  `json:"EnabledChannels"`
	EnableAllChannels                bool   `json:"EnableAllChannels"`
	EnabledFolders                   []any  `json:"EnabledFolders"`
	EnableAllFolders                 bool   `json:"EnableAllFolders"`
	InvalidLoginAttemptCount         int    `json:"InvalidLoginAttemptCount"`
	LoginAttemptsBeforeLockout       int    `json:"LoginAttemptsBeforeLockout"`
	MaxActiveSessions                int    `json:"MaxActiveSessions"`
	EnablePublicSharing              bool   `json:"EnablePublicSharing"`
	BlockedMediaFolders              []any  `json:"BlockedMediaFolders"`
	BlockedChannels                  []any  `json:"BlockedChannels"`
	RemoteClientBitrateLimit         int    `json:"RemoteClientBitrateLimit"`
	AuthenticationProviderID         string `json:"AuthenticationProviderId"`
	PasswordResetProviderID          string `json:"PasswordResetProviderId"`
	SyncPlayAccess                   string `json:"SyncPlayAccess"`
}

type CommonUser struct {
	Name                      string                   `json:"Name"`
	ServerID                  string                   `json:"ServerId"`
	ID                        string                   `json:"Id"`
	PrimaryImageTag           string                   `json:"PrimaryImageTag,omitempty"`
	HasPassword               bool                     `json:"HasPassword"`
	HasConfiguredPassword     bool                     `json:"HasConfiguredPassword"`
	HasConfiguredEasyPassword bool                     `json:"HasConfiguredEasyPassword"`
	EnableAutoLogin           bool                     `json:"EnableAutoLogin"`
	LastLoginDate             time.Time                `json:"LastLoginDate"`
	LastActivityDate          time.Time                `json:"LastActivityDate"`
	Configuration             CommonUsersConfiguration `json:"Configuration"`
	Policy                    CommonUsersPolicy        `json:"Policy"`
}

type CommonItemUserData struct {
	PlayedPercentage      float64   `json:"PlayedPercentage"`
	PlaybackPositionTicks int64     `json:"PlaybackPositionTicks"`
	PlayCount             int       `json:"PlayCount"`
	IsFavorite            bool      `json:"IsFavorite"`
	LastPlayedDate        time.Time `json:"LastPlayedDate"`
	Played                bool      `json:"Played"`
	Key                   string    `json:"Key"`
	ItemID                string    `json:"ItemId"`
}

type CommonItem struct {
	Name                    string             `json:"Name"`
	ServerID                string             `json:"ServerId"`
	ID                      string             `json:"Id"`
	Path                    string             `json:"Path"`
	HasSubtitles            bool               `json:"HasSubtitles"`
	Container               string             `json:"Container"`
	PremiereDate            time.Time          `json:"PremiereDate"`
	ChannelID               any                `json:"ChannelId"`
	RunTimeTicks            int64              `json:"RunTimeTicks"`
	ProductionYear          int                `json:"ProductionYear"`
	IsFolder                bool               `json:"IsFolder"`
	Type                    string             `json:"Type"`
	UserData                CommonItemUserData `json:"UserData"`
	PrimaryImageAspectRatio float64            `json:"PrimaryImageAspectRatio"`
	VideoType               string             `json:"VideoType"`
	ImageTags               struct {
		Primary string `json:"Primary"`
	} `json:"ImageTags"`
	/*

		BackdropImageTags []any `json:"BackdropImageTags"`
		ImageBlurHashes   struct {
			Primary struct {
				Four1B26B069Aaad042A8293C6Fe41D9738 string `json:"41b26b069aaad042a8293c6fe41d9738"`
			} `json:"Primary"`
		} `json:"ImageBlurHashes"`
	*/
	LocationType string `json:"LocationType"`
	MediaType    string `json:"MediaType"`
}

type CommonItemList struct {
	Items            []CommonItem `json:"Items"`
	TotalRecordCount int          `json:"TotalRecordCount"`
	StartIndex       int          `json:"StartIndex"`
}

/*
"MediaStreams": [
                {
                    "Codec": "hevc",
                    "TimeBase": "1/1000",
                    "VideoRange": "SDR",
                    "VideoRangeType": "SDR",
                    "AudioSpatialFormat": "None",
                    "DisplayTitle": "1080p HEVC SDR",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "BitRate": 2480084,
                    "BitDepth": 10,
                    "RefFrames": 1,
                    "IsDefault": true,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 1040,
                    "Width": 1920,
                    "AverageFrameRate": 23.976025,
                    "RealFrameRate": 23.976025,
                    "ReferenceFrameRate": 23.976025,
                    "Profile": "Main 10",
                    "Type": "Video",
                    "AspectRatio": "1.85:1",
                    "Index": 0,
                    "IsExternal": false,
                    "IsTextSubtitleStream": false,
                    "SupportsExternalStream": false,
                    "PixelFormat": "yuv420p10le",
                    "Level": 120,
                    "IsAnamorphic": false
                },
                {
                    "Codec": "aac",
                    "Language": "eng",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedDefault": "Default",
                    "LocalizedExternal": "External",
                    "DisplayTitle": "English - AAC - Stereo",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "ChannelLayout": "stereo",
                    "BitRate": 184134,
                    "Channels": 2,
                    "SampleRate": 48000,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Profile": "LC",
                    "Type": "Audio",
                    "Index": 1,
                    "IsExternal": false,
                    "IsTextSubtitleStream": false,
                    "SupportsExternalStream": false,
                    "Level": 0
                },
                {
                    "Codec": "aac",
                    "Language": "jpn",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedDefault": "Default",
                    "LocalizedExternal": "External",
                    "DisplayTitle": "Japanese - AAC - Stereo",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "ChannelLayout": "stereo",
                    "BitRate": 174185,
                    "Channels": 2,
                    "SampleRate": 48000,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Profile": "LC",
                    "Type": "Audio",
                    "Index": 2,
                    "IsExternal": false,
                    "IsTextSubtitleStream": false,
                    "SupportsExternalStream": false,
                    "Level": 0
                },
                {
                    "Codec": "PGSSUB",
                    "Language": "eng",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "English - PGSSUB",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 1080,
                    "Width": 1920,
                    "Type": "Subtitle",
                    "Index": 3,
                    "IsExternal": false,
                    "IsTextSubtitleStream": false,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "PGSSUB",
                    "Language": "jpn",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Japanese - PGSSUB",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 1080,
                    "Width": 1920,
                    "Type": "Subtitle",
                    "Index": 4,
                    "IsExternal": false,
                    "IsTextSubtitleStream": false,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "ara",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Arabic - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 5,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "dan",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Danish - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 6,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "nld",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Dutch - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 7,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "fas",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Persian - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 8,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "fra",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "French - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 9,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "deu",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "German - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 10,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "ass",
                    "Language": "vie",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Vietnamese - ASS",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 11,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "ass",
                    "Language": "hun",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Hungarian - ASS",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 12,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "ass",
                    "Language": "ita",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Italian - ASS",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 13,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "ass",
                    "Language": "pol",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Polish - ASS",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 14,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "ass",
                    "Language": "spa",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Spanish - ASS",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 15,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "swe",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Swedish - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 16,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "subrip",
                    "Language": "tha",
                    "TimeBase": "1/1000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "LocalizedUndefined": "Undefined",
                    "LocalizedDefault": "Default",
                    "LocalizedForced": "Forced",
                    "LocalizedExternal": "External",
                    "LocalizedHearingImpaired": "Hearing Impaired",
                    "DisplayTitle": "Thai - SUBRIP",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 0,
                    "Width": 0,
                    "Type": "Subtitle",
                    "Index": 17,
                    "IsExternal": false,
                    "IsTextSubtitleStream": true,
                    "SupportsExternalStream": true,
                    "Level": 0
                },
                {
                    "Codec": "mjpeg",
                    "ColorSpace": "bt470bg",
                    "TimeBase": "1/90000",
                    "VideoRange": "Unknown",
                    "VideoRangeType": "Unknown",
                    "AudioSpatialFormat": "None",
                    "IsInterlaced": false,
                    "IsAVC": false,
                    "BitDepth": 8,
                    "RefFrames": 1,
                    "IsDefault": false,
                    "IsForced": false,
                    "IsHearingImpaired": false,
                    "Height": 314,
                    "Width": 222,
                    "RealFrameRate": 90000,
                    "ReferenceFrameRate": 90000,
                    "Profile": "Baseline",
                    "Type": "EmbeddedImage",
                    "AspectRatio": "111:157",
                    "Index": 18,
                    "IsExternal": false,
                    "IsTextSubtitleStream": false,
                    "SupportsExternalStream": false,
                    "PixelFormat": "yuvj420p",
                    "Level": -99,
                    "IsAnamorphic": false
                }
            ],
*/

type CommonStream struct {
	Codec                    string  `json:"Codec"`
	CodecTag                 string  `json:"CodecTag,omitempty"`
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
}
