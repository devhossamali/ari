// file generated by eventgen

package ari

import (
	"encoding/json"

	"github.com/rotisserie/eris"
)

// EventTypes enumerates the list of event types
type EventTypes struct {
	All                      string
	ApplicationMoveFailed    string
	ApplicationReplaced      string
	BridgeAttendedTransfer   string
	BridgeBlindTransfer      string
	BridgeCreated            string
	BridgeDestroyed          string
	BridgeMerged             string
	BridgeVideoSourceChanged string
	ChannelCallerID          string
	ChannelConnectedLine     string
	ChannelCreated           string
	ChannelDestroyed         string
	ChannelDialplan          string
	ChannelDtmfReceived      string
	ChannelEnteredBridge     string
	ChannelHangupRequest     string
	ChannelHold              string
	ChannelLeftBridge        string
	ChannelStateChange       string
	ChannelTalkingFinished   string
	ChannelTalkingStarted    string
	ChannelUnhold            string
	ChannelUserevent         string
	ChannelVarset            string
	ContactInfo              string
	ContactStatusChange      string
	DeviceStateChanged       string
	Dial                     string
	EndpointStateChange      string
	MissingParams            string
	Peer                     string
	PeerStatusChange         string
	PlaybackContinuing       string
	PlaybackFinished         string
	PlaybackStarted          string
	RecordingFailed          string
	RecordingFinished        string
	RecordingStarted         string
	StasisEnd                string
	StasisStart              string
	TextMessageReceived      string
}

// Events is the instance for grabbing event types
var Events EventTypes

func init() {
	Events.All = "all"
	Events.ApplicationMoveFailed = "ApplicationMoveFailed"
	Events.ApplicationReplaced = "ApplicationReplaced"
	Events.BridgeAttendedTransfer = "BridgeAttendedTransfer"
	Events.BridgeBlindTransfer = "BridgeBlindTransfer"
	Events.BridgeCreated = "BridgeCreated"
	Events.BridgeDestroyed = "BridgeDestroyed"
	Events.BridgeMerged = "BridgeMerged"
	Events.BridgeVideoSourceChanged = "BridgeVideoSourceChanged"
	Events.ChannelCallerID = "ChannelCallerId"
	Events.ChannelConnectedLine = "ChannelConnectedLine"
	Events.ChannelCreated = "ChannelCreated"
	Events.ChannelDestroyed = "ChannelDestroyed"
	Events.ChannelDialplan = "ChannelDialplan"
	Events.ChannelDtmfReceived = "ChannelDtmfReceived"
	Events.ChannelEnteredBridge = "ChannelEnteredBridge"
	Events.ChannelHangupRequest = "ChannelHangupRequest"
	Events.ChannelHold = "ChannelHold"
	Events.ChannelLeftBridge = "ChannelLeftBridge"
	Events.ChannelStateChange = "ChannelStateChange"
	Events.ChannelTalkingFinished = "ChannelTalkingFinished"
	Events.ChannelTalkingStarted = "ChannelTalkingStarted"
	Events.ChannelUnhold = "ChannelUnhold"
	Events.ChannelUserevent = "ChannelUserevent"
	Events.ChannelVarset = "ChannelVarset"
	Events.ContactInfo = "ContactInfo"
	Events.ContactStatusChange = "ContactStatusChange"
	Events.DeviceStateChanged = "DeviceStateChanged"
	Events.Dial = "Dial"
	Events.EndpointStateChange = "EndpointStateChange"
	Events.MissingParams = "MissingParams"
	Events.Peer = "Peer"
	Events.PeerStatusChange = "PeerStatusChange"
	Events.PlaybackContinuing = "PlaybackContinuing"
	Events.PlaybackFinished = "PlaybackFinished"
	Events.PlaybackStarted = "PlaybackStarted"
	Events.RecordingFailed = "RecordingFailed"
	Events.RecordingFinished = "RecordingFinished"
	Events.RecordingStarted = "RecordingStarted"
	Events.StasisEnd = "StasisEnd"
	Events.StasisStart = "StasisStart"
	Events.TextMessageReceived = "TextMessageReceived"

}

