CREATE TABLE device_details
    (
        device_id       BIGINT,
        time            TIMESTAMP WITH TIME ZONE,
        name            VARCHAR NOT NULL,
        customer_id     BIGINT NOT NULL,
        customer_name   VARCHAR NOT NULL,
        PRIMARY KEY(device_id, time)
    );

CREATE TABLE device_heartbeats
    (
        device_id       BIGINT,
        time            TIMESTAMP WITH TIME ZONE,
        is_healthy      BOOLEAN NOT NULL,
        customer_id     BIGINT NOT NULL,
        customer_name   VARCHAR NOT NULL,
        PRIMARY KEY(device_id, time)
    );
