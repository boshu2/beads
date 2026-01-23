// Package v1alpha1 contains API Schema definitions for hermes.olympus v1alpha1
// +kubebuilder:object:generate=true
// +groupName=hermes.olympus
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ============================================================================
// BeadStore
// Git-synced issue tracker database
// ============================================================================

// BeadStoreSpec defines the desired state of BeadStore
type BeadStoreSpec struct {
	// Repository is the git repository URL for beads sync
	// +kubebuilder:validation:Required
	Repository string `json:"repository"`

	// Branch is the git branch to sync
	// +kubebuilder:default=main
	// +optional
	Branch string `json:"branch,omitempty"`

	// Prefix is the beads issue prefix for this store
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`^[a-z]{2,4}$`
	Prefix string `json:"prefix"`

	// SyncInterval is how often to sync with git
	// +kubebuilder:default="5m"
	// +optional
	SyncInterval string `json:"syncInterval,omitempty"`

	// SecretRef references git credentials
	// +optional
	SecretRef *SecretKeySelector `json:"secretRef,omitempty"`
}

// SecretKeySelector references a secret key.
// NOTE: This type is intentionally duplicated across mt-olympus.io API groups
// (athena, hephaestus, hermes) because kubebuilder requires types to be
// local to generate deepcopy methods. See api/common/v1alpha1/go/types.go
// for the canonical definition.
type SecretKeySelector struct {
	// Name is the secret name
	Name string `json:"name"`

	// Key is the key within the secret
	Key string `json:"key"`
}

// BeadStorePhase represents the current phase of the BeadStore
// +kubebuilder:validation:Enum=Initializing;Syncing;Ready;Error
type BeadStorePhase string

const (
	BeadStorePhaseInitializing BeadStorePhase = "Initializing"
	BeadStorePhaseSyncing      BeadStorePhase = "Syncing"
	BeadStorePhaseReady        BeadStorePhase = "Ready"
	BeadStorePhaseError        BeadStorePhase = "Error"
)

// BeadStoreStatus defines the observed state of BeadStore
type BeadStoreStatus struct {
	// Phase is the current phase
	// +optional
	Phase BeadStorePhase `json:"phase,omitempty"`

	// LastSyncAt is the last successful sync time
	// +optional
	LastSyncAt *metav1.Time `json:"lastSyncAt,omitempty"`

	// LastCommit is the git commit SHA of last sync
	// +optional
	LastCommit string `json:"lastCommit,omitempty"`

	// IssueCount is the total number of issues in store
	// +optional
	IssueCount int32 `json:"issueCount,omitempty"`

	// OpenCount is the number of open issues
	// +optional
	OpenCount int32 `json:"openCount,omitempty"`

	// Message is a human-readable status message
	// +optional
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Prefix",type=string,JSONPath=`.spec.prefix`
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Issues",type=integer,JSONPath=`.status.issueCount`
// +kubebuilder:printcolumn:name="Open",type=integer,JSONPath=`.status.openCount`
// +kubebuilder:printcolumn:name="Last Sync",type=date,JSONPath=`.status.lastSyncAt`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// BeadStore is the Schema for the beadstores API
type BeadStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BeadStoreSpec   `json:"spec,omitempty"`
	Status BeadStoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BeadStoreList contains a list of BeadStore
type BeadStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BeadStore `json:"items"`
}

// ============================================================================
// Issue
// Work item in a BeadStore (read-only view of beads issue)
// ============================================================================

// IssueSpec defines the desired state of Issue
type IssueSpec struct {
	// StoreRef references the BeadStore this issue belongs to
	// +kubebuilder:validation:Required
	StoreRef string `json:"storeRef"`

	// IssueID is the full beads issue ID (e.g., "at-1234")
	// +kubebuilder:validation:Required
	IssueID string `json:"issueId"`
}

// IssueType represents the type of issue
// +kubebuilder:validation:Enum=task;bug;epic;feature;research
type IssueType string

const (
	IssueTypeTask     IssueType = "task"
	IssueTypeBug      IssueType = "bug"
	IssueTypeEpic     IssueType = "epic"
	IssueTypeFeature  IssueType = "feature"
	IssueTypeResearch IssueType = "research"
)

// IssueStatusType represents the status of an issue
// +kubebuilder:validation:Enum=open;in_progress;blocked;closed
type IssueStatusType string

const (
	IssueStatusOpen       IssueStatusType = "open"
	IssueStatusInProgress IssueStatusType = "in_progress"
	IssueStatusBlocked    IssueStatusType = "blocked"
	IssueStatusClosed     IssueStatusType = "closed"
)

// IssuePriority represents issue priority
// +kubebuilder:validation:Enum=P0;P1;P2;P3
type IssuePriority string

