CREATE TRIGGER IF NOT EXISTS before_update BEFORE UPDATE ON users
FOR EACH ROW
BEGIN
    -- SET @TimeMillies = ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000);
    SET @TimeMillies = UNIX_TIMESTAMP();

    IF OLD.create_at IS NOT NULL AND NEW.create_at <> OLD.create_at THEN
        SIGNAL SQLSTATE "45000"
        SET MESSAGE_TEXT = "The `create_at` column cannot be written when updating";
    END IF;

    IF OLD.update_at IS NOT NULL AND NEW.update_at <> OLD.update_at THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = "Cannot set the `update_at` column, the `update_at` column will be created automatically";
    ELSE
        SET NEW.update_at = @TimeMillies;
    END IF;

    IF NEW.delete_at = 0 AND NEW.delete_at <> OLD.delete_at THEN
        SET NEW.update_at = @TimeMillies;
        SET NEW.create_at = @TimeMillies;
    END IF;
    SET @TimeMillies = NULL;
END;