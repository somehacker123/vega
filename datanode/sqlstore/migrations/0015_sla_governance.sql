-- +goose Up

ALTER TYPE proposal_error ADD VALUE 'PROPOSAL_ERROR_MISSING_SLA_PARAMS';
ALTER TYPE proposal_error ADD VALUE 'PROPOSAL_ERROR_INVALID_SLA_PARAMS';