const (
	IssuePriorityP0 IssuePriority = "P0"
	IssuePriorityP1 IssuePriority = "P1"
	IssuePriorityP2 IssuePriority = "P2"
	IssuePriorityP3 IssuePriority = "P3"
)

// IssueStatus defines the observed state of Issue
type IssueStatus struct {
	// Title is the issue title
	// +optional
	Title string `json:"title,omitempty"`

	// Type is the issue type
	// +optional
	Type IssueType `json:"type,omitempty"`

	// State is the current issue state (open, in_progress, blocked, closed)
	// +optional
	State IssueStatusType `json:"state,omitempty"`

	// Priority is the issue priority
	// +optional
	Priority IssuePriority `json:"priority,omitempty"`

	// Owner is who owns this issue
	// +optional
	Owner string `json:"owner,omitempty"`

	// ParentID is the parent issue ID (for epics)
	// +optional
	ParentID string `json:"parentId,omitempty"`

	// Blockers lists issue IDs that block this issue
	// +optional
	Blockers []string `json:"blockers,omitempty"`

	// CreatedAt is when the issue was created
	// +optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`

	// UpdatedAt is when the issue was last updated
	// +optional
	UpdatedAt *metav1.Time `json:"updatedAt,omitempty"`

	// ClosedAt is when the issue was closed
	// +optional
	ClosedAt *metav1.Time `json:"closedAt,omitempty"`

	// Description is the issue description
	// +optional
	Description string `json:"description,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Title",type=string,JSONPath=`.status.title`,priority=0
// +kubebuilder:printcolumn:name="Type",type=string,JSONPath=`.status.type`
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Priority",type=string,JSONPath=`.status.priority`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// Issue is the Schema for the issues API
type Issue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IssueSpec   `json:"spec,omitempty"`
	Status IssueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IssueList contains a list of Issue
type IssueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Issue `json:"items"`
}

// ============================================================================
// Epic
// Work grouping in a BeadStore (read-only view of beads epic)
// ============================================================================

// EpicSpec defines the desired state of Epic
type EpicSpec struct {
	// StoreRef references the BeadStore this epic belongs to
	// +kubebuilder:validation:Required
	StoreRef string `json:"storeRef"`

	// EpicID is the full beads issue ID for this epic
	// +kubebuilder:validation:Required
	EpicID string `json:"epicId"`
}

// EpicPhase represents the current phase of the Epic
// +kubebuilder:validation:Enum=Planning;Active;Blocked;Completed
type EpicPhase string

const (
	EpicPhasePlanning  EpicPhase = "Planning"
	EpicPhaseActive    EpicPhase = "Active"
	EpicPhaseBlocked   EpicPhase = "Blocked"
	EpicPhaseCompleted EpicPhase = "Completed"
)

// EpicStatus defines the observed state of Epic
type EpicStatus struct {
	// Title is the epic title
	// +optional
	Title string `json:"title,omitempty"`

	// Phase is the computed epic phase
	// +optional
	Phase EpicPhase `json:"phase,omitempty"`

	// TotalChildren is the total number of child issues
	// +optional
	TotalChildren int32 `json:"totalChildren,omitempty"`

	// CompletedChildren is the count of closed child issues
	// +optional
	CompletedChildren int32 `json:"completedChildren,omitempty"`

	// InProgressChildren is the count of in_progress child issues
	// +optional
	InProgressChildren int32 `json:"inProgressChildren,omitempty"`

	// BlockedChildren is the count of blocked child issues
	// +optional
	BlockedChildren int32 `json:"blockedChildren,omitempty"`

	// ReadyChildren is the count of ready (unblocked) child issues
	// +optional
	ReadyChildren int32 `json:"readyChildren,omitempty"`

	// ProgressPercent is the completion percentage
	// +optional
	ProgressPercent int32 `json:"progressPercent,omitempty"`

	// Owner is who owns this epic
	// +optional
	Owner string `json:"owner,omitempty"`

	// CreatedAt is when the epic was created
	// +optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`

	// UpdatedAt is when the epic was last updated
	// +optional
	UpdatedAt *metav1.Time `json:"updatedAt,omitempty"`

	// Description is the epic description
	// +optional
	Description string `json:"description,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Title",type=string,JSONPath=`.status.title`,priority=0
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Progress",type=string,JSONPath=`.status.progressPercent`
// +kubebuilder:printcolumn:name="Ready",type=integer,JSONPath=`.status.readyChildren`
// +kubebuilder:printcolumn:name="Total",type=integer,JSONPath=`.status.totalChildren`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// Epic is the Schema for the epics API
type Epic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EpicSpec   `json:"spec,omitempty"`
	Status EpicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EpicList contains a list of Epic
type EpicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Epic `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BeadStore{}, &BeadStoreList{})
	SchemeBuilder.Register(&Issue{}, &IssueList{})
	SchemeBuilder.Register(&Epic{}, &EpicList{})
}
