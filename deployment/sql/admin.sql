-- 状态表
CREATE TABLE status (
    id SMALLINT NOT NULL,
    name VARCHAR(20) NOT NULL,
    -- 约束
    CONSTRAINT pk_status PRIMARY KEY (id),
    CONSTRAINT uk_status_name UNIQUE (name)
);

COMMENT ON TABLE status IS '状态表';

COMMENT ON COLUMN status.id IS '状态ID';

COMMENT ON COLUMN status.name IS '状态名称';

-- 状态
INSERT INTO
    status (id, name)
VALUES (1, 'active'),
    (2, 'disabled'),
    (3, 'locked');

-- 账号表

CREATE TABLE accounts (
    id UUID NOT NULL,
    username VARCHAR(50) NOT NULL,
    nickname VARCHAR(50),
    roles SMALLINT[] NOT NULL,
    avatar TEXT NOT NULL,
    email VARCHAR(100),
    phone VARCHAR(20),
    pwd TEXT NOT NULL,
    slogan VARCHAR(255),
    status SMALLINT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    -- 约束
    CONSTRAINT pk_accounts PRIMARY KEY (id),
    CONSTRAINT uk_accounts_username UNIQUE (username),
    -- 外键约束：状态必须存在于 status 表
    CONSTRAINT fk_accounts_status FOREIGN KEY (status) REFERENCES status (id)
);
-- 账号状态索引
CREATE INDEX idx_accounts_status ON accounts (status);

COMMENT ON TABLE accounts IS '账号表';

COMMENT ON COLUMN accounts.id IS '账号ID(主键)';

COMMENT ON COLUMN accounts.username IS '用户名(唯一)';

COMMENT ON COLUMN accounts.nickname IS '昵称(可空)';

COMMENT ON COLUMN accounts.roles IS '角色ID数组(非空)';

COMMENT ON COLUMN accounts.avatar IS '头像URL(非空)';

COMMENT ON COLUMN accounts.email IS '邮箱(可空)';

COMMENT ON COLUMN accounts.phone IS '手机号(可空)';

COMMENT ON COLUMN accounts.pwd IS '密码哈希(非空)';

COMMENT ON COLUMN accounts.slogan IS '个性签名(非空)';

COMMENT ON COLUMN accounts.status IS '账号状态（非空）';

COMMENT ON COLUMN accounts.created_at IS '创建时间(非空)';

COMMENT ON COLUMN accounts.updated_at IS '更新时间(可空)';

-- 角色表
CREATE TABLE roles (
    id UUID NOT NULL,
    role VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    parent_id UUID,
    permissions UUID [],
    status SMALLINT NOT NULL DEFAULT 1,
    CONSTRAINT pk_roles PRIMARY KEY (id),
    CONSTRAINT uk_roles_role UNIQUE (role),
    CONSTRAINT uk_roles_name UNIQUE (name),
    -- 父子角色外键约束：父角色存在子角色时禁止删除（ON DELETE RESTRICT）
    CONSTRAINT fk_roles_parent FOREIGN KEY (parent_id) REFERENCES roles (id) ON DELETE RESTRICT
);
-- 角色表索引
CREATE INDEX idx_roles_status ON roles (status);

COMMENT ON TABLE roles IS '角色表';

COMMENT ON COLUMN roles.id IS '角色ID(主键)';

COMMENT ON COLUMN roles.role IS '角色标识(唯一)';

COMMENT ON COLUMN roles.name IS '角色名称(唯一)';

COMMENT ON COLUMN roles.parent_id IS '父级角色ID(根角色为空)';

COMMENT ON COLUMN roles.permissions IS '权限(可空)';

COMMENT ON COLUMN roles.status IS '状态';

INSERT INTO
    roles (id, role, name, parent_id)
VALUES (
        'a80b3690-50be-463f-81fa-eb135c0a84ae',
        'superadmin',
        '超级管理员',
        NULL
    );


