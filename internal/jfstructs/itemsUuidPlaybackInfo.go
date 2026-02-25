package jfstructs

type RequestItemsUuidPlaybackInfo struct {
	UserID                              string `json:"UserId"`
	StartTimeTicks                      int64  `json:"StartTimeTicks"`
	IsPlayback                          bool   `json:"IsPlayback"`
	AutoOpenLiveStream                  bool   `json:"AutoOpenLiveStream"`
	AudioStreamIndex                    string `json:"AudioStreamIndex"`
	SubtitleStreamIndex                 string `json:"SubtitleStreamIndex"`
	MediaSourceID                       string `json:"MediaSourceId"`
	MaxStreamingBitrate                 int    `json:"MaxStreamingBitrate"`
	AlwaysBurnInSubtitleWhenTranscoding bool   `json:"AlwaysBurnInSubtitleWhenTranscoding"`
	DeviceProfile                       struct {
		MaxStreamingBitrate              int `json:"MaxStreamingBitrate"`
		MaxStaticBitrate                 int `json:"MaxStaticBitrate"`
		MusicStreamingTranscodingBitrate int `json:"MusicStreamingTranscodingBitrate"`
		DirectPlayProfiles               []struct {
			Container  string `json:"Container"`
			Type       string `json:"Type"`
			VideoCodec string `json:"VideoCodec,omitempty"`
			AudioCodec string `json:"AudioCodec,omitempty"`
		} `json:"DirectPlayProfiles"`
		TranscodingProfiles []struct {
			Container              string `json:"Container"`
			Type                   string `json:"Type"`
			AudioCodec             string `json:"AudioCodec"`
			Context                string `json:"Context"`
			Protocol               string `json:"Protocol"`
			MaxAudioChannels       string `json:"MaxAudioChannels"`
			MinSegments            string `json:"MinSegments,omitempty"`
			BreakOnNonKeyFrames    bool   `json:"BreakOnNonKeyFrames,omitempty"`
			EnableAudioVbrEncoding bool   `json:"EnableAudioVbrEncoding,omitempty"`
			VideoCodec             string `json:"VideoCodec,omitempty"`
		} `json:"TranscodingProfiles"`
		ContainerProfiles []any `json:"ContainerProfiles"`
		CodecProfiles     []struct {
			Type       string `json:"Type"`
			Codec      string `json:"Codec,omitempty"`
			Conditions []struct {
				Condition  string `json:"Condition"`
				Property   string `json:"Property"`
				Value      string `json:"Value"`
				IsRequired bool   `json:"IsRequired"`
			} `json:"Conditions"`
		} `json:"CodecProfiles"`
		SubtitleProfiles []struct {
			Format string `json:"Format"`
			Method string `json:"Method"`
		} `json:"SubtitleProfiles"`
		ResponseProfiles []struct {
			Type      string `json:"Type"`
			Container string `json:"Container"`
			MimeType  string `json:"MimeType"`
		} `json:"ResponseProfiles"`
	} `json:"DeviceProfile"`
}

type ResponseItemsUuidPlaybackInfoMediaSource struct {
	Protocol  string `json:"Protocol"`
	ID        string `json:"Id"`
	Path      string `json:"Path"`
	Type      string `json:"Type"`
	Container string `json:"Container"`
	Size      int    `json:"Size"`
	Name      string `json:"Name"`
	IsRemote  bool   `json:"IsRemote"`
	ETag      string `json:"ETag"`
	/*
		RunTimeTicks                        int64  `json:"RunTimeTicks"`
	*/
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
	/*
		Bitrate             int   `json:"Bitrate"`
		RequiredHTTPHeaders struct {
		} `json:"RequiredHttpHeaders"`
		TranscodingSubProtocol     string `json:"TranscodingSubProtocol"`
	*/
	DefaultAudioStreamIndex int `json:"DefaultAudioStreamIndex"`
	/*
		DefaultSubtitleStreamIndex int    `json:"DefaultSubtitleStreamIndex"`
	*/
	HasSegments bool `json:"HasSegments"`
}

type ResponseItemsUuidPlaybackInfo struct {
	MediaSources  []ResponseItemsUuidPlaybackInfoMediaSource `json:"MediaSources"`
	PlaySessionID string                                     `json:"PlaySessionId"`
}
