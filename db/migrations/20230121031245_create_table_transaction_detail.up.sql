CREATE TABLE tb_transaction_detail
(
    transaction_detail_id varchar(36),
    sub_total_price       bigint,
    price                 bigint,
    quantity              int,
    transaction_id        varchar(36),
    product_id            varchar(36),
    PRIMARY KEY (transaction_detail_id),
    CONSTRAINT fk_tb_product_transaction_details FOREIGN KEY (product_id) REFERENCES tb_product (product_id) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT fk_tb_transaction_transaction_details FOREIGN KEY (transaction_id) REFERENCES tb_transaction (transaction_id) ON DELETE CASCADE ON UPDATE CASCADE
)