package service

import (
	"context"
	v1 "github.com/the-zion/matrix-core/api/creation/service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *CreationService) GetLeaderBoard(ctx context.Context, _ *emptypb.Empty) (*v1.GetLeaderBoardReply, error) {
	reply := &v1.GetLeaderBoardReply{Board: make([]*v1.GetLeaderBoardReply_Board, 0)}
	boardList, err := s.cc.GetLeaderBoard(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range boardList {
		reply.Board = append(reply.Board, &v1.GetLeaderBoardReply_Board{
			Id:   item.Id,
			Uuid: item.Uuid,
			Mode: item.Mode,
		})
	}
	return reply, nil
}

func (s *CreationService) GetCollections(ctx context.Context, req *v1.GetCollectionsReq) (*v1.GetCollectionsReply, error) {
	reply := &v1.GetCollectionsReply{Collections: make([]*v1.GetCollectionsReply_Collections, 0)}
	collections, err := s.cc.GetCollections(ctx, req.Uuid, req.Page)
	if err != nil {
		return nil, err
	}
	for _, item := range collections {
		reply.Collections = append(reply.Collections, &v1.GetCollectionsReply_Collections{
			Name:      item.Name,
			Introduce: item.Name,
		})
	}
	return reply, nil
}

func (s *CreationService) GetCollectionsByVisitor(ctx context.Context, req *v1.GetCollectionsReq) (*v1.GetCollectionsReply, error) {
	reply := &v1.GetCollectionsReply{Collections: make([]*v1.GetCollectionsReply_Collections, 0)}
	collections, err := s.cc.GetCollectionsByVisitor(ctx, req.Uuid, req.Page)
	if err != nil {
		return nil, err
	}
	for _, item := range collections {
		reply.Collections = append(reply.Collections, &v1.GetCollectionsReply_Collections{
			Name:      item.Name,
			Introduce: item.Name,
		})
	}
	return reply, nil
}

func (s *CreationService) GetCollectionsCount(ctx context.Context, req *v1.GetCollectionsCountReq) (*v1.GetCollectionsCountReply, error) {
	count, err := s.cc.GetCollectionsCount(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	return &v1.GetCollectionsCountReply{
		Count: count,
	}, nil
}

func (s *CreationService) GetCollectionsVisitorCount(ctx context.Context, req *v1.GetCollectionsCountReq) (*v1.GetCollectionsCountReply, error) {
	count, err := s.cc.GetCollectionsVisitorCount(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	return &v1.GetCollectionsCountReply{
		Count: count,
	}, nil
}

func (s *CreationService) CreateCollections(ctx context.Context, req *v1.CreateCollectionsReq) (*emptypb.Empty, error) {
	err := s.cc.CreateCollections(ctx, req.Uuid, req.Name, req.Introduce, req.Auth)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
