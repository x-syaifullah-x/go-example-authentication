CREATE TRIGGER IF NOT EXISTS before_insert BEFORE INSERT ON users
FOR EACH ROW 
    BEGIN
        -- SET @TimeMillies = ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000);
        SET @TimeMillies = UNIX_TIMESTAMP();
        IF NEW.create_at IS NOT NULL THEN
            SIGNAL SQLSTATE '45000'    
            SET MESSAGE_TEXT = "Cannot set the `created_at` column, the `created_at` column will be created automatically";
        ELSE
         SET NEW.create_at = @TimeMillies;
        END IF;

        IF NEW.update_at IS NOT NULL THEN
            SIGNAL SQLSTATE "45000"
            SET MESSAGE_TEXT = "The `update_at` column cannot be set when inserting data";
        ELSE
         SET NEW.update_at = @TimeMillies;
        END IF;
        
        IF NEW.delete_at IS NOT NULL THEN
            SIGNAL SQLSTATE "45000"
            SET MESSAGE_TEXT = "The `delete_at` column cannot be set when inserting data";
        ELSE
         SET NEW.delete_at = 0;
        END IF;
    END;