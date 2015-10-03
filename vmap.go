// Package vmap implement IAB's VMAP 1.0.1 (http://www.iab.net/guidelines/508676/digitalvideo/vsuite/vmap)
package vmap

import "github.com/rs/vast"

// VMAP is the root <VMAP> tag
type VMAP struct {
	// The version of the VMAP spec (should be 1.0)
	Version string `xml:"version,attr"`
	// Zero or more <AdBreak> child elements
	AdBreaks []AdBreak `xml:"AdBreak"`
	// Can be used to express additional information not supported in the VMAP specification.
	Extensions *Extensions `xml:",omitempty"`
}

// AdBreak represents a single ad break, but may allow for multiple ads.
type AdBreak struct {
	// Represents the timing of the ad break. Values of this attribute can be represented
	// in one of the four ways:
	//
	//  - time:￼￼￼￼￼￼￼￼￼￼￼￼ in the format hh:mm:ss or hh:mm:ss.mmm where .mmm is milliseconds
	//    and is optional. The time values is offset from the start of the video
	//    content to the placement of the ad break in the video content timeline.
	//  - percentage: if the duration of the video content is unknown, a percentage
	//    (in the format n% where "n" is a value from 0-100) an be entered and
	//    represents a percentage of the total video content duration from the start
	//    up to the point where the ad break should be entered.
	//  - start/end: for ad breaks taht are inserted at the very start or end of
	//    the video content, the value "start" or "end" can be entered.
	//  - position: In cases where the timing of the ad breaks is unknown (such as
	//    with live content), positional values can be entered in the format #m
	//    where "m" is an integer of 1 or greater and represents the ad break
	//    opportunity. For example, an ad break to be inserted at the first
	//    opportunity for an ad break would enter the value #1. Position values can
	//    only be honored if no other offset values are provided.
	//
	// An ad break may contain an identical time offset as another ad break and is
	// common when a linear ad is followed by a nonlinear ad. Also a VMAP response
	// can contain a mix of offset value types; however, when a mix is values provided,
	// any position value can be ignored.
	TimeOffset Offset `xml:"timeOffset,attr"`
	// Identifies whether the ad break allows "linear", "nonlinear" or "display" ads.
	// Display break types map to VAST companion ads. If more than one type is allowed,
	// they can be entered using a comma between each (no spaces). For example
	// "linear,nonlinear" can be entered. This attribute ensures that only intended ad
	// types are accepted, that the video player displays ad breaks appropriate for the
	// viewr controls and that the video player can optimize video content playback
	// dependent on the ad types being displayed (such as pausing content at the start
	// of a linear ad to ensure precise timing).
	BreakType string `xml:"breakType,attr"`
	// An optional string identifier for the ad break.
	BreakID string `xml:"breakId,attr,omitempty"`
	// An option used to distribute ad breaks equally spaced apart from one another
	// along a linear timeline. If used, the value is time in the format hh:mm:ss or
	// HH:MM:SS.mmm and indicates that the video player should repeat the same
	// <AdBreak> break (using the same <AdSource>) at time offsets equal to the duration
	// value of this attribute. Should a conflict occur where the duration of an
	// ad break overlaps with a repeating ad break, the ad break scheduled to play first
	// should take precedence while the overlapping ad break is ignored. Since an
	// <adSource> can be a VAST Wrapper to an ad server or ad network, the ads played
	// in a repeated ad break may not be the same at each point.
	RepeatAfter vast.Duration `xml:"repeatAfter,attr,omitempty"`
	// Provides the player with either an inline ad response or a reference to an ad response.
	AdSource *AdSource `xml:",omitempty"`
	// Defines event tracking URLs
	TrackingEvents []Tracking `xml:"TrackingEvents>Tracking,omitempty"`
	// Can be used to express additional information not supported in the VMAP specification.
	Extensions *Extensions `xml:",omitempty"`
}

// AdSource provides the player with either an inline ad response or reference to an ad response.
type AdSource struct {
	// Ad identifier for the ad source
	ID string `xml:"id,attr,omitempty"`
	// Indicates whether a VAST ad pod or multple buffet of ads can be served into an ad break.
	// If not specified, the video player accepts playing multple ads in an ad break. The video
	// player may choose to ignore non-VAST ad pods.
	AllowMultipleAds *bool `xml:"allowMultipleAds,attr,omitempty"`
	// Indicates whether the video player should honor the redirects within an ad response. If
	// not specified, the video player may choose whether it will honor redirects.
	FollowRedirects *bool `xml:"followRedirects,attr,omitempty"`
	// Contains an embedded VAST response.
	VASTAdData   *vast.VAST `xml:"VASTAdData>VAST,omitempty"`
	AdTagURI     *AdTagURI
	CustomAdData *CustomAdData
}

// AdTagURI references an ad response from another system.
type AdTagURI struct {
	// Can be vast, vast1, vast2, vast3 or any string identifying a proprietary template.
	TemplateType string `xml:"templateType,attr,omitempty"`
	URI          string `xml:",chardata"`
}

// CustomAdData is an arbitrary string data that represents a non-VAST ad response
type CustomAdData struct {
	// Can be vast, vast1, vast2, vast3 or any string identifying a proprietary template.
	TemplateType string `xml:"templateType,attr,omitempty"`
	Data         string `xml:",chardata"`
}

// Tracking defines an event tracking URL
type Tracking struct {
	// The name of the event to track for the element. Can be one of breakStart, breakEnd or error.
	Event string `xml:"event,attr"`
	URI   string `xml:",chardata"`
}

// Extensions defines extensions
type Extensions struct {
	Extensions []Extension `xml:"Extension,omitempty"`
}

// Extension represent aribtrary XML provided by the platform to extend the VAST response
type Extension struct {
	// The type of the extension. The type must be globaly unique. A URI is recommended.
	Type string `xml:"type,attr,omitempty"`
	// The XML content of the extension. Extension XML must use it's own namespace.
	Data []byte `xml:",innerxml"`
}