// DecodeEvent converts a JSON-encoded event to an ARI event.
func DecodeEvent(data []byte) (Event, error) {
	// Decode the event type
	var typer Message
	err := json.Unmarshal(data, &typer)
	if err != nil {
		return nil, eris.Wrap(err, "failed to decode type")
	}
	if typer.Type == "" {
		return nil, eris.New("no type found")
	}

	switch typer.Type {
	case Events.ApplicationMoveFailed:
		var e ApplicationMoveFailed
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ApplicationReplaced:
		var e ApplicationReplaced
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.BridgeAttendedTransfer:
		var e BridgeAttendedTransfer
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.BridgeBlindTransfer:
		var e BridgeBlindTransfer
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.BridgeCreated:
		var e BridgeCreated
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.BridgeDestroyed:
		var e BridgeDestroyed
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.BridgeMerged:
		var e BridgeMerged
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.BridgeVideoSourceChanged:
		var e BridgeVideoSourceChanged
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelCallerID:
		var e ChannelCallerID
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelConnectedLine:
		var e ChannelConnectedLine
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelCreated:
		var e ChannelCreated
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelDestroyed:
		var e ChannelDestroyed
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelDialplan:
		var e ChannelDialplan
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelDtmfReceived:
		var e ChannelDtmfReceived
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelEnteredBridge:
		var e ChannelEnteredBridge
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelHangupRequest:
		var e ChannelHangupRequest
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelHold:
		var e ChannelHold
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelLeftBridge:
		var e ChannelLeftBridge
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelStateChange:
		var e ChannelStateChange
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelTalkingFinished:
		var e ChannelTalkingFinished
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelTalkingStarted:
		var e ChannelTalkingStarted
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelUnhold:
		var e ChannelUnhold
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelUserevent:
		var e ChannelUserevent
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ChannelVarset:
		var e ChannelVarset
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ContactInfo:
		var e ContactInfo
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.ContactStatusChange:
		var e ContactStatusChange
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.DeviceStateChanged:
		var e DeviceStateChanged
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.Dial:
		var e Dial
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.EndpointStateChange:
		var e EndpointStateChange
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.MissingParams:
		var e MissingParams
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.Peer:
		var e Peer
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.PeerStatusChange:
		var e PeerStatusChange
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.PlaybackContinuing:
		var e PlaybackContinuing
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.PlaybackFinished:
		var e PlaybackFinished
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.PlaybackStarted:
		var e PlaybackStarted
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.RecordingFailed:
		var e RecordingFailed
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.RecordingFinished:
		var e RecordingFinished
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.RecordingStarted:
		var e RecordingStarted
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.StasisEnd:
		var e StasisEnd
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.StasisStart:
		var e StasisStart
		err = json.Unmarshal(data, &e)
		return &e, err
	case Events.TextMessageReceived:
		var e TextMessageReceived
		err = json.Unmarshal(data, &e)
		return &e, err

	}
	return nil, eris.New("unhandled type: " + typer.Type)
}

// ApplicationMoveFailed - "Notification that trying to move a channel to another Stasis application failed."
type ApplicationMoveFailed struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Args        []string    `json:"args"` // Arguments to the application
	Channel     ChannelData `json:"channel"`
	Destination string      `json:"destination"`
}

