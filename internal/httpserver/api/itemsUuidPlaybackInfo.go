package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"opal/internal/jfstructs"
	"opal/internal/librarymgmt"
	"opal/internal/mediamgmt"
	"path"
	"strconv"
	"strings"
)

func EndpointItemsUuidPlaybackInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !isHeaderAuthTokenValid(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	var req jfstructs.RequestItemsUuidPlaybackInfo
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	media := librarymgmt.LibTreeMap[req.MediaSourceID]
	if media == nil || media.Type != "movie" {
		http.Error(w, "MediaSourceID does not resolve to a valid media uuid", http.StatusNotFound)
		return
	}

	libraryRecord := librarymgmt.AllLibrariesMap[media.RootUuid]

	if libraryRecord == nil {
		http.Error(w, "Library not found", http.StatusNotFound)
		return
	}

	mediaFilePath := path.Join(libraryRecord.Path, media.Path)
	mediaProbe, err := mediamgmt.Probe(mediaFilePath)

	if err != nil {
		http.Error(w, "ffprobe failed", http.StatusInternalServerError)
		return
	}

	//TODO: should do this in the treenode struct
	mS := strings.Split(media.Path, ".")
	if len(mS) > 1 {
		mediaProbe.Format.FormatName = mS[len(mS)-1]
	}

	AudioStreamIndexInt, errA := strconv.Atoi(req.AudioStreamIndex)
	SubtitleStreamIndexInt, errS := strconv.Atoi(req.SubtitleStreamIndex)

	if errA != nil || errS != nil {
		http.Error(w, "Unabled to convert req.AudioStreamIndex or req.SubtitleStreamIndex to integers", http.StatusBadRequest)
		return
	}

	//TODO: assumes video is first stream in container
	if len(mediaProbe.Streams) == 0 || len(mediaProbe.Streams) < (AudioStreamIndexInt+1) || len(mediaProbe.Streams) < (SubtitleStreamIndexInt+1) {
		http.Error(w, "Stream selection out of bounds", http.StatusBadRequest)
		return
	}

	videoStream := mediaProbe.Streams[0]
	audioStream := mediaProbe.Streams[AudioStreamIndexInt]

	//TODO: add subtitle support
	//var subtitleStream mediamgmt.Stream

	//if SubtitleStreamIndexInt > -1 {
	//	subtitleStream = mediaProbe.Streams[SubtitleStreamIndexInt]
	//}

	if canDirectPlay(mediaProbe, &videoStream, &audioStream, &req) {
		serveDirectplay(w, media, mediaProbe, AudioStreamIndexInt)
		return
	} else {
		serveHls()
		http.Error(w, "HLS is WIP", http.StatusNotImplemented)
		return
	}
}

func canDirectPlay(
	mProbe *mediamgmt.FfprobeResult,
	videoS *mediamgmt.Stream,
	audioS *mediamgmt.Stream,
	//	subtitleS *mediamgmt.Stream,
	req *jfstructs.RequestItemsUuidPlaybackInfo) bool {

	for _, dPP := range req.DeviceProfile.DirectPlayProfiles {
		if dPP.Type != "Video" {
			continue
		}

		if dPP.Container != "" &&
			!strings.Contains(dPP.Container, mProbe.Format.FormatName) &&
			!strings.Contains(mProbe.Format.FormatName, dPP.Container) {
			continue
		}

		if dPP.VideoCodec != "" && !strings.Contains(dPP.VideoCodec, videoS.CodecName) {
			continue
		}

		if dPP.AudioCodec != "" && !strings.Contains(dPP.AudioCodec, audioS.CodecName) {
			continue
		}

		return true
	}

	//TODO: need to check codec profiles

	return false
}

func serveDirectplay(w http.ResponseWriter, media *librarymgmt.TreeNode, mediaProbe *mediamgmt.FfprobeResult, AudioStreamIndexInt int) {
	mPSize, err := strconv.Atoi(mediaProbe.Format.Size)

	if err != nil {
		http.Error(w, "ffprobe error: failed to convert mediaProve.Format.Size to integer", http.StatusBadRequest)
	}

	resMediaSource := jfstructs.ResponseItemsUuidPlaybackInfoMediaSource{
		Protocol:  "File",
		ID:        media.Uuid,
		Path:      fmt.Sprintf("/Libraries/%s%s", media.RootUuid, media.Path),
		Type:      "Default",
		Container: mediaProbe.Format.FormatName,
		Size:      mPSize,
		Name:      media.Name,
		IsRemote:  true,
		ETag:      "STUBBED",
		/*
			RunTimeTicks                        int64  `json:"RunTimeTicks"`
		*/
		ReadAtNativeFramerate:               false,
		IgnoreDts:                           false,
		IgnoreIndex:                         false,
		GenPtsInput:                         false,
		SupportsTranscoding:                 false,
		SupportsDirectStream:                false,
		SupportsDirectPlay:                  true,
		IsInfiniteStream:                    false,
		UseMostCompatibleTranscodingProfile: false,
		RequiresOpening:                     false,
		RequiresClosing:                     false,
		RequiresLooping:                     false,
		SupportsProbing:                     true,
		VideoType:                           "VideoFile",
		/*
			MediaStreams []CommonStream `json:"MediaStreams"`
		*/
		MediaAttachments: []any{},
		Formats:          []any{},
		/*
			Bitrate             int   `json:"Bitrate"`
			RequiredHTTPHeaders struct {
			} `json:"RequiredHttpHeaders"`
			TranscodingSubProtocol     string `json:"TranscodingSubProtocol"`

			DefaultSubtitleStreamIndex int    `json:"DefaultSubtitleStreamIndex"`
		*/
		DefaultAudioStreamIndex: AudioStreamIndexInt,
		HasSegments:             false,
	}

	for _, stream := range mediaProbe.Streams {
		switch stream.CodecType {
		case "video":
			resMediaSource.MediaStreams = append(resMediaSource.MediaStreams, buildMediaStreamVideo(stream))
		case "audio":
			resMediaSource.MediaStreams = append(resMediaSource.MediaStreams, buildMediaStreamAudio(stream))
		case "subtitle":
			resMediaSource.MediaStreams = append(resMediaSource.MediaStreams, buildMediaStreamSubtitle(stream))
		}
	}

	res := jfstructs.ResponseItemsUuidPlaybackInfo{
		PlaySessionID: "STUBBED",
	}
	res.MediaSources = append(res.MediaSources, resMediaSource)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func serveHls() {
	//TODO: implement hls
}
