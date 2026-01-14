-- 状态表
CREATE TABLE status (
    id SMALLINT NOT NULL,
    name VARCHAR(20) NOT NULL,

    CONSTRAINT pk_status PRIMARY KEY (id),
    CONSTRAINT uk_status_name UNIQUE (name)
);

COMMENT ON TABLE status IS '状态表';

COMMENT ON COLUMN status.id IS '状态ID';

COMMENT ON COLUMN status.name IS '状态名称';

-- 状态
INSERT INTO status (id, name) VALUES 
(1, 'active'),
(2, 'disabled'),
(3, 'locked'),
(4, 'delete');

-- 性别表
CREATE TABLE gender (
    id SMALLINT NOT NULL,
    name VARCHAR(20) NOT NULL,

    CONSTRAINT pk_gender PRIMARY KEY (id),
    CONSTRAINT uk_gender_name UNIQUE (name)
);

COMMENT ON TABLE gender IS '性别表';

COMMENT ON COLUMN gender.id IS '性别ID';

COMMENT ON COLUMN gender.name IS '性别名称';

-- 性别
INSERT INTO gender (id, name) VALUES 
(1, '保密'),
(2, '男'),
(3, '女');

-- 账号表
CREATE TABLE accounts (
    id UUID NOT NULL,
    username VARCHAR(50) NOT NULL,
    nickname VARCHAR(50),
    avatar TEXT NOT NULL,
    email VARCHAR(100),
    phone VARCHAR(20),
    pwd TEXT NOT NULL,
    salt VARCHAR(50),
    slogan VARCHAR(255),
    gender SMALLINT NOT NULL DEFAULT 1,
    status SMALLINT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    create_by UUID NOT NULL,
    updated_at TIMESTAMPTZ,
    update_by UUID,

    CONSTRAINT pk_accounts PRIMARY KEY (id),
    CONSTRAINT uk_accounts_username UNIQUE (username),
    CONSTRAINT fk_accounts_status FOREIGN KEY (status) REFERENCES status (id),
    CONSTRAINT fk_accounts_gender FOREIGN KEY (gender) REFERENCES gender (id)
);
-- 账号状态索引
CREATE INDEX idx_accounts_status ON accounts (status);

COMMENT ON TABLE accounts IS '账号表';

COMMENT ON COLUMN accounts.id IS '账号ID(主键)';

COMMENT ON COLUMN accounts.username IS '用户名(唯一)';

COMMENT ON COLUMN accounts.nickname IS '昵称(可空)';

COMMENT ON COLUMN accounts.avatar IS '头像URL(非空)';

COMMENT ON COLUMN accounts.email IS '邮箱(可空)';

COMMENT ON COLUMN accounts.phone IS '手机号(可空)';

COMMENT ON COLUMN accounts.pwd IS '密码哈希(非空)';

COMMENT ON COLUMN accounts.salt IS '盐(非空)';

COMMENT ON COLUMN accounts.slogan IS '个性签名(非空)';

COMMENT ON COLUMN accounts.status IS '账号状态（非空）';

COMMENT ON COLUMN accounts.gender IS '用户性别（非空）';

COMMENT ON COLUMN accounts.created_at IS '创建时间(非空)';

COMMENT ON COLUMN accounts.create_by IS '创建人(非空)';

COMMENT ON COLUMN accounts.update_by IS '修改人(非空)'

COMMENT ON COLUMN accounts.updated_at IS '更新时间(可空)';


-- 角色权限类型表
CREATE TABLE role_perm_type (
    id SMALLINT NOT NULL,
    name VARCHAR(20) NOT NULL,

    CONSTRAINT pk_rule_type PRIMARY KEY (id),
    CONSTRAINT uk_rule_type_name UNIQUE (name)
);

COMMENT ON TABLE role_perm_type IS '规则表';

COMMENT ON COLUMN role_perm_type.id IS '规则ID';

COMMENT ON COLUMN role_perm_type.name IS '规则名称';

-- 规则类型
INSERT INTO role_perm_type (id, name) VALUES 
(1, '服务'),
(2, '模块'),
(3, '菜单'),
(4, '菜单项'),
(5, '操作');

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

INSERT INTO roles (id, role, name, parent_id) VALUES (
        'a80b3690-50be-463f-81fa-eb135c0a84ae',
        'superadmin',
        '超级管理员',
        NULL
    );

-- 权限表
CREATE TABLE permissions (
    id UUID NOT NULL,
    parent_id UUID,
    type SMALLINT NOT NULL,
    title VARCHAR(50) NOT NULL,
    perm VARCHAR NOT NULL,
    route VARCHAR(50) NOT NULL,
    icon VARCHAR(50),
    desc VARCHAR(255),
    weight SMALLINT NOT NULL DEFAULT 0,
    status SMALLINT NOT NULL DEFAULT 1,
    CONSTRAINT pk_permissions PRIMARY KEY (id),
    CONSTRAINT uk_permissions_name UNIQUE (name),
    CONSTRAINT fk_permissions_type FOREIGN KEY (type) REFERENCES role_perm_type (id),
    CONSTRAINT fk_permissions_parent FOREIGN KEY (parent_id) REFERENCES permissions (id) ON DELETE RESTRICT
)

COMMENT ON TABLE permissions IS '权限表';

COMMENT ON COLUMN permissions.id IS '权限ID(主键)';

COMMENT ON COLUMN permissions.parent_id IS '父级权限ID(根权限为空)';

COMMENT ON COLUMN permissions.type IS '权限类型';

COMMENT ON COLUMN permissions.title IS '权限名称';

COMMENT ON COLUMN permissions.perm IS '权限标识';

COMMENT ON COLUMN permissions.route IS '权限路由';

COMMENT ON COLUMN permissions.icon IS '权限图标';

COMMENT ON COLUMN permissions.desc IS '权限描述';

COMMENT ON COLUMN permissions.weight IS '权重';

COMMENT ON COLUMN permissions.status IS '状态';

-- 角色权限表
CREATE TABLE role_permissions (
    role_id UUID NOT NULL,
    permission_id UUID NOT NULL,
    CONSTRAINT pk_role_permissions PRIMARY KEY (role_id, permission_id),
    CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE,
    CONSTRAINT fk_role_permissions_permission FOREIGN KEY (permission_id) REFERENCES permissions (id) ON DELETE CASCADE
);

COMMENT ON TABLE role_permissions IS '角色权限表';

COMMENT ON COLUMN role_permissions.role_id IS '角色ID';

COMMENT ON COLUMN role_permissions.permission_id IS '权限ID';

-- 用户角色表
CREATE TABLE user_roles (
    user_id UUID NOT NULL,
    role_id UUID NOT NULL,
    CONSTRAINT pk_user_roles PRIMARY KEY (user_id, role_id),
    CONSTRAINT fk_user_roles_user FOREIGN KEY (user_id) REFERENCES accounts (id) ON DELETE CASCADE,
    CONSTRAINT fk_user_roles_role FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE
);
