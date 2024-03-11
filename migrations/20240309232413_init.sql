-- +goose Up
-- +goose StatementBegin
create table pins
(
    user_id bigint                not null,
    msg_id  bigint                not null,
    is_active boolean default false not null,
    constraint pins_pk
        primary key (user_id, msg_id)
);

create index pins_is_active_index
    on pins (is_active);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
