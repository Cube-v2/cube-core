package service

import (
	"context"
	"github.com/the-zion/matrix-core/api/bff/interface/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *BffService) GetMessageNotification(ctx context.Context, _ *emptypb.Empty) (*v1.GetMessageNotificationReply, error) {
	notification, err := s.mc.GetMessageNotification(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetMessageNotificationReply{
		Timeline:   notification.Timeline,
		Comment:    notification.Comment,
		SubComment: notification.SubComment,
		System:     notification.SystemNotification,
	}, nil
}

func (s *BffService) GetMailBoxLastTime(ctx context.Context, _ *emptypb.Empty) (*v1.GetMailBoxLastTimeReply, error) {
	mailbox, err := s.mc.GetMailBoxLastTime(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetMailBoxLastTimeReply{
		Time: mailbox.Time,
	}, nil
}

func (s *BffService) GetMessageSystemNotification(ctx context.Context, req *v1.GetMessageSystemNotificationReq) (*v1.GetMessageSystemNotificationReply, error) {
	reply := &v1.GetMessageSystemNotificationReply{List: make([]*v1.GetMessageSystemNotificationReply_List, 0)}
	notificationList, err := s.mc.GetMessageSystemNotification(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	for _, item := range notificationList {
		reply.List = append(reply.List, &v1.GetMessageSystemNotificationReply_List{
			Id:               item.Id,
			ContentId:        item.ContentId,
			CreatedAt:        item.CreatedAt,
			NotificationType: item.NotificationType,
			Title:            item.Title,
			Uuid:             item.Uuid,
			Label:            item.Label,
			Result:           item.Result,
			Section:          item.Section,
			Text:             item.Text,
		})
	}
	return reply, nil
}

func (s *BffService) SetMailBoxLastTime(ctx context.Context, req *v1.SetMailBoxLastTimeReq) (*emptypb.Empty, error) {
	err := s.mc.SetMailBoxLastTime(ctx, req.Time)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *BffService) RemoveMailBoxCommentCount(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.mc.RemoveMailBoxCommentCount(ctx)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *BffService) RemoveMailBoxSubCommentCount(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.mc.RemoveMailBoxSubCommentCount(ctx)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
