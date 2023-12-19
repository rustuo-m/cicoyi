USE cicoyi_user;

CREATE TABLE IF NOT EXISTS user
(
    id         BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3),
    name       LONGTEXT,
    password   LONGTEXT,
    email      LONGTEXT,
    phone      LONGTEXT,
    gender     LONGTEXT,
    birth      LONGTEXT
);

CREATE TABLE IF NOT EXISTS org
(
    id          BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    created_at  DATETIME(3),
    updated_at  DATETIME(3),
    deleted_at  DATETIME(3),
    name        LONGTEXT,
    description LONGTEXT,
    user_id     BIGINT UNSIGNED,
    FOREIGN KEY (user_id) REFERENCES user (id)
);