// ApplicationReplaced - "Notification that another WebSocket has taken over for an application.An application may only be subscribed to by a single WebSocket at a time. If multiple WebSockets attempt to subscribe to the same application, the newer WebSocket wins, and the older one receives this event."
type ApplicationReplaced struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`
}

// BridgeAttendedTransfer - "Notification that an attended transfer has occurred."
type BridgeAttendedTransfer struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	DestinationApplication     string      `json:"destination_application"`      // Application that has been transferred into
	DestinationBridge          string      `json:"destination_bridge"`           // Bridge that survived the merge result
	DestinationLinkFirstLeg    ChannelData `json:"destination_link_first_leg"`   // First leg of a link transfer result
	DestinationLinkSecondLeg   ChannelData `json:"destination_link_second_leg"`  // Second leg of a link transfer result
	DestinationThreewayBridge  BridgeData  `json:"destination_threeway_bridge"`  // Bridge that survived the threeway result
	DestinationThreewayChannel ChannelData `json:"destination_threeway_channel"` // Transferer channel that survived the threeway result
	DestinationType            string      `json:"destination_type"`             // How the transfer was accomplished
	IsExternal                 bool        `json:"is_external"`                  // Whether the transfer was externally initiated or not
	ReplaceChannel             ChannelData `json:"replace_channel,omitempty"`    // The channel that is replacing transferer_first_leg in the swap
	Result                     string      `json:"result"`                       // The result of the transfer attempt
	TransferTarget             ChannelData `json:"transfer_target,omitempty"`    // The channel that is being transferred to
	Transferee                 ChannelData `json:"transferee,omitempty"`         // The channel that is being transferred
	TransfererFirstLeg         ChannelData `json:"transferer_first_leg"`         // First leg of the transferer
	TransfererFirstLegBridge   BridgeData  `json:"transferer_first_leg_bridge"`  // Bridge the transferer first leg is in
	TransfererSecondLeg        ChannelData `json:"transferer_second_leg"`        // Second leg of the transferer
	TransfererSecondLegBridge  BridgeData  `json:"transferer_second_leg_bridge"` // Bridge the transferer second leg is in
}

// BridgeBlindTransfer - "Notification that a blind transfer has occurred."
type BridgeBlindTransfer struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Bridge         BridgeData  `json:"bridge"`                    // The bridge being transferred
	Channel        ChannelData `json:"channel"`                   // The channel performing the blind transfer
	Context        string      `json:"context"`                   // The context transferred to
	Exten          string      `json:"exten"`                     // The extension transferred to
	IsExternal     bool        `json:"is_external"`               // Whether the transfer was externally initiated or not
	ReplaceChannel ChannelData `json:"replace_channel,omitempty"` // The channel that is replacing transferer when the transferee(s) can not be transferred directly
	Result         string      `json:"result"`                    // The result of the transfer attempt
	Transferee     ChannelData `json:"transferee,omitempty"`      // The channel that is being transferred
}

// BridgeCreated - "Notification that a bridge has been created."
type BridgeCreated struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Bridge BridgeData `json:"bridge"`
}

// BridgeDestroyed - "Notification that a bridge has been destroyed."
type BridgeDestroyed struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Bridge BridgeData `json:"bridge"`
}

// BridgeMerged - "Notification that one bridge has merged into another."
type BridgeMerged struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Bridge     BridgeData `json:"bridge"`
	BridgeFrom BridgeData `json:"bridge_from"`
}

// BridgeVideoSourceChanged - "Notification that the source of video in a bridge has changed."
type BridgeVideoSourceChanged struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Bridge           BridgeData `json:"bridge"`
	OldVideoSourceId string     `json:"old_video_source_id,omitempty"`
}

// ChannelCallerID - "Channel changed Caller ID."
type ChannelCallerID struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	CallerPresentation    int         `json:"caller_presentation"`     // The integer representation of the Caller Presentation value.
	CallerPresentationTxt string      `json:"caller_presentation_txt"` // The text representation of the Caller Presentation value.
	Channel               ChannelData `json:"channel"`                 // The channel that changed Caller ID.
}

// ChannelConnectedLine - "Channel changed Connected Line."
type ChannelConnectedLine struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Channel ChannelData `json:"channel"` // The channel whose connected line has changed.
}

// ChannelCreated - "Notification that a channel has been created."
type ChannelCreated struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Channel ChannelData `json:"channel"`
}

// ChannelDestroyed - "Notification that a channel has been destroyed."
type ChannelDestroyed struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Cause    int         `json:"cause"`     // Integer representation of the cause of the hangup
	CauseTxt string      `json:"cause_txt"` // Text representation of the cause of the hangup
	Channel  ChannelData `json:"channel"`
}

// ChannelDialplan - "Channel changed location in the dialplan."
type ChannelDialplan struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Channel         ChannelData `json:"channel"`           // The channel that changed dialplan location.
	DialplanApp     string      `json:"dialplan_app"`      // The application about to be executed.
	DialplanAppData string      `json:"dialplan_app_data"` // The data to be passed to the application.
}

// ChannelDtmfReceived - "DTMF received on a channel.This event is sent when the DTMF ends. There is no notification about the start of DTMF"
type ChannelDtmfReceived struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Channel    ChannelData `json:"channel"`     // The channel on which DTMF was received
	Digit      string      `json:"digit"`       // DTMF digit received (0-9, A-E, # or *)
	DurationMs int         `json:"duration_ms"` // Number of milliseconds DTMF was received
}

// ChannelEnteredBridge - "Notification that a channel has entered a bridge."
type ChannelEnteredBridge struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Bridge  BridgeData  `json:"bridge"`
	Channel ChannelData `json:"channel"`
}

// ChannelHangupRequest - "A hangup was requested on the channel."
type ChannelHangupRequest struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Cause   int         `json:"cause"`   // Integer representation of the cause of the hangup.
	Channel ChannelData `json:"channel"` // The channel on which the hangup was requested.
	Soft    bool        `json:"soft"`    // Whether the hangup request was a soft hangup request.
}

// ChannelHold - "A channel initiated a media hold."
type ChannelHold struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Channel    ChannelData `json:"channel"`              // The channel that initiated the hold event.
	Musicclass string      `json:"musicclass,omitempty"` // The music on hold class that the initiator requested.
}

// ChannelLeftBridge - "Notification that a channel has left a bridge."
type ChannelLeftBridge struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Bridge  BridgeData  `json:"bridge"`
	Channel ChannelData `json:"channel"`
}

// ChannelStateChange - "Notification of a channel's state change."
type ChannelStateChange struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Channel ChannelData `json:"channel"`
}

// ChannelTalkingFinished - "Talking is no longer detected on the channel."
type ChannelTalkingFinished struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Channel  ChannelData `json:"channel"`  // The channel on which talking completed.
	Duration int         `json:"duration"` // The length of time, in milliseconds, that talking was detected on the channel
}

// ChannelTalkingStarted - "Talking was detected on the channel."
type ChannelTalkingStarted struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Channel ChannelData `json:"channel"` // The channel on which talking started.
}

// ChannelUnhold - "A channel initiated a media unhold."
type ChannelUnhold struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Channel ChannelData `json:"channel"` // The channel that initiated the unhold event.
}

// ChannelUserevent - "User-generated event with additional user-defined fields in the object."
type ChannelUserevent struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Bridge    BridgeData   `json:"bridge,omitempty"`   // A bridge that is signaled with the user event.
	Channel   ChannelData  `json:"channel,omitempty"`  // A channel that is signaled with the user event.
	Endpoint  EndpointData `json:"endpoint,omitempty"` // A endpoint that is signaled with the user event.
	Eventname string       `json:"eventname"`          // The name of the user event.
	Userevent interface{}  `json:"userevent"`          // Custom Userevent data
}

// ChannelVarset - "Channel variable changed."
type ChannelVarset struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Channel  ChannelData `json:"channel,omitempty"` // The channel on which the variable was set.If missing, the variable is a global variable.
	Value    string      `json:"value"`             // The new value of the variable.
	Variable string      `json:"variable"`          // The variable that changed.
}

// ContactInfo - "Detailed information about a contact on an endpoint."
type ContactInfo struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Aor           string `json:"aor"`                      // The Address of Record this contact belongs to.
	ContactStatus string `json:"contact_status"`           // The current status of the contact.
	RoundtripUsec string `json:"roundtrip_usec,omitempty"` // Current round trip time, in microseconds, for the contact.
	Uri           string `json:"uri"`                      // The location of the contact.
}

// ContactStatusChange - "The state of a contact on an endpoint has changed."
type ContactStatusChange struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	ContactInfo ContactInfo  `json:"contact_info"`
	Endpoint    EndpointData `json:"endpoint"`
}

// DeviceStateChanged - "Notification that a device state has changed."
type DeviceStateChanged struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	DeviceState DeviceStateData `json:"device_state"` // Device state object
}

// Dial - "Dialing state has changed."
type Dial struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Caller     ChannelData `json:"caller,omitempty"`     // The calling channel.
	Dialstatus string      `json:"dialstatus"`           // Current status of the dialing attempt to the peer.
	Dialstring string      `json:"dialstring,omitempty"` // The dial string for calling the peer channel.
	Forward    string      `json:"forward,omitempty"`    // Forwarding target requested by the original dialed channel.
	Forwarded  ChannelData `json:"forwarded,omitempty"`  // Channel that the caller has been forwarded to.
	Peer       ChannelData `json:"peer"`                 // The dialed channel.
}

// EndpointStateChange - "Endpoint state changed."
type EndpointStateChange struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Endpoint EndpointData `json:"endpoint"`
}

// MissingParams - "Error event sent when required params are missing."
type MissingParams struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Params []string `json:"params"` // A list of the missing parameters
}

// Peer - "Detailed information about a remote peer that communicates with Asterisk."
type Peer struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Address    string `json:"address,omitempty"` // The IP address of the peer.
	Cause      string `json:"cause,omitempty"`   // An optional reason associated with the change in peer_status.
	PeerStatus string `json:"peer_status"`       // The current state of the peer. Note that the values of the status are dependent on the underlying peer technology.
	Port       string `json:"port,omitempty"`    // The port of the peer.
	Time       string `json:"time,omitempty"`    // The last known time the peer was contacted.
}

// PeerStatusChange - "The state of a peer associated with an endpoint has changed."
type PeerStatusChange struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Endpoint EndpointData `json:"endpoint"`
	Peer     Peer         `json:"peer"`
}

// PlaybackContinuing - "Event showing the continuation of a media playback operation from one media URI to the next in the list."
type PlaybackContinuing struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Playback PlaybackData `json:"playback"` // Playback control object
}

// PlaybackFinished - "Event showing the completion of a media playback operation."
type PlaybackFinished struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Playback PlaybackData `json:"playback"` // Playback control object
}

// PlaybackStarted - "Event showing the start of a media playback operation."
type PlaybackStarted struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Playback PlaybackData `json:"playback"` // Playback control object
}

// RecordingFailed - "Event showing failure of a recording operation."
type RecordingFailed struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Recording LiveRecordingData `json:"recording"` // Recording control object
}

// RecordingFinished - "Event showing the completion of a recording operation."
type RecordingFinished struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Recording LiveRecordingData `json:"recording"` // Recording control object
}

// RecordingStarted - "Event showing the start of a recording operation."
type RecordingStarted struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Recording LiveRecordingData `json:"recording"` // Recording control object
}

// StasisEnd - "Notification that a channel has left a Stasis application."
type StasisEnd struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Channel ChannelData `json:"channel"`
}

// StasisStart - "Notification that a channel has entered a Stasis application."
type StasisStart struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Args           []string    `json:"args"` // Arguments to the application
	Channel        ChannelData `json:"channel"`
	ReplaceChannel ChannelData `json:"replace_channel,omitempty"`
}

// TextMessageReceived - "A text message was received from an endpoint."
type TextMessageReceived struct {
	EventData `json:",inline"`

	// Header describes any transport-related metadata
	Header Header `json:"-"`

	Endpoint EndpointData    `json:"endpoint,omitempty"`
	Message  TextMessageData `json:"message"`
}
