CREATE TABLE tb_user_role
(
    user_role_id varchar(36),
    role         varchar(10),
    username     varchar(100),
    PRIMARY KEY (user_role_id),
    CONSTRAINT fk_tb_user_user_roles FOREIGN KEY (username) REFERENCES tb_user (username) ON DELETE CASCADE ON UPDATE CASCADE
)