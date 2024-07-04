BEGIN;

    CREATE TABLE PROFILE (
        pass_serie VARCHAR(4) NOT NULL,
        pass_number VARCHAR(6) NOT NULL,
        name VARCHAR(255) NOT NULL,
        surname VARCHAR(255) NOT NULL,
        patronymic VARCHAR(255),
        address VARCHAR(255) NOT NULL,
        PRIMARY KEY (pass_serie, pass_number)
    );

    COMMIT;
END;