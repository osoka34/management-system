create table if not exists requirements (
    id UUID primary key,
    title varchar(255) not null,
    status_id bigint not null,
    -- 0 - undefined, 
    -- 1 - create,
    -- 2 - closed,
    -- 3 - in progress,
    -- 4 - done,
    description text not null,
    executor_id UUID not null,
    project_id UUID not null,
    specification_id UUID not null,
    created_at timestamp not null,
    updated_at timestamp not null
)
