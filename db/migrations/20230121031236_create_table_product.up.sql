CREATE TABLE tb_product
(
    product_id varchar(36),
    name       varchar(100),
    price      bigint,
    quantity   int,
    PRIMARY KEY (product_id),
    INDEX      idx_tb_product_name ( name)
)