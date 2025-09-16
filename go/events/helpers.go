// Copyright 2025 The StrataSTOR Authors and Contributors
// SPDX-License-Identifier: Apache-2.0

package events

import (
	"encoding/json"
	"time"

	"github.com/stratastor/toggle-rodent-proto/proto"
)

// EventBuilder provides a fluent interface for building events
type EventBuilder struct {
	event *proto.Event
}

// NewEvent creates a new event builder
func NewEvent(eventID, eventType string, level proto.EventLevel, category proto.EventCategory, source string) *EventBuilder {
	return &EventBuilder{
		event: &proto.Event{
			EventId:   eventID,
			EventType: eventType,
			Level:     level,
			Category:  category,
			Source:    source,
			Timestamp: time.Now().UnixMilli(),
			Metadata:  make(map[string]string),
		},
	}
}

// WithPayload sets the event payload (JSON-marshaled)
func (eb *EventBuilder) WithPayload(payload interface{}) *EventBuilder {
	if payload != nil {
		if payloadBytes, err := json.Marshal(payload); err == nil {
			eb.event.Payload = payloadBytes
		}
	}
	return eb
}

// WithMetadata adds metadata key-value pairs
func (eb *EventBuilder) WithMetadata(metadata map[string]string) *EventBuilder {
	if metadata != nil {
		for k, v := range metadata {
			eb.event.Metadata[k] = v
		}
	}
	return eb
}

// WithMeta adds a single metadata key-value pair
func (eb *EventBuilder) WithMeta(key, value string) *EventBuilder {
	eb.event.Metadata[key] = value
	return eb
}

// WithTimestamp sets a custom timestamp
func (eb *EventBuilder) WithTimestamp(timestamp time.Time) *EventBuilder {
	eb.event.Timestamp = timestamp.UnixMilli()
	return eb
}

// Build returns the constructed event
func (eb *EventBuilder) Build() *proto.Event {
	return eb.event
}

// Helper functions for creating common metadata maps

// SystemMeta creates metadata for system events
func SystemMeta(component, action string) map[string]string {
	return map[string]string{
		MetaComponent: component,
		MetaAction:    action,
		MetaModule:    "system",
	}
}

// StorageMeta creates metadata for storage events
func StorageMeta(component, action, resourceType, resourceName string) map[string]string {
	return map[string]string{
		MetaComponent:    component,
		MetaAction:       action,
		MetaModule:       "storage",
		MetaResourceType: resourceType,
		MetaResourceName: resourceName,
	}
}

// NetworkMeta creates metadata for network events
func NetworkMeta(component, action, iface string) map[string]string {
	meta := map[string]string{
		MetaComponent: component,
		MetaAction:    action,
		MetaModule:    "network",
	}
	if iface != "" {
		meta[MetaInterface] = iface
	}
	return meta
}

// SecurityMeta creates metadata for security events
func SecurityMeta(component, action, user string) map[string]string {
	meta := map[string]string{
		MetaComponent: component,
		MetaAction:    action,
		MetaModule:    "security",
	}
	if user != "" {
		meta[MetaUser] = user
	}
	return meta
}

// ServiceMeta creates metadata for service events
func ServiceMeta(service, action string) map[string]string {
	return map[string]string{
		MetaComponent: "service-manager",
		MetaAction:    action,
		MetaModule:    "service",
		MetaService:   service,
	}
}

// IdentityMeta creates metadata for identity events
func IdentityMeta(component, action, domain, user string) map[string]string {
	meta := map[string]string{
		MetaComponent: component,
		MetaAction:    action,
		MetaModule:    "identity",
	}
	if domain != "" {
		meta[MetaDomain] = domain
	}
	if user != "" {
		meta[MetaUser] = user
	}
	return meta
}

// AccessMeta creates metadata for access control events
func AccessMeta(component, action, resourceType, resourceName string) map[string]string {
	return map[string]string{
		MetaComponent:    component,
		MetaAction:       action,
		MetaModule:       "access",
		MetaResourceType: resourceType,
		MetaResourceName: resourceName,
	}
}

// SharingMeta creates metadata for file sharing events
func SharingMeta(component, action, shareType, shareName string) map[string]string {
	meta := map[string]string{
		MetaComponent:    component,
		MetaAction:       action,
		MetaModule:       "sharing",
		MetaResourceType: shareType,
	}
	if shareName != "" {
		meta[MetaShareName] = shareName
	}
	return meta
}

// Category helper functions

// GetCategoryString returns the string representation of an EventCategory
func GetCategoryString(category proto.EventCategory) string {
	switch category {
	case proto.EventCategory_EVENT_CATEGORY_SYSTEM:
		return "system"
	case proto.EventCategory_EVENT_CATEGORY_STORAGE:
		return "storage"
	case proto.EventCategory_EVENT_CATEGORY_NETWORK:
		return "network"
	case proto.EventCategory_EVENT_CATEGORY_SECURITY:
		return "security"
	case proto.EventCategory_EVENT_CATEGORY_SERVICE:
		return "service"
	case proto.EventCategory_EVENT_CATEGORY_IDENTITY:
		return "identity"
	case proto.EventCategory_EVENT_CATEGORY_ACCESS:
		return "access"
	case proto.EventCategory_EVENT_CATEGORY_SHARING:
		return "sharing"
	default:
		return "unspecified"
	}
}

// GetLevelString returns the string representation of an EventLevel
func GetLevelString(level proto.EventLevel) string {
	switch level {
	case proto.EventLevel_EVENT_LEVEL_INFO:
		return "info"
	case proto.EventLevel_EVENT_LEVEL_WARN:
		return "warn"
	case proto.EventLevel_EVENT_LEVEL_ERROR:
		return "error"
	case proto.EventLevel_EVENT_LEVEL_CRITICAL:
		return "critical"
	default:
		return "unspecified"
	}
}

// ParseEventCategory parses a string to EventCategory
func ParseEventCategory(category string) proto.EventCategory {
	switch category {
	case "system":
		return proto.EventCategory_EVENT_CATEGORY_SYSTEM
	case "storage":
		return proto.EventCategory_EVENT_CATEGORY_STORAGE
	case "network":
		return proto.EventCategory_EVENT_CATEGORY_NETWORK
	case "security":
		return proto.EventCategory_EVENT_CATEGORY_SECURITY
	case "service":
		return proto.EventCategory_EVENT_CATEGORY_SERVICE
	case "identity":
		return proto.EventCategory_EVENT_CATEGORY_IDENTITY
	case "access":
		return proto.EventCategory_EVENT_CATEGORY_ACCESS
	case "sharing":
		return proto.EventCategory_EVENT_CATEGORY_SHARING
	default:
		return proto.EventCategory_EVENT_CATEGORY_UNSPECIFIED
	}
}

// ParseEventLevel parses a string to EventLevel
func ParseEventLevel(level string) proto.EventLevel {
	switch level {
	case "info":
		return proto.EventLevel_EVENT_LEVEL_INFO
	case "warn", "warning":
		return proto.EventLevel_EVENT_LEVEL_WARN
	case "error":
		return proto.EventLevel_EVENT_LEVEL_ERROR
	case "critical", "crit":
		return proto.EventLevel_EVENT_LEVEL_CRITICAL
	default:
		return proto.EventLevel_EVENT_LEVEL_UNSPECIFIED
	}
}