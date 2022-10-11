// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnknownError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MessageErrorReason_UNKNOWN_ERROR.String() && e.Code == 500
}

func ErrorUnknownError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MessageErrorReason_UNKNOWN_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsAccessUserMedalFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MessageErrorReason_ACCESS_USER_MEDAL_FAILED.String() && e.Code == 500
}

func ErrorAccessUserMedalFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MessageErrorReason_ACCESS_USER_MEDAL_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetMailboxLastTimeFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MessageErrorReason_GET_MAILBOX_LAST_TIME_FAILED.String() && e.Code == 500
}

func ErrorGetMailboxLastTimeFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MessageErrorReason_GET_MAILBOX_LAST_TIME_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsGetMessageNotificationFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MessageErrorReason_GET_MESSAGE_NOTIFICATION_FAILED.String() && e.Code == 500
}

func ErrorGetMessageNotificationFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MessageErrorReason_GET_MESSAGE_NOTIFICATION_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsSetMailboxLastTimeFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MessageErrorReason_SET_MAILBOX_LAST_TIME_FAILED.String() && e.Code == 500
}

func ErrorSetMailboxLastTimeFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MessageErrorReason_SET_MAILBOX_LAST_TIME_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsRemoveMailboxCommentFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MessageErrorReason_REMOVE_MAILBOX_COMMENT_FAILED.String() && e.Code == 500
}

func ErrorRemoveMailboxCommentFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MessageErrorReason_REMOVE_MAILBOX_COMMENT_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsRemoveMailboxSystemNotificationFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == MessageErrorReason_REMOVE_MAILBOX_SYSTEM_NOTIFICATION_FAILED.String() && e.Code == 500
}

func ErrorRemoveMailboxSystemNotificationFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, MessageErrorReason_REMOVE_MAILBOX_SYSTEM_NOTIFICATION_FAILED.String(), fmt.Sprintf(format, args...))
}
