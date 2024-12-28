package domain

type SubscriptionStatus string

const (
	SubscriptionStatusActive            SubscriptionStatus = "active"
	SubscriptionStatusCanceled          SubscriptionStatus = "canceled"
	SubscriptionStatusIncomplete        SubscriptionStatus = "incomplete"
	SubscriptionStatusIncompleteExpired SubscriptionStatus = "incomplete_expired"
	SubscriptionStatusPastDue           SubscriptionStatus = "past_due"
	SubscriptionStatusPaused            SubscriptionStatus = "paused"
	SubscriptionStatusTrialing          SubscriptionStatus = "trialing"
	SubscriptionStatusUnpaid            SubscriptionStatus = "unpaid"
)

// Subscription represents a client subscription.
type Subscription struct {
	ID        string
	ProductID string
	Status    SubscriptionStatus
}
