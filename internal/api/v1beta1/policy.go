package v1beta1

import (
	"context"
	"errors"

	"github.com/raystack/frontier/core/role"

	"github.com/raystack/frontier/core/audit"
	"github.com/raystack/frontier/internal/bootstrap/schema"

	"google.golang.org/protobuf/types/known/structpb"

	"github.com/raystack/frontier/pkg/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/raystack/frontier/core/namespace"
	"github.com/raystack/frontier/core/policy"
	frontierv1beta1 "github.com/raystack/frontier/proto/v1beta1"
)

type PolicyService interface {
	Get(ctx context.Context, id string) (policy.Policy, error)
	List(ctx context.Context, f policy.Filter) ([]policy.Policy, error)
	Create(ctx context.Context, pol policy.Policy) (policy.Policy, error)
	Delete(ctx context.Context, id string) error
	ListForUser(ctx context.Context, userID, objectNamespace, objectID string) ([]role.Role, error)
}

var grpcPolicyNotFoundErr = status.Errorf(codes.NotFound, "policy doesn't exist")

func (h Handler) ListPolicies(ctx context.Context, request *frontierv1beta1.ListPoliciesRequest) (*frontierv1beta1.ListPoliciesResponse, error) {
	logger := grpczap.Extract(ctx)
	var policies []*frontierv1beta1.Policy

	policyList, err := h.policyService.List(ctx, policy.Filter{
		OrgID:       request.GetOrgId(),
		PrincipalID: request.GetUserId(),
		ProjectID:   request.GetProjectId(),
		RoleID:      request.GetRoleId(),
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, grpcInternalServerError
	}

	for _, p := range policyList {
		policyPB, err := transformPolicyToPB(p)
		if err != nil {
			logger.Error(err.Error())
			return nil, grpcInternalServerError
		}

		policies = append(policies, policyPB)
	}

	return &frontierv1beta1.ListPoliciesResponse{Policies: policies}, nil
}

func (h Handler) CreatePolicy(ctx context.Context, request *frontierv1beta1.CreatePolicyRequest) (*frontierv1beta1.CreatePolicyResponse, error) {
	logger := grpczap.Extract(ctx)

	var metaDataMap metadata.Metadata
	var err error
	if request.GetBody().GetMetadata() != nil {
		metaDataMap = metadata.Build(request.GetBody().GetMetadata().AsMap())
	}

	resourceType, resourceID, err := schema.SplitNamespaceAndResourceID(request.GetBody().GetResource())
	if err != nil {
		return nil, ErrNamespaceSplitNotation
	}
	principalType, principalID, err := schema.SplitNamespaceAndResourceID(request.GetBody().GetPrincipal())
	if err != nil {
		return nil, ErrNamespaceSplitNotation
	}

	newPolicy, err := h.policyService.Create(ctx, policy.Policy{
		RoleID:        request.GetBody().GetRoleId(),
		ResourceID:    resourceID,
		ResourceType:  resourceType,
		PrincipalID:   principalID,
		PrincipalType: principalType,
		Metadata:      metaDataMap,
	})
	if err != nil {
		logger.Error(err.Error())
		switch {
		case errors.Is(err, role.ErrInvalidID):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, policy.ErrInvalidDetail):
			return nil, grpcBadBodyError
		default:
			return nil, grpcInternalServerError
		}
	}

	policyPB, err := transformPolicyToPB(newPolicy)
	if err != nil {
		logger.Error(err.Error())
		return nil, grpcInternalServerError
	}

	audit.GetAuditor(ctx, schema.PlatformOrgID.String()).
		LogWithAttrs(audit.PolicyCreatedEvent, audit.Target{
			ID:   newPolicy.ResourceID,
			Type: newPolicy.ResourceType,
		}, map[string]string{
			"role_id":        newPolicy.RoleID,
			"principal_id":   newPolicy.PrincipalID,
			"principal_type": newPolicy.PrincipalType,
		})
	return &frontierv1beta1.CreatePolicyResponse{Policy: policyPB}, nil
}

func (h Handler) GetPolicy(ctx context.Context, request *frontierv1beta1.GetPolicyRequest) (*frontierv1beta1.GetPolicyResponse, error) {
	logger := grpczap.Extract(ctx)

	fetchedPolicy, err := h.policyService.Get(ctx, request.GetId())
	if err != nil {
		logger.Error(err.Error())
		switch {
		case errors.Is(err, policy.ErrNotExist),
			errors.Is(err, policy.ErrInvalidUUID),
			errors.Is(err, policy.ErrInvalidID):
			return nil, grpcPolicyNotFoundErr
		default:
			return nil, grpcInternalServerError
		}
	}

	policyPB, err := transformPolicyToPB(fetchedPolicy)
	if err != nil {
		logger.Error(err.Error())
		return nil, grpcInternalServerError
	}

	return &frontierv1beta1.GetPolicyResponse{Policy: policyPB}, nil
}

func (h Handler) UpdatePolicy(ctx context.Context, request *frontierv1beta1.UpdatePolicyRequest) (*frontierv1beta1.UpdatePolicyResponse, error) {
	// not implemented
	return &frontierv1beta1.UpdatePolicyResponse{}, status.Errorf(codes.Unimplemented, "unsupported at the moment")
}

func (h Handler) DeletePolicy(ctx context.Context, request *frontierv1beta1.DeletePolicyRequest) (*frontierv1beta1.DeletePolicyResponse, error) {
	logger := grpczap.Extract(ctx)
	err := h.policyService.Delete(ctx, request.GetId())
	if err != nil {
		logger.Error(err.Error())
		switch {
		case errors.Is(err, policy.ErrNotExist),
			errors.Is(err, policy.ErrInvalidID),
			errors.Is(err, policy.ErrInvalidUUID):
			return nil, grpcPolicyNotFoundErr
		case errors.Is(err, policy.ErrInvalidDetail),
			errors.Is(err, namespace.ErrNotExist):
			return nil, grpcBadBodyError
		case errors.Is(err, policy.ErrConflict):
			return nil, grpcConflictError
		default:
			return nil, grpcInternalServerError
		}
	}

	audit.GetAuditor(ctx, schema.PlatformOrgID.String()).Log(audit.PolicyDeletedEvent, audit.Target{
		ID:   request.GetId(),
		Type: "app/policy",
	})
	return &frontierv1beta1.DeletePolicyResponse{}, nil
}

func transformPolicyToPB(policy policy.Policy) (*frontierv1beta1.Policy, error) {
	var metadata *structpb.Struct
	var err error
	if len(policy.Metadata) > 0 {
		metadata, err = structpb.NewStruct(policy.Metadata)
		if err != nil {
			return nil, err
		}
	}

	pbPol := &frontierv1beta1.Policy{
		Id:        policy.ID,
		RoleId:    policy.RoleID,
		Resource:  schema.JoinNamespaceAndResourceID(policy.ResourceType, policy.ResourceID),
		Principal: schema.JoinNamespaceAndResourceID(policy.PrincipalType, policy.PrincipalID),
		Metadata:  metadata,
	}
	if !policy.CreatedAt.IsZero() {
		pbPol.CreatedAt = timestamppb.New(policy.CreatedAt)
	}
	if !policy.UpdatedAt.IsZero() {
		pbPol.UpdatedAt = timestamppb.New(policy.UpdatedAt)
	}
	return pbPol, nil
}
