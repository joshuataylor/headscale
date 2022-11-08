package headscale

import "time"

type Device struct {
	Addresses        []string
	ID               string
	User             string
	Name             string
	Hostname         string
	ClientVersion    string
	UpdateAvailable  bool
	OS               string
	Created          *time.Time
	LastSeen         *time.Time
	KeyExpiry        bool
	Expires          *time.Time
	Authorized       bool
	IsExternal       bool
	MachineKey       string
	NodeKey          string
	BlocksIncoming   bool
	EnabledRoutes    []string
	AdvertisedRoutes []string
	ClientConnectivity
}

type ClientConnectivity struct {
	Endpoints []string
	Derp      string
	Mapping   bool
	Latency   map[string]Latency
	ClientSupports
}

type Latency struct {
	LatencyMs float64
	preferred bool
}

type ClientSupports struct {
	HairPinning bool
	IPv6        bool
	PCP         bool
	PMP         bool
	UDP         bool
	UPNP        bool
}

func (h *Headscale) ListDevices() ([]Device, error) {
	devices := []Device{}
	if err := h.db.Preload("AuthKey").Preload("AuthKey.Namespace").Preload("Namespace").Find(&devices).Error; err != nil {
		return nil, err
	}

	return devices, nil
}
